package creational

import (
	"errors"
	"fmt"
)

// The Factory Method design pattern is used to group families of objects abstracted by a
// family object creator. The user works at the interface level instead of with concrete
// implementations. The responsability for creating new structures is delegated to the
// factory method.
func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized", m))
	}
}

type PaymentMethod interface {
	Pay(amount float32) string
}

// Our current implemented Payment methods are
const (
	Cash      = 1
	DebitCard = 2
)

type CashPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash", amount)
}

type DebitCardPM struct{}

func (d *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card", amount)
}
