package middlewares

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/DzulfiqarSiraj/go-backend/src/models"
	"github.com/DzulfiqarSiraj/go-backend/src/services"
	"github.com/KEINOS/go-argonize"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Auth() (*jwt.GinJWTMiddleware, error) {
	godotenv.Load()
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "go-backend",
		Key:         []byte(os.Getenv("APP_SECRET")),
		IdentityKey: "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			user := data.(*models.User)
			return jwt.MapClaims{
				"id":   user.Id,
				"role": user.Role,
			}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.User{
				Id:   int(claims["id"].(float64)),
				Role: claims["role"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			form := models.User{}
			err := c.ShouldBind(&form)

			if err != nil {
				return nil, err
			}

			found, err := models.FindOneUserByEmail(form.Email)
			if err != nil {
				return nil, err
			}

			decodedPassword, err := argonize.DecodeHashStr(found.Password)
			if err != nil {
				return nil, err
			}

			plainPassword := []byte(form.Password)

			if decodedPassword.IsValidPassword(plainPassword) {
				return &models.User{
					Id:   found.Id,
					Role: found.Role,
				}, nil
			} else {
				return nil, errors.New("invalid_password")
			}
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			user := data.(*models.User)
			if strings.HasPrefix(c.Request.URL.Path, "/customer") {
				if user.Role != "Customer" {
					return false
				}
			}

			if strings.HasPrefix(c.Request.URL.Path, "/admin") {
				if user.Role != "Staff Administrator" {
					return false
				}
			}
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			if strings.HasPrefix(c.Request.URL.Path, "/auth/login") {
				c.JSON(http.StatusBadRequest, &services.ResponseOnly{
					Success: false,
					Message: "Wrong Email or Password",
				})
				return
			}

			c.JSON(http.StatusUnauthorized, &services.ResponseOnly{
				Success: false,
				Message: "Forbidden Access",
			})
		},
		LoginResponse: func(c *gin.Context, code int, token string, time time.Time) {
			c.JSON(http.StatusOK, &services.Response{
				Success: true,
				Message: "Login Success",
				Results: struct {
					Token string `json:"token"`
				}{
					Token: token,
				},
			})
		},
	})

	if err != nil {
		return nil, err
	}
	return authMiddleware, nil
}
