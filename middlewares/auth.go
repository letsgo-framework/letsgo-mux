package middlewares

import (
	"net/http"

	"github.com/letsgo-framework/letsgo-mux/controllers"
)

// Auth middleware
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := controllers.Srv.ValidationBearerToken(r)
		if err != nil {
			http.Error(w, http.StatusText(401), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
