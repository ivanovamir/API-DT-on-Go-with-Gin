package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/ivanovamir/gin-test-4/config"
	"github.com/ivanovamir/gin-test-4/models"
	"github.com/joho/godotenv"
)

const (
	smtpHost = "smtp.gmail.com"
	smtpPort = "587"
	address  = smtpHost + ":" + smtpPort
)

func EmailStatusUpdate(c *gin.Context) {
	QueryArray := c.Request.URL.Query()
	email := QueryArray["email"]
	config.DB.Model([]models.Email_list_to_send{}).Where("Email = ?", email).Update("Can_to_send", false)
}

func SendEmailPhyz(email, name, phone, note string, products_title, address []string, solo_price_array []float32, products_count []int, logo, check, cart_image []string, id_order uint) {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	// Sender data.
	from := os.Getenv("Email")
	password := os.Getenv("App_Password")

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, err := template.ParseFiles("../templates/index_fizik.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	var message []string
	var cena []float32
	var kolvo []int
	var result float32
	var itogo float32

	for i := 0; i <= len(products_title)-1; i++ {
		message = append(message, products_title[i])
		kolvo = append(kolvo, products_count[i])
		cena = append(cena, solo_price_array[i]*float32(products_count[i]))
		itogo += solo_price_array[i] * float32(products_count[i])
	}
	result += itogo

	logo_1 := "https://deshevle-tut.ru/media/" + logo[0]
	check_1 := "https://deshevle-tut.ru/media/" + check[0]
	cart_image_1 := "https://deshevle-tut.ru/media/" + cart_image[0]

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Ваш заказ с сайта «Дешевле Тут» \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Id          uint
		Name        string
		Phone       string
		Email       string
		Products    []string
		Kolvo       []int
		Cena        []float32
		Result      float32
		Note        string
		Logo_image  string
		Check_image string
		Cart_image  string
		Address     string
	}{
		Id:          id_order,
		Name:        name,
		Phone:       phone,
		Email:       email,
		Products:    message,
		Kolvo:       kolvo,
		Cena:        cena,
		Result:      result,
		Note:        note,
		Logo_image:  logo_1,
		Check_image: check_1,
		Cart_image:  cart_image_1,
		Address:     address[0],
	})

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}

}

func SendEmailJurik(inn, email, note, manager_name, manager_phone, company_name string, products_title, address []string, solo_price_array []float32, products_count []int, logo, check, cart_image []string, id_order uint) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	// Sender data.
	from := os.Getenv("Email")
	password := os.Getenv("App_Password")

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, err := template.ParseFiles("../templates/index_jurik.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	var message []string
	var cena []float32
	var kolvo []int
	var result float32
	var itogo float32

	for i := 0; i <= len(products_title)-1; i++ {
		message = append(message, products_title[i])
		kolvo = append(kolvo, products_count[i])
		cena = append(cena, solo_price_array[i]*float32(products_count[i]))
		itogo += solo_price_array[i] * float32(products_count[i])
	}
	result += itogo

	logo_1 := "https://deshevle-tut.ru/media/" + logo[0]
	check_1 := "https://deshevle-tut.ru/media/" + check[0]
	cart_image_1 := "https://deshevle-tut.ru/media/" + cart_image[0]

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Ваш заказ с сайта «Дешевле Тут» \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Id            uint
		Inn           string
		Manager_name  string
		Manager_phone string
		Company_name  string
		Email         string
		Products      []string
		Kolvo         []int
		Cena          []float32
		Result        float32
		Note          string
		Logo_image    string
		Check_image   string
		Cart_image    string
		Address       string
	}{
		Id:            id_order,
		Inn:           inn,
		Manager_name:  manager_name,
		Manager_phone: manager_phone,
		Company_name:  company_name,
		Email:         email,
		Products:      message,
		Kolvo:         kolvo,
		Cena:          cena,
		Result:        result,
		Note:          note,
		Logo_image:    logo_1,
		Check_image:   check_1,
		Cart_image:    cart_image_1,
		Address:       address[0],
	})

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
}

type Emails struct {
	Key    string   `json:"key"`
	Emails []int    `json:"emails"`
	Body   []string `json:"body"`
}

func Email_list(c *gin.Context) {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	// Sender data.
	bodytext, error := ioutil.ReadAll(c.Request.Body)
	if error != nil {
		panic(error)
	}
	var k Emails
	error = json.Unmarshal(bodytext, &k)
	if error != nil {
		panic(error)
	}

	fmt.Print(k.Emails)
	if os.Getenv("Key") == k.Key {
		from := os.Getenv("Email")
		password := os.Getenv("App_Password")

		var emails_array []string

		config.DB.Model(&models.Email_list_to_send{}).Where("Can_to_send = ?", true).Pluck("Email", &emails_array)

		// Receiver email address.
		to := emails_array
		// smtp server configuration.
		smtpHost := "smtp.gmail.com"
		smtpPort := "587"

		// Authentication.
		auth := smtp.PlainAuth("", from, password, smtpHost)

		subject := "Subject: Дешевле ТУТ"
		// body := body_text.Body

		message := []byte(subject + k.Body[0])

		// Sending email.
		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
