package admin_controllers

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/DzulfiqarSiraj/go-backend/src/lib"
	"github.com/DzulfiqarSiraj/go-backend/src/models"
	"github.com/DzulfiqarSiraj/go-backend/src/services"
	"github.com/KEINOS/go-argonize"
	"github.com/gin-gonic/gin"
)

func ListAllUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	search := c.DefaultQuery("search", "")
	orderBy := c.DefaultQuery("orderBy", "id")
	// orderBy := c.DefaultQuery("orderBy", "id")
	offset := (page - 1) * limit
	result, err := models.FindAllUsers(search, orderBy, limit, offset)

	totalPage := int(math.Ceil(float64(result.Count) / float64(limit)))

	pageInfo := &services.PageInfo{
		Page:      page,
		Limit:     limit,
		TotalPage: totalPage,
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
		Message:  "List All Users",
		PageInfo: *pageInfo,
		Results:  result.Data,
	})
}

func DetailUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := models.FindOneUser(id)
	if err != nil {
		log.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "User Not Found",
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
		Message: "Detail User",
		Results: user,
	})
}

func CreateUser(c *gin.Context) {
	data := models.User{}
	emailInput := c.PostForm("email")
	passwordInput := c.PostForm("password")
	roleInput := c.PostForm("role")
	pictureInput, _, _ := c.Request.FormFile("picture")

	if emailInput == "" || passwordInput == "" {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Email or Password Must Not Be Empty",
		})
		return
	}

	existingUser, _ := models.FindOneUserByEmail(emailInput)

	if existingUser.Email == emailInput {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Email is Already Used",
		})
		return
	}

	if roleInput != "" {
		switch roleInput {
		case "Customer", "Staff Administrator", "Super Administrator":
			fmt.Println("Valid Role")
		default:
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Invalid Role",
			})
			return
		}
	} else {
		data.Role = "Customer"
	}

	c.ShouldBind(&data)

	// upload file
	if pictureInput != nil {
		data.Picture = lib.Upload(c, "picture", "users")
		if *data.Picture == "Invalid File Type" {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "File Type Must Be jpg/jpeg/png",
			})
			return
		}

		if *data.Picture == "Invalid File Size" {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "File Size Must Less than 1MB",
			})
			return
		}
	}
	// *upload file

	plain := []byte(data.Password)
	hash, err := argonize.Hash(plain)

	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Can't Generate Hashed Password",
		})
		return
	}

	data.Password = hash.String()

	user, err := models.CreateUser(data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "User Created Successfully",
		Results: user,
	})
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	emailInput := c.PostForm("email")
	fileInput, _, _ := c.Request.FormFile("picture")
	fmt.Println(&fileInput)

	existingUser, err := models.FindOneUser(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "User Not Found",
			})
			return
		}
		fmt.Println(existingUser)
	}

	if emailInput != "" {
		existingUser, _ := models.FindOneUserByEmail(emailInput)

		if emailInput == existingUser.Email {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Email is Already Used",
			})
			return
		}
	}

	data := models.User{}

	c.ShouldBind(&data)

	if data.Password != "" {
		plain := []byte(data.Password)
		hash, err := argonize.Hash(plain)
		if err != nil {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Can't Generate Hashed Password",
			})
			return
		}

		data.Password = hash.String()
	}

	// upload file
	if fileInput != nil {
		data.Picture = lib.Upload(c, "picture", "users")
		if *data.Picture == "Invalid File Type" {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "File Type Must Be jpg/jpeg/png",
			})
			return
		}

		if *data.Picture == "Invalid File Size" {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "File Size Must Less than 1MB",
			})
			return
		}

		if existingUser.Picture != nil {
			fileName := *existingUser.Picture
			fileDest := fmt.Sprintf("uploads/users/%v", fileName)
			fmt.Println(fileDest)
			os.Remove(fileDest)
		}
	}
	// *upload file

	data.Id = id

	user, err := models.UpdateUser(data)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Update User Successfully",
		Results: user,
	})
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := models.DeleteUser(id)
	if user.Picture != nil {
		fileName := *user.Picture
		fileDest := fmt.Sprintf("uploads/users/%v", fileName)
		fmt.Println(fileDest)
		os.Remove(fileDest)
	}
	fmt.Println(*user.Picture)
	if err != nil {
		log.Println(err)
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
		Message: "Delete User Successfully",
		Results: user,
	})
}
