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
	jwt
}

func GetEnv(key string) interface{} {
	if !viper.IsSet(key) {
		panic(fmt.Sprintf("The key %s is not defined", key))
	}
	return viper.Get(key)
}

func GetTokenAuth() *jwtauth.JWTAuth {
	tokenAuth := jwtauth.New("HS256", []byte(GetEnv("jwt_secret").(string)), nil)
	return tokenAuth
}

func init() {
	cfg := &conf{}
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath("../")
	viper.SetConfigFile(".env")
	// viper.SetConfigFile("../config/.env.default")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	
}
