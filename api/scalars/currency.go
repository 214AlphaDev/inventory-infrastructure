package scalars

import (
	"encoding/json"
	"errors"

	vo "github.com/214alphadev/inventory-infrastructure-go/api/value_objects"
)

type Currency struct {
	vo.Currency
}

func (Currency) ImplementsGraphQLType(name string) bool {
	return name == "Currency"
}

func (m *Currency) UnmarshalGraphQL(Currency interface{}) error {

	switch v := Currency.(type) {
	case *string:
		Currency, err := vo.NewCurrency(*v)
		if err != nil {
			return err
		}
		m.Currency = Currency
		return nil
	case string:
		Currency, err := vo.NewCurrency(v)
		if err != nil {
			return err
		}
		m.Currency = Currency
		return nil
	default:
		return errors.New("failed to unmarshal currency")
	}

}

func (m Currency) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Currency.Currency())
}
