package admin

import (
	admin_controllers "github.com/DzulfiqarSiraj/go-backend/src/controllers/admin"
	"github.com/gin-gonic/gin"
)

func CategoriesRouter(r *gin.RouterGroup) {
	r.GET("", admin_controllers.ListAllCategories)
	r.GET("/:id", admin_controllers.DetailCategory)
	r.POST("", admin_controllers.CreateCategory)
	r.PATCH("/:id", admin_controllers.UpdateCategory)
	r.DELETE("/:id", admin_controllers.DeleteCategory)
}
