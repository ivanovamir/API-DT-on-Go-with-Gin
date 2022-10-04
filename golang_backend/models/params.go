package models

type Params struct {
	Limit                     int    `json:"limit"`
	Page                      int    `json:"page"`
	Cat_id                    []int  `json:"cat_id" gorm:"type:int"`
	Spare_parts_unit          []int  `json:"spare_parts_unit" gorm:"type:int"`
	Search                    string `json:"search"`
	Spare_parts_units_feature []int  `json:"spare_parts_units_feature" gorm:"type:int"`
}
