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
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
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

	return r
}

func main() {
	// 初期化(環境変数の読み込み)
	infra.Initialize()

	// DB接続
	db := infra.SetupDB()

	// ルーターの設定 引数にDBを渡すことで、各レイヤー(サービス,リポジトリ,コントローラ)でDBを利用できる
	r := setupRouter(db)

	r.Run(":8080")
}

// itemRepository := repositories.NewItemMemoryRepository(items)
// items := []models.Item{
// 	{ID: 1, Name: "item1", Price: 100, Description: "This is item1", SoldOut: false},
// 	{ID: 2, Name: "item2", Price: 200, Description: "This is item2", SoldOut: false},
// 	{ID: 3, Name: "item3", Price: 300, Description: "This is item3", SoldOut: false},
// }
