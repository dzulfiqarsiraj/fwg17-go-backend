package admin

import (
	admin_controllers "github.com/DzulfiqarSiraj/go-backend/src/controllers/admin"
	"github.com/gin-gonic/gin"
)

func TagsRouter(r *gin.RouterGroup) {
	r.GET("", admin_controllers.ListAllTags)
	r.GET("/:id", admin_controllers.DetailTag)
	r.POST("", admin_controllers.CreateTag)
	r.PATCH("/:id", admin_controllers.UpdateTag)
	r.DELETE("/:id", admin_controllers.DeleteTag)
}
