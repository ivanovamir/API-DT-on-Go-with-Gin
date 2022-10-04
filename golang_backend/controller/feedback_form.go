package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivanovamir/gin-test-4/config"
	"github.com/ivanovamir/gin-test-4/models"
)

func FeedbackFormPost(c *gin.Context) {
	var Feedback models.Feedback_form
	c.ShouldBindJSON(&Feedback)
	config.DB.Create(&Feedback)
	c.JSON(http.StatusOK, gin.H{
		"data": &Feedback})
}
