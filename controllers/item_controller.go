package controllers

import (
	"net/http"
	"strconv"

	"gin-fleamarket/services"

	"github.com/gin-gonic/gin"
)

type ItemController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
}

type itemController struct {
	service services.IItemService
}

func NewItemController(service services.IItemService) ItemController {
	return &itemController{service: service}
}

// サービスのFindAll()メソッドを呼び出し、結果をJSON形式で返す
func (c *itemController) FindAll(ctx *gin.Context) {
	items, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": items})
}

// リクエストパラメータからIDを取得し、サービスのFindById()メソッドを呼び出し、結果をJSON形式で返す
func (c *itemController) FindById(ctx *gin.Context) {
	// strconv.ParseUint()文字列を整数に変換, 10進数, 64ビット
	itemId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	item, err := c.service.FindById(uint(itemId))
	if err != nil {
		if err.Error() == "item not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": item})
}
