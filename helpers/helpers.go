package helpers

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON sends response in json format
func RespondWithJSON(w http.ResponseWriter, code int, message interface{}) {
	response, _ := json.Marshal(message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
