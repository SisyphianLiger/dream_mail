package handlers

import (
	"log"
	"net/http"

	"github.com/SisyphianLiger/dream_mail/view"
	"github.com/labstack/echo/v4"
)

// Here we could add a DB connection etc
type Connection struct{}

func (cn *Connection) HandleEmailerShow(c echo.Context) error {
	return render(c, viewer.Show())
}

func (cn *Connection) SendMail(c echo.Context) error {

	e := Emailer{}

	snd, rec, err := ValidateEmails(c)
	if err != nil {
		log.Printf("snd is %s, rec is %s\n",snd, rec)
		return err
	}

	// Loading Up Message
	message := Message{}
	message.GetMessage(c)

	// Creating Payload
	e.LoadPayload(snd, rec, message)


	// LAST STEP REORGANIZE FAILURES
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
