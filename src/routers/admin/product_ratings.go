package admin

import (
	admin_controllers "github.com/DzulfiqarSiraj/go-backend/src/controllers/admin"
	"github.com/gin-gonic/gin"
)

func ProductRatingsRouter(r *gin.RouterGroup) {
	r.GET("", admin_controllers.ListAllProductRatings)
	r.GET("/:id", admin_controllers.DetailProductRating)
	r.POST("", admin_controllers.CreateProductRating)
	r.PATCH("/:id", admin_controllers.UpdateProductRating)
	r.DELETE("/:id", admin_controllers.DeleteProductRating)
}
