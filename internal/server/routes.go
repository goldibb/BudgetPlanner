package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173", "http://localhost:5137"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-HX-Request",
			"X-HX-Current-URL",
			"X-HX-Target",
		},
		AllowCredentials: true,
	}))
	// Books handlers

	r.GET("/health", s.healthHandler)
	r.GET("/", s.ShowSomething)
	return r
}
func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
func (s *Server) ShowSomething(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}
