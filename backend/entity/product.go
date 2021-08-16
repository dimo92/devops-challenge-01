package entity

type Product struct {
	SKU         uint64 `gorm:"primaryKey" json:"SKU"`
	Name        string `json:"Name"`
	Brand       string `json:"Brand"`
	Model       string `json:"Model"`
	Price       string `json:"Price"`
	Description string `json:"Description"`
}
