// Copyright © 2019 SIC! Software GmbH

package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sensu/sensu-go/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"sensu-sic-handler/handler"
	"sensu-sic-handler/recipient"
)

// eventCmd represents the "event" command
var eventCmd = &cobra.Command{
	Use:   "event",
	Short: "Handle event data",
	Run: func(cmd *cobra.Command, args []string) {
		event, err := loadEvent()
		if err != nil {
			terminateWithHelpAndMessage(cmd, fmt.Sprintf("failed to load event: %s", err.Error()))
			return
		}

		err = validateEvent(event)
		if err != nil {
			terminateWithHelpAndMessage(cmd, fmt.Sprintf("failed to validate event: %s", err.Error()))
			return
		}

		err = handleEvent(event)
		if err != nil {
			terminateWithHelpAndMessage(cmd, fmt.Sprintf("failed to handle event: %s", err.Error()))
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(eventCmd)

	eventCmd.PersistentFlags().String(
		"outputs",
		"mail,slack,xmpp",
		"The outputs to use, defaults to 'mail,slack,xmpp'")

	eventCmd.PersistentFlags().String(
		"annotation-prefix",
		"sic.software",
		"The annotation prefix to use")

	eventCmd.PersistentFlags().String(
		"smtp-address",
		"localhost",
		"The address of the SMTP server to use, defaults to localhost")

	eventCmd.PersistentFlags().String(
		"mail-from",
		"",
		"The sender address for emails")

	eventCmd.PersistentFlags().String(
		"slack-webhook-url",
		"",
		"The webhook url to send messages to")

	eventCmd.PersistentFlags().String(
		"slack-username",
		"sensu",
		"The username that messages will be sent as")

	eventCmd.PersistentFlags().String(
		"slack-icon-url",
		"http://s3-us-west-2.amazonaws.com/sensuapp.org/sensu.png",
		"A URL to an image to use as the user avatar")

	_ = viper.BindPFlag("outputs", eventCmd.PersistentFlags().Lookup("outputs"))
	_ = viper.BindPFlag("annotation-prefix", eventCmd.PersistentFlags().Lookup("annotation-prefix"))
	_ = viper.BindPFlag("smtp-address", eventCmd.PersistentFlags().Lookup("smtp-address"))
	_ = viper.BindPFlag("mail-from", eventCmd.PersistentFlags().Lookup("mail-from"))
	_ = viper.BindPFlag("slack-webhook-url", eventCmd.PersistentFlags().Lookup("slack-webhook-url"))
	_ = viper.BindPFlag("slack-username", eventCmd.PersistentFlags().Lookup("slack-username"))
	_ = viper.BindPFlag("slack-icon-url", eventCmd.PersistentFlags().Lookup("slack-icon-url"))
}

func loadEvent() (*types.Event, error) {
	event := &types.Event{}

	eventJSON, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return event, err
	}

	err = json.Unmarshal(eventJSON, event)
	if err != nil {
		return event, err
	}

	return event, nil
}

func validateEvent(event *types.Event) error {
	if event.Timestamp <= 0 {
		return errors.New("timestamp is missing or must be greater than zero")
	}

	if event.Entity == nil {
		return errors.New("entity is missing from event")
	}

	if !event.HasCheck() {
		return errors.New("check is missing from event")
	}

	if err := event.Entity.Validate(); err != nil {
		return err
	}

	if err := event.Check.Validate(); err != nil {
		return err
	}

	return nil
}

func handleEvent(event *types.Event) error {
	if event.Entity.Annotations == nil {
		return nil
	}

	if val, ok := event.Entity.Annotations[fmt.Sprintf("%s/recipients", viper.GetString("annotation-prefix"))]; ok {
		handlerConfig := &handler.Config{
			SMTPAddress:     viper.GetString("smtp-address"),
			MailFrom:        viper.GetString("mail-from"),
			SlackWebhookURL: viper.GetString("slack-webhook-url"),
			SlackUsername:   viper.GetString("slack-username"),
			SlackIconURL:    viper.GetString("slack-icon-url"),
		}

		recipients := recipient.Parse(redisClient, val)

		recipients = filterRecipients(recipients)

		err := handler.Handle(recipients, event, handlerConfig)
		if err != nil {
			return err
		}
	}

	return nil
}

func filterRecipients(recipients []*recipient.Recipient) []*recipient.Recipient {
	filtered := make([]*recipient.Recipient, 0)

	useMail, useSlack, useXMPP := false, false, false

	for _, output := range strings.Split(viper.GetString("outputs"), ",") {
		switch output {
		case "mail":
			useMail = true
		case "slack":
			useSlack = true
		case "xmpp":
			useXMPP = true
		}
	}

	for _, rcpt := range recipients {
		switch rcpt.Type {
		case recipient.HandlerTypeMail:
			if useMail {
				filtered = append(filtered, rcpt)
			}
		case recipient.HandlerTypeSlack:
			if useSlack {
				filtered = append(filtered, rcpt)
			}
		case recipient.HandlerTypeXMPP:
			if useXMPP {
				filtered = append(filtered, rcpt)
			}
		}
	}

	return filtered
}
