package customer

import (
	customer_controllers "github.com/DzulfiqarSiraj/go-backend/src/controllers/customer"
	"github.com/gin-gonic/gin"
)

func ProfileRouter(r *gin.RouterGroup) {
	r.GET("", customer_controllers.UserProfile)
	r.PATCH("", customer_controllers.UpdateProfile)
}
