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

func InvalidEmailSuffixes(t *testing.T) {
	testing := []string{
		"user@example.comm",
		"info@company.nt",
		"contact@nonprofit.og",
		"student@university.ed",
		"official@department.gv",
		"developer@startup.io0",
		"support@business.co.uk1",
		"service@provider.dd",
		"client@enterprise.frr",
		"customer@store.jp1",
		"sales@shop.au1",
		"help@assistance.c",
		"admin@website.itt",
		"info@company.nl1",
		"user@mail.r",
		"contact@swiss.chh",
		"support@empresa.es1",
		"info@svensk.s",
	}
	receiver := handlers.ReceiverEmail{}
	for _, e := range testing {
		if receiver.ValidEmail(e) == nil {
			t.Errorf("Valid Email failed on correct email: %s\n", e)
		}
	}
}

func ValidEmailSuffixes(t *testing.T) {
	testing := []string{
		"user@example.com",
		"info@company.net",
		"contact@nonprofit.org",
		"student@university.edu",
		"official@department.gov",
		"developer@startup.io",
		"support@business.co.uk",
		"service@provider.de",
		"client@enterprise.fr",
		"customer@store.jp",
		"sales@shop.au",
		"help@assistance.ca",
		"admin@website.it",
		"info@company.nl",
		"user@mail.ru",
		"contact@swiss.ch",
		"support@empresa.es",
		"info@svensk.se",
	}
	receiver := handlers.ReceiverEmail{}
	for _, e := range testing {
		if receiver.ValidEmail(e) == nil {
			t.Errorf("Valid Email failed on correct email: %s\n", e)
		}
	}
}

func TestValidSenderEmail(t *testing.T) {
	testing := []string{
		"user@example.com",
		"john.doe@company.org",
		"info@website.net",
		"test123@domain.co.uk",
		"first.last@subdomain.example.com",
		"user-name@provider.edu",
		"email2020@site.io",
		"support@business-name.com",
		"no_reply@newsletter.info",
		"contact.us@organization.gov",
	}

	for _, e := range testing {
		_, err := handlers.SplitAndCheck(e)
		if err != nil {
			t.Errorf("Split And Check Failed on %s\n", e)
		}
	}
}

func TestInvalidSenderEmail(t *testing.T) {
	testing := []string{
		"user@example@.com",
		"john.doe@company@.org",
		"@info@website.net",
		"@jtest123@domain.co.uk",
		"f@jirst.last@subdomain.example.com",
		"user-name@@jprovider.edu",
		"email2020@j@site.io",
		"support@bus@jiness-name.com",
		"no_reply@new@jsletter.info",
		"contact.us@or@jganization.gov",
	}

	for _, e := range testing {
		_, err := handlers.SplitAndCheck(e)
		if err == nil {
			t.Errorf("Split And Check Failed on %s\n", e)
		}
	}
}
