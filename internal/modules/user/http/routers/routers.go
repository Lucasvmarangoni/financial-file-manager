package routers

import (
	"net/http"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/management"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/services"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/http/handlers"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/infra/repositories"
	"github.com/streadway/amqp"

	"github.com/Lucasvmarangoni/logella/router"

	"github.com/Lucasvmarangoni/financial-file-manager/pkg/cache"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/queue"
	"github.com/go-chi/chi"
	"github.com/go-chi/httprate"
	"github.com/jackc/pgx/v5"
)

type UserRouter struct {
	Conn           *pgx.Conn
	userHandler    *handlers.UserHandler
	Router         *router.Router
	RabbitMQ       *queue.RabbitMQ
	MessageChannel chan amqp.Delivery
	Memcached      *cache.Memcached[*entities.User]
}

func NewUserRouter(
	conn *pgx.Conn,
	router *router.Router,
	rabbitMQ *queue.RabbitMQ,
	messageChannel chan amqp.Delivery,
	mencached *cache.Memcached[*entities.User],
) *UserRouter {
	u := &UserRouter{
		Conn:           conn,
		Router:         router,
		RabbitMQ:       rabbitMQ,
		MessageChannel: messageChannel,
		Memcached:      mencached,
	}
	u.userHandler = u.init()
	return u
}

func (u *UserRouter) init() *handlers.UserHandler {
	returnChannel := make(chan error)

	userRepository := repositories.NewUserRepository(u.Conn)
	userService := services.NewUserService(userRepository, u.RabbitMQ, u.MessageChannel, returnChannel, u.Memcached)
	userHandler := handlers.NewUserHandler(userService)

	userManagement := management.NewManagement(userRepository, u.RabbitMQ)

	go func() {
		userManagement.CreateManagement(u.MessageChannel, returnChannel)
	}()

	return userHandler
}

func (u *UserRouter) InitializeUserRoutes(r chi.Router) {
	prefix := "/authn"
	r.Route(prefix, func(r chi.Router) {
		u.Router.Method("POST").Prefix(prefix).InitializeRoute(r, "/create", u.userHandler.Create)
		u.Router.Method("GET").Prefix(prefix).InitializeRoute(r, "/google", u.userHandler.Oauth)
		r.Group(func(r chi.Router) {
			r.Use(httprate.Limit(
				10,
				60*time.Minute,
				httprate.WithKeyFuncs(httprate.KeyByIP, httprate.KeyByEndpoint),
				httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
					http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
				}),
			))
			u.Router.Method("POST").Prefix(prefix).InitializeRoute(r, "/jwt", u.userHandler.Authentication)
		})
	})
}

func (u *UserRouter) UserRoutes(r chi.Router) {
	prefix := "/user"
	r.Route(prefix, func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(httprate.Limit(
				10,
				60*time.Minute,
				httprate.WithKeyFuncs(httprate.KeyByIP, httprate.KeyByEndpoint),
				httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
					http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
				}),
			))
			u.Router.Method("GET").Prefix(prefix).InitializeRoute(r, "/me", u.userHandler.Me)
			u.Router.Method("PUT").Prefix(prefix).InitializeRoute(r, "/update", u.userHandler.Update)
		})

		r.Group(func(r chi.Router) {
			r.Use(httprate.Limit(
				5,
				60*time.Minute,
				httprate.WithKeyFuncs(httprate.KeyByIP, httprate.KeyByEndpoint),
				httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
					http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
				}),
			))
			u.Router.Method("DELETE").Prefix(prefix).InitializeRoute(r, "/del", u.userHandler.Delete)
		})
	})
}
