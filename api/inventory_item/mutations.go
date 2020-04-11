package inventory_item

import (
	"context"
	"errors"
	"time"

	am "github.com/214alphadev/community-authentication-middleware"
	"github.com/214alphadev/inventory-infrastructure-go/api/entities"
	"github.com/214alphadev/inventory-infrastructure-go/api/types"
)

func (i *InventoryItemResolver) AddItem(ctx context.Context, a struct{ Input AddInventoryItemInput }) (AddInventoryItemResponse, error) {

	memberID := am.GetAuthenticateMember(ctx)

	switch memberID {
	case nil:
		return AddInventoryItemResponse{
			error: errors.New("Unauthenticated"),
		}, nil
	default:

		member, err := i.community.GetMember(*memberID)
		if err != nil {
			return AddInventoryItemResponse{}, err
		}

		if !member.Verified {
			return AddInventoryItemResponse{
				error: errors.New("Unauthorized"),
			}, nil
		}

		itemEntity := entities.NewInventoryItemEntity(
			a.Input.Location,
			time.Now(),
			types.NewMemberResponse(member.ID, member.Username, member.Metadata),
			a.Input.Money.Amount(),
			a.Input.Money.Currency(),
			a.Input.Description,
		)

		err = i.inventory.Save(itemEntity)

		switch err {
		case nil:
			return AddInventoryItemResponse{
				inventoryItem: &itemEntity,
			}, nil
		default:
			return AddInventoryItemResponse{}, err
		}

	}
}

type AddInventoryItemResponse struct {
	inventoryItem *entities.InventoryItemEntity
	error         error
}

func (a AddInventoryItemResponse) Error() *string {
	switch a.error {
	case nil:
		return nil
	default:
		e := a.error.Error()
		return &e
	}
}

func (a AddInventoryItemResponse) InventoryItem() *entities.InventoryItemEntity {
	return a.inventoryItem
}
