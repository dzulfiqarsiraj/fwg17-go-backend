package admin

import (
	"github.com/DzulfiqarSiraj/go-backend/src/middlewares"
	"github.com/gin-gonic/gin"
)

func Combine(r *gin.RouterGroup) {
	authMiddleware, _ := middlewares.Auth()
	r.Use(authMiddleware.MiddlewareFunc())
	UserRouter(r.Group("/users"))
	ProductRouter(r.Group("/products"))
	ProductSizeRouter(r.Group("/product-size"))
	ProductVariantRouter(r.Group("/product-variant"))
	TagsRouter(r.Group("/tags"))
	ProductTagsRouter(r.Group("/product-tags"))
	ProductRatingsRouter(r.Group("/product-ratings"))
	CategoriesRouter(r.Group("/categories"))
	ProductCategoriesRouter(r.Group("/product-categories"))
	PromoRouter(r.Group("/promo"))
	OrdersRouter(r.Group("/orders"))
	OrderDetailsRouter(r.Group("/order-details"))
}
