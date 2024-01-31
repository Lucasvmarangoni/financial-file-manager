package middlewares

import (
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/Lucasvmarangoni/financial-file-manager/config"
	errors "github.com/Lucasvmarangoni/logella/err"
	"github.com/casbin/casbin/v2"
	"github.com/go-chi/jwtauth"
	"github.com/rs/zerolog/log"
)

type Authorization struct {
	enforcer *casbin.Enforcer
	policy   string
	model    string
	admins   []string
}

func NewAuthorization(policy string, model string) *Authorization {
	a := &Authorization{
		policy: policy,
		model:  model,
	}
	a.init()
	return a
}

func (a *Authorization) add(adminID string) string {
	content, err := os.ReadFile(a.policy)
	if err != nil {
		panic(errors.ErrCtx(err, "os.ReadFile"))
	}
	env_id := config.GetEnv(strings.ToLower(adminID)).(string)
	a.admins = append(a.admins, env_id)
	newContent := strings.ReplaceAll(string(content), adminID, env_id)
	return newContent
}

func (a *Authorization) init() {

	newContent := a.add("ADMIN_1")
	newContent = a.add("READ_1")

	err := os.WriteFile(a.policy, []byte(newContent), 0644)
	if err != nil {
		panic(errors.ErrCtx(err, "os.WriteFile"))
	}
	a.enforcer, err = casbin.NewEnforcer(a.model, a.policy)
	if err != nil {
		log.Print("Failed to create enforcer:", err)
	}
}

func (a *Authorization) Authorizer() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			var role string
			id := getSub(w, r)

			if !slices.Contains(a.admins, id) {
				role = "member"
			} else {
				role = id
			}
			log.Print(role)
			method := r.Method
			path := r.URL.Path

			if ok, _ := a.enforcer.Enforce(role, path, method); !ok {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func getSub(w http.ResponseWriter, r *http.Request) string {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		http.Error(w, "Failed to get JWT claims", http.StatusInternalServerError)
	}
	sub, ok := claims["sub"].(string)
	if !ok {
		http.Error(w, "Sub claim is missing or not a string", http.StatusInternalServerError)
	}
	return sub
}
