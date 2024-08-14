package tests

import (
	"github.com/SisyphianLiger/dream_mail/handler"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_No_Email_Is_Nil(t *testing.T) {
	emailer := handler.Emailer{}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("emailfrom=&emailto=&subject=&message="))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := emailer.ValidateSend(c); err == nil {
		t.Errorf("Function has incorrectly passed Sender email")
	}
}

func Test_To_Many_ATs_Is_Nil(t *testing.T) {

	emailer := handler.Emailer{}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("emailfrom=h@@@h&emailto=ryan.m.williams.12@gmail.com&subject=Test&message=Hello"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := emailer.ValidateSend(c); err == nil {
		t.Errorf("Function has incorrectly passed Sender email")
	}
}

func Test_Domain_Is_Wrong(t *testing.T) {
	emailer := handler.Emailer{}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("emailfrom=hello@wrongdomain.dk&emailto=ryan.m.williams.12@gmail.com&subject=Test&message=Hello"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := emailer.ValidateSend(c); err == nil {
		t.Errorf("Function has incorrectly passed Sender email")
	}
}

func Test_Not_An_Email_Is_Nil(t *testing.T) {
	emailer := handler.Emailer{}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("emailfrom=hello&emailto=ryan.m.williams.12@gmail.com&subject=Test&message=Hello"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := emailer.ValidateSend(c); err == nil {
		t.Errorf("Function has incorrectly passed Sender email")
	}
}
