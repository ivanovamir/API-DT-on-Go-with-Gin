package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivanovamir/gin-test-4/config"
	"github.com/ivanovamir/gin-test-4/models"
)

func CreateFeedback(c *gin.Context) {
	var feedback models.Feedback
	c.BindJSON(&feedback)
	config.DB.Create(&feedback)
	c.JSON(http.StatusOK, gin.H{
		"data": &feedback})
}

func GetFeedbackByProduct_id(c *gin.Context) {
	Feedbacks := []models.Feedback{}
	if err := config.DB.Where("product_id", c.Query("product_id")).Find(&Feedbacks).Error; err != nil {
	} else {
		c.JSON(http.StatusOK, gin.H{"data": &Feedbacks})
	}
}
