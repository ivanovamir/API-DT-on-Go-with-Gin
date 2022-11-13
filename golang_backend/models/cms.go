package models

type Slider struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	MainText  string `json:"main_text"`
	UpperText string `json:"upper_text"`
	DownText  string `json:"down_text"`
	Image     string `json:"image"`
}

type MiniSlider struct {
	Id         uint   `json:"id" gorm:"primaryKey"`
	UpperText  string `json:"upper_text"`
	MarkedText string `json:"marked_text"`
	MainText   string `json:"main_text"`
	Image      string `json:"image"`
}

type Links struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Link string `json:"link"`
}

type Site_link struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Link string `json:"site_link"`
}

type Address struct {
	Id      uint   `json:"id" gorm:"primaryKey"`
	Address string `json:"address"`
}

type Email struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	Logo_image  string `json:"logo_image" gorm:"column:logo_image"`
	Cart_image  string `json:"cart_image" gorm:"column:cart_image"`
	Check_image string `json:"check_image" gorm:"column:check_image"`
}

type About struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	MainText string `json:"main_text"`
	Image    string `json:"image"`
}

type Metricts struct {
	Id              uint   `json:"id" gorm:"primaryKey"`
	GoogleMetric    string `json:"google_metric" gorm:"column:google_metric"`
	YandexMetric    string `json:"yandex_metric" gorm:"column:yandex_metric"`
	TelegramBotLink string `json:"telegrambot_link" gorm:"column:telegrambot_link"`
}
