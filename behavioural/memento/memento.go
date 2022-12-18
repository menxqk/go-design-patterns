package memento

import "errors"

// The Memento design pattern is used to save the state of an object and allow
// it to undo actions and return to previous states.
type State struct {
	Description string
}

// Memento
type Memento struct {
	state State
}

// Originator
type Originator struct {
	state State
}

func (o *Originator) NewMemento() Memento {
	return Memento{state: o.state}
}

func (o *Originator) ExtractAndStoreState(m Memento) {
	o.state = m.state
}

// Caretaker
type CareTaker struct {
	mementoList []Memento
}

func (c *CareTaker) Add(m Memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *CareTaker) Memento(i int) (Memento, error) {
	if len(c.mementoList) < i || i < 0 {
		return Memento{}, errors.New("index not found")
	}
	return c.mementoList[i], nil
}
