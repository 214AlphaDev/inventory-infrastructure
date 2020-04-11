package scalars

import (
	"encoding/json"
	"errors"

	vo "github.com/214alphadev/inventory-infrastructure-go/api/value_objects"
)

type FirstName struct {
	vo.FirstName
}

func (FirstName) ImplementsGraphQLType(name string) bool {
	return name == "FirstName"
}

func (m *FirstName) UnmarshalGraphQL(FirstName interface{}) error {
	switch v := FirstName.(type) {
	case *string:
		FirstName, err := vo.NewFirstName(*v)
		if err != nil {
			return err
		}
		m.FirstName = FirstName
		return nil
	case string:
		FirstName, err := vo.NewFirstName(v)
		if err != nil {
			return err
		}
		m.FirstName = FirstName
		return nil
	default:
		return errors.New("failed to unmarshal firstname")
	}

}

func (m FirstName) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.FirstName.FirstName())
}
