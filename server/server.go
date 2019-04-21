package server

import (
	"time"

	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	Engine *gin.Engine
}

func New() *APIServer {
	server := &APIServer{
		Engine: gin.New(),
	}
	server.Engine.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339Nano, true))
	server.Engine.Use(gin.Recovery())
	return server
}

func (self *APIServer) Start() *APIServer {
	self.Engine.Run()
	return self
}
