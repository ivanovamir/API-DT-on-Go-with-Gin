package config

import (
	"github.com/ivanovamir/gin-test-4/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	db, err := gorm.Open(postgres.Open("postgres://postgres:vdnjvjnjhgDFGHJXCV56483754343@178.21.8.81/shop"), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		// Logger:                 logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Spare_parts_units{})
	db.AutoMigrate(&models.Spare_parts_units_feature{})
	db.AutoMigrate(&models.Feature{})
	db.AutoMigrate(&models.MiniSlider{}, &models.Categories{}, &models.Products{}, &models.Params{}, &models.Feedback{}, &models.Form{}, &models.Feedback_form{}, &models.News{}, &models.Slider{},
		&models.Slider{}, &models.Links{}, &models.Site_link{}, &models.Address{}, &models.Price_list{}, &models.Email{}, &models.Email_list_to_send{}, &models.Order{},
		&models.OrderProducts{}, &models.Services{}, &models.About{}, &models.Spare_parts_units{}, &models.Feature_validator{},
		&models.Product_features{}, &models.Metricts{},
	)

	DB = db
}
