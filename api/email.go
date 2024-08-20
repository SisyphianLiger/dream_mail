package api

import (
	"context"
	sp "github.com/SparkPost/gosparkpost"
	"github.com/labstack/echo/v4"
	"github.com/mailgun/mailgun-go/v4"
	"log"
	"time"
)

// Payload Struct used for
type Emailer struct {
	Senderemail   string
	Receiveremail string
	Subject       string
	Body          string
}

func (e *Emailer) SendMailGun(c echo.Context, api *ApiConfig) error {
	mg := mailgun.NewMailgun(api.DomainName, api.MailGunApi)

	// Because MailGun-EU
	mg.SetAPIBase("https://api.eu.mailgun.net/v3")

	message := mg.NewMessage(e.Senderemail, e.Subject, e.Body, e.Receiveremail)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, _, err := mg.Send(ctx, message)
	return err
}

func (e *Emailer) SendSparkMail(api *ApiConfig) error {

	cfg := &sp.Config{
		BaseUrl:    "https://api.sparkpost.com",
		ApiKey:     api.SparkPost,
		ApiVersion: 1,
	}

	var client sp.Client

	err := client.Init(cfg)

	if err != nil {
		log.Fatalf("SparkPost client init failed: %s\n", err)
		return err
	}

	tx := &sp.Transmission{
		// Could be used to send to multiple recipients
		Recipients: []string{e.Receiveremail},
		Content: sp.Content{
			Text:    e.Body,
			From:    e.Senderemail,
			Subject: e.Subject,
		},
	}

	_, _, send_err := client.Send(tx)
	if send_err != nil {
		log.Printf("Error Sending Email via SparkPost: %v\n", send_err)
		return send_err
	}

	return nil
}
