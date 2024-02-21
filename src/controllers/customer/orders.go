package customer_controllers

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/DzulfiqarSiraj/go-backend/src/models"
	"github.com/DzulfiqarSiraj/go-backend/src/services"
	"github.com/gin-gonic/gin"
)

func ListAllOrders(c *gin.Context) {
	data := c.MustGet("id").(*models.User)
	userId := data.Id
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "4"))
	offset := (page - 1) * limit
	result, err := models.FindAllOrders(userId, limit, offset)

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
	data := models.Order{}
	user := c.MustGet("id").(*models.User)
	userId := user.Id

	fullNameInput := c.PostForm("fullName")
	emailInput := c.PostForm("email")

	cartInfo, err := models.TotalPrice(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "No Product in Cart",
		})
	}

	if fullNameInput == "" {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Name Must Not Be Empty",
		})
		return
	}

	if emailInput == "" {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Email Must Not Be Empty",
		})
		return
	}

	orderTime := time.DateOnly

	data.OrderNumber = &orderTime

	// Total Price Count
	plainPrice := float64(cartInfo.TotalPrice)
	tax := 0.1
	taxPrice := plainPrice * tax
	total := plainPrice + taxPrice
	// ~Total Price Count

	c.ShouldBind(&data)

	data.UserId = &userId
	data.Tax = &tax
	data.Total = &total

	order, err := models.CreateOrder(data)

	// Update orderNumber
	orderNew := models.Order{}
	orderNumber := fmt.Sprintf(`#INV-%v-%v`, order.Id, time.DateOnly)

	orderNew.Id = order.Id
	orderNew.OrderNumber = &orderNumber

	models.UpdateOrderNumber(orderNew)

	// Update orderId at order details
	orderDetails := models.OrderDetail{}

	orderDetails.OrderId = &order.Id

	models.UpdateOrderDetailByOrderId(userId, orderDetails)
	models.DeleteAllCart(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Order Created Successfully",
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
