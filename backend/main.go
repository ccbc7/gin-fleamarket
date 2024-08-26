package main

import (
	"gin-fleamarket/controllers"
	"gin-fleamarket/infra"
	// "gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	// modelsパッケージ内のItem構造体のスライスを作成
	db := infra.SetupDB()
	// items := []models.Item{
	// 	{ID: 1, Name: "item1", Price: 100, Description: "This is item1", SoldOut: false},
	// 	{ID: 2, Name: "item2", Price: 200, Description: "This is item2", SoldOut: false},
	// 	{ID: 3, Name: "item3", Price: 300, Description: "This is item3", SoldOut: false},
	// }

	// itemRepository := repositories.NewItemMemoryRepository(items)
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	r := gin.Default()
	itemRouter := r.Group("/items")
	itemRouter.GET("", itemController.FindAll)
	itemRouter.GET("/:id", itemController.FindById)
	itemRouter.POST("", itemController.Create)
	itemRouter.PUT("/:id", itemController.Update)
	itemRouter.DELETE("/:id", itemController.Delete)
	r.Run(":8080")
}
