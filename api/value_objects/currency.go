package value_objects_local

import "errors"

type Currency struct {
	currency string
}

func (c Currency) Currency() string {
	return c.currency
}

func NewCurrency(currency string) (Currency, error) {

	isValidCurrencyMap := map[string]bool{
		"EUR": true,
		"USD": true,
	}

	if !isValidCurrencyMap[currency] {
		return Currency{}, errors.New("invalid currency")
	}

	return Currency{
		currency: currency,
	}, nil

}
