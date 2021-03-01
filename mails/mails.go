package mails

import (
	"fmt"
	"net/mail"
	"net/smtp"
)

// Sender provides a method to send mails.
type Sender struct {
	Addr string
	Auth smtp.Auth
	From mail.Address
}

// Mail to send.
type Mail struct {
	To      mail.Address
	Subject string
	Body    string
}

// NewSender creates a email sender
func NewSender(host, port, username, password string) *Sender {
	return &Sender{
		Addr: host + port,
		Auth: smtp.PlainAuth("", username, password, host),
		From: mail.Address{
			Name:    "Login magic link",
			Address: username,
		},
	}
}

// Send sends email
func (s *Sender) Send(mail Mail) error {
	headers := map[string]string{
		"From":         s.From.String(),
		"To":           mail.To.String(),
		"Subject":      mail.Subject,
		"MIME-Version": "1.0",
		"Content-Type": "text/html; charset=utf-8",
	}
	msg := ""
	for k, v := range headers {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	msg += "\r\n"
	msg += mail.Body

	return smtp.SendMail(
		s.Addr,
		s.Auth,
		s.From.Address,
		[]string{mail.To.Address},
		[]byte(msg))
}
