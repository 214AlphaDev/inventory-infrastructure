package value_objects_local

import (
	"errors"
	"regexp"
)

type FirstName struct {
	firstName string
}

func (f FirstName) FirstName() string {
	return f.firstName
}

func NewFirstName(firstName string) (FirstName, error) {

	validFirstNameSlice := []*regexp.Regexp{}

	latinFirstName, err := regexp.Compile(`^[a-zA-Z\x{00C0}-\x{00FF}]+(([',. -][a-zA-Z\x{00C0}-\x{00FF}])?[a-zA-Z\x{00C0}-\x{00FF}]*)*$`)

	if err != nil {
		return FirstName{}, errors.New("invalid regex")
	}

	validFirstNameSlice = append(validFirstNameSlice, latinFirstName)

	for _, validFirstName := range validFirstNameSlice {
		if validFirstName.MatchString(firstName) {
			return FirstName{firstName: firstName}, nil
		}
	}

	return FirstName{}, errors.New("invalid FirstName")

}
