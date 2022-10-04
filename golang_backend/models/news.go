package models

import "time"

type News struct {
	Id        uint      `json:"id" gorm:"primary_key, AUTO_INCREMENT"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
}
