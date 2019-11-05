package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/letsgo-framework/letsgo/controllers"
)

func AuthRouteHandler(r *mux.Router) {

	// Auth Init
	controllers.AuthInit()

	// var manager = manage.NewDefaultManager()
	// var srv = server.NewDefaultServer(manager)

	r.HandleFunc("/credentials", controllers.GetCredentials).Methods(http.MethodGet)
	r.HandleFunc("/login", controllers.GetToken).Methods(http.MethodGet)
	r.HandleFunc("/register", controllers.Register).Methods(http.MethodPost)

	auth := r.PathPrefix("/auth").Subrouter()

	auth.HandleFunc("", controllers.Verify).Methods(http.MethodGet)
	// r.Use(srv.HandleTokenVerify())

	// auth := r.Group("auth")
	// {
	// 	auth.Use(ginserver.HandleTokenVerify(config))
	// }
}
