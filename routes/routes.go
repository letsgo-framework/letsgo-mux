/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application.
| Enjoy building your API!
|
*/

package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/letsgo-framework/letsgo/controllers"
)

// PaveRoutes sets up all api routes
func PaveRoutes() *mux.Router {
	mr := mux.NewRouter()
	r := mr.PathPrefix("/api/v1").Subrouter()

	// CORS
	// r.HandleFunc("/foo", fooHandler).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions)
	r.Use(mux.CORSMethodMiddleware(r))

	// Start CRON
	// jobs.Run()
	AuthRouteHandler(r)
	r.HandleFunc("/", controllers.Greet).Methods(http.MethodGet)
	r.HandleFunc("/login", controllers.GetToken).Methods(http.MethodGet)
	// v1 := r.Group("/api/v1")
	// {
	// 	v1.GET("/", controllers.Greet)
	// 	auth := AuthRoutes(v1)
	// 	auth.GET("/", controllers.Verify)
	// }

	return r
}
