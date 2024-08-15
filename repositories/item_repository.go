package repositories

import "gin-fleamarket/models"

// インターフェースを定義
type IItemRepository interface {
	FindAll() (*[]models.Item, error)
}

// 構造体を定義 構造体とは、フィールドの集まりを定義した型
type ItemMemoryRepository struct {
	Items []models.Item
}

// コンストラクタ関数を定義 この関数は、ItemMemoryRepository構造体のポインタを返す
// ポインタとは、メモリ上のアドレスを指す変数
func NewItemMemoryRepository(items []models.Item) IItemRepository {
	return &ItemMemoryRepository{Items: items}
}

// メソッドを定義 このメソッドは、ItemMemoryRepository構造体のItemsフィールドを返す
func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.Items, nil
}
