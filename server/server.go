package server

import (
	"log"

	"github.com/Carmo-sousa/api/server/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	log.Fatal(router.Run(":" + s.port))
}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}
