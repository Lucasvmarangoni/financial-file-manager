package routers

import (
	"github.com/Lucasvmarangoni/financial-file-manager/config"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/http/handlers"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/infra/repositories"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/jackc/pgx/v5"
)

type UserRouter struct {
	Db          pgx.Tx
	Chi         *chi.Mux
	userHandler handlers.UserHandler
}

func NewUserRouter(db pgx.Tx, chi *chi.Mux) *UserRouter {
	return &UserRouter{
		Db:  db,
		Chi: chi,
	}
}

func (u *UserRouter) init() {
	userRepository := repositories.NewUserRepository(u.Db)
	tokenAuth := config.GetEnv("jwt.tokenAuth").(*jwtauth.JWTAuth)
	jwtExpiriesIn := config.GetEnv("jwt.expired_in").(int)
	u.userHandler = handlers.UserHandler{
		Repository:    userRepository,
		Jwt:           tokenAuth,
		JwtExpiriesIn: jwtExpiriesIn,
	}
}

func (u *UserRouter) InitializeUserRoutes() {
	u.Chi.Route("/user", func(r chi.Router) {
		r.Post("/", u.userHandler.Create)
		r.Post("/authn", u.userHandler.Authentication)
	})
}

func (u *UserRouter) UserRoutes() {
	u.Chi.Route("/user", func(r chi.Router) {
		r.Get("/", u.userHandler.Me)
		// r.Put("/update", u.userHandler.Update)
	})
}
