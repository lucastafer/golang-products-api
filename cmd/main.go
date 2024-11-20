package main

import (
	"golang-api/controller"
	"golang-api/db"
	"golang-api/repository"
	"golang-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init server HTTP, creating a new instance
	server := gin.Default()

	// DB Connection
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	// UseCases
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	// Controllers
	ProductController := controller.NewProductController(ProductUseCase)

	// Routes
	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductById)
	server.PUT("/product/:productId", ProductController.UpdateProduct)
	server.DELETE("/product/:productId", ProductController.DeleteProduct)

	// Run server, receiving port as param
	server.Run(":8000")
}
