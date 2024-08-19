package main

import (
	"gin-fleamarket/controllers"
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// modelsパッケージ内のItem構造体のスライスを作成
	items := []models.Item{
		{ID: 1, Name: "item1", Price: 100, Description: "This is item1", SoldOut: false},
		{ID: 2, Name: "item2", Price: 200, Description: "This is item2", SoldOut: false},
		{ID: 3, Name: "item3", Price: 300, Description: "This is item3", SoldOut: false},
	}

	itemRepository := repositories.NewItemMemoryRepository(items)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	r := gin.Default()
	r.GET("/items", itemController.FindAll)
	r.GET("/items/:id", itemController.FindById)
	r.Run("localhost:8080")
}
