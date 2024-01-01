package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/config"
	// "github.com/Lucasvmarangoni/financial-file-manager/internal/common/queue"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/infra/database"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/http/routers"
	// "github.com/Lucasvmarangoni/financial-file-manager/internal/rpc"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	// "github.com/streadway/amqp"
)

var db database.Config

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

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
		log.Fatal(err)
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
		return nil, err
	}
	defer dbConnection.Close(ctx)

	tx, err := dbConnection.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

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
