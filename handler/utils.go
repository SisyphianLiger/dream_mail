package handler

import (
	"errors"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func (e *Emailer) PrettyEmailerPrint() {
	log.Printf("Sender Email: %s\n", e.Senderemail)
	log.Printf("Receiver Email: %s\n", e.Receiveremail)
	log.Printf("Subject is: %s\n", e.Subject)
	log.Printf("Body: %s\n", e.Body)
}

func (cn *Connection) SendMailNoAPIs(c echo.Context) error {
	e := Emailer{}
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

	e.PrettyEmailerPrint()
	c.Response().Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.Response().WriteHeader(http.StatusOK)

	return nil
}
