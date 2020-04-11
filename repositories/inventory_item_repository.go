package inventory_item_repository

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/214alphadev/inventory-infrastructure-go/api/entities"
	types "github.com/214alphadev/inventory-infrastructure-go/api/types"
)

type InventoryItem struct {
	Location      *string   `gorm:"null"`
	FirstName     string    `gorm:"non null"`
	LastName      string    `gorm:"non null"`
	CreationDate  time.Time `gorm:"non null"`
	OfferedBy     string    `gorm:"non null"`
	MoneyAmount   uint32    `gorm:"non null"`
	MoneyCurrency string    `gorm:"non null"`
	Description   *string   `gorm:"null"`
}

type InventoryItemRepository struct {
	db *gorm.DB
}

func (i *InventoryItemRepository) GetItems(args types.InventoryItemsQuery) ([]InventoryItem, error) {
	fetchedItems := []InventoryItem{}
	err := i.db.Find(&fetchedItems).Where(`creation_date > ?`, args.Position).Limit(args.Limit).Find(&fetchedItems).Error
	if err != nil {
		return nil, err
	}
	return fetchedItems, nil
}

func (i *InventoryItemRepository) Save(item entities.InventoryItemEntity) error {
	itemGorm := InventoryItem{
		Location:      item.Location(),
		CreationDate:  item.CreationDate().Time,
		OfferedBy:     item.OfferedBy().ID().String(),
		MoneyAmount:   item.Money().Amount().Amount.Amount(),
		MoneyCurrency: item.Money().Currency().Currency.Currency(),
		Description:   item.Description(),
	}
	return i.db.Create(itemGorm).Error
}

func NewInventoryItemRepository(db *gorm.DB) (*InventoryItemRepository, error) {

	if err := db.AutoMigrate(&InventoryItem{}).Error; err != nil {
		return nil, err
	}

	return &InventoryItemRepository{
		db: db,
	}, nil
}
