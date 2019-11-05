package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/letsgo-framework/letsgo/database"
	letslog "github.com/letsgo-framework/letsgo/log"
	"github.com/letsgo-framework/letsgo/routes"
)

func main() {

	// Load env
	err := godotenv.Load()

	// Setup log writing
	letslog.InitLogFuncs()

	if err != nil {
		letslog.Error("Error loading .env file")
	} else {
		letslog.Info("env loaded")
	}

	database.Connect()

	r := routes.PaveRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	fmt.Println("Server is running on port", port)
	http.ListenAndServe(port, r)

	// if os.Getenv("SERVE_TLS") == "true" {
	// 	srv.RunTLS(port,os.Getenv("CERTIFICATE_LOCATION"),"KEY_FILE_LOCATION")
	// } else {
	// 	srv.Run(port)
	// }

}
