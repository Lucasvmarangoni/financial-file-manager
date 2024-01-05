package routers

import (
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

func (u *UserRouter) InitializeRoute(path string, handler http.HandlerFunc)  {

	switch strings.ToUpper(method) {
	case "POST":
		router.Post(path, handler)
	case "GET":
		router.Get(path, handler)
	case "PUT":
		router.Put(path, handler)
	case "PATH":
		router.Patch(path, handler)
	case "DELETE":
		router.Delete(path, handler)
	}
	log.Info().Str("context", "UserRouter").Msgf("Mapped - Initialized: (%s) /user%s ", method, path)
	
}

func (u *UserRouter) Method(m string) *UserRouter {
	method = m
	return u
}
