package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/letsgo-framework/letsgo/controllers"
)

func AuthRouteHandler(r *mux.Router) {

	// Auth Init
	controllers.AuthInit()

	// config := ginserver.Config{
	// 	ErrorHandleFunc: func(ctx *gin.Context, err error) {
	// 		helpers.RespondWithError(ctx, 401, "invalid access_token")
	// 	},
	// 	TokenKey: "github.com/go-oauth2/gin-server/access-token",
	// 	Skipper: func(_ *gin.Context) bool {
	// 		return false
	// 	},
	// }
	// var manager = manage.NewDefaultManager()
	// var srv = server.NewDefaultServer(manager)

	r.HandleFunc("/credentials", controllers.GetCredentials).Methods(http.MethodGet)
	r.HandleFunc("/login", controllers.GetToken).Methods(http.MethodGet)
	r.HandleFunc("/register", controllers.Register).Methods(http.MethodPost)

	// r.Use(srv.HandleTokenVerify())

	// auth := r.Group("auth")
	// {
	// 	auth.Use(ginserver.HandleTokenVerify(config))
	// }
}
