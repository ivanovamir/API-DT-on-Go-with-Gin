package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivanovamir/gin-test-4/config"
	"github.com/ivanovamir/gin-test-4/models"
)

func GetServiceById(c *gin.Context) {
	var service models.Services
	if err := config.DB.Where("id=?", c.Query("service_id")).First(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		service.Image = "https://deshevle-tut.ru/media/" + service.Image
		c.JSON(http.StatusOK, gin.H{
			"data": []models.Services{service}})
	}
}

func GetAllServices(q *models.Services, pagination *models.Params) (*[]models.Services, error) {
	var services []models.Services
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := config.DB.Limit(pagination.Limit).Offset(offset)
	result := queryBuider.Model(&models.Services{}).Where(q).Find(&services)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	for x := range services {
		services[x].Image = "https://deshevle-tut.ru/media/" + services[x].Image
	}
	return &services, nil
}

func GetAllServices_by_params(c *gin.Context) {
	pagination := GenerateMultiParams(c)
	var services models.Services
	ServicesList, err := GetAllServices(&services, &pagination)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"data": ServicesList,
	})

}
