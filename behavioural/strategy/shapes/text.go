package shapes

import (
	"github.com/menxqk/go-design-patterns/behavioural/strategy"
)

type TextSquare struct {
	strategy.PrintOutput
}

func (t *TextSquare) Print() error {
	t.Writer.Write([]byte("Square"))
	return nil
}
