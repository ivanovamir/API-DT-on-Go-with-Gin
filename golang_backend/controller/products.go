package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ivanovamir/gin-test-4/config"
	"github.com/ivanovamir/gin-test-4/models"
	ua "github.com/mileusna/useragent"
)

const (
	domain = "https://deshevle-tut.ru/data/photos/"
)

type Product_to_send struct {
	ID                string              `json:"ID"`
	Category          string              `json:"category"`
	Spare_parts_unit  string              `json:"spare_parts_unit"`
	Title             string              `json:"title"`
	Vendor_code       string              `json:"vendor_code"`
	Quantity          string              `json:"quantity"`
	Description       string              `json:"description"`
	Short_description string              `json:"short_description"`
	Price             string              `json:"price"`
	Image_original    string              `json:"image_original"`
	Image_128         string              `json:"image_128"`
	Image_432         string              `json:"image_432"`
	Features          []map[string]string `json:"features"`
}

func GetProductById(c *gin.Context) {

	Product := models.Products{}
	var Products []models.Products

	rand.Seed(time.Now().Unix())

	config.DB.Where("id=?", c.Query("prod_id")).Select("ID", "Category", "Title", "Spare_parts_unitsRefer", "Vendor_code", "Description", "Short_description", "Price", "Image_original", "Image_128", "Image_432").First(&Product)

	var id_s_true []int
	config.DB.Model(&Products).Where("category = ?", &Product.CategoriesRefer).Not("ID = (?)", &Product.ID).Pluck("ID", &id_s_true)

	var id_array []int

	for i := 0; i <= 4; i++ {
		randomIndex := rand.Intn(len(id_s_true))
		pick := id_s_true[randomIndex]
		id_array = append(id_array, pick)
	}
	config.DB.Select("ID", "Category", "Title", "Spare_parts_unitsRefer", "Vendor_code", "Description", "Short_description", "Price", "Image_original", "Image_128", "Image_432").Find(&Products, id_array)

	///////////Optomizate///////////
	var Product_features []models.Product_features

	var Spare_parts_units_feature []models.Spare_parts_units_feature

	var values_array, title_array, uint_array []string

	var id_spare_parts_units_feature []int

	config.DB.Model(&Product).Association("Product_features").Find(&Product_features)

	for x := range Product_features {
		values_array = append(values_array, Product_features[x].Value)
		id_spare_parts_units_feature = append(id_spare_parts_units_feature, Product_features[x].Spare_parts_units_featureRefer)
	}

	config.DB.Find(&Spare_parts_units_feature, id_spare_parts_units_feature)

	for x := range Spare_parts_units_feature {
		title_array = append(title_array, Spare_parts_units_feature[x].Feature_name)
		uint_array = append(uint_array, Spare_parts_units_feature[x].Unit)
	}

	for x := range Spare_parts_units_feature {
		title_array = append(title_array, Spare_parts_units_feature[x].Feature_name)
		uint_array = append(uint_array, Spare_parts_units_feature[x].Unit)
	}
	res_features := []map[string]string{}
	for x := range values_array {
		raw_map := map[string]string{
			"name_feature": title_array[x],
			"value":        values_array[x],
			"unit":         uint_array[x],
		}
		res_features = append(res_features, raw_map)
	}

	///////////Optomizate///////////

	var res_relatives []map[string]string
	for a := range Products {

		result_2 := map[string]string{

			"ID":                fmt.Sprint(Products[a].ID),
			"category":          fmt.Sprint(Products[a].CategoriesRefer),
			"spare_parts_unit":  fmt.Sprint(Products[a].Spare_parts_unitsRefer),
			"title":             Products[a].Title,
			"vendor_code":       Products[a].Vendor_code,
			"quantity":          "0",
			"description":       Products[a].Description,
			"short_description": Products[a].Short_description,
			"price":             fmt.Sprint(Products[a].Price),
			"image_original":    domain + Products[a].Image_original,
			"image_128":         domain + Products[a].Image_128,
			"image_432":         domain + Products[a].Image_432,
		}
		res_relatives = append(res_relatives, result_2)
	}

	prod_to_send := Product_to_send{
		fmt.Sprint(Product.ID),
		fmt.Sprint(Product.CategoriesRefer),
		fmt.Sprint(Product.Spare_parts_unitsRefer),
		Product.Title,
		Product.Vendor_code,
		"0",
		Product.Description,
		Product.Short_description,
		fmt.Sprint(Product.Price),
		domain + Product.Image_original,
		domain + Product.Image_128,
		domain + Product.Image_432,
		res_features,
	}

	ua := ua.Parse(c.Request.UserAgent())

	if ua.Bot {

		var count int64

		config.DB.Model(&models.Feedback{}).Where("product_id = ?", &Product.ID).Count(&count)

		var categories []models.Categories

		config.DB.Select("ID", "Title").Find(&categories)

		var price_lists []models.Price_list

		config.DB.Find(&price_lists)

		var social_links []models.Links

		config.DB.Find(&social_links)

		c.HTML(http.StatusOK, "product_detail.html", gin.H{
			"product_content":    []models.Products{Product},
			"products_content":   Products,
			"comments_count":     count,
			"categories_content": categories,
			"price_lists":        price_lists,
			"social_links":       social_links,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":      []Product_to_send{prod_to_send},
			"relatives": &res_relatives,
		})
	}
}

func GetAllIproducts(q *models.Products, pagination *models.Params) []models.Products {
	var prod []models.Products
	offset := (pagination.Page - 1) * pagination.Limit
	config.DB.Model(&models.Products{}).Where(q).Where("category IN (?)", pagination.Cat_id).Where("spare_parts_unit IN (?)", pagination.Spare_parts_unit).Where("can_to_view", true).Select("ID", "Category", "Title", "Spare_parts_unitsRefer", "Vendor_code", "Description", "Short_description", "Price", "Image_original", "Image_128", "Image_432").Limit(pagination.Limit).Offset(offset).Find(&prod)
	return prod
}

func GetAllProducts_by_multi_params(c *gin.Context) {
	pagination := GenerateMultiParams(c)
	var prod models.Products
	ProdLists := GetAllIproducts(&prod, &pagination)
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
