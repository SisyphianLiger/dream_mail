package handlers


// Verify The Emails

type SenderEmailer struct {
	SenderEmailer string
}

type ReceiverEmailer struct {
	ReceiverEmailer string
}

type Message struct {	
	Subject       string
	Body          string
}

type Payload struct {
	SenderEmailer string
	ReceiverEmailer string
	Subject       string
	Body          string
}

type SendEmail interface {
	SendMailGun() error
	SendSparkPost() error
}

func (se *SenderEmailer) ValidEmail() (SenderEmailer, error) {
	// Add logic to validate the email
	return SenderEmailer{},nil
}

func (re *ReceiverEmailer) ValidEmail() (ReceiverEmailer, error) {
	// Add logic to validate the email
	return ReceiverEmailer{}, nil
}




