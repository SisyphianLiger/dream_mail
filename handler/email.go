package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	sp "github.com/SparkPost/gosparkpost"
	"github.com/labstack/echo/v4"
	"github.com/mailgun/mailgun-go/v4"
)

const DOMAIN_NAME string = "dreamtest.dk"

type Emailer struct {
	Senderemail   string
	Receiveremail string
	Subject       string
	Body          string
}

func (e *Emailer) SendMail(c echo.Context) error {

	if val_err := e.ValidateSend(c); val_err != nil {
		return val_err
	}

	if mg_err := e.SendMailGun(c); mg_err != nil {
		log.Println("MailGun has failed to send trying with sparkpost...")
		if spark_err := e.SendSparkMail(); spark_err != nil {
			return c.HTML(http.StatusBadRequest, "<h1>Message Unsuccessful: Code not send with Mail Service Provider</h1>")
		}
	}

	c.Response().Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.Response().WriteHeader(http.StatusOK)
	return c.HTML(http.StatusOK, "<h1>Message Sent</h1>")
}

func (e *Emailer) ValidateSend(c echo.Context) error {

	// Need to check for domain (dreamtest.dk)
	e.Senderemail = c.FormValue("emailfrom")

	senderDomain := strings.Split(e.Senderemail, "@")

	if len(senderDomain) != 2 {
		return errors.New("The Sender Email is incorrectly specified, check for @'s")
	}

	if senderDomain[1] != "dreamtest.dk" {
		return errors.New("Sender Domain not correctly specified, please use @dreamtest.dk")
	}

	e.Receiveremail = c.FormValue("emailto")
	e.Subject = c.FormValue("subject")
	e.Body = c.FormValue("message")

	return nil
}

func (e *Emailer) SendMailGun(c echo.Context) error {
	MAIL_GUN_API_KEY := os.Getenv("MAIL_GUN_API_KEY")
	mg := mailgun.NewMailgun(DOMAIN_NAME, MAIL_GUN_API_KEY)

	// Because MailGun-EU
	mg.SetAPIBase("https://api.eu.mailgun.net/v3")

	message := mg.NewMessage(e.Senderemail, e.Subject, e.Body, e.Receiveremail)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, _, err := mg.Send(ctx, message)
	return err
}

func (e *Emailer) SendSparkMail() error {
	SPARK_POST_API := os.Getenv("SPARK_POST_API")

	cfg := &sp.Config{
		BaseUrl:    "https://api.sparkpost.com",
		ApiKey:     SPARK_POST_API,
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
