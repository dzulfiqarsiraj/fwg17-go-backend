package lib

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Upload(c *gin.Context, field string, dest string) *string {
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
		return &invalidType
	}

	if file.Size > 1048576 {
		return &invalidSize
	}

	fmt.Println(file.Header)
	fmt.Println(fileType)
	fmt.Println(file.Size)

	fileName := fmt.Sprintf("%v%v", uuid.NewString(), fileExt[fileType])
	fileDest := fmt.Sprintf("uploads/%v/%v", dest, fileName)

	c.SaveUploadedFile(file, fileDest)
	return &fileName
}
