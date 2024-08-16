package main

import (
	"log"

	"github.com/SisyphianLiger/dream_mail/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Print("Problem with .env file")
	}
	app := echo.New()

	// Remember this is the correct way
	emailer := handler.Connection{}

	app.GET("/", emailer.HandleEmailerShow)
	app.POST("/emailed", emailer.SendMail)

	app.Logger.Fatal(app.Start("localhost:9001"))
}
