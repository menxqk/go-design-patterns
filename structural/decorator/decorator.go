package structural

import (
	"errors"
	"fmt"
)

// The Decorator design pattern is used to decorate an already existing type with
// more functional features.
type HasIngredients interface {
	GetIngredients() (string, error)
}

// ---------------------------------
type Pizza struct{}

func (p *Pizza) GetIngredients() (string, error) {
	return "Pizza with the following ingredients:", nil
}

// Meat Decorator for HasIngredients(Pizza)
type Meat struct {
	HasIngredients HasIngredients
}

func (m *Meat) GetIngredients() (string, error) {
	if m.HasIngredients == nil {
		return "", errors.New("a Pizza is needed for adding Meat")
	}

	s, err := m.HasIngredients.GetIngredients()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s", s, "meat"), nil
}

// Onion Decorator for HasIngredients(Pizza)
type Onion struct {
	HasIngredients HasIngredients
}

func (o *Onion) GetIngredients() (string, error) {
	if o.HasIngredients == nil {
		return "", errors.New("a Pizza is needed for adding Onion")
	}

	s, err := o.HasIngredients.GetIngredients()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s", s, "onion"), nil
}
