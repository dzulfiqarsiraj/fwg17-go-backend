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

func ListAllPromo(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit
	result, err := models.FindAllPromo(limit, offset)

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
		Message:  "List All Promo",
		PageInfo: *pageInfo,
		Results:  result.Data,
	})
}

func DetailPromo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	promo, err := models.FindOnePromo(id)
	if err != nil {
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "Promo Not Found",
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
		Results: promo,
	})
}

func CreatePromo(c *gin.Context) {
	data := models.Promo{}
	nameInput := c.PostForm("name")

	if nameInput == "" {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Promo Name Must Not Be Empty",
		})
		return
	}
	existingPromo, _ := models.FindOnePromoByName(nameInput)
	existingPromoName := existingPromo.Name

	if *existingPromoName == nameInput {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Promo is Already Exist",
		})
		return
	}

	c.ShouldBind(&data)

	promo, err := models.CreatePromo(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Promo Created Succesfully",
		Results: promo,
	})
}

func UpdatePromo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	nameInput := c.PostForm("name")

	existingPromo, err := models.FindOnePromo(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Promo Not Found",
			})
			return
		}
		fmt.Println(existingPromo)
	}

	if nameInput != "" {
		existingPromo, _ := models.FindOnePromoByName(nameInput)

		if nameInput == *existingPromo.Name {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Promo is Already Exist",
			})
			return
		}
	}

	data := models.Promo{}

	c.ShouldBind(&data)

	data.Id = id

	promo, err := models.UpdatePromo(data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Update Promo Succesfully",
		Results: promo,
	})
}

func DeletePromo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	promo, err := models.DeletePromo(id)
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
		Results: promo,
	})
}
