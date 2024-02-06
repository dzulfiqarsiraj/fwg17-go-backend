package admin

import (
	admin_controllers "github.com/DzulfiqarSiraj/go-backend/src/controllers/admin"
	"github.com/gin-gonic/gin"
)

func ProductSizeRouter(r *gin.RouterGroup) {
	r.GET("", admin_controllers.ListAllProductSize)
	r.GET("/:id", admin_controllers.DetailProductize)
	// r.POST("", admin_controllers.CreateProduct)
	// r.PATCH("/:id", admin_controllers.UpdateProduct)
	// r.DELETE("/:id", admin_controllers.DeleteProduct)
}
