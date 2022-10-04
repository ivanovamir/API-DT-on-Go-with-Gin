package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivanovamir/gin-test-4/config"
	"github.com/ivanovamir/gin-test-4/models"
	ua "github.com/mileusna/useragent"
)

func GetNewById(c *gin.Context) {
	var new models.News

	if err := config.DB.Where("id=?", c.Query("new_id")).First(&new).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	ua := ua.Parse(c.Request.UserAgent())

	if ua.Bot {
		var categories []models.Categories

		config.DB.Select("ID", "Title").Find(&categories)

		var price_lists []models.Price_list

		config.DB.Find(&price_lists)

		var social_links []models.Links

		config.DB.Find(&social_links)

		var news_list []models.News

		config.DB.Not("ID = ?", c.Query("new_id")).Select("ID", "Title").Limit(5).Find(&news_list)

		c.HTML(http.StatusOK, "new_detail.html", gin.H{
			"new_content":        []models.News{new},
			"news_content":       news_list,
			"categories_content": categories,
			"price_lists":        price_lists,
			"social_links":       social_links,
		})
	} else {
		new.Image = "https://deshevle-tut.ru/media/" + new.Image
		c.JSON(http.StatusOK, gin.H{
			"data": []models.News{new}})
	}

}

func GetAllNews(q *models.News, pagination *models.Params) (*[]models.News, error) {
	var news []models.News
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := config.DB.Limit(pagination.Limit).Offset(offset)
	result := queryBuider.Model(&models.News{}).Where(q).Order("created_at DESC").Find(&news)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	for x := range news {
		news[x].Image = "https://deshevle-tut.ru/media/" + news[x].Image
	}
	return &news, nil
}

func GetAllNews_by_params(c *gin.Context) {
	pagination := GenerateMultiParams(c)
	var news models.News
	NewsLists, err := GetAllNews(&news, &pagination)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"data": NewsLists,
	})

}
