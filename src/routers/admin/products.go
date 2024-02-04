package admin

import (
	admin_controllers "github.com/DzulfiqarSiraj/go-backend/src/controllers/admin"
	"github.com/gin-gonic/gin"
)

func ProductRouter(r *gin.RouterGroup) {
	r.GET("", admin_controllers.ListAllProducts)
	r.GET("/:id", admin_controllers.DetailProduct)
	r.POST("", admin_controllers.CreateProduct)
	r.PATCH("/:id", admin_controllers.UpdateProduct)
	r.DELETE("/:id", admin_controllers.DeleteProduct)
}
