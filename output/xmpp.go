// Copyright © 2019 SIC! Software GmbH

package output

import (
	"errors"

	"github.com/mattn/go-xmpp"

	"sensu-sic-handler/recipient"
)

var xmppMessageTemplate = "{{ .FormattedMessage }}"

// XMPP handles XMPP recipients (recipient.HandlerTypeXMPP)
func XMPP(recipient *recipient.Recipient, event *ExtendedEvent, config *Config) (rerr error) {
	if len(config.XMPPServer) == 0 {
		return errors.New("hostname is empty")
	}

	if len(config.XMPPUsername) == 0 {
		return errors.New("username is empty")
	}

	if len(config.XMPPPassword) == 0 {
		return errors.New("password is empty")
	}

	clientOptions := xmpp.Options{
		Host:     config.XMPPServer,
		User:     config.XMPPUsername,
		Password: config.XMPPPassword,
		NoTLS:    true,
	}

	client, err := clientOptions.NewClient()
	if err != nil {
		return err
	}

	msg, err := resolveTemplate(xmppMessageTemplate, event)
	if err != nil {
		return err
	}

	switch recipient.Args["type"] {
	case "user":
		err = xmppSendUser(client, recipient.Args["user"], msg)
	case "muc":
		err = xmppSendMUC(client, recipient.Args["room"], msg)
	}

	if err != nil {
		return err
	}

	return nil
}

func xmppSendUser(client *xmpp.Client, remote string, msg string) error {
	_, err := client.Send(xmpp.Chat{Remote: remote, Type: "chat", Text: msg})
	if err != nil {
		return err
	}

	return nil
}

func xmppSendMUC(client *xmpp.Client, remote string, msg string) error {
	_, err := client.JoinMUCNoHistory(remote, "sensu")
	if err != nil {
		return err
	}

	_, err = client.Send(xmpp.Chat{Remote: remote, Type: "groupchat", Text: msg})
	if err != nil {
		return err
	}

	return nil
}
