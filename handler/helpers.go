// Copyright © 2019 SIC! Software GmbH

package handler

import (
	"bytes"
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
