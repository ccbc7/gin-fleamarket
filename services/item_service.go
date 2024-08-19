package services

import (
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
)

// インターフェースを定義
type IItemService interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
}

// 構造体を定義
type ItemService struct {
	repository repositories.IItemRepository
}

// コンストラクタを定義
func NewItemService(repository repositories.IItemRepository) IItemService {
	return &ItemService{repository: repository}
}

// メソッドを定義
func (s *ItemService) FindAll() (*[]models.Item, error) {
	return s.repository.FindAll()
}

func (s *ItemService) FindById(itemId uint) (*models.Item, error) {
	return s.repository.FindById(itemId)
}
