package routers

import (
	controllers "github.com/DzulfiqarSiraj/go-backend/src/controllers"
	"github.com/gin-gonic/gin"
)

func ProductRouter(r *gin.RouterGroup) {
	r.GET("", controllers.ListAllProducts)
}
