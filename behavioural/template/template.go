package template

import "strings"

// The Template design pattern is used to break down an algorithm into a series
// of steps, turn theses steps into methods, and execute these methods inside
// a method that organizes the execution of the algorithm (teplate method).
type Template interface {
	firstStep() string
	thirdStep() string
	ExecuteAlgorithm(MessageRetriever) string // the MessageRetriever is used in the secon step
}

type MessageRetriever interface {
	Message() string
}

type TemplateImplementation struct{}

func (t *TemplateImplementation) firstStep() string {
	return "hello"
}

func (t *TemplateImplementation) thirdStep() string {
	return "template"
}

// execution of the algorithm
func (t *TemplateImplementation) ExecuteAlgorithm(m MessageRetriever) string {
	var buf strings.Builder
	buf.WriteString(t.firstStep())           // first step of the algorithm
	buf.WriteString(" " + m.Message() + " ") // second step of the algorithm
	buf.WriteString(t.thirdStep())           // third step of the algorithm
	return buf.String()
}

type AnonymousTemplateImplementation struct{}

func (a *AnonymousTemplateImplementation) firstStep() string {
	return "hello"
}

func (a *AnonymousTemplateImplementation) thirdStep() string {
	return "template"
}

func (a *AnonymousTemplateImplementation) ExecuteAlgorithm(f func() string) string {
	var buf strings.Builder
	buf.WriteString(a.firstStep())   // first step of the algorithm
	buf.WriteString(" " + f() + " ") // second step of the algorithm
	buf.WriteString(a.thirdStep())   // third step of the algorithm
	return buf.String()
}

type TemplateAdapter struct {
	myFunc func() string
}

func (a *TemplateAdapter) Message() string {
	if a.myFunc != nil {
		return a.myFunc()
	}
	return ""
}
func MessageRetrieverAdapter(f func() string) MessageRetriever {
	return &TemplateAdapter{myFunc: f}
}
