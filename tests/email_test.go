package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/SisyphianLiger/dream_mail/api/handlers"
	"github.com/labstack/echo/v4"
)

func TestPostRequestForRC(t *testing.T) {

	e := echo.New()
	email := handlers.Connection{}

	var wg sync.WaitGroup

	e.GET("/", email.HandleEmailerShow)
	e.POST("/emailed", email.SendMailNoAPIs)

	// Creating ctx that will be used to stop the server
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Start server
	go func() {
		if err := e.Start("localhost:9001"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// May need to change this number depending on computer hardware
	for i := 0; i < 500000; i++ {

		// Here we test the Payloads
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			req := httptest.NewRequest(http.MethodPost, "/emailed", strings.NewReader("emailfrom="+strconv.Itoa(i)+"@dreamtest.dk&emailto=&subject=&message="))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if sendErr := email.SendMailNoAPIs(c); sendErr != nil {
				t.Errorf("Failed to add email info to struct")
			}
		}(i)
	}
	// Shutdown the server gracefully
	wg.Wait()
	stop()
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if errSd := e.Shutdown(ctx); errSd != nil {
		t.Errorf("Could not shutdown properly")
	}

}

func TestNoEmailIsNil(t *testing.T) {
	emailer := handlers.Emailer{}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("emailfrom=&emailto=&subject=&message="))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := emailer.ValidateSend(c); err == nil {
		t.Errorf("Function has incorrectly passed Sender email")
	}
}

func TestToManyATsIsNil(t *testing.T) {

	emailer := handlers.Emailer{}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("emailfrom=h@@@h&emailto=ryan.m.williams.12@gmail.com&subject=Test&message=Hello"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := emailer.ValidateSend(c); err == nil {
		t.Errorf("Function has incorrectly passed Sender email")
	}
}

func TestDomainIsWrong(t *testing.T) {
	emailer := handlers.Emailer{}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("emailfrom=hello@wrongdomain.dk&emailto=ryan.m.williams.12@gmail.com&subject=Test&message=Hello"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := emailer.ValidateSend(c); err == nil {
		t.Errorf("Function has incorrectly passed Sender email")
	}
}

func TestNotAnEmailIsNil(t *testing.T) {
	emailer := handlers.Emailer{}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("emailfrom=hello&emailto=ryan.m.williams.12@gmail.com&subject=Test&message=Hello"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := emailer.ValidateSend(c); err == nil {
		t.Errorf("Function has incorrectly passed Sender email")
	}
}
