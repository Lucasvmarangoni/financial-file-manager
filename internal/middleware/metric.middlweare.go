package middlewares

import (
	"net/http"
	"strconv"

	"github.com/Lucasvmarangoni/financial-file-manager/pkg/metric"

)

func Metrics(mService metric.UseCase) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            appMetric := metric.NewHTTP(r.URL.Path, r.Method)
            appMetric.Started()

            rw := &responseWriterWithStatus{ResponseWriter: w}

            next.ServeHTTP(rw, r)

            appMetric.Finished()
            status := rw.StatusCode() 
            appMetric.StatusCode = strconv.Itoa(status)
            mService.SaveHTTP(appMetric)
        })
    }
}


type responseWriterWithStatus struct {
    http.ResponseWriter
    status int
}

func (rw *responseWriterWithStatus) WriteHeader(status int) {
    rw.status = status
    rw.ResponseWriter.WriteHeader(status)
}

func (rw *responseWriterWithStatus) Write(b []byte) (int, error) {
    return rw.ResponseWriter.Write(b)
}

func (rw *responseWriterWithStatus) StatusCode() int {
    return rw.status
}