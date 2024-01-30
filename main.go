package main

import (
	"github.com/DzulfiqarSiraj/go-backend/src/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.Combine(r)
	r.Run()
}
