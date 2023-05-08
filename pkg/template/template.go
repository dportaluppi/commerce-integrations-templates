package template

import (
	"bytes"
	"encoding/json"
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
	// TODO: Add any custom validation logic here.
	return nil
}

// Render renders the template with the provided data.
func (t *Template) Render(data map[string]any) (map[string]any, error) {
	contentBytes, err := json.Marshal(t.Content)
	if err != nil {
		return nil, err
	}
	contentString := string(contentBytes)

	tmpl, err := template.New(t.Name).Delims("[[", "]]").Funcs(sprig.TxtFuncMap()).Parse(contentString)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		return nil, err
	}

	return result, nil
}
