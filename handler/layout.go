package handler

import (
	"github.com/SisyphianLiger/dream_mail/view"
	"github.com/labstack/echo/v4"
)

func (e *Emailer) HandleEmailerShow(c echo.Context) error {
	return render(c, viewer.Show())
}
