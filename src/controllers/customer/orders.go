package customer_controllers

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/DzulfiqarSiraj/go-backend/src/models"
	"github.com/DzulfiqarSiraj/go-backend/src/services"
	"github.com/gin-gonic/gin"
)

type Size struct {
	AdditionalPrice int    `json:"additionalPrice"`
	Id              int    `json:"id"`
	Size            string `json:"size"`
}

type Variant struct {
	AdditionalPrice int    `json:"additionalPrice"`
	Id              int    `json:"id"`
	Variant         string `json:"variant"`
}

type Product struct {
	BasePrice     int     `json:"basePrice"`
	Category      string  `json:"category"`
	CreatedAt     string  `json:"createdAt"`
	Description   string  `json:"description"`
	Discount      float64 `json:"discount"`
	Id            int     `json:"id"`
	Image         string  `json:"image"`
	IsRecommended bool    `json:"isRecommended"`
	Name          string  `json:"name"`
	Sizes         Size    `json:"sizes"`
	Tag           string  `json:"tag"`
	UpdatedAt     string  `json:"updatedAt"`
	Variants      Variant `json:"variants"`
}

type ProductData struct {
	Product  Product `json:"product"`
	Size     Size    `json:"size"`
	Variant  Variant `json:"variant"`
	Quantity int     `json:"quantity"`
	UniqueId string  `json:"uniqueId"`
}

type Cart struct {
	CartData      []ProductData     `json:"cartData"`
	CustData      map[string]string `json:"custData"`
	ShippingData  string            `json:"shippingData"`
	ShippingPrice int               `json:"shippingPrice"`
}

func ListAllOrders(c *gin.Context) {
	data := c.MustGet("id").(*models.User)
	userId := data.Id
	status := c.DefaultQuery("status", "Awaiting Payment")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "4"))
	offset := (page - 1) * limit
	result, err := models.FindAllOrders(userId, status, limit, offset)

	pageInfo := &services.PageInfo{
		CurrentPage: page,
		Limit:       limit,
		TotalPage:   int(math.Ceil(float64(result.Count) / float64(limit))),
		TotalData:   result.Count,
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
		Message:  "List All Orders",
		PageInfo: *pageInfo,
		Results:  result.Data,
	})
}

func DetailOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	order, err := models.FindOneOrder(id)
	if err != nil {
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "Order Not Found",
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
		Message: "Detail Order",
		Results: order,
	})
}

func CreateOrder(c *gin.Context) {
	var cartData Cart
	orderData := models.Order{}
	user := c.MustGet("id").(*models.User)
	userId := user.Id

	c.ShouldBind(&cartData)

	// grandTotal Calculation
	var total float64

	for i := 0; i < len(cartData.CartData); i++ {
		discount := cartData.CartData[i].Product.Discount
		basePrice := float64(cartData.CartData[i].Product.BasePrice)
		sizeAdditionalPrice := float64(cartData.CartData[i].Size.AdditionalPrice)
		variantAdditionalPrice := float64(cartData.CartData[i].Variant.AdditionalPrice)
		quantity := float64(cartData.CartData[i].Quantity)
		total = total + (((basePrice - (basePrice * discount)) + variantAdditionalPrice + sizeAdditionalPrice) * quantity)
	}

	grandTotal := total + float64(cartData.ShippingPrice) + (total * 0.05)

	var randNumber string
	arrRandNumber := rand.Perm(6)
	for i := 0; i < len(arrRandNumber); i++ {
		randNumber = randNumber + fmt.Sprintf("%v", arrRandNumber[i])
	}

	var customerData map[string]string = cartData.CustData
	fullName := customerData["fullname"]
	email := customerData["email"]
	address := customerData["address"]
	tax := 0.05
	status := "Awaiting Payment"

	// Mapping Order Data
	orderData.UserId = &userId
	orderData.OrderNumber = &randNumber
	orderData.FullName = &fullName
	orderData.Email = &email
	orderData.Tax = &tax
	orderData.GrandTotal = &grandTotal
	orderData.DeliveryAddress = &address
	orderData.Status = &status
	orderData.Shipping = &cartData.ShippingData

	// Create Order
	order, err := models.CreateOrder(orderData)
	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Can't Create Order",
		})
		return
	}

	// Find Order Id
	existOrder, err := models.FindOneOrderByOrderNumber(randNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, &services.ResponseOnly{
			Success: false,
			Message: "Order Not Found",
		})
		return
	}
	existOrderId := existOrder.Id

	// Add Data Order to Order Detail
	for i := 0; i < len(cartData.CartData); i++ {
		var orderDetailData models.OrderDetail
		var basePrice float64 = float64(cartData.CartData[i].Product.BasePrice)
		sizePrice := (cartData.CartData[i].Size.AdditionalPrice)
		variantPrice := (cartData.CartData[i].Variant.AdditionalPrice)
		discount := cartData.CartData[i].Product.Discount
		quantity := cartData.CartData[i].Quantity
		subTotal := ((basePrice - (basePrice * discount)) + float64(sizePrice) + float64(variantPrice)) * float64(quantity)

		// 	// Mapping Data to Order Detail
		orderDetailData.UserId = &userId
		orderDetailData.OrderId = &existOrderId
		orderDetailData.ProductId = &cartData.CartData[i].Product.Id
		orderDetailData.ProductSizeId = &cartData.CartData[i].Size.Id
		orderDetailData.ProductVariantId = &cartData.CartData[i].Variant.Id
		orderDetailData.Quantity = &quantity
		orderDetailData.SubTotal = &subTotal

		_, err := models.CreateOrderDetail(orderDetailData)
		if err != nil {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Can't Create Order Detail",
			})
			return
		}
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Ok",
		Results: order,
	})
}

func UpdateOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	existingOrder, err := models.FindOneOrder(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Order Not Found",
			})
			return
		}
		fmt.Println(existingOrder)
	}

	data := models.Order{}

	c.ShouldBind(&data)

	data.Id = id

	order, err := models.UpdateOrder(data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Update Order Succesfully",
		Results: order,
	})
}

func DeleteOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	order, err := models.DeleteOrder(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "No Data",
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
		Message: "Delete Order Succesfully",
		Results: order,
	})
}
