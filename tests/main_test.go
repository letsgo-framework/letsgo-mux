package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/letsgo-framework/letsgo/controllers"
	"github.com/letsgo-framework/letsgo/database"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// func TestMain(m *testing.M) {
// 	// Setup log writing
// 	letslog.InitLogFuncs()
// 	err := godotenv.Load("../.env.testing")
// 	database.TestConnect()

// 	database.DB.Drop(context.Background())

// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	s := TestSuite{
// 		srv: routes.PaveRoutes(),
// 	}
// 	go s.srv.Run(os.Getenv("PORT"))

// 	os.Exit(m.Run())
// }

func TestHelloWorld(t *testing.T) {
	requestURL := "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", requestURL, nil)

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.Greet)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDatabaseTestConnection(t *testing.T) {
	database.TestConnect()
	err := database.Client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		t.Fatal(err)
	}
}
func TestDatabaseConnection(t *testing.T) {
	database.Connect()
	err := database.Client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		t.Fatal(err)
	}
}
