package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivanovamir/gin-test-4/config"
	"github.com/ivanovamir/gin-test-4/models"
)

const (
	media_url = "https://deshevle-tut.ru/media/"
)

func GetAllSliders(c *gin.Context) {
	var sliders []models.Slider

	if err := config.DB.Find(&sliders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		for x := range sliders {
			sliders[x].Image = media_url + sliders[x].Image
		}
		c.JSON(http.StatusOK, gin.H{"data": sliders})
	}

}

func GetAllLinks(c *gin.Context) {
	Links := []models.Links{}
	if err := config.DB.Find(&Links).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"data": &Links})
	}
}

func GetAllSite_link(c *gin.Context) {
	Site_link := []models.Site_link{}
	if err := config.DB.Find(&Site_link).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"data": &Site_link})
	}
}

func GetAllAddress(c *gin.Context) {
	Address := []models.Address{}
	if err := config.DB.Find(&Address).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"data": &Address})
	}
}

func GetMiniSlider(c *gin.Context) {
	var slider models.MiniSlider
	if err := config.DB.First(&slider).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		slider.Image = media_url + slider.Image
		c.JSON(http.StatusOK, gin.H{"data": []models.MiniSlider{slider}})
	}
}

func GetAllAbout(c *gin.Context) {
	var about []models.About

	if err := config.DB.Find(&about).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		for x := range about {
			about[x].Image = media_url + about[x].Image
		}
		c.JSON(http.StatusOK, gin.H{"data": about})
	}
}
