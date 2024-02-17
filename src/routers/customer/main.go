package customer

import (
	"github.com/DzulfiqarSiraj/go-backend/src/middlewares"
	"github.com/gin-gonic/gin"
)

func Combine(r *gin.RouterGroup) {
	authMiddleware, _ := middlewares.Auth()
	r.Use(authMiddleware.MiddlewareFunc())
	ProfileRouter(r.Group("/profile"))
	ProductRouter(r.Group("/products"))
}
