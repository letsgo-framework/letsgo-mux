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
	"github.com/letsgo-framework/letsgo-mux/controllers"
	"github.com/letsgo-framework/letsgo-mux/middlewares"
)

// PaveRoutes sets up all api routes
func PaveRoutes() *mux.Router {
	mr := mux.NewRouter()
	r := mr.PathPrefix("/api/v1").Subrouter()

	// CORS
	r.Use(mux.CORSMethodMiddleware(r))

	// Start CRON
	// jobs.Run()

	// greeter
	r.HandleFunc("", controllers.Greet).Methods(http.MethodGet)

	// Register Auth routes
	auth := AuthRouteHandler(r)
	auth.Use(middlewares.Auth)
	auth.HandleFunc("", controllers.Verify).Methods(http.MethodGet)

	return r
}
