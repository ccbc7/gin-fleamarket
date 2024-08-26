package repositories

import (
	"errors"

	"gin-fleamarket/models"

	"gorm.io/gorm"
)

// インターフェースを定義
type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
	Update(updatedItem models.Item) (*models.Item, error)
	Delete(itemId uint) error
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

func (r *ItemMemoryRepository) Create(newItem models.Item) (*models.Item, error) {
	newItem.ID = uint(len(r.items) + 1)
	r.items = append(r.items, newItem)
	return &newItem, nil
}

// 更新
func (r *ItemMemoryRepository) Update(updatedItem models.Item) (*models.Item, error) {
	for i, v := range r.items {
		if v.ID == updatedItem.ID {
			r.items[i] = updatedItem
			return &r.items[i], nil
		}
	}
	return nil, errors.New("unexpected error")
}

// 削除
func (r *ItemMemoryRepository) Delete(itemId uint) error {
	// i, v := range r.items でスライスのインデックスと要素を取得
	for i, v := range r.items {
		if v.ID == itemId {
			// append()関数でスライスの要素を削除 ...演算子でスライスの要素を展開
			r.items = append(r.items[:i], r.items[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}

/*
* DBを使う場合のリポジトリの実装
 */
type ItemDBRepository struct {
	db *gorm.DB
}

// 作成
func (r *ItemDBRepository) Create(newItem models.Item) (*models.Item, error) {
	result := r.db.Create(&newItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newItem, nil
}

// Delete implements IItemRepository.
func (i *ItemDBRepository) Delete(itemId uint) error {
	panic("unimplemented")
}

// FindAll implements IItemRepository.
func (i *ItemDBRepository) FindAll() (*[]models.Item, error) {
	panic("unimplemented")
}

// FindById implements IItemRepository.
func (i *ItemDBRepository) FindById(itemId uint) (*models.Item, error) {
	panic("unimplemented")
}

// Update implements IItemRepository.
func (i *ItemDBRepository) Update(updatedItem models.Item) (*models.Item, error) {
	panic("unimplemented")
}

func NewItemRepository(db *gorm.DB) IItemRepository {
	return &ItemDBRepository{db: db}
}
