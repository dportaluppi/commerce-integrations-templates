// Package template defines the structure of the Template entity and its associated operations.
package template

// Service is an interface that defines the methods for managing
// the business logic related to templates.
type Service interface {
	Create(name string, content map[string]any) (*Template, error)
	Get(name string) (*Template, error)
	Render(name string, data map[string]any) (map[string]any, error)
}

type service struct {
	repo Repository
}

// NewService creates a new template service.
func NewService(repo Repository) Service {
	return &service{repo}
}

// Create creates a new template.
func (s *service) Create(name string, content map[string]any) (*Template, error) {
	t := &Template{
		Name:    name,
		Content: content,
	}

	if err := t.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.Save(t); err != nil {
		return nil, err
	}

	return t, nil
}

// Get gets a template by name.
func (s *service) Get(name string) (*Template, error) {
	return s.repo.FindByName(name)
}

// Render renders a template with the provided data.
func (s *service) Render(name string, data map[string]any) (map[string]any, error) {
	t, err := s.repo.FindByName(name)
	if err != nil {
		return nil, err
	}

	return t.Render(data)
}
