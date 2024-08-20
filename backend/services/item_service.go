package services

import (
	"gin-fleamarket/dto"
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
)

// インターフェースを定義
type IItemService interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	Create(createItemInput dto.CreateItemInput) (*models.Item, error)
	Update(itemId uint, updateItemInput dto.UpdateItemInput) (*models.Item, error)
	Delete(itemId uint) error
}

// 構造体を定義
type ItemService struct {
	repository repositories.IItemRepository
}

// コンストラクタを定義
func NewItemService(repository repositories.IItemRepository) IItemService {
	return &ItemService{repository: repository}
}

// 全ての商品を取得
func (s *ItemService) FindAll() (*[]models.Item, error) {
	return s.repository.FindAll()
}

// IDで商品を取得
func (s *ItemService) FindById(itemId uint) (*models.Item, error) {
	return s.repository.FindById(itemId)
}

// 作成
func (s *ItemService) Create(createItemInput dto.CreateItemInput) (*models.Item, error) {
	newItem := models.Item{
		Name:        createItemInput.Name,
		Price:       createItemInput.Price,
		Description: createItemInput.Description,
		SoldOut:     false,
	}
	// リポジトリ層のCreateメソッドを呼び出し、作成処理を行う
	return s.repository.Create(newItem)
}

// 更新
func (s *ItemService) Update(itemId uint, updateItemInput dto.UpdateItemInput) (*models.Item, error) {
	targetItem, err := s.FindById(itemId)
	if err != nil {
		return nil, err
	}

	if updateItemInput.Name != nil {
		targetItem.Name = *updateItemInput.Name
	}
	if updateItemInput.Price != nil {
		targetItem.Price = *updateItemInput.Price
	}
	if updateItemInput.Description != nil {
		targetItem.Description = *updateItemInput.Description
	}
	if updateItemInput.SoldOut != nil {
		targetItem.SoldOut = *updateItemInput.SoldOut
	}

	// リポジトリ層のUpdateメソッドを呼び出し、更新処理を行う
	return s.repository.Update(*targetItem)
}

// 削除
func (s *ItemService) Delete(itemId uint) error {
	return s.repository.Delete(itemId)
}
