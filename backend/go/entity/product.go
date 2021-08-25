package entity

type Product struct {
	ID          uint   `json:"id" gorm:"primaryKey,autoIncrement, not null"`
	SKU         string `json:"sku" gorm:"unique"`
	Name        string `json:"name"`
	Brand       string `json:"brand"`
	Model       string `json:"model"`
	Price       string `json:"price"`
	Description string `json:"description"`
}
