package main

import (
	productController "github.com/d3sc/go-rest-api/controllers/productController"
	"github.com/d3sc/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// inisialisasi route
	r := gin.Default()

	// mengkoneksikan ke database
	models.ConnectDatabase()

	// mengatur route
	r.GET("/api/products", productController.Index)
	r.GET("/api/product/:id", productController.Show)
	r.POST("/api/product", productController.Create)
	r.PUT("/api/product/:id", productController.Update)
	r.DELETE("/api/product", productController.Delete)

	// menjalankan route
	r.Run()
}
