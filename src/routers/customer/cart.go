package customer

import (
	customer_controllers "github.com/DzulfiqarSiraj/go-backend/src/controllers/customer"
	"github.com/gin-gonic/gin"
)

func CartRouter(r *gin.RouterGroup) {
	r.GET("", customer_controllers.ListAllCarts)
	r.GET("/:id", customer_controllers.DetailCart)
	r.POST("/", customer_controllers.AddToCart)
	r.PATCH("/:id", customer_controllers.UpdateCart)
	r.DELETE("/:id", customer_controllers.DeleteCart)
}
