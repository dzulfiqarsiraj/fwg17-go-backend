package controllers

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/DzulfiqarSiraj/go-backend/src/models"
	"github.com/gin-gonic/gin"
)

type pageInfo struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	LastPage  int `json:"lastPage"`
	TotalData int `json:"totalData"`
}

type responseList struct {
	Success  bool        `json:"success"`
	Message  string      `json:"message"`
	PageInfo pageInfo    `json:"pageInfo"`
	Results  interface{} `json:"results"`
}
type response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Results interface{} `json:"results"`
}

type responseOnly struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type User struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	FullName    string `json:"fullName"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
	Role        string `json:"role"`
	Picture     string `json:"picture"`
}

func ListAllUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
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
	result, err := models.FindAllUsers(limit, offset)

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
		Message:  "List All Users",
		PageInfo: pageInfo,
		Results:  result.Data,
	})
}

func DetailUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := models.FindOneUser(id)
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
		Message: "Detail User",
		Results: user,
	})
}

func CreateUser(c *gin.Context) {
	data := models.User{}

	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, &responseOnly{
			Success: false,
			Message: "Invalid Input",
		})
		return
	}

	user, err := models.CreateUser(data)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, &responseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "User Created Successfully",
		Results: user,
	})
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data := models.User{}

	c.Bind(&data)
	data.Id = id

	user, err := models.UpdateUser(data)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, &responseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "Update User Successfully",
		Results: user,
	})
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := models.DeleteUser(id)
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
		Message: "Delete User",
		Results: user,
	})
}
