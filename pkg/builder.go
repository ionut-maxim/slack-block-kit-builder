package builder

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/slack-go/slack"
	"sigs.k8s.io/yaml"
)

func BuildBlocks(i interface{}, path string) (*slack.Blocks, error) {
	var blocks slack.Blocks
	var templateContent bytes.Buffer

	if path == "" {
		return nil, fmt.Errorf("template path must not be empty")
	}

	t, err := template.ParseFiles(path)
	if err != nil {
		return nil, err
	}

	if err := t.Execute(&templateContent, i); err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(templateContent.Bytes(), &blocks)
	if err != nil {
		return nil, err
	}

	return &blocks, nil
}
