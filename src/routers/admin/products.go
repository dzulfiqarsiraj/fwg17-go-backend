package admin

import (
	controllers "github.com/DzulfiqarSiraj/go-backend/src/controllers/admin"
	"github.com/gin-gonic/gin"
)

func ProductRouter(r *gin.RouterGroup) {
	r.GET("", controllers.ListAllProducts)
	r.GET("/:id", controllers.DetailProduct)
	r.POST("", controllers.CreateProduct)
}
