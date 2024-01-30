package routers

import (
	"github.com/DzulfiqarSiraj/go-backend/src/controllers"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	r.GET("", controllers.ListAllUsers)
	r.GET("/:id", controllers.DetailUser)
}
