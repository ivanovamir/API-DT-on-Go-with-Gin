package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivanovamir/gin-test-4/config"
	"github.com/ivanovamir/gin-test-4/models"
)

func GetPrice_List(c *gin.Context) {
	var Price_list []models.Price_list
	if err := config.DB.Find(&Price_list).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		for x := range Price_list {
			Price_list[x].Price_1 = "https://deshevle-tut.ru/media/" + Price_list[x].Price_1
			Price_list[x].Price_2 = "https://deshevle-tut.ru/media/" + Price_list[x].Price_2
		}
		c.JSON(http.StatusOK, gin.H{"data": &Price_list})
	}
}
