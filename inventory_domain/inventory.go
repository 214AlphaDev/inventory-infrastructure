package inventory_domain

import (
	"github.com/214alphadev/inventory-infrastructure-go/api/entities"
	"github.com/214alphadev/inventory-infrastructure-go/api/types"
	repository "github.com/214alphadev/inventory-infrastructure-go/repositories"
)

type Inventory struct {
	inventoryService *inventoryService
}

func (i *Inventory) GetItems(args types.InventoryItemsQuery) ([]*entities.InventoryItemEntity, error) {
	return i.inventoryService.inventoryRepository.GetItems(args)
}

func (i *Inventory) Save(item entities.InventoryItemEntity) error {
	return i.inventoryService.inventoryRepository.Save(item)
}

type InventoryInterface interface {
	Save(item entities.InventoryItemEntity) error
	GetItems(args types.InventoryItemsQuery) ([]repository.InventoryItem, error)
}

func NewInventory(inventoryRepository InventoryItemRepositoryInterface) *Inventory {
	return &Inventory{
		inventoryService: &inventoryService{
			inventoryRepository: inventoryRepository,
		},
	}
}
