package gin_utils

import (
	"github.com/gin-gonic/gin"
)

// abortAndRespondError is the preferred way to respond to the caller with failure
func AbortAndRespondError(context *gin.Context, httpCode int, err error) {
	context.AbortWithError(httpCode, err)
	context.IndentedJSON(httpCode, gin.H{
		"error": err.Error(),
	})
}
