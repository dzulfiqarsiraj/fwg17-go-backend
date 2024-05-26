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

func ListAllProductCategories(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "4"))
	offset := (page - 1) * limit
	result, err := models.FindAllProductCategories(limit, offset)

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
		Message:  "List All Product Categories",
		PageInfo: *pageInfo,
		Results:  result.Data,
	})
}

func DetailProductCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	productCategory, err := models.FindOneProductCategory(id)
	if err != nil {
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "Product Category Not Found",
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
		Message: "Detail Product Category",
		Results: productCategory,
	})
}

func CreateProductCategory(c *gin.Context) {
	data := models.ProductCategory{}
	productIdInput, err := strconv.Atoi(c.PostForm("productId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Product Id Must Not Be Empty",
		})
		return
	} else {
		fmt.Println(productIdInput)
	}

	categoryIdInput, err := strconv.Atoi(c.PostForm("categoryId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Product Id Must Not Be Empty",
		})
		return
	} else {
		fmt.Println(categoryIdInput)
	}

	c.ShouldBind(&data)

	productCategory, err := models.CreateProductCategory(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Product Category Created Successfully",
		Results: productCategory,
	})
}

func UpdateProductCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data := models.ProductCategory{}

	c.ShouldBind(&data)

	data.Id = id

	productCategory, err := models.UpdateProductCategory(data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Update Product Category Succesfully",
		Results: productCategory,
	})
}

func DeleteProductCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	productCategory, err := models.DeleteProductCategory(id)
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
		Message: "Delete Product Category Succesfully",
		Results: productCategory,
	})
}
