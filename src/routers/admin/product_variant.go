package admin

import (
	admin_controllers "github.com/DzulfiqarSiraj/go-backend/src/controllers/admin"
	"github.com/gin-gonic/gin"
)

func ProductVariantRouter(r *gin.RouterGroup) {
	r.GET("", admin_controllers.ListAllProductVariant)
	r.GET("/:id", admin_controllers.DetailProductVariant)
	r.POST("", admin_controllers.CreateProductVariant)
	r.PATCH("/:id", admin_controllers.UpdateProductVariant)
	r.DELETE("/:id", admin_controllers.DeleteProductVariant)
}
