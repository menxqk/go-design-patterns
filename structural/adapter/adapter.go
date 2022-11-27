package structural

import "fmt"

// The Adapter design pattern is used to enable the use of old code via a new signature.
// An old interface can be called using an adapter that implements a new interface and can access
// the old functionality.
type LegacyPrinter interface {
	Print(s string) string
}

type MyLegacyPrinter struct{}

func (l *MyLegacyPrinter) Print(s string) string {
	newMsg := fmt.Sprintf("Legacy Printer: %s", s)
	fmt.Println(newMsg)
	return newMsg
}

// -----------------------------------------
type NewPrinter interface {
	PrintStored() string
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (p *PrinterAdapter) PrintStored() string {
	if p.OldPrinter != nil {
		return p.OldPrinter.Print(fmt.Sprintf("Adapter: %s", p.Msg))
	}
	return p.Msg
}
