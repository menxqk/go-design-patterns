package strategy

import (
	"io"
)

// The Strategy design pattern is used to hide algorithms behind an interface.
// The algorithms achieve the same functionality in a different way.

type PrintStrategy interface {
	Print() error
	SetLog(io.Writer)
	SetWriter(io.Writer)
}

type PrintOutput struct {
	Writer    io.Writer
	LogWriter io.Writer
}

func (p *PrintOutput) SetWriter(w io.Writer) {
	p.Writer = w
}

func (p *PrintOutput) SetLog(w io.Writer) {
	p.LogWriter = w
}
