package models

type Email_list_to_send struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	Email       string `json:"email"`
	Can_to_send bool   `json:"can_to_send" gorm:"default:true"`
}
