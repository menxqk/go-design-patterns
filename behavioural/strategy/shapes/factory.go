package shapes

import (
	"fmt"
	"os"

	"github.com/menxqk/go-design-patterns/behavioural/strategy"
)

const (
	TEXT_STRATEGY  = "text"
	IMAGE_STRATEGY = "image"
)

func NewPrinter(s string) (strategy.PrintStrategy, error) {
	switch s {
	case TEXT_STRATEGY:
		return &TextSquare{
			PrintOutput: strategy.PrintOutput{
				Writer:    os.Stdout,
				LogWriter: os.Stdout,
			}}, nil
	case IMAGE_STRATEGY:
		return &ImageSquare{
			PrintOutput: strategy.PrintOutput{
				Writer:    nil,
				LogWriter: os.Stdout,
			}}, nil
	default:
		return nil, fmt.Errorf("strategy '%s' not found", s)
	}
}
