package main

import (
	"log"

	"github.com/Ellie-Yen/go_scraping_house/routes"
)

func main() {
	// add routes
	r := routes.SetupRouter()

	port := "8080"
	host := "http://0.0.0.0:" + port
	log.Println("Server running on", "\x1b]8;;"+host+"\x07"+""+host+"\x1b]8;;\x07"+"\u001b[0m")
	r.Run(":" + port)
}
