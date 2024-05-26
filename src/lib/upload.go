package lib

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func CloudinaryConfig() *cloudinary.Cloudinary {
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	cloudAPI := os.Getenv("CLOUDINARY_API_KEY")
	cloudAPISecret := os.Getenv("CLOUDINARY_API_SECRET")

	cloud, _ := cloudinary.NewFromParams(cloudName, cloudAPI, cloudAPISecret)
	return cloud
}

func Upload(c *gin.Context, field string, dest string) (*string, error) {
	godotenv.Load()

	cloud := CloudinaryConfig()

	var ctx = context.Background()

	file, _ := c.FormFile(field)
	invalidType := "Invalid File Type"
	invalidSize := "Invalid File Size"

	fileExt := map[string]string{
		"image/jpg":  ".jpg",
		"image/jpeg": ".jpeg",
		"image/png":  ".png",
	}

	fileType := file.Header["Content-Type"][0]

	_, typeExists := fileExt[fileType]

	if !typeExists {
		return &invalidType, c.Err()
	}

	if file.Size > 1048576 {
		return &invalidSize, c.Err()
	}

	fileName := fmt.Sprintf("%v", uuid.NewString())
	fileDestCloudinary := fmt.Sprintf("cov-shop/%v/%v", dest, fileName)

	// Upload file to Cloudinary
	resp, err := cloud.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: fileDestCloudinary})
	if err != nil {
		fmt.Println("error")
	}

	return &resp.SecureURL, c.Err()
}
