package customer_controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/DzulfiqarSiraj/go-backend/src/lib"
	"github.com/DzulfiqarSiraj/go-backend/src/models"
	"github.com/DzulfiqarSiraj/go-backend/src/services"
	"github.com/KEINOS/go-argonize"
	"github.com/gin-gonic/gin"
)

func UserProfile(c *gin.Context) {
	data := c.MustGet("id").(*models.User)
	id := data.Id
	fmt.Println(id)

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

func UpdateProfile(c *gin.Context) {
	customer := c.MustGet("id").(*models.User)
	id := customer.Id

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
