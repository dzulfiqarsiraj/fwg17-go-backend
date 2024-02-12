package admin

import (
	admin_controllers "github.com/DzulfiqarSiraj/go-backend/src/controllers/admin"
	"github.com/gin-gonic/gin"
)

func ProductTagsRouter(r *gin.RouterGroup) {
	r.GET("", admin_controllers.ListAllProductTags)
	r.GET("/:id", admin_controllers.DetailProductTag)
	r.POST("", admin_controllers.CreateProductTag)
	r.PATCH("/:id", admin_controllers.UpdateProductTag)
	r.DELETE("/:id", admin_controllers.DeleteProductTag)
}
