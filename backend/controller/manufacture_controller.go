package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosshi/gym-machine-management/backend/model"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func CreateManufacturer(c *gin.Context) {
	var manufacturer model.Manufacturer

	// JSON をパース
	if err := c.ShouldBindJSON(&manufacturer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON",
		})
		return
	}

	// ID を生成
	id, err := gonanoid.New()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate ID",
		})
		return
	}
	manufacturer.ID = id

	// DB に保存（例: config.DB.Create(&manufacturer)）
	// if err := config.DB.Create(&manufacturer).Error; err != nil {
	//     c.JSON(http.StatusInternalServerError, gin.H{
	//         "error": "Failed to create manufacturer",
	//     })
	//     return
	// }

	// 成功レスポンス
	c.JSON(http.StatusCreated, gin.H{
		"id": manufacturer.ID,
	})
}
