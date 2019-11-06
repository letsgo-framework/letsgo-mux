/*
 |--------------------------------------------------------------------------
 | Authentication Controller
 |--------------------------------------------------------------------------
 |
 | GetCredentials works on oauth2 Client Credentials Grant and returns CLIENT_ID, CLIENT_SECRET
 | GetToken takes CLIENT_ID, CLIENT_SECRET, grant_type, scope and returns access_token and some other information
*/

package controllers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/letsgo-framework/letsgo-mux/database"
	"github.com/letsgo-framework/letsgo-mux/helpers"
	letslog "github.com/letsgo-framework/letsgo-mux/log"
	"github.com/letsgo-framework/letsgo-mux/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

var clientStore = store.NewClientStore()
var manager = manage.NewDefaultManager()

var Srv = server.NewDefaultServer(manager)

// AuthInit initializes authentication
func AuthInit() {
	cfg := &manage.Config{
		// access token expiration time
		AccessTokenExp: time.Hour * 2,
		// refresh token expiration time
		RefreshTokenExp: time.Hour * 24 * 7,
		// whether to generate the refreshing token
		IsGenerateRefresh: true,
	}

	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	manager.SetPasswordTokenCfg(cfg)

	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	manager.MapClientStorage(clientStore)

	Srv.SetAllowGetAccessRequest(true)
	Srv.SetClientInfoHandler(server.ClientFormHandler)

	Srv.SetPasswordAuthorizationHandler(login)

	err := clientStore.Set("client@letsgo", &models.Client{
		ID:     "client@letsgo",
		Secret: "Va4a8bFFhTJZdybnzyhjHjj6P9UVh7UL",
		Domain: "http://localhost:8080",
	})

	if err != nil {
		letslog.Error(err.Error())
	}
}

// GetCredentials sends client credentials
func GetCredentials(w http.ResponseWriter, r *http.Request) {
	clientId := uuid.New().String()
	clientSecret := uuid.New().String()
	err := clientStore.Set(clientId, &models.Client{
		ID:     clientId,
		Secret: clientSecret,
		Domain: "http://localhost:8000",
	})
	if err != nil {
		letslog.Error(err.Error())
	}
	helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"CLIENT_ID": clientId, "CLIENT_SECRET": clientSecret})
}

// GetToken sends access_token
func GetToken(w http.ResponseWriter, r *http.Request) {
	Srv.HandleTokenRequest(w, r)
}

// Verify accessToken with client
func Verify(w http.ResponseWriter, r *http.Request) {
	ti, exists := Srv.ValidationBearerToken(r)
	if exists == nil {
		helpers.RespondWithJSON(w, http.StatusOK, ti)
		return
	}
	helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Authentication unsuccessful"})
}

// register
func Register(w http.ResponseWriter, r *http.Request) {
	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Transform into RequestBody struct
	a := types.User{}
	err = json.Unmarshal(raw, a)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	collection := database.UserCollection()

	a.Password, _ = generateHash(a.Password)
	a.Id = primitive.NewObjectID()

	if err != nil {
		letslog.Error(err.Error())
		return
	}
	_, err = collection.InsertOne(ctx, a)
	if err != nil {
		letslog.Error(err.Error())
		return
	}
	helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Registration Successful"})
}

// Generate a salted hash for the input string
func generateHash(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

// Compare string to generated hash
func compare(hash string, s string) error {
	incoming := []byte(s)
	existing := []byte(hash)

	return bcrypt.CompareHashAndPassword(existing, incoming)
}

func login(username, password string) (userID string, err error) {

	collection := database.UserCollection()

	user := types.User{}
	err = collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)

	if err != nil {
		letslog.Error(err.Error())
		return userID, err
	}
	loginError := compare(user.Password, password)

	if loginError != nil {
		letslog.Error(loginError.Error())
		return userID, err
	} else {
		userID = user.Id.Hex()
		return userID, err
	}
}
