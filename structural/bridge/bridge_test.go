package structural

import (
	"strings"
	"testing"
)

func TestPrintAPI1(t *testing.T) {
	api1 := PrinterImpl1{}

	err := api1.PrintMessage("Hello")
	if err != nil {
		t.Errorf("Error trying to use the API1 implementation: Message: %s", err.Error())
	}
}

func TestPrintAPI2(t *testing.T) {
	api2 := PrinterImpl2{}

	err := api2.PrintMessage("Hello")
	if err != nil {
		expectedErrorMessage := "you need to pass an io.Writer to PrinterImpl2"
		if !strings.Contains(err.Error(), expectedErrorMessage) {
			t.Errorf("Error message was not correct.\n Actual: %s\n Expected: %s", err.Error(), expectedErrorMessage)
		}
	}

	myWriter := MyWriter{}
	api2 = PrinterImpl2{
		Writer: &myWriter,
	}
	expectedMessage := "Hello"
	err = api2.PrintMessage(expectedMessage)
	if err != nil {
		t.Errorf("Error trying to use the API2 implementation: %s", err.Error())
	}
	if myWriter.Msg != expectedMessage {
		t.Fatalf("API2 did not write correctly on the io.Writer.\n Actual: %s\n Expected: %s", myWriter.Msg, expectedMessage)
	}
}

func TestNormalPrinter(t *testing.T) {
	expectedMessage := "Hello io.Writer"
	normal := NormalPrinter{
		Msg:     expectedMessage,
		Printer: &PrinterImpl1{},
	}

	err := normal.Print()
	if err != nil {
		t.Error(err.Error())
	}

	myWriter := MyWriter{}
	normal = NormalPrinter{
		Msg: expectedMessage,
		Printer: &PrinterImpl2{
			Writer: &myWriter,
		},
	}
	err = normal.Print()
	if err != nil {
		t.Error(err.Error())
	}
	if myWriter.Msg != expectedMessage {
		t.Errorf("The expected message on the io.Writer does not match actual.\n Actual: %s\n Expected: %s", myWriter.Msg, expectedMessage)
	}
}

func TestPacktPrinter(t *testing.T) {
	passedMessage := "Hello io.Writer"
	expectedMessage := "Message from Packt: Hello io.Writer"
	packt := PacktPrinter{
		Msg:     passedMessage,
		Printer: &PrinterImpl1{},
	}

	err := packt.Print()
	if err != nil {
		t.Error(err.Error())
	}

	myWriter := MyWriter{}
	packt = PacktPrinter{
		Msg: passedMessage,
		Printer: &PrinterImpl2{
			Writer: &myWriter,
		},
	}
	err = packt.Print()
	if err != nil {
		t.Error(err.Error())
	}
	if myWriter.Msg != expectedMessage {
		t.Errorf("The expected message on the io.Writer does not match actual.\n Actual: %s\n Expected: %s", myWriter.Msg, expectedMessage)
	}
}
