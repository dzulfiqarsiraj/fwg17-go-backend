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

func ListAllOrderDetails(c *gin.Context) {
	// data := c.MustGet("id").(*models.User)
	orderId, _ := strconv.Atoi(c.Query("orderId"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "4"))
	offset := (page - 1) * limit
	result, err := models.FindAllOrderDetails(orderId, limit, offset)
	// fmt.Println(orderId)

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
