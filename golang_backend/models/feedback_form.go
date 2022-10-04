package models

import "gorm.io/gorm"

type Feedback_form struct {
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email"`
	Topic   string `json:"topic"`
	Message string `json:"message"`
}
