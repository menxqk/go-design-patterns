package chain_of_responsibility

import (
	"fmt"
	"io"
	"strings"
)

// The Chain of Responsibility design pattern is used to create a chain
// of functions that execute in order and have a single responsibility.
type ChainLogger interface {
	Next(string)
}

type FirstLogger struct {
	NextChain ChainLogger
}

func (l *FirstLogger) Next(s string) {
	fmt.Printf("First logger: %s\n", s)
	if l.NextChain != nil {
		l.NextChain.Next(s)
	}
}

type SecondLogger struct {
	NextChain ChainLogger
}

func (l *SecondLogger) Next(s string) {
	if strings.Contains(strings.ToLower(s), "hello") {
		fmt.Printf("Second logger: %s\n", s)
		if l.NextChain != nil {
			l.NextChain.Next(s)
		}
		return
	}
	fmt.Print("Finishing in second logger\n\n")
}

type WriterLogger struct {
	NextChain ChainLogger
	Writer    io.Writer
}

func (l *WriterLogger) Next(s string) {
	if l.Writer != nil {
		l.Writer.Write([]byte(fmt.Sprintf("WriterLogger: %s", s)))
	}
	if l.NextChain != nil {
		l.NextChain.Next(s)
	}
}

type ClosureChain struct {
	NextChain ChainLogger
	Closure   func(string)
}

func (c *ClosureChain) Next(s string) {
	if c.Closure != nil {
		c.Closure(s)
	}
	if c.NextChain != nil {
		c.Next(s)
	}
}
