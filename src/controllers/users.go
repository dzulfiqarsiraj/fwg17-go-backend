package controllers

import (
	"net/http"
	"strconv"

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
	Id       int    `json:"id" form:"id"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ListAllUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	c.JSON(http.StatusOK, responseList{
		Success: true,
		Message: "List All Users",
		PageInfo: pageInfo{
			Page: page,
		},
		Results: []User{
			{
				Id:       1,
				Email:    "admin@mail.com",
				Password: "1234",
			},
			{
				Id:       2,
				Email:    "guest@mail.com",
				Password: "1234",
			},
		},
	})
}

func DetailUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "Detail User",
		Results: User{
			Id:       id,
			Email:    "admin@mail.com",
			Password: "1234",
		},
	})
}

func CreateUser(c *gin.Context) {
	user := User{}
	c.Bind(&user)
	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "User Created Successfully",
		Results: user,
	})
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, &responseOnly{
	// 		Success: false,
	// 		Message: "Required Param id",
	// 	})
	// }
	user := User{}

	c.Bind(&user)
	user.Id = id
	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "Update User Successfully",
		Results: user,
	})
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "User Deleted Successfully",
		Results: User{
			Id:       id,
			Email:    "guest@mail.com",
			Password: "1234",
		},
	})
}
