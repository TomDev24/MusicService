package server

import (
	"github.com/gin-gonic/gin"

	"github.com/TomDev24/MusicServiceGo/internal/database"
	"github.com/TomDev24/MusicServiceGo/internal/config"
)

type Server struct {
	router		*gin.Engine
	db			*database.Database
	cfg			*config.Config
}

func NewServer(cfg *config.Config, db *database.Database) *Server{
	r := gin.Default()

	return &Server{
		router: r,
		db:		db,
		cfg:	cfg,
	}
}

func (s *Server) registerRoutes() {
	s.router.GET("/song", s.getAll)
	s.router.GET("/song/:id", s.getById)
	s.router.POST("/song", s.create)
	s.router.DELETE("/song/:id", s.deleteById)
	s.router.PUT("/song/:id", s.editById)
}

func (s *Server) Start() error {
	s.registerRoutes()
	return s.router.Run(s.cfg.Serv.Port)
}
