package inventory_item

const Schema = `
schema {
	query: Query
	mutation: Mutation
}

# An username consists of a sequence of a-z, 0-9 and _. Note that __ is not allowed.
scalar UUIDV4
# Currency represents a valid currency string representation, such as "USD" or "EUR"
scalar Currency

scalar Amount

scalar Time

scalar FirstName

scalar LastName

scalar Limit

scalar Username

scalar EmailAddress

scalar Ed25519PublicKey

scalar ProfilePicture

type Query {
	FetchItems(Position: Time!, Limit: Limit!) : FetchInventoryItemsResponse!
}

enum FetchInventoryItemsResponseError {
	Unauthenticated
	Unauthorized
}

type FetchInventoryItemsResponse {
   Error: FetchInventoryItemsResponseError
   FetchedItems: [InventoryItemEntity]!
}

type Mutation {
	AddItem(input: AddInventoryItemInput!) : AddInventoryItemResponse!
}


type ProperName {
    FirstName: FirstName!
    LastName: LastName!
}

input MoneyInput {
    AmountInput: Amount!
    CurrencyInput: Currency!
}

type Money {
	Amount: Amount!
	Currency: Currency!
}


input AddInventoryItemInput {
	Location: String
	Money: MoneyInput!
	Description: String
}

type MetadataEntity {
	ProperName:   ProperName!
	ProfileImage: ProfilePicture
}

type Member {
	ID: UUIDV4!
	CreatedAt: Time!
	VerifiedEmailAddress:  Boolean!
	Username:              Username!
	EmailAddress:          EmailAddress!
	Metadata:              MetadataEntity!
	MemberAccessPublicKey: Ed25519PublicKey
	AccessTokenID:         UUIDV4
	Admin:                 Boolean!
	Verified: Boolean!
}

type MemberResponse {
	ID: UUIDV4!
	Username: Username!
	Metadata: MetadataEntity!
}

type InventoryItemEntity {
	Location: String
	CreationDate: Time!
	OfferedBy: MemberResponse!
	Money: Money!
	Description: String
}


type AddInventoryItemResponse {
		InventoryItem: InventoryItemEntity
		Error: AddInventoryItemResponseError
}

enum AddInventoryItemResponseError {
	Unauthenticated
	Unauthorized
}
`
