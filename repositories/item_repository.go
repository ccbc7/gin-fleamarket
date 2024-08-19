package repositories

import (
	"errors"

	"gin-fleamarket/models"
)

// インターフェースを定義
type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
}

// 構造体を定義 構造体とは、フィールドの集まりを定義した型
type ItemMemoryRepository struct {
	items []models.Item
}

// コンストラクタを定義
func NewItemMemoryRepository(items []models.Item) IItemRepository {
	return &ItemMemoryRepository{items: items}
}

// メソッドを定義 このメソッドは、ItemMemoryRepository構造体のItemsフィールドを返す
func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}

// 構造体のItemsフィールドからIDを検索して返す
func (r *ItemMemoryRepository) FindById(itemId uint) (*models.Item, error) {
	for _, v := range r.items {
		if v.ID == itemId {
			return &v, nil
		}
	}
	return nil, errors.New("item not found")
}
