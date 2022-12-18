package visitor

import (
	"fmt"
	"io"
)

// The Visitor design pattern is used to separate the logic needed to work with
// a specific object and the object itself.The logic is implemented by a "visitor"
// called by the specific object.
type Visitor interface {
	VisitA(*MessageA)
	VisitB(*MessageB)
}

type MessageVisitor struct{}

func (mv *MessageVisitor) VisitA(m *MessageA) {
	m.Output.Write([]byte(fmt.Sprintf("A: %s (Visited A)", m.Msg)))
}

func (mv *MessageVisitor) VisitB(m *MessageB) {
	m.Output.Write([]byte(fmt.Sprintf("B: %s (Visited B)", m.Msg)))
}

type Visitable interface {
	Accept(Visitor)
}

type MessageA struct {
	Msg    string
	Output io.Writer
	v      Visitor
}

func (a *MessageA) Accept(v Visitor) {
	a.v = v
}

func (a *MessageA) Print() {
	if a.v != nil {
		a.v.VisitA(a)
	}
}

type MessageB struct {
	Msg    string
	Output io.Writer
	v      Visitor
}

func (b *MessageB) Accept(v Visitor) {
	b.v = v
}

func (b *MessageB) Print() {
	if b.v != nil {
		b.v.VisitB(b)
	}
}
