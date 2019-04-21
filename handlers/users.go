package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostUserHandler(context *gin.Context) {
	context.Status(http.StatusOK)
}
