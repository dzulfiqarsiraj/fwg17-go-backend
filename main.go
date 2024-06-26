package main

import (
	"net/http"

	"github.com/DzulfiqarSiraj/go-backend/src/routers"
	"github.com/DzulfiqarSiraj/go-backend/src/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:5174", "https://cov-shop.netlify.app", "https://cov-go.netlify.app", "http://143.110.156.215:5174"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "Accept", "access-control-allow-origin", "access-control-allow-headers"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	routers.Combine(r)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, &services.ResponseOnly{
			Success: false,
			Message: "Resource Not Found...",
		})
	})
	r.Run("0.0.0.0:8181")
}
