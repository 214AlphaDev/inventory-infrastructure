package inventory_infrastructure_go

import (
	gql "github.com/graph-gophers/graphql-go"
	cd "github.com/214alphadev/community-domain-go"
	gqlh "github.com/214alphadev/graphql-handler"
	"github.com/214alphadev/inventory-infrastructure-go/api/inventory_item"
	"github.com/214alphadev/inventory-infrastructure-go/inventory_domain"
)

func NewInventoryApi(community cd.CommunityInterface, logger gqlh.Logger, inventory inventory_domain.InventoryInterface) (*gqlh.Handler, error) {

	resolver := inventory_item.NewInventoryItemResolver(inventory, community)

	schema, err := gql.ParseSchema(inventory_item.Schema, resolver)
	if err != nil {
		return nil, err
	}

	return gqlh.NewHandler(schema, logger)

}
