package main

import (
	"github.com/SisyphianLiger/dream_mail/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Print("Problem with .env file")
	}
	app := echo.New()

	emailer := handler.Emailer{}

	app.GET("/", emailer.HandleEmailerShow)
	app.POST("/emailed", emailer.SendMail)
	app.Logger.Fatal(app.Start("localhost:9001"))
}
