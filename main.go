package main

// * Template file for testing API. NOT EDIT

import (
	wwf "api/wwf/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api := r.Group("/api/wwf")
	wwf.RegisterRoutes(api)

	r.Run(":3000")
}
