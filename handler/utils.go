package handler

import (
	"fmt"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func (e *Emailer) PrettyEmailerPrint() {
	fmt.Printf("Sender Email: %s\n", e.Senderemail)
	fmt.Printf("Receiver Email: %s\n", e.Receiveremail)
	fmt.Printf("Subject is: %s\n", e.Subject)
	fmt.Printf("Body: %s\n", e.Body)
}
