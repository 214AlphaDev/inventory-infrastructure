package inventory_item

import (
	cd "github.com/214alphadev/community-domain-go"
	"github.com/214alphadev/inventory-infrastructure-go/api/input_types"
	inventory "github.com/214alphadev/inventory-infrastructure-go/inventory_domain"
)

type InventoryItemResolver struct {
	inventory inventory.InventoryInterface
	community cd.CommunityInterface
}

type AddInventoryItemInput struct {
	Location    *string
	Money       input_types.Money
	Description *string
}

func NewInventoryItemResolver(inventory inventory.InventoryInterface, community cd.CommunityInterface) *InventoryItemResolver {
	return &InventoryItemResolver{
		inventory: inventory,
		community: community,
	}
}
