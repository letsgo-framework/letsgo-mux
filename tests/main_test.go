package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
	letslog "github.com/letsgo-framework/letsgo-mux/log"
	"github.com/letsgo-framework/letsgo-mux/routes"
	"github.com/letsgo-framework/letsgo-mux/types"

	"github.com/letsgo-framework/letsgo-mux/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	. "gopkg.in/check.v1"
)

type TestInsert struct {
	Name string `form:"name" binding:"required" json:"name" bson:"name"`
}

func TestMain(m *testing.M) {
	// Setup log writing
	letslog.InitLogFuncs()
	database.TestConnect()
	err := godotenv.Load("../.env.testing")

	database.DB.Drop(context.Background())

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := routes.PaveRoutes()

	port := os.Getenv("PORT")

	fmt.Println("Server is running on port", port)
	go http.ListenAndServe(port, r)

	os.Exit(m.Run())
}

func TestGetEnv(t *testing.T) {
	dbPort := os.Getenv("DATABASE_PORT")
	fmt.Printf("db port %s", dbPort)

	if dbPort != "27017" {
		t.Fatal("Port Incorrect")
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

func TestCredentials(t *testing.T) {
	requestURL := "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/credentials"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", requestURL, nil)

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
}

func TestHelloWorld(t *testing.T) {
	requestURL := "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", requestURL, nil)

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
}

func TestTokenSuccess(t *testing.T) {
	requestURL := "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/credentials"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", requestURL, nil)

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	responseData, _ := ioutil.ReadAll(resp.Body)
	var credResponse types.CredentialResponse
	json.Unmarshal(responseData, &credResponse)

	requestURL = "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/login?grant_type=client_credentials&client_id=" + credResponse.CLIENT_ID + "&client_secret=" + credResponse.CLIENT_SECRET + "&scope=read"
	// fmt.Println(credResponse.CLIENT_ID)
	// fmt.Println(credResponse.CLIENT_SECRET)

	req, _ = http.NewRequest("GET", requestURL, nil)

	resp, err = client.Do(req)
	if err != nil {
		t.Fatal(err)

	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		t.Fatal("TEST FAILED :: Token Success")
	}
}

func TestTokenFail(t *testing.T) {
	requestURL := "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/login"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", requestURL, nil)

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)

	}
	defer resp.Body.Close()
	if resp.StatusCode != 401 {
		fmt.Println(resp.StatusCode)

		t.Fatal("TEST FAILED :: Token Fail")
	}
}

func TestAccessTokenSuccess(t *testing.T) {
	requestURL := "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/credentials"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", requestURL, nil)

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	responseData, _ := ioutil.ReadAll(resp.Body)
	var credResponse types.CredentialResponse
	json.Unmarshal(responseData, &credResponse)

	requestURL = "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/login?grant_type=client_credentials&client_id=" + credResponse.CLIENT_ID + "&client_secret=" + credResponse.CLIENT_SECRET + "&scope=read"

	req, _ = http.NewRequest("GET", requestURL, nil)

	resp, err = client.Do(req)
	if err != nil {
		t.Fatal(err)

	}
	defer resp.Body.Close()

	respData, _ := ioutil.ReadAll(resp.Body)
	var tokenResponse types.TokenResponse
	json.Unmarshal(respData, &tokenResponse)

	requestURL = "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/auth?access_token=" + tokenResponse.AccessToken

	req, _ = http.NewRequest("GET", requestURL, nil)

	resp, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		t.Fatal("TEST FAILED :: Access Token Success")
	}
}

func TestAccessTokenFail(t *testing.T) {
	requestURL := "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/credentials/"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", requestURL, nil)

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	responseData, _ := ioutil.ReadAll(resp.Body)
	var credResponse types.CredentialResponse
	json.Unmarshal(responseData, &credResponse)

	requestURL = "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/login?grant_type=client_credentials&client_id=" + credResponse.CLIENT_ID + "&client_secret=" + credResponse.CLIENT_SECRET + "&scope=read"

	req, _ = http.NewRequest("GET", requestURL, nil)

	resp, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	respData, _ := ioutil.ReadAll(resp.Body)
	var tokenResponse types.TokenResponse
	json.Unmarshal(respData, &tokenResponse)

	requestURL = "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/auth?access_token=mywrongaccesstoken"

	req, _ = http.NewRequest("GET", requestURL, nil)

	resp, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 401 {
		fmt.Println(resp.StatusCode)
		t.Fatal("TEST FAILED :: Access Token Fail !")
	}
}

func TestDBInsert(t *testing.T) {
	database.TestConnect()
	input := TestInsert{Name: "testname"}
	collection := database.DB.Collection("test_collection")
	_, err := collection.InsertOne(context.Background(), input)
	if err != nil {
		t.Fatal(err)
	}
	result := TestInsert{}
	err = collection.FindOne(context.Background(), bson.M{"name": "testname"}).Decode(&result)
	if err != nil {
		t.Fatal(err)
	}
	if result != input {
		t.Fatal("Insertion Unsuccesful!")
	}
}

//Test User-registartion
func Test1UserRegistration(t *testing.T) {
	data := types.User{
		Name:     "Letsgo User",
		Username: "letsgoUs3r",
		Password: "qwerty",
	}

	requestURL := "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/register"
	client := &http.Client{}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(data)
	req, _ := http.NewRequest("POST", requestURL, b)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	respData, _ := ioutil.ReadAll(resp.Body)
	var user types.User
	json.Unmarshal(respData, &user)
	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		t.Fatal("Registration Unsuccesful!")
	}
}

func Test2UserLoginPasswordGrant(t *testing.T) {
	requestURL := "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/login?grant_type=password&client_id=client@letsgo&client_secret=Va4a8bFFhTJZdybnzyhjHjj6P9UVh7UL&scope=read&username=letsgoUs3r&password=qwerty"
	req, _ := http.NewRequest("GET", requestURL, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		letslog.Debug(err.Error())
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatal("Login Unsuccesful!")
	}
}

func Test(t *testing.T) {
	TestingT(t)
}
