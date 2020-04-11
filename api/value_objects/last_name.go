package value_objects_local

import (
	"errors"
	"regexp"
)

type LastName struct {
	lastName string
}

func (l LastName) LastName() string {
	return l.lastName
}

func NewLastName(lastName string) (LastName, error) {

	validLastNameSlice := []*regexp.Regexp{}

	latinLastName, err := regexp.Compile(`^[a-zA-Z\x{00C0}-\x{00FF}]+(([',. -][a-zA-Z\x{00C0}-\x{00FF}])?[a-zA-Z\x{00C0}-\x{00FF}]*)*$`)

	if err != nil {
		return LastName{}, errors.New("invalid regex")
	}

	validLastNameSlice = append(validLastNameSlice, latinLastName)

	for _, validLastName := range validLastNameSlice {
		if validLastName.MatchString(lastName) {
			return LastName{lastName: lastName}, nil
		}
	}

	return LastName{}, errors.New("invalid LastName")

}
