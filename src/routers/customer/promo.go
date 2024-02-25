package customer

import (
	customer_controllers "github.com/DzulfiqarSiraj/go-backend/src/controllers/customer"
	"github.com/gin-gonic/gin"
)

func PromoRouter(r *gin.RouterGroup) {
	r.GET("", customer_controllers.ListAllPromo)
	r.GET("/:id", customer_controllers.DetailPromo)
	r.POST("", customer_controllers.CreatePromo)
	r.PATCH("/:id", customer_controllers.UpdatePromo)
	r.DELETE("/:id", customer_controllers.DeletePromo)
}
