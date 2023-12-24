package config

import (
	"fmt"

	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)



type conf struct {
	database
	rabbitMQ
	grpc
	storage
	tokenAuth *jwtauth.JWTAuth
	jWTSecret string
}

func GetEnv(key string) interface{} {
	if !viper.IsSet(key) {
		panic(fmt.Sprintf("The key %s is not defined", key))
	}
	return viper.Get(key)
}

func init() {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	cfg.tokenAuth = jwtauth.New("HS256", []byte(cfg.jWTSecret), nil)
}
