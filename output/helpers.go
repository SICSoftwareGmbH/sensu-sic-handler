// Copyright © 2019 SIC! Software GmbH

package output

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	sensu "github.com/sensu/sensu-go/types"
)

func resolveTemplate(templateValue string, event *sensu.Event) (string, error) {
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

func chomp(s string) string {
	return strings.Trim(strings.Trim(strings.Trim(s, "\n"), "\r"), "\r\n")
}

func messageStatus(event *sensu.Event) string {
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

func eventKey(event *sensu.Event) string {
	return fmt.Sprintf("%s/%s", event.Entity.Name, event.Check.Name)
}

func eventSummary(event *sensu.Event, maxLength int) string {
	output := chomp(event.Check.Output)

	if len(event.Check.Output) > maxLength {
		output = output[0:maxLength] + "..."
	}

	return fmt.Sprintf("%s:%s", eventKey(event), output)
}

func formattedMessage(event *sensu.Event) string {
	return fmt.Sprintf("%s - %s", formattedEventAction(event), eventSummary(event, 100))
}
