package types

import "time"

type InventoryItemsQuery struct {
	Position time.Time
	Limit    uint32
}
