package admin

import (
	admin_controllers "github.com/DzulfiqarSiraj/go-backend/src/controllers/admin"
	"github.com/gin-gonic/gin"
)

func ProductCategoriesRouter(r *gin.RouterGroup) {
	r.GET("", admin_controllers.ListAllProductCategories)
	r.GET("/:id", admin_controllers.DetailProductCategory)
	r.POST("", admin_controllers.CreateProductCategory)
	r.PATCH("/:id", admin_controllers.UpdateProductCategory)
	r.DELETE("/:id", admin_controllers.DeleteProductCategory)
}
