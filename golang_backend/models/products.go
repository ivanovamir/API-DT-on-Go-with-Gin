package models

import (
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	CategoriesRefer        int64              `json:"category" gorm:"column:category"`
	Spare_parts_unitsRefer int64              `json:"spare_parts_unit" gorm:"column:spare_parts_unit"`
	Title                  string             `json:"title" gorm:"column:title"`
	Vendor_code            string             `json:"vendor_code" gorm:"column:vendor_code"`
	Description            string             `json:"description" gorm:"column:description"`
	Short_description      string             `json:"short_description" gorm:"column:short_description"`
	Price                  float32            `json:"price" gorm:"column:price"`
	Image_original         string             `json:"image_original" gorm:"column:image_original"`
	Image_128              string             `json:"image_128" gorm:"column:image_128"`
	Image_432              string             `json:"image_432" gorm:"column:image_432"`
	Feedback               []Feedback         `gorm:"foreignKey:FeedbackRefer" json:"feedbacks"`
	Can_to_view            bool               `json:"can_to_view" gorm:"default:true"`
	Product_featuresRefer  []Product_features `gorm:"foreignKey:ProductRefer"`
	Product_features       []Product_features `gorm:"many2many:features"`
	Features_products      uint               `json:"features_products" gorm:"column:features_products;unique"`
}

type Feature struct {
	ID int64 `json:"ID" gorm:"primaryKey"`
}
