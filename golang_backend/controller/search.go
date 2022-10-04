package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivanovamir/gin-test-4/config"
	"github.com/ivanovamir/gin-test-4/models"
)

func GetAllResultSearchByParams(q *models.Products, pagination *models.Params) []models.Products {
	var prod []models.Products
	offset := (pagination.Page - 1) * pagination.Limit
	config.DB.Debug().Model(&models.Products{}).Where("title LIKE ?", "%"+pagination.Search+"%").Or("vendor_code LIKE ?", "%"+pagination.Search+"%").Select("ID", "Category", "Title", "Spare_parts_unitsRefer", "Vendor_code", "Description", "Short_description", "Price", "Image_original", "Image_128", "Image_432").Limit(pagination.Limit).Offset(offset).Find(&prod)
	return prod
}

func GetAllResultSearch(c *gin.Context) {
	pagination := GenerateMultiParams(c)
	var prod models.Products
	ProdLists := GetAllResultSearchByParams(&prod, &pagination)
	var res_already []map[string]string
	for a := range ProdLists {
		result := map[string]string{
			"ID":                fmt.Sprint(ProdLists[a].ID),
			"category":          fmt.Sprint(ProdLists[a].CategoriesRefer),
			"spare_parts_unit":  fmt.Sprint(ProdLists[a].Spare_parts_unitsRefer),
			"title":             ProdLists[a].Title,
			"vendor_code":       ProdLists[a].Vendor_code,
			"quantity":          "0",
			"description":       ProdLists[a].Description,
			"short_description": ProdLists[a].Short_description,
			"price":             fmt.Sprint(ProdLists[a].Price),
			"image_original":    domain + ProdLists[a].Image_original,
			"image_128":         domain + ProdLists[a].Image_128,
			"image_432":         domain + ProdLists[a].Image_432,
		}
		res_already = append(res_already, result)
	}
	emptySlice := []string{}
	if res_already == nil {
		c.JSON(http.StatusOK, gin.H{"data": &emptySlice})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": &res_already})
	}

}
