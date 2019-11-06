package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/letsgo-framework/letsgo-mux/controllers"
)

func AuthRouteHandler(r *mux.Router) *mux.Router {

	// Auth Init
	controllers.AuthInit()

	r.HandleFunc("/credentials", controllers.GetCredentials).Methods(http.MethodGet)
	r.HandleFunc("/login", controllers.GetToken).Methods(http.MethodGet)
	r.HandleFunc("/register", controllers.Register).Methods(http.MethodPost)

	auth := r.PathPrefix("/auth").Subrouter()

	return auth
	// r.Use(srv.HandleTokenVerify())

	// auth := r.Group("auth")
	// {
	// 	auth.Use(ginserver.HandleTokenVerify(config))
	// }
}
