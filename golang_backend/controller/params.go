package controller

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ivanovamir/gin-test-4/config"
	"github.com/ivanovamir/gin-test-4/models"
)

func GenerateMultiParams(c *gin.Context) models.Params {

	var spare_parts_unit_ids []models.Spare_parts_units

	var id_s_spare_parts_unit_ids_there []int
	config.DB.Model(&spare_parts_unit_ids).Pluck("ID", &id_s_spare_parts_unit_ids_there)

	var categories_ids []models.Categories

	var categories_ids_ids_there []int
	config.DB.Model(&categories_ids).Pluck("ID", &categories_ids_ids_there)

	limit := 0
	page := 1
	search := ""
	var cat_id []int
	var spare_parts_unit []int
	spare_parts_unit = id_s_spare_parts_unit_ids_there
	cat_id = categories_ids_ids_there
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
		case "cat_id":
			cat_id = toIntArray(queryValue)
		case "spare_parts_unit":
			spare_parts_unit = toIntArray(queryValue)
		case "s":
			search = queryValue
		}

	}
	return models.Params{
		Limit:            limit,
		Page:             page,
		Cat_id:           cat_id,
		Spare_parts_unit: spare_parts_unit,
		Search:           search,
	}

}
func toIntArray(str string) []int {
	chunks := strings.Split(str, ",")

	var res []int
	for _, c := range chunks {
		i, _ := strconv.Atoi(c) // error handling ommitted for concision
		res = append(res, i)
	}

	return res
}
