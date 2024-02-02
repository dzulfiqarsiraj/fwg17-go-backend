package routers

import (
	"github.com/DzulfiqarSiraj/go-backend/src/routers/admin"
	"github.com/gin-gonic/gin"
)

func Combine(r *gin.Engine) {
	admin.Combine(r.Group("/admin"))
}
