package server

import "github.com/gin-gonic/gin"

type Server struct {
	handler *Handler
	engine  *gin.Engine
}

func NewServer(handler *Handler, engine *gin.Engine) *Server {
	return &Server{handler: handler, engine: engine}
}

func (s *Server) MapRoutes() {
	g := s.engine.Group("/v1")

	g.PUT("/shark", s.handler.ConfigureShark())
	g.PUT("/prey", s.handler.ConfigurePrey())
	g.POST("/simulate", s.handler.SimulateHunt())
}

func (s *Server) Run() error {
	return s.engine.Run(":8080")
}
