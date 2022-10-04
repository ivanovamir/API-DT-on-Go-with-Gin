package models

import (
	"gorm.io/gorm"
)

type OrderProducts struct {
	gorm.Model
	Order     Order
	OrderID   uint `json:"order_id"`
	Product   Products
	ProductID int `json:"product_id"`
	Count     int `json:"count"`
}

type Order struct {
	gorm.Model
	OrderProducts []OrderProducts `gorm:"foreignKey:OrderID" json:"products"`
	User_id       string          `json:"user_id"`
	Form          Form            `json:"form"`
	FormID        int
	Note          string `json:"note"`
}

type Form struct {
	gorm.Model
	///////////////UR////////////////
	Company_name  string `json:"company_name"`
	Manager_phone string `json:"manager_phone"`
	Manager_name  string `json:"manager_name"`
	Inn           string `json:"inn"`
	//////////////PHIZ///////////////
	Name  string `json:"name"`
	Phone string `json:"phone"`
	//////////////SAME///////////////
	Email string `json:"email" gorm:"column:email"`
}
