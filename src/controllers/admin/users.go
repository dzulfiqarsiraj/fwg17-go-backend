package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/DzulfiqarSiraj/go-backend/src/models"
	"github.com/KEINOS/go-argonize"
	"github.com/gin-gonic/gin"
)

type pageInfo struct {
	Page int `json:"page"`
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
	Id          int         `json:"id"`
	Email       string      `json:"email"`
	Password    string      `json:"password"`
	FullName    interface{} `json:"fullName"`
	PhoneNumber interface{} `json:"phoneNumber"`
	Address     interface{} `json:"address"`
	Role        *string     `json:"role"`
	Picture     interface{} `json:"picture"`
	CreatedAt   *time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time  `json:"updatedAt"`
}

func ListAllUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	users, err := models.FindAllUsers()
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, &responseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, responseList{
		Success: true,
		Message: "List All Users",
		PageInfo: pageInfo{
			Page: page,
		},
		Results: users,
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
	emailInput := c.PostForm("email")
	passwordInput := c.PostForm("password")

	if emailInput == "" || passwordInput == "" {
		c.JSON(http.StatusBadRequest, &responseOnly{
			Success: false,
			Message: "Email or Password Must Not Be Empty",
		})
		return
	}

	existingEmail, _ := models.FindOneUserByEmail(emailInput)

	if existingEmail.Email == emailInput {
		c.JSON(http.StatusBadRequest, &responseOnly{
			Success: false,
			Message: "Email is Already Used",
		})
		return
	}

	c.Bind(&data)

	plain := []byte(data.Password)
	hash, err := argonize.Hash(plain)

	if err != nil {
		c.JSON(http.StatusBadRequest, &responseOnly{
			Success: false,
			Message: "Can't Generate Hashed Password",
		})
		return
	}

	data.Password = hash.String()

	user, err := models.CreateUser(data)
	if err != nil {
		log.Fatalln(err)
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
