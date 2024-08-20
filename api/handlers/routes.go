package handlers

import (
	"github.com/SisyphianLiger/dream_mail/view"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

// Here we could add a DB connection etc
type Connection struct{}

type Response struct {
	Success  bool     `json:"success"`
	Messages []string `json:"messages"`
}

func (cn *Connection) HandleEmailerShow(c echo.Context) error {
	return render(c, viewer.Show())
}

func (cn *Connection) SendMail(c echo.Context) error {

	e := Emailer{}

	snd, rec, err := ValidateEmails(c)
	if err != nil {
		errorMsg := []string{err.Error()}
		return c.JSON(http.StatusBadRequest, Response{
			Success:  false,
			Messages: errorMsg,
		})
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
			errorMsg := []string{"MailGunFailure: " + mg_err.Error(), "SparkPost Error: " + spark_err.Error()}
			return c.JSON(http.StatusServiceUnavailable, Response{
				Success:  false,
				Messages: errorMsg,
			})
		}
	}
	return c.JSON(http.StatusOK, Response{
		Success:  true,
		Messages: []string{"Message Send Successfully!"},
	})
}
