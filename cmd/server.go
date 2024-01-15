package main

import (
	"context"
	"net/http"
	"strconv"

	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/config"
	_ "github.com/Lucasvmarangoni/financial-file-manager/docs"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/errors"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/http"
	logger "github.com/Lucasvmarangoni/financial-file-manager/pkg/log"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/queue"
	"github.com/streadway/amqp"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/infra/database"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/http/routers"

	// "github.com/Lucasvmarangoni/financial-file-manager/internal/rpc"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger"
)

var db database.Config

func init() {
	logger.Config()

	db.DbName = config.GetEnv("database_name").(string)
	db.Port = config.GetEnv("database_port").(string)
	db.User = config.GetEnv("database_user").(string)
	db.Password = config.GetEnv("database_password").(string)
	db.SSLMode = config.GetEnv("database_ssl_mode").(string)
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

	tx, err := Database(ctx)
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("Failed exec Database")
	}
	// rpc.Connect()

	r := chi.NewRouter()

	messageChannel, rabbitMQ, ch := Queues()
	defer ch.Close()
	Http(tx, r, messageChannel, rabbitMQ, ch)
	

	err = http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("Failed server listen")
	}

}

func Database(ctx context.Context) (pgx.Tx, error) {
	dbConnection, err := db.Connect(ctx)
	if err != nil {
		return nil, errors.NewError(err, "db.Connect")
	}

	tx, err := dbConnection.Begin(ctx)
	if err != nil {
		dbConnection.Close(ctx)
		return nil, errors.NewError(err, "dbConnection.Begin")
	}

	repo := database.NewTableRepository(tx)
	err = repo.InitTables(ctx)
	if err != nil {
		return nil, errors.NewError(err, "repo.InitTables")
	}
	return tx, nil
}

func Http(tx pgx.Tx, r *chi.Mux, messageChannel chan amqp.Delivery, rabbitMQ *queue.RabbitMQ, ch *amqp.Channel) {
	tokenAuth := config.GetTokenAuth()
	jwtExpiresInStr := config.GetEnv("jwt_expiredIn").(string)
	jwtExpiresIn, err := strconv.Atoi(jwtExpiresInStr)
	if err != nil {
		jwtExpiresIn = 50
		log.Warn().Err(errors.NewError(err, "strconv.Atoi")).Str("Source", "server.go").Str("Func", "Rest").Msg("Failed to convert jwtExpiresIn into int. Default value has been applied.")
	}

	router := router.NewRouter()
	userRouter := routers.NewUserRouter(tx, router, rabbitMQ, messageChannel)

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", config.GetTokenAuth()))
	r.Use(middleware.WithValue("JwtExpiresIn", jwtExpiresIn))
	userRouter.InitializeUserRoutes(r)

	r.Route("/", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
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
