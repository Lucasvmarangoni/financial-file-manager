package middlewares

import (
	"fmt"
	"net/http"

	"github.com/Lucasvmarangoni/logella/err"
	"github.com/corazawaf/coraza/v3"
	"github.com/rs/zerolog/log"
)

func WAF() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			waf, err := coraza.NewWAF(coraza.NewWAFConfig().
				WithDirectivesFromFile("config/coraza/coraza.conf"))

			if err != nil {
				errors.PanicErr(err, "coraza.NewWAF")
			}

			tx := waf.NewTransaction()

			// tx.ProcessConnection("172.22.0.1", 8000, "172.22.0.1", 443)

			tx.ProcessURI(r.URL.Path, r.Method, r.Proto)

			if it := tx.ProcessRequestHeaders(); it != nil {
				switch it.Action {
				case "deny":
					statusCode := it.Status
					if statusCode < 100 || statusCode > 599 {
						statusCode = http.StatusBadRequest
					}
					w.WriteHeader(statusCode)
					errorMessage := "Error to request headers processing."
					log.Error().Msgf("ID: %d - %s", it.RuleID, errorMessage)
					http.Error(w, errorMessage, http.StatusInternalServerError)
					return
				}
			}

			if it, err := tx.ProcessRequestBody(); err != nil || it != nil {
				if err != nil {
					http.Error(w, fmt.Sprintf("Error to request body processing: %s", err), http.StatusInternalServerError)
					return
				}

				switch it.Action {
				case "deny":
					w.WriteHeader(it.Status)
					log.Error().Msgf("ID: %d", it.RuleID)
					w.Write([]byte("Some error message"))
					return
				}
			}

			tx = waf.NewTransaction()
			defer tx.ProcessLogging()

			// if it := tx.Interruption(); it != nil {
			// 	switch it.Action {
			// 	case "deny":
			// 		w.WriteHeader(it.Status)
			// 		w.Write([]byte("Some error message"))
			// 		return
			// 	}
			// }
			next.ServeHTTP(w, r)
		})
	}
}
