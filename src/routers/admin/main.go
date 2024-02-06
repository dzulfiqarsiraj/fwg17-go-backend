package admin

import "github.com/gin-gonic/gin"

func Combine(r *gin.RouterGroup) {
	UserRouter(r.Group("/users"))
	ProductRouter(r.Group("/products"))
	ProductSizeRouter(r.Group("/product-size"))
}
