package main

import (
	"context"
	"net/http"

	"time"

	_ "github.com/Lucasvmarangoni/financial-file-manager/api"
	"github.com/Lucasvmarangoni/financial-file-manager/config"

	"github.com/Lucasvmarangoni/financial-file-manager/pkg/cache"
	logger "github.com/Lucasvmarangoni/financial-file-manager/pkg/log"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/queue"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/infra/database"
	middlewares "github.com/Lucasvmarangoni/financial-file-manager/internal/middleware"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/http/routers"
	errors "github.com/Lucasvmarangoni/logella/err"
	"github.com/Lucasvmarangoni/logella/router"
	"github.com/streadway/amqp"

	crdbpgx "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"

	// "github.com/Lucasvmarangoni/financial-file-manager/internal/rpc"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/jackc/pgx/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

var db database.Config

func init() {
	logger.Config()

	db.DbName = config.GetEnvString("database", "name")
	db.Port = config.GetEnvString("database", "port")
	db.User = config.GetEnvString("database", "user")
	db.Password = config.GetEnvString("database", "password")
	db.SSLMode = config.GetEnvString("database", "ssl_mode")
}

// @title           Financial File Manager
// @version         1.0
// @description
// @termsOfService  http://swagger.io/terms/

// @contact.name   Lucas V Marangoni
// @contact.url    https://lucasvmarangoni.vercel.app/
// @contact.email  lucasvm.ti@gmail.com

// @license.name   MIT

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	conn, err := Database(ctx)
	errors.FailOnErrLog(err, "Database(ctx)", "Failed exec Database")
	// rpc.Connect()

	r := chi.NewRouter()
	mc := Cache[*entities.User]()

	messageChannel, rabbitMQ, ch := Queues()
	defer ch.Close()
	Http(conn, r, messageChannel, rabbitMQ, ch, mc)

	err = http.ListenAndServeTLS(":8000", "/app/nginx/cert.pem", "/app//nginx/key.pem", r)
	errors.FailOnErrLog(err, "http.ListenAndServe", "Failed server listen")
}

func Database(ctx context.Context) (*pgx.Conn, error) {
	conn, err := db.Connect(ctx)
	if err != nil {
		return nil, errors.ErrCtx(err, "db.Connect")
	}

	err = crdbpgx.ExecuteTx(ctx, conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		repo := database.NewTableRepository(tx)
		return repo.InitTables(ctx)
	})
	if err != nil {
		return nil, errors.ErrCtx(err, "repo.InitTables")
	}
	return conn, nil
}

func Http(
	conn *pgx.Conn,
	r *chi.Mux,
	messageChannel chan amqp.Delivery,
	rabbitMQ *queue.RabbitMQ,
	ch *amqp.Channel,
	mc *cache.Memcached[*entities.User],
) {
	tokenAuth := config.GetTokenAuth()
	jwtExpiresIn := config.GetEnvInt("jwt", "expiredIn")
	// jwtExpiresIn, err := strconv.Atoi(jwtExpiresInStr)
	// if err != nil {
	// 	jwtExpiresIn = 50
	// 	log.Warn().Err(errors.ErrCtx(err, "strconv.Atoi")).Str("Source", "server.go").Str("Func", "Rest").Msg("Failed to convert jwtExpiresIn into int. Default value has been applied.")
	// }
	redisPassword := config.GetEnvString("password", "redis")

	mw := middlewares.NewAuthorization("config/casbin/policy.csv", "config/casbin/model.conf")

	router := router.NewRouter()
	userRouter := routers.NewUserRouter(conn, router, rabbitMQ, messageChannel, mc)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://192.168.96.1"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", config.GetTokenAuth()))
	r.Use(middleware.WithValue("JwtExpiresIn", jwtExpiresIn))

	userRouter.InitializeUserRoutes(r)

	r.Route("/", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Use(mw.Authorizer())
		r.Use(middlewares.NewUserRateLimit("redis:6379", redisPassword).Handler())
		userRouter.UserRoutes(r)

	})
	userRouter.Router.Method("GET").Prefix("").InitializeRoute(r, "/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))
}

func Queues() (chan amqp.Delivery, *queue.RabbitMQ, *amqp.Channel) {
	messageChannel := make(chan amqp.Delivery)
	rabbitMQ := queue.NewRabbitMQ()
	ch := rabbitMQ.Connect()

	return messageChannel, rabbitMQ, ch
}

func Cache[T cache.Entity]() *cache.Memcached[T] {
	mc := cache.NewMemcached[T]("localhost:11211", "localhost:11212")
	return mc
}
