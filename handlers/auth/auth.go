package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	userModel "github.com/tsheaff/api/models/user"
)

type postAuthSignupBody struct {
	Email    userModel.Email         `json:"email"`
	Password userModel.PlainPassword `json:"password"`

	// TODO: additional user fields?
}

func PostAuthSignupHandler(context *gin.Context) {
	body := &postAuthSignupBody{}
	context.BindJSON(body)

	user, err := userModel.CreateUser(body.Email, body.Password)
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}

	context.IndentedJSON(http.StatusOK, gin.H{
		"user": user,
	})
}

type postAuthLoginBody struct {
	Email    userModel.Email         `json:"email"`
	Password userModel.PlainPassword `json:"password"`
}

func PostAuthLoginHandler(context *gin.Context) {
	body := &postAuthLoginBody{}
	context.BindJSON(body)

	user, err := userModel.FetchUserByEmailMatchingPassword(body.Email, body.Password)
	if err != nil {
		context.AbortWithError(http.StatusUnauthorized, err)
	}

	context.IndentedJSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func getUserTokenFromHeader(context *gin.Context) (userModel.SecureToken, error) {
	token := context.Request.Header.Get("x-user-token")
	if len(token) == 0 {
		return "", errors.New("missing x-user-token")
	}
	return userModel.SecureToken(token), nil
}

func RequireUserMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token, err := getUserTokenFromHeader(context)
		if err != nil {
			context.AbortWithError(http.StatusUnauthorized, err)
		}
		user, err := userModel.FetchUserByToken(token)
		if err != nil {
			context.AbortWithError(http.StatusUnauthorized, err)
		}
		err = user.Validate()
		if err != nil {
			context.AbortWithError(http.StatusUnauthorized, err)
		}

		context.Set("user", user)
		context.Next()
	}
}
