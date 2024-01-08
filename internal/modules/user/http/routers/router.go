package routers

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"
)

func (u *UserRouter) InitializeRoute(r chi.Router, path string, handler http.HandlerFunc) {

	switch strings.ToUpper(u.method) {
	case "POST":
		r.Post(path, handler)
	case "GET":
		r.Get(path, handler)
	case "PUT":
		r.Put(path, handler)
	case "PATH":
		r.Patch(path, handler)
	case "DELETE":
		r.Delete(path, handler)
	}
	log.Info().Str("context", "UserRouter").Msgf("Mapped - Initialized: (%s) /user%s ", u.method, path)
}

func (u *UserRouter) Method(m string) *UserRouter {
	u.method = m
	return u
}
