package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostSessionHandler(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{
		"session": true,
	})
}
