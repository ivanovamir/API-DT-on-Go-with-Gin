package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivanovamir/gin-test-4/config"
	"github.com/ivanovamir/gin-test-4/models"
	// ua "github.com/mileusna/useragent"
)

func GetAllCategories(c *gin.Context) {

	// ua := ua.Parse(c.Request.UserAgent())
	// fmt.Print(ua)

	var Categories []models.Categories
	emptySlice := []string{}
	if err := config.DB.Find(&Categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": &emptySlice})
	} else {
		var res_already []map[string]string
		for a := range Categories {
			result := map[string]string{
				"ID":         fmt.Sprint(Categories[a].ID),
				"title":      Categories[a].Title,
				"countitems": "0",
			}
			res_already = append(res_already, result)
		}
		c.JSON(http.StatusOK, gin.H{"data": &res_already})
	}
}

func GetCategoriesById(c *gin.Context) {

	var cat models.Categories
	emptySlice := []string{}
	if err := config.DB.Where("id=?", c.Query("cat_id")).First(&cat).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": &emptySlice})
	} else {
		var res_already []map[string]string
		result := map[string]string{
			"ID":         fmt.Sprint(cat.ID),
			"title":      cat.Title,
			"countitems": "0",
		}
		res_already = append(res_already, result)
		c.JSON(http.StatusOK, gin.H{"data": &res_already})
	}
}

func GetAllSpare_parts_units(c *gin.Context) {
	var Spare_parts_units []models.Spare_parts_units
	emptySlice := []string{}
	if err := config.DB.Find(&Spare_parts_units).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": &emptySlice})
	} else {
		var res_already []map[string]string
		for a := range Spare_parts_units {
			result := map[string]string{
				"ID":         fmt.Sprint(Spare_parts_units[a].ID),
				"title":      Spare_parts_units[a].Title,
				"countitems": "0",
			}
			res_already = append(res_already, result)
		}
		c.JSON(http.StatusOK, gin.H{"data": &res_already})

	}
}

func GetSpare_parts_unitsById(c *gin.Context) {
	var Spare_parts_units models.Spare_parts_units
	emptySlice := []string{}
	if err := config.DB.Where("id=?", c.Query("spare_parts_unit_id")).First(&Spare_parts_units).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": &emptySlice})
	} else {
		var res_already []map[string]string
		result := map[string]string{
			"ID":         fmt.Sprint(Spare_parts_units.ID),
			"title":      Spare_parts_units.Title,
			"countitems": "0",
		}
		res_already = append(res_already, result)
		c.JSON(http.StatusOK, gin.H{"data": &res_already})
	}

}
