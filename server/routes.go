package server

import (
	"github.com/tsheaff/api/handlers"
)

func (self *APIServer) RegisterRoutes() {
	self.Engine.POST("/auth/signup", handlers.PostAuthSignupHandler)
	self.Engine.POST("/auth/login", handlers.PostAuthLoginHandler)
}
