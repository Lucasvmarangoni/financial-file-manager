package config

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type conf struct {
	database
	rabbitMQ
	grpc
	storage
	jwt
	authz
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
	
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("It was not possible to obtain the path to the config.go file")
	}
	dir := filepath.Dir(filename)	
	envPath := filepath.Join(dir, "../.env") 

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.SetConfigFile(envPath)
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
