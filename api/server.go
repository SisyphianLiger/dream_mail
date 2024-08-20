package api

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

/*
	TODOS
	1. FLATTEN THIS FOLDER!
	2. SET UP GODOTENV to EMAILER
*/

// Here we could add a DB connection etc
type ApiConfig struct{
	MailGunApi string
	SparkPost string
	DomainName string
}

func StartServer() {

	err := godotenv.Load()

	if err != nil {
		log.Print("Problem with .env file")
	}
	app := echo.New()

	// Remember this is the correct way
	emailer := ApiConfig{}
	emailer.SetValidEnv()

	app.GET("/", emailer.HandleEmailerShow)
	app.POST("/emailed", emailer.SendMail)

	app.Logger.Fatal(app.Start("localhost:9001"))
}

func (api *ApiConfig) SetValidEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Environment Variables not accessable")
	}
	
	mg := os.Getenv("MAIL_GUN_API_KEY")
	if mg == "" {
		log.Fatal("Unable to Load MailGun API")
	}

	sp := os.Getenv("SPARK_POST_API")
	if sp == "" {
		log.Fatal("Unable to Load MailGun API")
	}

	dn := os.Getenv("DOMAIN_NAME")
	if dn == "" {
		log.Fatal("Unable to Load MailGun API")
	}

	api.MailGunApi = mg
	api.SparkPost = sp
	api.DomainName = dn
}


