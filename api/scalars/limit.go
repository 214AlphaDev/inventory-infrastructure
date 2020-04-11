package scalars

import (
	"encoding/json"
	"errors"

	vo "github.com/214alphadev/inventory-infrastructure-go/api/value_objects"
)

type Limit struct {
	vo.Limit
}

func (Limit) ImplementsGraphQLType(name string) bool {
	return name == "Limit"
}

func (m *Limit) UnmarshalGraphQL(Limit interface{}) error {
	switch v := Limit.(type) {
	case *int32:
		Limit, err := vo.NewLimit(*v)
		if err != nil {
			return err
		}
		m.Limit = Limit
		return nil
	case int32:
		Limit, err := vo.NewLimit(v)
		if err != nil {
			return err
		}
		m.Limit = Limit
		return nil
	default:
		return errors.New("failed to unmarshal limit")
	}

}

func (m Limit) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Limit.Limit())
}
