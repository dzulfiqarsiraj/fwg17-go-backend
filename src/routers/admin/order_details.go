package admin

import (
	admin_controllers "github.com/DzulfiqarSiraj/go-backend/src/controllers/admin"
	"github.com/gin-gonic/gin"
)

func OrderDetailsRouter(r *gin.RouterGroup) {
	r.GET("", admin_controllers.ListAllOrderDetails)
	r.GET("/:id", admin_controllers.DetailOrderDetail)
	r.POST("", admin_controllers.CreateOrderDetail)
	r.PATCH("/:id", admin_controllers.UpdateOrderDetail)
	r.DELETE("/:id", admin_controllers.DeleteOrderDetail)
}
