package models

type Categories struct {
	ID       int64      `json:"ID" gorm:"primaryKey"`
	Title    string     `json:"title"`
	Products []Products `gorm:"foreignKey:CategoriesRefer" json:"products"`
}

type Spare_parts_units struct {
	ID                        int64                       `json:"ID" gorm:"primaryKey"`
	Title                     string                      `json:"title"`
	Spare_parts_units         []Products                  `gorm:"foreignKey:Spare_parts_unitsRefer"`
	Spare_parts_units_feature []Spare_parts_units_feature `gorm:"foreignKey:Spare_parts_unitsRefer"`
	Feature_validator         []Feature_validator         `gorm:"foreignKey:Spare_parts_unitsRefer"`
}

type Spare_parts_units_feature struct {
	ID                     int64               `json:"ID" gorm:"primaryKey"`
	Feature_name           string              `json:"title"`
	Feature_filter_name    string              `json:"filter_name"`
	Unit                   string              `json:"unit"`
	Spare_parts_unitsRefer int                 `json:"spare_parts_units" gorm:"column:spare_parts_units"`
	Feature_validator      []Feature_validator `gorm:"foreignKey:Spare_parts_units_featureRefer"`
	Product_features       []Product_features  `gorm:"foreignKey:Spare_parts_units_featureRefer"`
}

type Feature_validator struct {
	ID                             int64  `json:"ID" gorm:"primaryKey"`
	Valid_feature_value            string `json:"valid_feature_value"`
	Spare_parts_unitsRefer         int    `json:"spare_parts_units" gorm:"column:spare_parts_units"`
	Spare_parts_units_featureRefer int    `json:"spare_parts_units_feature" gorm:"column:feature_key"`
}

type Product_features struct {
	ID                             int64  `json:"ID" gorm:"primaryKey"`
	Value                          string `json:"value" gorm:"column:value"`
	ProductRefer                   int    `json:"product" gorm:"column:product"`
	Spare_parts_units_featureRefer int    `json:"spare_parts_units_feature" gorm:"column:feature"`
}
