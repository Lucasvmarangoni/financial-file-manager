package routers

import (
	"net/http"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/config"
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
	Memcached_1    *cache.Memcached[bool]
}

func NewUserRouter(
	conn *pgx.Conn,
	router *router.Router,
	rabbitMQ *queue.RabbitMQ,
	messageChannel chan amqp.Delivery,
	mencached *cache.Memcached[*entities.User],
	memcached_1 *cache.Memcached[bool],
) *UserRouter {
	u := &UserRouter{
		Conn:           conn,
		Router:         router,
		RabbitMQ:       rabbitMQ,
		MessageChannel: messageChannel,
		Memcached:      mencached,
		Memcached_1:    memcached_1,
	}
	u.userHandler = u.init()
	return u
}

func (u *UserRouter) init() *handlers.UserHandler {
	userRepository := repositories.NewUserRepository(u.Conn)
	userService := services.NewUserService(userRepository, u.RabbitMQ, u.MessageChannel, u.Memcached, *u.Memcached_1)
	userManagement := management.NewManagement(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	u.RabbitMQ.Consume(u.MessageChannel, config.GetEnvString("rabbitMQ", "routingkey_userCreate"))
	for i := 0; i < config.GetEnvInt("concurrency", "create_management"); i++ {
		go userManagement.CreateManagement(u.MessageChannel)
	}
	return userHandler
}

func (u *UserRouter) InitializeUserRoutes(r chi.Router) {
	prefix := "/authn"
	r.Route(prefix, func(r chi.Router) {
		u.Router.Method("POST").Prefix(prefix).InitializeRoute(r, "/create", u.userHandler.Create)
		r.Group(func(r chi.Router) {
			r.Use(httprate.Limit(
				2,
				60*time.Minute,
				httprate.WithKeyFuncs(httprate.KeyByRealIP, httprate.KeyByEndpoint),
				httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
					http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
				}),
			))
			u.Router.Method("POST").Prefix(prefix).InitializeRoute(r, "/jwt", u.userHandler.Authentication)
		})

	})
}

func (u *UserRouter) UserRoutes(r chi.Router) {

	prefix := "/totp"
	r.Route(prefix, func(r chi.Router) {
		u.Router.Method("GET").Prefix(prefix).InitializeRoute(r, "/generate", u.userHandler.TwoFactorAuthn)
		u.Router.Method("POST").Prefix(prefix).InitializeRoute(r, "/verify/{is_validate}", u.userHandler.TwoFactorVerify)
		u.Router.Method("PATCH").Prefix(prefix).InitializeRoute(r, "/disable", u.userHandler.TwoFactorDisable)
	})

	prefix = "/user"
	r.Route(prefix, func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(httprate.Limit(
				5,
				30*time.Minute,
				httprate.WithKeyFuncs(httprate.KeyByRealIP, httprate.KeyByEndpoint),
				httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
					http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
				}),
			))
			u.Router.Method("GET").Prefix(prefix).InitializeRoute(r, "/me", u.userHandler.Me)
			u.Router.Method("PUT").Prefix(prefix).InitializeRoute(r, "/update", u.userHandler.Update)
		})

		r.Group(func(r chi.Router) {
			r.Use(httprate.Limit(
				1,
				60*time.Minute,
				httprate.WithKeyFuncs(httprate.KeyByRealIP, httprate.KeyByEndpoint),
				httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
					http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
				}),
			))
			u.Router.Method("DELETE").Prefix(prefix).InitializeRoute(r, "/del", u.userHandler.Delete)
		})
	})
}
