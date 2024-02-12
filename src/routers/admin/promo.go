package admin

import (
	admin_controllers "github.com/DzulfiqarSiraj/go-backend/src/controllers/admin"
	"github.com/gin-gonic/gin"
)

func PromoRouter(r *gin.RouterGroup) {
	r.GET("", admin_controllers.ListAllPromo)
	r.GET("/:id", admin_controllers.DetailPromo)
	r.POST("", admin_controllers.CreatePromo)
	r.PATCH("/:id", admin_controllers.UpdatePromo)
	r.DELETE("/:id", admin_controllers.DeletePromo)
}
