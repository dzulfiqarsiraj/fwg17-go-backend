package controllers

import (
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/DzulfiqarSiraj/go-backend/src/models"
	"github.com/DzulfiqarSiraj/go-backend/src/services"
	"github.com/gin-gonic/gin"
)

func ListAllProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	keyword := c.DefaultQuery("keyword", "")
	category := c.DefaultQuery("category", "")
	orderBy := c.DefaultQuery("orderBy", "id")

	offset := (page - 1) * limit
	result, err := models.FindAllProducts(category, keyword, orderBy, limit, offset)

	totalPage := int(math.Ceil(float64(result.Count) / float64(limit)))
	var nextPage any
	var prevPage any

	if (int(page) + 1) <= totalPage {
		nextPage = int(page) + 1
	} else {
		nextPage = nil
	}

	if (int(page) - 1) > 0 {
		prevPage = int(page) - 1
	} else {
		prevPage = nil
	}

	pageInfo := &services.PageInfo{
		CurrentPage: page,
		NextPage:    nextPage,
		PrevPage:    prevPage,
		Limit:       limit,
		TotalPage:   totalPage,
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
		Message:  "List All Products",
		PageInfo: *pageInfo,
		Results:  result.Data,
	})
}

func DetailProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := models.FindOneProductDetailed(id)

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

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Detail Product",
		Results: product,
	})
}
