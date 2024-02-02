package controllers

import (
	"database/sql"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/DzulfiqarSiraj/go-backend/src/models"
	"github.com/gin-gonic/gin"
)

type Product struct {
	Id           int          `json:"id"`
	Name         string       `json:"name" form:"name"`
	BasePrice    string       `json:"basePrice" form:"basePrice"`
	Description  string       `json:"description" form:"description"`
	Image        string       `son:"image" form:"image"`
	IsBestSeller string       `json:"isBestSeller" form:"isBestSeller"`
	Discount     string       `json:"discount" form:"discount"`
	CreatedAt    time.Time    ` json:"createdAt"`
	UpdatedAt    sql.NullTime `json:"updatedAt"`
}

func ListAllProducts(c *gin.Context) {
	page, _ := strconv.Atoi((c.DefaultQuery("page", "1")))
	if page < 1 {
		c.JSON(http.StatusBadRequest, &responseOnly{
			Success: false,
			Message: "No Such Page",
		})
		return
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	if limit < 1 {
		c.JSON(http.StatusBadRequest, &responseOnly{
			Success: false,
			Message: "Limit Must Be At Least 1",
		})
		return
	}
	offset := (page - 1) * limit
	result, err := models.FindAllProducts(limit, offset)

	pageInfo := pageInfo{
		Page:      page,
		Limit:     limit,
		LastPage:  int(math.Ceil(float64(result.Count) / float64(limit))),
		TotalData: result.Count,
	}

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, &responseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, responseList{
		Success:  true,
		Message:  "List All Products",
		PageInfo: pageInfo,
		Results:  result.Data,
	})
}
