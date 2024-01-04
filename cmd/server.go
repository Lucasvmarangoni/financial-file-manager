package main

import (
	"context"
	"net/http"
	"strconv"

	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/config"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/errors"
	logger "github.com/Lucasvmarangoni/financial-file-manager/pkg/log"

	// "github.com/Lucasvmarangoni/financial-file-manager/internal/common/queue"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/infra/database"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/http/routers"

	// "github.com/Lucasvmarangoni/financial-file-manager/internal/rpc"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/jackc/pgx/v5"

	"github.com/rs/zerolog/log"
	// "github.com/streadway/amqp"
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

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tx, err := Database(ctx)
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("Failed exec Database")
	}

	// rpc.Connect()

	// messageChannel := make(chan amqp.Delivery)

	// rabbitMQ := queue.NewRabbitMQ()
	// ch := rabbitMQ.Connect()
	// defer ch.Close()
	// rabbitMQ.Consume(messageChannel)

	r := chi.NewRouter()
	Rest(tx, r)

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

func Rest(tx pgx.Tx, r *chi.Mux) {
	tokenAuth := config.GetTokenAuth()
	jwtExpiresInStr := config.GetEnv("jwt_expiredIn").(string)
	jwtExpiresIn, err := strconv.Atoi(jwtExpiresInStr)
	if err != nil {
		jwtExpiresIn = 50
		log.Warn().Err(errors.NewError(err, "strconv.Atoi")).Str("Source", "server.go").Str("Func", "Rest").Msg("Failed to convert jwtExpiresIn into int. Default value has been applied.")
	}
	userRouter := routers.NewUserRouter(tx, r, jwtExpiresIn, tokenAuth)

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", config.GetTokenAuth()))
	r.Use(middleware.WithValue("JwtExperesIn", jwtExpiresIn))

	userRouter.InitializeUserRoutes()

	r.Route("/api", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		userRouter.UserRoutes(r)
	})
}
