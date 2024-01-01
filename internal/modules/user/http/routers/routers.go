package routers

import (
	"strconv"

	"github.com/Lucasvmarangoni/financial-file-manager/config"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/services"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/http/handlers"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/infra/repositories"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5"
)

type UserRouter struct {
	Db          pgx.Tx
	Chi         *chi.Mux
	userHandler *handlers.UserHandler
}

func NewUserRouter(db pgx.Tx, chi *chi.Mux) *UserRouter {
	return &UserRouter{
		Db:  db,
		Chi: chi,
	}
}

func Init(db pgx.Tx) *handlers.UserHandler {
	tokenAuth := config.GetTokenAuth()
	jwtExpiriesIn, err := strconv.Atoi(config.GetEnv("jwt_expiredIn").(string))
	if err != nil {
		panic(err)
	}
	userRepository := repositories.NewUserRepository(db)	
	createService := services.NewCreateService(userRepository)
	userHandler := handlers.NewUserHandler(createService, tokenAuth, jwtExpiriesIn)
	return userHandler
}

func (u *UserRouter) InitializeUserRoutes() {
	u.userHandler = Init(u.Db)
	u.Chi.Route("/user", func(r chi.Router) {
		r.Post("/", u.userHandler.Create)
		// r.Post("/authn", u.userHandler.Authentication)
	})
}

func (u *UserRouter) UserRoutes(r chi.Router) {
	r.Get("/me", u.userHandler.Me)
	// r.Put("/update", u.userHandler.Update)
}
