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
