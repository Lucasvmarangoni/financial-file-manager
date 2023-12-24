// FIZ ESSA LIB ANTES DE APRENDER A UTILIZAR A VIPER / ESTOU MANTENDO NO CÓDIGO APENAS PARA REGISTRO
// package config

// import (
// 	"fmt"
// 	"os"
// 	"strconv"

// 	"github.com/joho/godotenv"
// 	zerolog "github.com/rs/zerolog/log"
// )

// type Env struct {
// 	Database
// 	RabbitMQ
// 	Grpc
// 	Storage	
// }

// func (env *Env) Config() {

// 	if err := godotenv.Load(); err != nil {
// 		zerolog.Fatal().Err(err).Str("file", ".env").Msg("Error loading .env file")
// 		os.Exit(1)
// 	}

// 	database := env.Database
// 	verifyString(database.Name, "DATABASE_NAME")
// 	verifyString(database.User, "DATABASE_USER")
// 	verifyString(database.Password, "DATABASE_PASSWORD")
// 	verifyString(database.Ssl_mode, "DATABASE_SSL_MODE")
// 	verifyString(database.Port, "DATABASE_PORT")

// 	rabbitMQ := env.RabbitMQ
// 	verifyString(rabbitMQ.User, "RABBITMQ_DEFAULT_USER")
// 	verifyString(rabbitMQ.Password, "RABBITMQ_DEFAULT_PASS")
// 	verifyString(rabbitMQ.Host, "RABBITMQ_DEFAULT_HOST")
// 	verifyString(rabbitMQ.Port, "RABBITMQ_DEFAULT_PORT")
// 	verifyString(rabbitMQ.Vhost, "RABBITMQ_DEFAULT_VHOST")
// 	verifyString(rabbitMQ.Consumer_name, "RABBITMQ_CONSUMER_NAME")
// 	verifyString(rabbitMQ.Consumer_queue_name, "RABBITMQ_CONSUMER_QUEUE_NAME")
// 	verifyString(rabbitMQ.Notification_ex, "RABBITMQ_NOTIFICATION_EX")
// 	verifyString(rabbitMQ.Notification_routing_key, "RABBITMQ_NOTIFICATION_ROUTING_KEY")
// 	verifyString(rabbitMQ.Dlx, "RABBITMQ_DLX")

// 	grpc := env.Grpc
// 	verifyInt(grpc.Port, "GRPC_PORT")

// 	storage := env.Storage
// 	verifyString(storage.Local_path, "LOCAL_STORAGE_PATH")
// 	verifyString(storage.Local_file, "LOCAL_FILE")

// 	management := env.Management
// 	verifyInt(management.Concurrency, "CONCURRENCY_WORKERS")
// }

// func verifyString(variable string, env string) error {
// 	variable = os.Getenv(env)
// 	if variable == "" {
// 		err := fmt.Errorf("%s is not set or is not a valid string", env)
// 		zerolog.Error().Msg(err.Error())
// 		return err
// 	}
// 	return nil
// }

// func verifyInt(variable int, env string) error {
// 	var err error
// 	variable, err = strconv.Atoi(os.Getenv(env))
// 	if err != nil {
// 		zerolog.Error().Msgf("%s is not set or is not a valid integer", env)
// 		return err
// 	}
// 	return nil
// }

// func verifyBool(variable bool, env string) error {
// 	var err error
// 	variable, err = strconv.ParseBool(os.Getenv(env))
// 	if err != nil {
// 		zerolog.Error().Msgf("%s is not set or is not a valid boolean", env)
// 		return err
// 	}
// 	return nil
// }
