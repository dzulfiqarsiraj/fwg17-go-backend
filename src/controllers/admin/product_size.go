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

func ListAllProductSize(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit
	result, err := models.FindAllProductSize(limit, offset)

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
		Message:  "List All Products",
		PageInfo: *pageInfo,
		Results:  result.Data,
	})
}

func DetailProductize(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	productSize, err := models.FindOneProductSize(id)
	if err != nil {
		log.Println(err)
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

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Detail Product Size",
		Results: productSize,
	})
}

func CreateProductSize(c *gin.Context) {
	data := models.ProductSize{}
	sizeInput := c.PostForm("size")

	if sizeInput == "" {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Size Name Must Not Be Empty",
		})
		return
	}
	existingProduct, _ := models.FindOneProductSizeBySize(sizeInput)
	existingProductSize := existingProduct.Size

	if *existingProductSize == sizeInput {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Product Size is Already Exist",
		})
		return
	}

	c.ShouldBind(&data)

	productSize, err := models.CreateProductSize(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Product Size Created Successfully",
		Results: productSize,
	})
}

func UpdateProductSize(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sizeInput := c.PostForm("size")

	existingProductSize, err := models.FindOneProductSize(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Product Size Not Found",
			})
			return
		}
		fmt.Println(existingProductSize)
	}

	if sizeInput != "" {
		existingProductSize, _ := models.FindOneProductSizeBySize(sizeInput)

		if sizeInput == *existingProductSize.Size {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Size is Already Used",
			})
			return
		}
	}

	data := models.ProductSize{}

	c.ShouldBind(&data)

	data.Id = id

	productSize, err := models.UpdateProductSize(data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Update Product Size Succesfully",
		Results: productSize,
	})
}
func DeleteProductSize(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	productSize, err := models.DeleteProductSize(id)
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
		Message: "Delete Product Size Succesfully",
		Results: productSize,
	})
}
