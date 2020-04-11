package entities

import (
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/214alphadev/inventory-infrastructure-go/api/input_types"
	"github.com/214alphadev/inventory-infrastructure-go/api/scalars"
	"github.com/214alphadev/inventory-infrastructure-go/api/types"
)

type InventoryItemEntity struct {
	location     *string
	creationDate time.Time
	offeredBy    types.MemberResponse
	money        input_types.Money
	description  *string
}

func NewInventoryItemEntity(location *string, creationDate time.Time, offeredBy types.MemberResponse, moneyAmount scalars.Amount, moneyCurrency scalars.Currency, description *string) InventoryItemEntity {

	return InventoryItemEntity{
		location:     location,
		creationDate: creationDate,
		offeredBy:    offeredBy,
		money:        input_types.NewMoney(moneyAmount, moneyCurrency),
		description:  description,
	}
}

func (i InventoryItemEntity) Location() *string {
	return i.location
}

func (i InventoryItemEntity) CreationDate() graphql.Time {
	return graphql.Time{
		Time: i.creationDate,
	}
}

func (i InventoryItemEntity) OfferedBy() types.MemberResponse {
	return i.offeredBy

}

func (i InventoryItemEntity) Money() input_types.Money {
	return i.money

}

func (i InventoryItemEntity) Description() *string {
	return i.description
}
