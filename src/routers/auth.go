package routers

import (
	controllers "github.com/DzulfiqarSiraj/go-backend/src/controllers/auth"
	"github.com/DzulfiqarSiraj/go-backend/src/middlewares"
	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.RouterGroup) {
	authMiddleware, _ := middlewares.Auth()
	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/register", controllers.Register)
	r.POST("/forgot-password", controllers.ForgotPassword)
}
