package config

import "github.com/go-chi/jwtauth"

type database struct {
	name     string `mapstructure:"DATABASE_NAME"`
	user     string `mapstructure:"DATABASE_USER"`
	password string `mapstructure:"DATABASE_PASSWORD"`
	ssl_mode string `mapstructure:"DATABASE_SSL_MODE"`
	port     string `mapstructure:"DATABASE_PORT"`
}

type jwt struct {
	secret     string `mapstructure:"JWT_SECRET"`
	expired_in int    `mapstructure:"JWT_EXPIREDIN"`
	tokenAuth *jwtauth.JWTAuth	
}

type rabbitMQ struct {
	user                     string `mapstructure:"RABBITMQ_DEFAULT_USER"`
	password                 string `mapstructure:"RABBITMQ_DEFAULT_PASS"`
	host                     string `mapstructure:"RABBITMQ_DEFAULT_HOST"`
	port                     string `mapstructure:"RABBITMQ_DEFAULT_PORT"`
	vhost                    string `mapstructure:"RABBITMQ_DEFAULT_VHOST"`
	consumer_name            string `mapstructure:"RABBITMQ_CONSUMER_NAME"`
	consumer_queue_name      string `mapstructure:"RABBITMQ_CONSUMER_QUEUE_NAME"`
	notification_ex          string `mapstructure:"RABBITMQ_NOTIFICATION_EX"`
	notification_routing_key string `mapstructure:"RABBITMQ_NOTIFICATION_ROUTING_KEY"`
	dlx                      string `mapstructure:"RABBITMQ_DLX"`
}

type grpc struct {
	port int `mapstructure:"GRPC_PORT"`
}

type storage struct {
	local_path string `mapstructure:"LOCAL_STORAGE_PATH"`
	local_file string `mapstructure:"LOCAL_FILE"`
}
