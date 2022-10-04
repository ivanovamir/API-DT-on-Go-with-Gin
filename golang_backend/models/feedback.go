package models

import (
	"time"
)

type Feedback struct {
	ID            uint      `json:"ID" gorm:"primaryKey"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Body          string    `json:"body"`
	FeedbackRefer int64     `json:"product_id" gorm:"column:product_id"`
	Created_At    time.Time `json:"created_at"`
}
