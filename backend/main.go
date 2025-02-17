package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gosshi/gym-machine-management/docs" // swag によって生成される docs パッケージ

	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

func main() {
	r := gin.Default()

	// Swagger UI を /swagger/*any で提供
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API エンドポイントの例
	r.GET("/api/v1/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, world!"})
	})

	log.Println("Server started at :8080")
	r.Run(":8080")
}