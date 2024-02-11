package admin

import (
	"github.com/DzulfiqarSiraj/go-backend/src/middlewares"
	"github.com/gin-gonic/gin"
)

func Combine(r *gin.RouterGroup) {
	authMiddleware, _ := middlewares.Auth()
	r.Use(authMiddleware.MiddlewareFunc())
	UserRouter(r.Group("/users"))
	ProductRouter(r.Group("/products"))
	ProductSizeRouter(r.Group("/product-size"))
	ProductVariantRouter(r.Group("/product-variant"))
}
