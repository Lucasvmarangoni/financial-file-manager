package main

import (
	"context"
	"net/http"

	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/config"
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
		log.Fatal().Err(err).Str("File", "server.go").Str("Method", "Database").Msg("Failed to exec Database")
	}

	// rpc.Connect()

	// messageChannel := make(chan amqp.Delivery)

	// rabbitMQ := queue.NewRabbitMQ()
	// ch := rabbitMQ.Connect()
	// defer ch.Close()
	// rabbitMQ.Consume(messageChannel)

	r := chi.NewRouter()
	Web(r, tx)

	http.ListenAndServe(":8000", r)
}

func Database(ctx context.Context) (pgx.Tx, error) {
	dbConnection, err := db.Connect(ctx)
	if err != nil {
		return nil, logger.NewError(err, "db.Connect")
	}

	tx, err := dbConnection.Begin(ctx)
	if err != nil {
		dbConnection.Close(ctx)
		return nil, logger.NewError(err, "dbConnection.Begin")
	}

	repo := database.NewTableRepository(tx)
	err = repo.InitTables(ctx)
	if err != nil {
		return nil, logger.NewError(err, "repo.InitTables")
	}
	return tx, nil
}

func Web(r *chi.Mux, tx pgx.Tx) {
	tokenAuth := config.GetTokenAuth()

	userRouter := routers.NewUserRouter(tx, r)

	r.Use(middleware.Logger)

	userRouter.InitializeUserRoutes()

	r.Route("/api", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		userRouter.UserRoutes(r)
	})
}
