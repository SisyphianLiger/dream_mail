package main

import (
	"github.com/SisyphianLiger/dream_mail/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	emailer := handler.Emailer{}

	app.GET("/", emailer.HandleEmailerShow)
	app.POST("/emailed", emailer.SendMail)
	app.Logger.Fatal(app.Start("localhost:9001"))
}
