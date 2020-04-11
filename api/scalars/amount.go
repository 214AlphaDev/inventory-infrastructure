package scalars

import (
	"encoding/json"
	"errors"

	vo "github.com/214alphadev/inventory-infrastructure-go/api/value_objects"
)

type Amount struct {
	vo.Amount
}

func (Amount) ImplementsGraphQLType(name string) bool {
	return name == "Amount"
}

func (m *Amount) UnmarshalGraphQL(Amount interface{}) error {
	switch v := Amount.(type) {
	case *int32:
		Amount, err := vo.NewAmount(*v)
		if err != nil {
			return err
		}
		m.Amount = Amount
		return nil
	case int32:
		Amount, err := vo.NewAmount(v)
		if err != nil {
			return err
		}
		m.Amount = Amount
		return nil
	default:
		return errors.New("failed to unmarshal amount")
	}

}

func (m Amount) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Amount.Amount())
}
