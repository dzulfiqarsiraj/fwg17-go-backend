package routers

import (
	"github.com/DzulfiqarSiraj/go-backend/src/routers/admin"
	"github.com/gin-gonic/gin"
)

func Combine(r *gin.Engine) {
	AuthRouter(r.Group("/auth"))
	admin.Combine(r.Group("/admin"))
}
