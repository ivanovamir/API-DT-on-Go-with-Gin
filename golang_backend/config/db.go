package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ivanovamir/gin-test-4/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := gorm.Open(postgres.Open(fmt.Sprintf("postgres://%s:%s@%s/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Silent),
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
