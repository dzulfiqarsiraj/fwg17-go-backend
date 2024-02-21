package customer_controllers

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/DzulfiqarSiraj/go-backend/src/models"
	"github.com/DzulfiqarSiraj/go-backend/src/services"
	"github.com/gin-gonic/gin"
)

func ListAllCarts(c *gin.Context) {
	data := c.MustGet("id").(*models.User)
	userId := data.Id
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit
	result, err := models.FindAllCarts(userId, limit, offset)

	pageInfo := &services.PageInfo{
		Page:      page,
		Limit:     limit,
		TotalPage: int(math.Ceil(float64(result.Count) / float64(limit))),
		TotalData: result.Count,
	}

	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success:  true,
		Message:  "List All Carts",
		PageInfo: *pageInfo,
		Results:  result.Data,
	})
}

func DetailCart(c *gin.Context) {
	// Get user id via token payload
	user := c.MustGet("id").(*models.User)
	userId := user.Id
	id, _ := strconv.Atoi(c.Param("id"))

	order, err := models.FindOneCart(id, userId)
	if err != nil {
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "Cart Not Found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Detail Cart",
		Results: order,
	})
}

func AddToCart(c *gin.Context) {
	// Get user id via token payload
	user := c.MustGet("id").(*models.User)
	userId := user.Id

	// Get body request
	productIdInput, _ := strconv.Atoi(c.PostForm("productId"))
	productSizeIdInput, _ := strconv.Atoi(c.PostForm("productSizeId"))
	productVariantIdInput, _ := strconv.Atoi(c.PostForm("productVariantId"))
	quantityInput, _ := strconv.Atoi(c.PostForm("quantity"))

	// Validate product id
	product, err := models.FindOneProduct(productIdInput)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "Product Not Found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	// Validate product size id
	productSize, err := models.FindOneProductSize(productSizeIdInput)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "Product Size Not Found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	// Validate product variant id
	productVariant, err := models.FindOneProductVariant(productVariantIdInput)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "Product Variant Not Found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	// Calculate total price
	productBasePrice := float64(*product.BasePrice)
	productDiscount := product.Discount
	productSizeAdditionalPrice := float64(*productSize.AdditionalPrice)
	productVariantAdditionalPrice := float64(*productVariant.AdditionalPrice)
	totalProductPrice := ((productBasePrice - (productBasePrice * *productDiscount)) + productSizeAdditionalPrice + productVariantAdditionalPrice) * float64(quantityInput)

	// Cart struct
	dataCart := models.Cart{}

	dataCart.ProductId = &product.Id
	dataCart.ProductSizeId = &productSize.Id
	dataCart.ProductVariantId = &productVariant.Id
	dataCart.Quantity = &quantityInput
	dataCart.Total = &totalProductPrice
	dataCart.UserId = &userId

	// Order Detail struct
	dataOrderDetail := models.OrderDetail{}

	dataOrderDetail.ProductId = &product.Id
	dataOrderDetail.ProductSizeId = &productSize.Id
	dataOrderDetail.ProductVariantId = &productVariant.Id
	dataOrderDetail.Quantity = &quantityInput
	dataOrderDetail.UserId = &userId

	// Add data to query
	cart, err := models.CreateCart(dataCart)
	models.CreateOrderDetail(dataOrderDetail)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Add Product to Cart Successfuully",
		Results: cart,
	})
}

func UpdateCart(c *gin.Context) {
	var productSize models.ProductSize
	var productVariant models.ProductVariant
	var productSizeAddPrice float64
	var productVariantAddPrice float64
	var quantityProduct int

	// Get user id via token payload
	user := c.MustGet("id").(*models.User)
	userId := user.Id

	// Get query param
	id, _ := strconv.Atoi(c.Param("id"))

	// Find existing cart
	existingCart, err := models.FindOneCart(id, userId)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Cart Not Found",
			})
			return
		}
		fmt.Println(existingCart)
	}

	product, _ := models.FindOneProduct(*existingCart.ProductId)

	// Get body request
	productSizeIdInput, err := strconv.Atoi(c.DefaultPostForm("productSizeId", ""))

	if productSizeIdInput == 0 || err != nil {
		productSize, _ = models.FindOneProductSize(*existingCart.ProductSizeId)
		productSizeAddPrice = float64(*productSize.AdditionalPrice)
	} else {
		productSize, _ = models.FindOneProductSize(productSizeIdInput)
		productSizeAddPrice = float64(*productSize.AdditionalPrice)
	}

	productVariantIdInput, err := strconv.Atoi(c.DefaultPostForm("productVariantId", ""))

	if productVariantIdInput == 0 || err != nil {
		productVariant, _ = models.FindOneProductVariant(*existingCart.ProductVariantId)
		productVariantAddPrice = float64(*productVariant.AdditionalPrice)
	} else {
		productVariant, _ = models.FindOneProductVariant(productVariantIdInput)
		productVariantAddPrice = float64(*productVariant.AdditionalPrice)
	}

	quantityInput, err := strconv.Atoi(c.DefaultPostForm("quantity", ""))

	if quantityInput == 0 || err != nil {
		quantityProduct = *existingCart.Quantity
	} else {
		quantityProduct = quantityInput
	}

	// Calculate total price
	productBasePrice := float64(*product.BasePrice)
	productDiscount := product.Discount
	productSizeAdditionalPrice := productSizeAddPrice
	productVariantAdditionalPrice := productVariantAddPrice
	totalProductPrice := ((productBasePrice - (productBasePrice * *productDiscount)) + productSizeAdditionalPrice + productVariantAdditionalPrice) * float64(quantityProduct)

	// Cart struct
	dataCart := models.Cart{}

	// Order Detail struct
	dataOrderDetail := models.OrderDetail{}

	c.ShouldBind(&dataCart)
	c.ShouldBind(&dataOrderDetail)

	dataCart.Id = id
	dataCart.UserId = &userId
	dataOrderDetail.ProductId = &product.Id
	dataOrderDetail.UserId = &userId

	if productSizeIdInput != 0 {
		dataCart.ProductSizeId = &productSizeIdInput
		dataOrderDetail.ProductSizeId = &productSizeIdInput
	}

	if productVariantIdInput != 0 {
		dataCart.ProductVariantId = &productVariantIdInput
		dataOrderDetail.ProductVariantId = &productVariantIdInput
	}

	if quantityInput != 0 {
		dataCart.Quantity = &quantityInput
		dataOrderDetail.Quantity = &quantityInput
	}

	dataCart.Total = &totalProductPrice

	cart, err := models.UpdateCart(dataCart)
	models.UpdateOrderDetail(dataOrderDetail)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Update Cart Succesfully",
		Results: cart,
	})
}

func DeleteCart(c *gin.Context) {
	// Get user id via token payload
	user := c.MustGet("id").(*models.User)
	userId := user.Id

	// Get cart id from body request
	id, _ := strconv.Atoi(c.Param("id"))

	// Validate cart id
	_, err := models.FindOneCart(id, userId)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "Data Cart Not Found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	// Delete data by query
	cart, err := models.DeleteCart(id, userId)
	models.DeleteOrderDetail(*cart.ProductId, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Delete Cart Succesfully",
		Results: cart,
	})
}
