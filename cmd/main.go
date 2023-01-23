package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	mr := manage.Routes(router.Group)

	router.Run()
}
