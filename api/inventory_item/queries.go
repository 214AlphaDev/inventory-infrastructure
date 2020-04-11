package inventory_item

import (
	"context"
	"errors"

	"github.com/graph-gophers/graphql-go"
	uuid "github.com/satori/go.uuid"
	am "github.com/214alphadev/community-authentication-middleware"
	"github.com/214alphadev/inventory-infrastructure-go/api/entities"
	"github.com/214alphadev/inventory-infrastructure-go/api/scalars"
	"github.com/214alphadev/inventory-infrastructure-go/api/types"
	vo_local "github.com/214alphadev/inventory-infrastructure-go/api/value_objects"
)

type FetchInventoryItemsResponse struct {
	error        error
	fetchedItems []*entities.InventoryItemEntity
}

type inventoryItemsQuery struct {
	Position graphql.Time
	Limit    scalars.Limit
}

func (f FetchInventoryItemsResponse) Error() *string {
	if f.error == nil {
		return nil
	}
	e := f.error.Error()
	return &e
}

func (f FetchInventoryItemsResponse) FetchedItems() []*entities.InventoryItemEntity {
	return f.fetchedItems
}

func (i *InventoryItemResolver) FetchItems(ctx context.Context, args inventoryItemsQuery) (FetchInventoryItemsResponse, error) {

	memberID := am.GetAuthenticateMember(ctx)

	switch memberID {

	case nil:
		return FetchInventoryItemsResponse{error: errors.New("Unauthenticated")}, nil
	default:

		member, err := i.community.GetMember(*memberID)
		if err != nil {
			return FetchInventoryItemsResponse{}, err
		}

		if !member.Verified {
			return FetchInventoryItemsResponse{error: errors.New("Unauthorized")}, err
		}

		query := types.InventoryItemsQuery{
			Position: args.Position.Time,
			Limit:    args.Limit.Limit.Limit(),
		}

		fetchedItemsEntities := []*entities.InventoryItemEntity{}

		fetchedItems, err := i.inventory.GetItems(query)
		if err != nil {
			return FetchInventoryItemsResponse{}, err
		}

		for _, fetchedItem := range fetchedItems {

			uuid, err := uuid.FromString(fetchedItem.OfferedBy)
			if err != nil {
				return FetchInventoryItemsResponse{}, err
			}

			member, err := i.community.GetMember(uuid)
			if err != nil {
				return FetchInventoryItemsResponse{}, err
			}

			amount, err := vo_local.NewAmount(int32(fetchedItem.MoneyAmount))
			if err != nil {
				return FetchInventoryItemsResponse{}, err
			}
			currency, err := vo_local.NewCurrency(fetchedItem.MoneyCurrency)
			if err != nil {
				return FetchInventoryItemsResponse{}, err
			}

			fetchedItemEntity := entities.NewInventoryItemEntity(
				fetchedItem.Location,
				fetchedItem.CreationDate,
				types.NewMemberResponse(member.ID, member.Username, member.Metadata),
				scalars.Amount{Amount: amount},
				scalars.Currency{Currency: currency},
				fetchedItem.Description,
			)
			fetchedItemsEntities = append(fetchedItemsEntities, &fetchedItemEntity)
		}

		return FetchInventoryItemsResponse{fetchedItems: fetchedItemsEntities}, nil

	}

}
