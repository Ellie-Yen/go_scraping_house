package routes

import (
	"github.com/Ellie-Yen/go_scraping_house/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controllers.HouseList)

	return r
}
