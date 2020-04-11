package input_types

import (
	"github.com/214alphadev/inventory-infrastructure-go/api/scalars"
)

type ProperName struct {
	FirstNameInput scalars.FirstName
	LastNameInput  scalars.LastName
}

func (p ProperName) FirstName() scalars.FirstName {
	return p.FirstNameInput
}

func (p ProperName) LastName() scalars.LastName {
	return p.LastNameInput
}

func NewProperName(firstName scalars.FirstName, lastName scalars.LastName) ProperName {
	return ProperName{firstName, lastName}
}
