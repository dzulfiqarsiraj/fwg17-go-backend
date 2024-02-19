package customer

import (
	customer_controllers "github.com/DzulfiqarSiraj/go-backend/src/controllers/customer"
	"github.com/gin-gonic/gin"
)

func CartRouter(r *gin.RouterGroup) {
	r.GET("", customer_controllers.ListAllCarts)
	r.GET("/:id", customer_controllers.DetailCart)
	r.DELETE("/:id", customer_controllers.DeleteCart)
}
