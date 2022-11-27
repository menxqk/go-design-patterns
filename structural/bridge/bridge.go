package structural

import (
	"errors"
	"fmt"
	"io"
)

// The Bridge design pattern is used to decouple an object from what it does.
// This results in the decoupling of the abstraction from its implementation,
// allowing both to vary independently.
type PrinterAPI interface {
	PrintMessage(string) error
}

// ---------------------------------
type PrinterImpl1 struct{}

func (p *PrinterImpl1) PrintMessage(msg string) error {
	fmt.Printf("%s\n", msg)
	return nil
}

// ---------------------------------
type PrinterImpl2 struct {
	Writer io.Writer
}

func (p *PrinterImpl2) PrintMessage(msg string) error {
	if p.Writer == nil {
		return errors.New("you need to pass an io.Writer to PrinterImpl2")
	}
	fmt.Fprintf(p.Writer, "%s", msg)
	return nil
}

type MyWriter struct {
	Msg string
}

func (m *MyWriter) Write(p []byte) (int, error) {
	n := len(p)
	if n > 0 {
		m.Msg = string(p)
		return n, nil
	}
	return n, errors.New("content received on Writer was empty")
}

// ---------------------------------
type PrinterAbstraction interface {
	Print() error
}

type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (n *NormalPrinter) Print() error {
	n.Printer.PrintMessage(n.Msg)
	return nil
}

// ---------------------------------
type PacktPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (p *PacktPrinter) Print() error {
	p.Printer.PrintMessage(fmt.Sprintf("Message from Packt: %s", p.Msg))
	return nil
}
