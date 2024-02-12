package admin

import (
	admin_controllers "github.com/DzulfiqarSiraj/go-backend/src/controllers/admin"
	"github.com/gin-gonic/gin"
)

func OrdersRouter(r *gin.RouterGroup) {
	r.GET("", admin_controllers.ListAllOrders)
	r.GET("/:id", admin_controllers.DetailOrder)
	r.POST("", admin_controllers.CreateOrder)
	r.PATCH("/:id", admin_controllers.UpdateOrder)
	r.DELETE("/:id", admin_controllers.DeleteOrder)
}
