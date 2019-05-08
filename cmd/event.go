// Copyright © 2019 SIC! Software GmbH

package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sensu/sensu-go/types"
	"github.com/spf13/cobra"

	"sensu-sic-handler/handler"
	"sensu-sic-handler/recipient"
)

var (
	annotationPrefix string
	handlerConfig    = &handler.Config{}
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

	eventCmd.PersistentFlags().StringVar(&annotationPrefix,
		"annotation-prefix",
		os.Getenv("EVENT_ANNOTATION_PREFIX"),
		"The annotation prefix to use, defaults to value of EVENT_ANNOTATION_PREFIX env variable")
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

	if val, ok := event.Entity.Annotations[fmt.Sprintf("%s/recipients", annotationPrefix)]; ok {
		recipients := recipient.Parse(redisClient, val)

		err := handler.Handle(recipients, event, handlerConfig)
		if err != nil {
			return err
		}
	}

	return nil
}
