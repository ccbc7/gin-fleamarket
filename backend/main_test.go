package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"gin-fleamarket/infra"
	"gin-fleamarket/models"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("../.env.test"); err != nil {
		log.Println("Error loading .env.test file:", err) // 変更
	}

	code := m.Run()

	os.Exit(code)
}

func setupTestData(db *gorm.DB) {
	items := []models.Item{
		{Name: "テストアイテム1", Price: 1000, Description: "", SoldOut: false, UserID: 1},
		{Name: "テストアイテム2", Price: 1000, Description: "テスト２", SoldOut: false, UserID: 1},
		{Name: "テストアイテム3", Price: 1000, Description: "テスト３", SoldOut: false, UserID: 2},
	}

	users := []models.User{
		{Email: "test1@example.com", Password: "test1pass"},
		{Email: "test2@example.com", Password: "test2pass"},
	}

	for _, user := range users {
		db.Create(&user)
	}

	for _, item := range items {
		db.Create(&item)
	}
}

func setup() *gin.Engine {
	db := infra.SetupDB()
	db.AutoMigrate(&models.User{}, &models.Item{})

	setupTestData(db)
	router := setupRouter(db)

	return router
}

func TestFindAll(t *testing.T) {
	// テスト用のデータをセットアップ
	router := setup()

	//HTTPレスポンスを記録するオブジェクトを作成
	w := httptest.NewRecorder()

	// NewRequestを使ってリクエストを作成
	req, _ := http.NewRequest("GET", "/items", nil)

	// ServeHTTPメソッドを使ってリクエストを実行
	router.ServeHTTP(w, req)

	// resの型を定義しているだけで、中身は空のmap
	var res map[string][]models.Item

	// レスポンスのボディをJSON形式からGoのデータ構造に変換
	json.Unmarshal(w.Body.Bytes(), &res)

	// レスポンスのステータスコードが200であることを確認
	assert.Equal(t, http.StatusOK, w.Code)

	// レスポンスのボディに含まれるデータの数が3であることを確認
	assert.Equal(t, 3, len(res["data"]))
	fmt.Println(res)
}
