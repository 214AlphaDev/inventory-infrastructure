package scalars

import (
	"encoding/json"
	"errors"

	vo "github.com/214alphadev/inventory-infrastructure-go/api/value_objects"
)

type LastName struct {
	vo.LastName
}

func (LastName) ImplementsGraphQLType(name string) bool {
	return name == "LastName"
}

func (m *LastName) UnmarshalGraphQL(LastName interface{}) error {
	switch v := LastName.(type) {
	case *string:
		LastName, err := vo.NewLastName(*v)
		if err != nil {
			return err
		}
		m.LastName = LastName
		return nil
	case string:
		LastName, err := vo.NewLastName(v)
		if err != nil {
			return err
		}
		m.LastName = LastName
		return nil
	default:
		return errors.New("failed to unmarshal lastname")
	}

}

func (m LastName) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.LastName.LastName())
}
