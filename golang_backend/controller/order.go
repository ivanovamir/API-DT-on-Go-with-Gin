package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivanovamir/gin-test-4/config"
	"github.com/ivanovamir/gin-test-4/models"
	"gorm.io/gorm/clause"
)

func CreateOrder(c *gin.Context) {

	var order models.Order
	var products []models.Products
	var Email models.Email
	var add models.Address
	c.ShouldBindJSON(&order)
	config.DB.Preload(clause.Associations).Create(&order)
	c.JSON(http.StatusOK, gin.H{
		"data": &order})

	var logo []string
	var check []string
	var cart_image []string
	var address []string

	config.DB.Model(&Email).Pluck("Logo_image", &logo)
	config.DB.Model(&Email).Pluck("Cart_image", &check)
	config.DB.Model(&Email).Pluck("Check_image", &cart_image)
	config.DB.Model(&add).Pluck("Address", &address)

	email := order.Form.Email
	phone := order.Form.Phone
	name := order.Form.Name
	manager_phone := order.Form.Manager_phone
	company_name := order.Form.Company_name
	manager_name := order.Form.Manager_name
	inn := order.Form.Inn
	id_order := order.ID
	note := order.Note

	var products_id []int
	var products_count []int

	for x := range order.OrderProducts {
		products_count = append(products_count, order.OrderProducts[x].Count)
	}

	for x := range order.OrderProducts {
		products_id = append(products_id, order.OrderProducts[x].ProductID)
	}

	config.DB.Select("title", "price").Where("Id IN (?)", products_id).Find(&products)
	length := len(products)
	var products_title []string
	var solo_price_array []float32
	for i := 0; i <= length-1; i++ {
		x := products[i].Title
		z := products[i].Price
		products_title = append(products_title, x)
		solo_price_array = append(solo_price_array, z)
	}

	if order.Form.Company_name == "" {
		SendEmailPhyz(email, name, phone, note, products_title, address, solo_price_array, products_count, logo, check, cart_image, id_order)
	} else {
		SendEmailJurik(inn, email, note, manager_name, manager_phone, company_name, products_title, address, solo_price_array, products_count, logo, check, cart_image, id_order)
	}
	var email_list_db models.Email_list_to_send
	if err := config.DB.Model(&models.Email_list_to_send{}).Where("Email = ?", email).First(&email_list_db).Error; err != nil {
		email_list := models.Email_list_to_send{Email: email}
		config.DB.Create(&email_list)
	}
}
