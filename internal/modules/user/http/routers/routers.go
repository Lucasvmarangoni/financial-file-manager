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
	userHandler   *handlers.UserHandler
	jwtExpiriesIn int
	tokenAuth     *jwtauth.JWTAuth	
}

var (
	router chi.Router
	method string
)

func NewUserRouter(db pgx.Tx, chi *chi.Mux, jwtExpiriesIn int, tokenAuth *jwtauth.JWTAuth) *UserRouter {
	u := &UserRouter{
		Db:            db,
		Chi:           chi,
		jwtExpiriesIn: jwtExpiriesIn,
		tokenAuth:     tokenAuth,
	}
	u.userHandler = u.init()
	return u
}

func (u *UserRouter) init() *handlers.UserHandler {
	userRepository := repositories.NewUserRepository(u.Db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService, u.tokenAuth, u.jwtExpiriesIn)
	return userHandler
}

func (u *UserRouter) InitializeUserRoutes() {
	u.Chi.Route("/user", func(r chi.Router) {
		router = r
		u.Method("POST").InitializeRoute("/", u.userHandler.Create)
		u.Method("POST").InitializeRoute("/authn", u.userHandler.Authentication)
	})
}

func (u *UserRouter) UserRoutes(r chi.Router) {
	u.Method("GET").InitializeRoute("/me", u.userHandler.Me)
	// r.Put("/update", u.userHandler.Update)
}
