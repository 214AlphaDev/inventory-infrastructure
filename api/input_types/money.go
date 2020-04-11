package input_types

import (
	"github.com/214alphadev/inventory-infrastructure-go/api/scalars"
)

type Money struct {
	AmountInput   scalars.Amount
	CurrencyInput scalars.Currency
}

func (m Money) Amount() scalars.Amount {
	return m.AmountInput
}
func (m Money) Currency() scalars.Currency {
	return m.CurrencyInput
}

func NewMoney(amount scalars.Amount, currency scalars.Currency) Money {
	return Money{amount, currency}
}
