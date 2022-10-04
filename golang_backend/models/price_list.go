package models

type Price_list struct {
	Id      uint   `json:"id" gorm:"primaryKey"`
	Price_1 string `json:"price_1"`
	Price_2 string `json:"price_2"`
}
