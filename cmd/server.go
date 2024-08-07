package main

import (
	"context"
	"net/http"

	"time"

	_ "github.com/Lucasvmarangoni/financial-file-manager/api"
	"github.com/Lucasvmarangoni/financial-file-manager/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/Lucasvmarangoni/financial-file-manager/pkg/cache"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/metric"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/queue"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/infra/database"
	middlewares "github.com/Lucasvmarangoni/financial-file-manager/internal/middleware"

	// observability_routers "github.com/Lucasvmarangoni/financial-file-manager/internal/modules/observability/routers"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	user_routers "github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/http/routers"
	logger "github.com/Lucasvmarangoni/logella/config/log"
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
	logger.ConfigDefault()

	db.DbName = config.GetEnvString("database", "name")
	db.Port = config.GetEnvString("database", "port")
	db.User = config.GetEnvString("database", "user")
	db.Password = config.ReadSecretString(config.GetEnvString("database", "password"))
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
	mc_1 := Cache[bool]()

	messageChannel, rabbitMQ, ch := Queues()
	defer ch.Close()

	Http(conn, r, messageChannel, rabbitMQ, ch, mc, mc_1)

	certFile, keyFile := "/run/secrets/cert.pem", "/run/secrets/key.pem"

	err = http.ListenAndServeTLS(":8000", certFile, keyFile, r)
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
	mc_1 *cache.Memcached[bool],
) {
	tokenAuth := config.GetTokenAuth()
	jwtExpiresIn := config.GetEnvInt("jwt", "expiredIn")
	redisPassword := config.ReadSecretString(config.GetEnvString("password", "redis"))

	mw := middlewares.NewAuthorization("config/casbin/policy.csv", "config/casbin/model.conf")

	router := router.NewRouter()
	userRouter := user_routers.NewUserRouter(conn, router, rabbitMQ, messageChannel, mc, mc_1)
	// ObservabilityRouter := observability_routers.NewObservability(router)

	metricService := Metric()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://localhost"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Use(middlewares.WAF())
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middlewares.Metrics(metricService))
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
	userRouter.Router.Method("GET").Prefix("").InitializeRoute(r, "/docs/*", httpSwagger.Handler(httpSwagger.URL("https://localhost:443/docs/doc.json")))

	router.Method("POST").Prefix("").InitializeRoute(r, "/metric", func(w http.ResponseWriter, r *http.Request) {
		promhttp.Handler().ServeHTTP(w, r)
	})
	router.Method("GET").Prefix("").InitializeRoute(r, "/exporter", func(w http.ResponseWriter, r *http.Request) {
		promhttp.Handler().ServeHTTP(w, r)
	})

}

func Queues() (chan amqp.Delivery, *queue.RabbitMQ, *amqp.Channel) {
	messageChannel := make(chan amqp.Delivery)
	rabbitMQ := queue.NewRabbitMQ()
	ch := rabbitMQ.Connect()

	return messageChannel, rabbitMQ, ch
}

func Cache[T cache.Entity]() *cache.Memcached[T] {
	mc := cache.NewMemcached[T]("memcached-1:11211")
	return mc
}

func Metric() *metric.Service {
	metricService, err := metric.NewPrometheusService()
	if err != nil {
		errors.FailOnErrLog(err, "metric.NewPrometheusService", "")
	}
	appMetric := metric.NewCLI("search")
	appMetric.Started()

	appMetric.Finished()
	err = metricService.SaveCLI(appMetric)
	if err != nil {
		errors.FailOnErrLog(err, "metricService.SaveCLI", "")
	}

	return metricService
}
