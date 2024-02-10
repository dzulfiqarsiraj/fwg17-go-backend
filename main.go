package main

import (
	"net/http"

	"github.com/DzulfiqarSiraj/go-backend/src/routers"
	"github.com/DzulfiqarSiraj/go-backend/src/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.Combine(r)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, &services.ResponseOnly{
			Success: false,
			Message: "Resource Not Found...",
		})
	})
	r.Run(":8888")
}
