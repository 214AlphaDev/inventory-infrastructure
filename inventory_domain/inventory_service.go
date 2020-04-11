package inventory_domain

import (
	"github.com/214alphadev/inventory-infrastructure-go/api/entities"
	"github.com/214alphadev/inventory-infrastructure-go/api/types"
)

type inventoryService struct {
	inventoryRepository InventoryItemRepositoryInterface
}

type InventoryItemRepositoryInterface interface {
	GetItems(args types.InventoryItemsQuery) ([]*entities.InventoryItemEntity, error)
	Save(item entities.InventoryItemEntity) error
}
