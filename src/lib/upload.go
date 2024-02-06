package lib

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Upload(c *gin.Context, field string, dest string) *string {
	file, _ := c.FormFile(field)

	fileExt := map[string]string{
		"image/jpg":  ".jpg",
		"image/jpeg": ".jpeg",
		"image/png":  ".png",
	}

	fileType := file.Header["Content-Type"][0]
	log.Println(file.Header["Content-Type"][0])

	fileName := fmt.Sprintf("%v%v", uuid.NewString(), fileExt[fileType])
	fileDest := fmt.Sprintf("uploads/%v/%v", dest, fileName)

	c.SaveUploadedFile(file, fileDest)
	return &fileName
}
