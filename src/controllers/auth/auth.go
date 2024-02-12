package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DzulfiqarSiraj/go-backend/src/lib"
	"github.com/DzulfiqarSiraj/go-backend/src/models"
	"github.com/DzulfiqarSiraj/go-backend/src/services"
	"github.com/KEINOS/go-argonize"
	"github.com/gin-gonic/gin"
)

type FormResetPassword struct {
	Email           string `form:"email"`
	Otp             string `form:"otp"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirmPassword"`
}

func ForgotPassword(c *gin.Context) {
	form := FormResetPassword{}
	c.ShouldBind(&form)
	if form.Email != "" {
		found, _ := models.FindOneUserByEmail(form.Email)
		if found.Id == 0 {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Email Not Found",
			})
			return
		}

		formResetPassword := models.FormResetPassword{
			Otp:   lib.RandomNumberStr(6),
			Email: found.Email,
		}
		models.CreateResetPassword(formResetPassword)
		// here
		fmt.Println(formResetPassword.Otp)
		// Send OTP to Email Process
		c.JSON(http.StatusOK, &services.ResponseOnly{
			Success: true,
			Message: "OTP Has Been Sent to Your Email",
		})
		return

	}

	if form.Otp != "" {
		found, _ := models.FindOneRPByOtp(form.Otp)
		if found.Id == 0 {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Failed to Reset Password",
			})
			return
		}

		if form.Password != form.ConfirmPassword {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "Confirm Password Doesn't Match",
			})
			return
		}

		foundUser, _ := models.FindOneUserByEmail(found.Email)
		hashedPassword, _ := argonize.Hash([]byte(form.Password))

		data := models.User{
			Id:       foundUser.Id,
			Password: hashedPassword.String(),
		}

		updated, _ := models.UpdateUser(data)
		message := fmt.Sprintf("Reset Password for %v is Success", updated.Email)
		c.JSON(http.StatusOK, &services.ResponseOnly{
			Success: true,
			Message: message,
		})
		models.DeleteResetPassword(found.Id)
		return
	}
	c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
		Success: false,
		Message: "Internal Server Error",
	})
}

func Login(c *gin.Context) {

}

func Register(c *gin.Context) {
	form := models.User{}

	err := c.ShouldBind(&form)

	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Invalid",
		})
		return
	}

	existingUser, _ := models.FindOneUserByEmail(form.Email)

	if existingUser.Email == form.Email {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Email is Already Used",
		})
		return
	}

	plainPassword := []byte(form.Password)
	hashPassword, err := argonize.Hash(plainPassword)

	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Can't Generate Hashed Password",
		})
		return
	}

	form.Password = hashPassword.String()
	form.Role = "Customer"

	user, err := models.CreateUser(form)
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "Register Success",
		Results: user,
	})
}
