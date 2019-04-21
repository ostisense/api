package server

import (
	"github.com/tsheaff/api/handlers/auth"
	"github.com/tsheaff/api/handlers/session"
)

func (self *APIServer) RegisterRoutes() {
	self.Engine.POST("/auth/signup", auth.PostAuthSignupHandler)
	self.Engine.POST("/auth/login", auth.PostAuthLoginHandler)

	requireUserGroup := self.Engine.Group("/")
	requireUserGroup.Use(auth.RequireUserMiddleware())
	{
		requireUserGroup.POST("/session", session.PostSessionHandler)
	}
}
