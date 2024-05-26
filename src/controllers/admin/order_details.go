package admin_controllers

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

func ListAllOrderDetails(c *gin.Context) {
	data := c.MustGet("id").(*models.User)
	userId := data.Id
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "4"))
	offset := (page - 1) * limit
	result, err := models.FindAllOrderDetails(userId, limit, offset)

	pageInfo := &services.PageInfo{
		CurrentPage: page,
		Limit:       limit,
		TotalPage:   int(math.Ceil(float64(result.Count) / float64(limit))),
		TotalData:   result.Count,
	}

	lastOrderId, _ := models.FindMaxIdOrder()
	fmt.Println(&lastOrderId.Max)

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
		Message:  "List All Order Details",
		PageInfo: *pageInfo,
		Results:  result.Data,
	})
}

func DetailOrderDetail(c *gin.Context) {
	data := c.MustGet("id").(*models.User)
	userId := data.Id
	id, _ := strconv.Atoi(c.Param("id"))

	orderDetail, err := models.FindOneOrderDetail(id, userId)
	if err != nil {
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "Order Detail Not Found",
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
		Message: "Detail Order Detail",
		Results: orderDetail,
	})
}

func UpdateOrderDetail(c *gin.Context) {
	userId, _ := strconv.Atoi(c.PostForm("userId"))
	id, _ := strconv.Atoi(c.Param("id"))

	existingOrderDetail, err := models.FindOneOrderDetail(id, userId)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Order Detail Not Found",
			})
			return
		}
		fmt.Println(existingOrderDetail)
	}

	data := models.OrderDetail{}

	c.ShouldBind(&data)

	data.Id = id

	orderDetail, err := models.UpdateOrderDetail(data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Update Order Detail Succesfully",
		Results: orderDetail,
	})
}

func DeleteOrderDetail(c *gin.Context) {
	userId, _ := strconv.Atoi(c.PostForm("userId"))
	id, _ := strconv.Atoi(c.Param("id"))

	orderDetail, err := models.DeleteOrderDetail(id, userId)
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
		Message: "Delete Order Detail Succesfully",
		Results: orderDetail,
	})
}
