package main

import (
	"gin-fleamarket/controllers"
	"gin-fleamarket/infra"
	"gin-fleamarket/middlewares"

	// "gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"

	"github.com/gin-contrib/cors"
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

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	// ルーターの作成
	r := gin.Default()

	// CORSの設定
	r.Use(cors.Default())

	// ルーティンググループの作成
	itemRouter := r.Group("/items")
	itemRouterWithAuth := r.Group("/items", middlewares.AuthMiddleware(authService))
	authRouter := r.Group("/auth")

	// ルーティングの設定
	itemRouter.GET("", itemController.FindAll)

	itemRouterWithAuth.GET("/:id", itemController.FindById)
	itemRouterWithAuth.POST("", itemController.Create)
	itemRouterWithAuth.PUT("/:id", itemController.Update)
	itemRouterWithAuth.DELETE("/:id", itemController.Delete)

	authRouter.POST("/signup", authController.SignUp)
	authRouter.POST("/login", authController.Login)

	r.Run(":8080")
}
