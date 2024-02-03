package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/DzulfiqarSiraj/go-backend/src/models"
	"github.com/gin-gonic/gin"
)

func ListAllProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	products, err := models.FindAllProducts()
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, &responseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &responseList{
		Success: true,
		Message: "List All Products",
		PageInfo: pageInfo{
			Page: page,
		},
		Results: products,
	})
}

func DetailProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := models.FindOneProduct(id)
	if err != nil {
		log.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &responseOnly{
				Success: false,
				Message: "Product Not Found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, &responseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "Detail Product",
		Results: product,
	})
}

func CreateProduct(c *gin.Context) {
	data := models.Product{}
	nameInput := c.PostForm("name")

	if nameInput == "" {
		c.JSON(http.StatusBadRequest, &responseOnly{
			Success: false,
			Message: "Name Must Not Be Empty",
		})
		return
	}

	existingProduct, _ := models.FindOneProductByName(nameInput)

	if existingProduct.Name == &nameInput {
		c.JSON(http.StatusBadRequest, &responseOnly{
			Success: false,
			Message: "Name is Already Exist",
		})
		return
	}

	c.Bind(&data)

	product, err := models.CreateProduct(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &responseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "Product Created Successfully",
		Results: product,
	})
}

func UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	nameInput := c.PostForm("name")

	if nameInput != "" {
		existingProduct, _ := models.FindOneProductByName(nameInput)

		if nameInput == *existingProduct.Name {
			c.JSON(http.StatusBadRequest, &responseOnly{
				Success: false,
				Message: "Name is Already Used",
			})
			return
		}
	}

	data := models.Product{}

	c.Bind(&data)

	data.Id = id

	product, err := models.UpdateProduct(data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &responseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "Update Product Successfully",
		Results: product,
	})
}

func DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := models.DeleteProduct(id)
	if err != nil {
		log.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &responseOnly{
				Success: false,
				Message: "No Data",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, &responseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "Delete Product",
		Results: product,
	})
}
