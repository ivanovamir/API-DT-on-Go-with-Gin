package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivanovamir/gin-test-4/controller"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Routes(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	spare_parts_unit := router.Group("/spare_parts_unit")
	{
		spare_parts_unit.GET("/", controller.GetAllSpare_parts_units)
	}
	router.GET("/spare_parts_unit_detail", controller.GetSpare_parts_unitsById)

	categories := router.Group("/categories")
	{

		categories.GET("/", controller.GetAllCategories)
	}

	router.GET("/category_detail", controller.GetCategoriesById)
	products := router.Group("/products")
	{
		products.GET("/", controller.GetAllProducts_by_multi_params)
	}

	router.GET("/product_detail", controller.GetProductById)
	feedback := router.Group("/feedback")
	{
		feedback.POST("/", controller.CreateFeedback)
		feedback.GET("/", controller.GetFeedbackByProduct_id)
	}
	order := router.Group("/order")
	{
		order.POST("/", controller.CreateOrder)
	}
	feedback_form := router.Group("/feedback_form")
	{
		feedback_form.POST("/", controller.FeedbackFormPost)
	}
	news := router.Group("/news")
	{
		news.GET("/", controller.GetAllNews_by_params)
	}
	router.GET("/new_detail", controller.GetNewById)
	services := router.Group("/services")
	{
		services.GET("/", controller.GetAllServices_by_params)
	}
	router.GET("/service_detail", controller.GetServiceById)
	cms := router.Group("/cms")
	{
		cms.GET("/sliders", controller.GetAllSliders)
		cms.GET("/links", controller.GetAllLinks)
		cms.GET("/site_link", controller.GetAllSite_link)
		cms.GET("/address", controller.GetAllAddress)
		cms.GET("/about", controller.GetAllAbout)
	}
	search := router.Group("/search")
	{
		search.GET("/", controller.GetAllResultSearch)
	}
	price_list := router.Group("/price_list")
	{
		price_list.GET("/", controller.GetPrice_List)
	}
	email_list := router.Group("/email_sender")
	{
		email_list.POST("/", controller.Email_list)

	}
	router.GET("/email_update", controller.EmailStatusUpdate)

	router.GET("/:id", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/:id/:id", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

}
