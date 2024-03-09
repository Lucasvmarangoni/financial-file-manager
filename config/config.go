package config

import (
	"fmt"
	"path/filepath"
	"runtime"

	errors "github.com/Lucasvmarangoni/logella/err"
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
	password
}

func GetEnvString(strucT, field string) string {
	key := fmt.Sprintf("%s_%s", strucT, field)
	errors.PanicBool(viper.IsSet(key), fmt.Sprintf("The key %s is not defined", key))
	return viper.GetString(key)
}

func GetEnvInt(strucT, field string) int {
	key := fmt.Sprintf("%s_%s", strucT, field)
	errors.PanicBool(viper.IsSet(key), fmt.Sprintf("The key %s is not defined", key))
	return viper.GetInt(key)
}

func GetEnvBool(strucT, field string) bool {
	key := fmt.Sprintf("%s_%s", strucT, field)
	errors.PanicBool(viper.IsSet(key), fmt.Sprintf("The key %s is not defined", key))
	return viper.GetBool(key)
}

func GetTokenAuth() *jwtauth.JWTAuth {
	tokenAuth := jwtauth.New("HS256", []byte(GetEnvString("jwt", "secret")), nil)
	return tokenAuth
}

func init() {

	cfg := &conf{}

	_, filename, _, ok := runtime.Caller(0)
	errors.PanicBool(ok, "It was not possible to obtain the path to the config.go file")

	dir := filepath.Dir(filename)
	envPath := filepath.Join(dir, "../.env")

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.SetConfigFile(envPath)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	errors.PanicErr(err, "viper.ReadInConfig")

	err = viper.Unmarshal(&cfg)
	errors.PanicErr(err, "viper.Unmarshal")
}
