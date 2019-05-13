// Copyright © 2019 SIC! Software GmbH

package output

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	sensu "github.com/sensu/sensu-go/types"
)

func extendedEventFromEvent(event *sensu.Event) *ExtendedEvent {
	return &ExtendedEvent{
		Event:            event,
		Status:           messageEventStatus(event),
		EventAction:      formattedEventAction(event),
		EventKey:         eventKey(event),
		Output:           formattedEventOutput(event, 100),
		FullOutput:       event.Check.Output,
		FormattedMessage: formattedMessage(event),
	}
}

func resolveTemplate(templateValue string, event *ExtendedEvent) (string, error) {
	var resolved bytes.Buffer

	tmpl, err := template.New("tmpl").Parse(templateValue)
	if err != nil {
		return "", err
	}

	err = tmpl.Execute(&resolved, *event)
	if err != nil {
		return "", err
	}

	return resolved.String(), nil
}

func messageEventStatus(event *sensu.Event) string {
	switch event.Check.Status {
	case 0:
		return "Resolved"
	case 2:
		return "Critical"
	default:
		return "Warning"
	}
}

func formattedEventAction(event *sensu.Event) string {
	switch event.Check.Status {
	case 0:
		return "RESOLVED"
	default:
		return "ALERT"
	}
}

func formattedEventOutput(event *sensu.Event, maxLength int) string {
	output := strings.Trim(strings.Trim(strings.Trim(event.Check.Output, "\n"), "\r"), "\r\n")

	if len(event.Check.Output) > maxLength {
		output = output[0:maxLength] + "..."
	}

	return output
}

func eventKey(event *sensu.Event) string {
	return fmt.Sprintf("%s/%s", event.Entity.Name, event.Check.Name)
}

func formattedMessage(event *sensu.Event) string {
	return fmt.Sprintf("[%s] %s - %s", formattedEventAction(event), eventKey(event), formattedEventOutput(event, 100))
}
