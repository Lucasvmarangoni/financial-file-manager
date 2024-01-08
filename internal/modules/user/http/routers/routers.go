package routers

import (
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/services"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/http/handlers"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/infra/repositories"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/jackc/pgx/v5"
)

type UserRouter struct {
	Db            pgx.Tx
	Chi           *chi.Mux
	method        string
	userHandler   *handlers.UserHandler
	jwtExpiriesIn int
}

func NewUserRouter(db pgx.Tx, chi *chi.Mux, jwtExpiriesIn int, tokenAuth *jwtauth.JWTAuth) *UserRouter {
	u := &UserRouter{
		Db:            db,
		Chi:           chi,
		jwtExpiriesIn: jwtExpiriesIn,
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

func (u *UserRouter) InitializeUserRoutes() {
	u.Chi.Route("/user", func(r chi.Router) {
		u.Method("POST").InitializeRoute(r, "/", u.userHandler.Create)
		u.Method("POST").InitializeRoute(r, "/authn", u.userHandler.Authentication)
	})
}

func (u *UserRouter) UserRoutes(r chi.Router) {
	u.Method("GET").InitializeRoute(r, "/me", u.userHandler.Me)
	r.Route("/user", func(r chi.Router) {
		u.Method("PUT").InitializeRoute(r, "/update", u.userHandler.Me)
	})
}
