package routers

import (
	controllers "github.com/DzulfiqarSiraj/go-backend/src/controllers/auth"
	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.RouterGroup) {
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)
	r.POST("/forgot-password", controllers.ForgotPassword)
}
