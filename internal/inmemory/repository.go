package inmemory

import (
	"sync"

	"github.com/dportaluppi/commerce-integrations-templates/pkg/template"
)

// Repository represents an in-memory implementation of the template.Repository interface.
type Repository struct {
	templates map[string]*template.Template
	mu        sync.Mutex
}

// NewRepository creates a new in-memory template repository.
func NewRepository() *Repository {
	return &Repository{
		templates: make(map[string]*template.Template),
	}
}

// Save saves a template in the in-memory repository.
func (r *Repository) Save(t *template.Template) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.templates[t.Name] = t
	return nil
}

// FindByName finds a template by name in the in-memory repository.
func (r *Repository) FindByName(name string) (*template.Template, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	t, ok := r.templates[name]
	if !ok {
		return nil, template.ErrTemplateNotFound
	}

	return t, nil
}
