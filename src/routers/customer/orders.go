package customer

import (
	customer_controllers "github.com/DzulfiqarSiraj/go-backend/src/controllers/customer"
	"github.com/gin-gonic/gin"
)

func OrdersRouter(r *gin.RouterGroup) {
	r.GET("", customer_controllers.ListAllOrders)
	r.GET("/:id", customer_controllers.DetailOrder)
	r.POST("", customer_controllers.CreateOrder)
	r.PATCH("/:id", customer_controllers.UpdateOrder)
	r.DELETE("/:id", customer_controllers.DeleteOrder)
}
