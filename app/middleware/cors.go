package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

func Cors(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler
	return handleCORS(handler)
}

func Middleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("middleware", r.URL)
        h.ServeHTTP(w, r)
    })
}
