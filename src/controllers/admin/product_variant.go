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

func ListAllProductVariant(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit
	result, err := models.FindAllProductVariant(limit, offset)

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
		Message:  "List All Product Variant",
		PageInfo: *pageInfo,
		Results:  result.Data,
	})
}

func DetailProductVariant(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	productVariant, err := models.FindOneProductVariant(id)
	if err != nil {
		fmt.Println(err)
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

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Detail Product Variant",
		Results: productVariant,
	})
}

func CreateProductVariant(c *gin.Context) {
	data := models.ProductVariant{}
	nameInput := c.PostForm("name")

	if nameInput == "" {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Variant Name Must Not Be Empty",
		})
		return
	}
	existingProductVariant, _ := models.FindOneProductVariantByName(nameInput)
	existingProductVariantName := existingProductVariant.Name

	if *existingProductVariantName == nameInput {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Product Variant is Already Exist",
		})
		return
	}

	c.ShouldBind(&data)

	productVariant, err := models.CreateProductVariant(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Product Variant Created Succesfully",
		Results: productVariant,
	})
}

func UpdateProductVariant(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	nameInput := c.PostForm("name")

	existingProductVariant, err := models.FindOneProductVariant(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Product Variant Not Found",
			})
			return
		}
		fmt.Println(existingProductVariant)
	}

	if nameInput != "" {
		existingProductVariant, _ := models.FindOneProductVariantByName(nameInput)

		if nameInput == *existingProductVariant.Name {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Variant is Already Exist",
			})
			return
		}
	}

	data := models.ProductVariant{}

	c.ShouldBind(&data)

	data.Id = id

	productVariant, err := models.UpdateProductVariant(data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Update Product Variant Succesfully",
		Results: productVariant,
	})
}

func DeleteProductVariant(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	productVariant, err := models.DeleteProductVariant(id)
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
		Message: "Delete Product Variant Succesfully",
		Results: productVariant,
	})
}
