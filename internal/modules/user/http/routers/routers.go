package routers

import (
	"net/http"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/services"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/http/handlers"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/infra/repositories"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/httprate"
	"github.com/jackc/pgx/v5"
)

type UserRouter struct {
	Db            pgx.Tx
	userHandler   *handlers.UserHandler
	jwtExpiriesIn int
	Router        *router.Router
}

func NewUserRouter(db pgx.Tx, router *router.Router) *UserRouter {
	u := &UserRouter{
		Db: db,
		Router: router,
	}
	u.userHandler = u.init()
	return u
}

func (u *UserRouter) init() *handlers.UserHandler {
	userRepository := repositories.NewUserRepository(u.Db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	return userHandler
}

func (u *UserRouter) InitializeUserRoutes(r chi.Router) {
	r.Route("/authn", func(r chi.Router) {
		u.Router.Method("POST").Prefix("/authn").InitializeRoute(r, "/create", u.userHandler.Create)
		r.Group(func(r chi.Router) {
			r.Use(httprate.Limit(
				5,
				60*time.Minute,
				httprate.WithKeyFuncs(httprate.KeyByIP, httprate.KeyByEndpoint),
				httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
					http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
				}),
			))
			u.Router.Method("POST").Prefix("/authn").InitializeRoute(r, "/jwt", u.userHandler.Authentication)
		})
	})
}

func (u *UserRouter) UserRoutes(r chi.Router) {
	r.Route("/user", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(httprate.Limit(
				10,
				60*time.Minute,
				httprate.WithKeyFuncs(httprate.KeyByIP, httprate.KeyByEndpoint),
				httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
					http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
				}),
			))
			u.Router.Method("GET").Prefix("/user").InitializeRoute(r, "/me", u.userHandler.Me)
			u.Router.Method("PUT").Prefix("/user").InitializeRoute(r, "/update", u.userHandler.Update)
		})

		r.Group(func(r chi.Router) {
			r.Use(httprate.Limit(
				3,
				60*time.Minute,
				httprate.WithKeyFuncs(httprate.KeyByIP, httprate.KeyByEndpoint),
				httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
					http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
				}),
			))
			u.Router.Method("DELETE").Prefix("/user").InitializeRoute(r, "/del", u.userHandler.Delete)
			u.Router.Method("PATCH").Prefix("/user").InitializeRoute(r, "/authz/{id}", u.userHandler.AdminAuthz)
		})
	})
}
