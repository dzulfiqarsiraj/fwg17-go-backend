package customer

import (
	customer_controllers "github.com/DzulfiqarSiraj/go-backend/src/controllers/customer"
	"github.com/gin-gonic/gin"
)

func OrderDetailsRouter(r *gin.RouterGroup) {
	r.GET("", customer_controllers.ListAllOrderDetails)
	r.GET("/:id", customer_controllers.DetailOrderDetail)
	r.POST("", customer_controllers.CreateOrderDetail)
	r.PATCH("/:id", customer_controllers.UpdateOrderDetail)
	r.DELETE("/:id", customer_controllers.DeleteOrderDetail)
}
