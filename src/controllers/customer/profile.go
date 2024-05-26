package customer_controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/DzulfiqarSiraj/go-backend/src/lib"
	"github.com/DzulfiqarSiraj/go-backend/src/models"
	"github.com/DzulfiqarSiraj/go-backend/src/services"
	"github.com/KEINOS/go-argonize"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

func UserProfile(c *gin.Context) {
	data := c.MustGet("id").(*models.User)
	id := data.Id

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

	data := models.User{}

	c.ShouldBind(&data)

	fileInput, _ := c.FormFile("pictures")

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

	// upload file -start
	if fileInput != nil {
		data.Pictures, _ = lib.Upload(c, "pictures", "users")
		if *data.Pictures == "Invalid File Type" {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "File Type Must Be jpg/jpeg/png",
			})
			return
		}

		if *data.Pictures == "Invalid File Size" {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "File Size Must Less than 1MB",
			})
			return
		}

		if existingUser.Pictures != nil {
			// Delete src picture from local storage - start
			// fileName := *existingUser.Pictures
			// fileDest := fmt.Sprintf("cov-shop/users/%v", fileName)
			// os.Remove(fileDest)
			// Delete src picture from local storage - end

			// Delete src picture from cloudinary - start
			existingPicture := strings.Split(*existingUser.Pictures, "/")
			slicedExistingPicture := strings.Split(existingPicture[9], ".")
			existingPublicId := "cov-shop/users/" + slicedExistingPicture[0]
			fmt.Println(existingPublicId)
			cloud := lib.CloudinaryConfig()
			var ctx = context.Background()
			resp, err := cloud.Upload.Destroy(ctx, uploader.DestroyParams{
				PublicID: existingPublicId,
			})

			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(resp.Result)
			// Delete src picture from cloudinary - start
		}
	}
	// upload file - end

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
