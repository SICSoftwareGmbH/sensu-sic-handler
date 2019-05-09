// Copyright © 2019 SIC! Software GmbH

package handler

import (
	"bytes"
	"errors"
	"net/mail"
	"net/smtp"

	sensu "github.com/sensu/sensu-go/types"

	"sensu-sic-handler/recipient"
)

var (
	mailSubjectTemplate = "[Sensu] {{.Entity.Name}}/{{.Check.Name}}: {{.Check.State}}"
	mailBodyTemplate    = "{{.Check.Output}}"
)

// HandleMail handles mail recipients (recipient.HandlerTypeMail)
func HandleMail(recipient *recipient.Recipient, event *sensu.Event, config *Config) error {
	if len(config.MailFrom) == 0 {
		return errors.New("from email is empty")
	}

	fromAddress, err := mail.ParseAddress(config.MailFrom)
	if err != nil {
		return err
	}

	if len(recipient.Args["mail"]) == 0 {
		return errors.New("to email is empty")
	}

	toAddress, err := mail.ParseAddress(recipient.Args["mail"])
	if err != nil {
		return err
	}

	subject, err := resolveTemplate(mailSubjectTemplate, event)
	if err != nil {
		return err
	}

	body, err := resolveTemplate(mailBodyTemplate, event)
	if err != nil {
		return err
	}

	msg := []byte("From: " + fromAddress.String() + "\r\n" +
		"To: " + toAddress.String() + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	conn, err := smtp.Dial(config.SMTPAddress)
	if err != nil {
		return err
	}
	defer conn.Close()

	conn.Mail(fromAddress.Address)
	conn.Rcpt(toAddress.Address)

	data, err := conn.Data()
	if err != nil {
		return err
	}
	defer data.Close()

	buffer := bytes.NewBuffer(msg)
	if _, err := buffer.WriteTo(data); err != nil {
		return err
	}

	return nil
}
