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

func ListAllProductTags(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit
	result, err := models.FindAllProductTags(limit, offset)

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
		Message:  "List All Product Tags",
		PageInfo: *pageInfo,
		Results:  result.Data,
	})
}

func DetailProductTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	productTag, err := models.FindOneProductTag(id)
	if err != nil {
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "Product Tag Not Found",
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
		Message: "Detail Product Tag",
		Results: productTag,
	})
}

func CreateProductTag(c *gin.Context) {
	data := models.ProductTag{}

	c.ShouldBind(&data)

	productTag, err := models.CreateProductTag(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Product Tag Created Succesfully",
		Results: productTag,
	})
}

func UpdateProductTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data := models.ProductTag{}

	c.ShouldBind(&data)

	data.Id = id

	productTag, err := models.UpdateProductTag(data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Update Product Tag Succesfully",
		Results: productTag,
	})
}

func DeleteProductTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	productTag, err := models.DeleteProductTag(id)
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
		Message: "Delete Product Tag Succesfully",
		Results: productTag,
	})
}
