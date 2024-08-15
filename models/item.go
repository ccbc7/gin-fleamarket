package models

type Item struct {
	// uintは符号なし整数型
	ID          uint
	Name        string
	Price       uint
	Description string
	SoldOut     bool
}
