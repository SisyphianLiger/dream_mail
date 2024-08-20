package api

import (
	"github.com/SisyphianLiger/dream_mail/api/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
)

func StartServer() {

	err := godotenv.Load()

	if err != nil {
		log.Print("Problem with .env file")
	}
	app := echo.New()

	// Remember this is the correct way
	emailer := handlers.Connection{}

	app.GET("/", emailer.HandleEmailerShow)
	app.POST("/emailed", emailer.SendMail)

	app.Logger.Fatal(app.Start("localhost:9001"))
}
