package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ilhamhaikal/film_go/internal/handlers"
	"github.com/ilhamhaikal/film_go/internal/middleware"
	"github.com/ilhamhaikal/film_go/internal/repository"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
	db     *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{
		router: gin.Default(),
		db:     db,
	}
}

func (s *Server) SetupRoutes() {
	// Add health check endpoint
	s.router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"message": "API is running",
		})
	})

	cinemaRepo := repository.NewCinemaRepository(s.db)
	cinemaHandler := handlers.NewCinemaHandler(cinemaRepo)

	auth := s.router.Group("/auth")
	{
		auth.POST("/register", cinemaHandler.Register)
		auth.POST("/login", cinemaHandler.Login)
	}

	api := s.router.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/bioskop", cinemaHandler.CreateCinema)
		api.GET("/bioskop", cinemaHandler.GetAllCinemas)
		api.GET("/bioskop/:id", cinemaHandler.GetCinema)
		api.PUT("/bioskop/:id", cinemaHandler.UpdateCinema)
		api.DELETE("/bioskop/:id", cinemaHandler.DeleteCinema)
	}
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
