package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RespondWithError creates response for error
func RespondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, message interface{}) {
	response, _ := json.Marshal(message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
