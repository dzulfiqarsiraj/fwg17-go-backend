package customer

import (
	customer_controllers "github.com/DzulfiqarSiraj/go-backend/src/controllers/customer"
	"github.com/gin-gonic/gin"
)

func ProductRouter(r *gin.RouterGroup) {
	r.GET("", customer_controllers.ListAllProducts)
	r.GET("/:id", customer_controllers.DetailProduct)
}
