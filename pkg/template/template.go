package template

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
)

// Template represents a template entity.
type Template struct {
	Name    string
	Content map[string]any
}

// Validate validates the template entity.
func (t *Template) Validate() error {
	// Add any custom validation logic here.
	return nil
}

// Render renders the template with the provided data.
func (t *Template) Render(data map[string]any) (string, error) {
	contentBytes, err := json.Marshal(t.Content)
	if err != nil {
		return "", err
	}
	contentString := string(contentBytes)

	tmpl, err := template.New(t.Name).Funcs(sprig.TxtFuncMap()).Parse(contentString)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func replaceCustomTags(content string) string {
	replacedContent := strings.ReplaceAll(content, "{%", "{{")
	replacedContent = strings.ReplaceAll(replacedContent, "%}", "}}")
	return replacedContent
}
