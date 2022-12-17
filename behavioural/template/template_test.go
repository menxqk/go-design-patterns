package template

import (
	"strings"
	"testing"
)

type TestStruct struct {
	TemplateImplementation
}

func (t *TestStruct) Message() string {
	return "world"
}

func TestTemplate_ExecuteAlgorithm(t *testing.T) {
	t.Run("Using interfaces", func(t *testing.T) {
		s := &TestStruct{}
		res := s.ExecuteAlgorithm(s)
		expected := "world"
		if !strings.Contains(res, expected) {
			t.Errorf("expected string '%s' was not found on returned string: '%s'", expected, res)
		}
	})

	t.Run("Using anonymous functions", func(t *testing.T) {
		m := new(AnonymousTemplateImplementation)
		res := m.ExecuteAlgorithm(func() string {
			return "world"
		})
		expected := " world "
		if !strings.Contains(res, expected) {
			t.Errorf("expected string '%s' was not found on returned string: '%s'", expected, res)
		}
	})

	t.Run("Using anonymous functions adapted to an interface", func(t *testing.T) {
		messageRetriever := MessageRetrieverAdapter(func() string {
			return "world"
		})

		if messageRetriever == nil {
			t.Fatal("messageRetriver should not be nil")
		}

		template := TemplateImplementation{}
		res := template.ExecuteAlgorithm(messageRetriever)
		expected := " world "
		if !strings.Contains(res, expected) {
			t.Errorf("expected string '%s' was not found on returned string: '%s'", expected, res)
		}

	})
}
