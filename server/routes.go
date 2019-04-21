package server

import (
	"github.com/tsheaff/api/handlers"
)

func (self *APIServer) RegisterRoutes() {
	self.Engine.POST("/user", handlers.PostUserHandler)
}
