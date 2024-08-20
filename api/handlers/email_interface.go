package handlers

import (
	"errors"
	"strings"

	"github.com/labstack/echo/v4"
)

// Verify The Emails

type SenderEmail struct {
	SenderEmailer string
}

type ReceiverEmail struct {
	ReceiverEmailer string
}

type Message struct {
	Subject string
	Body    string
}

type SendEmail interface {
	SendMailGun() error
	SendSparkPost() error
}

func ValidateEmails(c echo.Context) (SenderEmail, ReceiverEmail, error) {

	sender := SenderEmail{}
	receiver := ReceiverEmail{}

	s := strings.ToLower(c.FormValue("emailfrom"))
	r := strings.ToLower(c.FormValue("emailto"))
	// Need to check for domain (dreamtest.dk)
	sendEmail, sErr := SplitAndCheck(s)
	if sErr != nil {
		return SenderEmail{}, ReceiverEmail{}, sErr
	}

	recEmail, rErr := SplitAndCheck(r)
	if rErr != nil {
		return SenderEmail{}, ReceiverEmail{}, rErr
	}
	// TODO FIX THIS PART
	if sendvErr := sender.ValidEmail(sendEmail); sendvErr != nil {
		return SenderEmail{}, ReceiverEmail{}, sendvErr
	}
	if recvErr := receiver.ValidEmail(recEmail); recvErr != nil {
		return SenderEmail{}, ReceiverEmail{}, recvErr
	}

	// Now we can add them
	sender.SenderEmailer = c.FormValue("emailfrom")
	receiver.ReceiverEmailer = c.FormValue("emailto")

	return sender, receiver, nil
}

func (re *ReceiverEmail) ValidEmail(Estr string) error {
	// Add logic to validate the email
	tlds := []string{".com", ".net", ".org", ".edu", ".gov", ".io", ".co.uk",
		".de", ".fr", ".jp", ".au", ".ca", ".it", ".nl", ".ru", ".ch", ".es",
		".se"}

	for _, e := range tlds {
		if strings.HasSuffix(Estr, e) {
			return nil
		}
	}
	return errors.New("Incorrect top level domain please check the tld")
}

func (se *SenderEmail) ValidEmail(Estr string) error {

	// Add logic to validate the email
	if !strings.HasSuffix(Estr, "@dreamtest.dk") {
		return errors.New("Sender Domain not correctly specified, please use @dreamtest.dk")
	}

	// Here we can specify more logic i.e. can only send from account email
	return nil
}

func SplitAndCheck(email string) (string, error) {

	emailChk := strings.Split(email, "@")

	if len(emailChk) != 2 {
		return "", errors.New("The Sender Email is incorrectly specified, check for @'s")
	}

	return email, nil
}
