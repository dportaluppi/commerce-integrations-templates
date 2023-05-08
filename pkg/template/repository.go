package template

// Repository defines the interface for template storage.
type Repository interface {
	Save(*Template) error
	FindByName(string) (*Template, error)
	Delete(*Template) error
}
