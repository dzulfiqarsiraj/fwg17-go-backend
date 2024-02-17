package routers

import (
	"github.com/DzulfiqarSiraj/go-backend/src/routers/admin"
	"github.com/DzulfiqarSiraj/go-backend/src/routers/customer"
	"github.com/gin-gonic/gin"
)

func Combine(r *gin.Engine) {
	ProductRouter(r.Group("/products"))
	AuthRouter(r.Group("/auth"))
	admin.Combine(r.Group("/admin"))
	customer.Combine(r.Group("/customer"))
}
