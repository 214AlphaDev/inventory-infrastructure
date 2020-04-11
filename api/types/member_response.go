package types

import (
	"github.com/graph-gophers/graphql-go"
	uuid "github.com/satori/go.uuid"
	community "github.com/214alphadev/community-domain-go"
	"github.com/214alphadev/community-domain-go/value_objects"
	vo "github.com/214alphadev/community-domain-go/value_objects"
	"github.com/214alphadev/inventory-infrastructure-go/api/scalars"
	vo_local "github.com/214alphadev/inventory-infrastructure-go/api/value_objects"
)

type MetadataEntityExtemded struct {
	Metadata community.MetadataEntity
}

type ProperName struct {
	firstName scalars.FirstName
	lastName  scalars.LastName
}

func (p ProperName) FirstName() scalars.FirstName {
	return p.firstName
}

func (p ProperName) LastName() scalars.LastName {
	return p.lastName
}

func (m MetadataEntityExtemded) ProperName() ProperName {
	firstName, _ := vo_local.NewFirstName(m.Metadata.ProperName.FirstName())
	lastName, _ := vo_local.NewLastName(m.Metadata.ProperName.LastName())
	return ProperName{
		firstName: scalars.FirstName{FirstName: firstName},
		lastName:  scalars.LastName{LastName: lastName},
	}
}

func (m MetadataEntityExtemded) ProfileImage() *scalars.ProfilePicture {
	if m.Metadata.ProfileImage == nil {
		return &scalars.ProfilePicture{Base64String: value_objects.Base64String{}}
	}
	return &scalars.ProfilePicture{Base64String: *m.Metadata.ProfileImage}
}

type MemberResponse struct {
	id       uuid.UUID
	username vo.Username
	metadata community.MetadataEntity
}

func NewMemberResponse(id uuid.UUID, username vo.Username, metadata community.MetadataEntity) MemberResponse {
	return MemberResponse{
		id:       id,
		username: username,
		metadata: metadata,
	}
}

func (m MemberResponse) ID() scalars.UUIDV4 {
	return scalars.UUIDV4{UUID: m.id}
}

func (m MemberResponse) Username() scalars.Username {
	return scalars.Username{Username: m.username}
}

func (m MemberResponse) Metadata() MetadataEntityExtemded {
	return MetadataEntityExtemded{Metadata: m.metadata}
}

func (m MemberResponse) ProfileImage() scalars.ProfilePicture {
	return scalars.ProfilePicture{Base64String: *m.metadata.ProfileImage}
}

type MemberEntityExtended struct {
	Member community.MemberEntity
}

func (m MemberEntityExtended) ID() scalars.UUIDV4 {
	return scalars.UUIDV4{UUID: m.Member.ID}
}

func (m MemberEntityExtended) CreatedAt() graphql.Time {
	return graphql.Time{Time: m.Member.CreatedAt}
}

func (m MemberEntityExtended) VerifiedEmailAddress() bool {
	return m.Member.VerifiedEmailAddress
}

func (m MemberEntityExtended) Username() scalars.Username {
	return scalars.Username{Username: m.Member.Username}
}

func (m MemberEntityExtended) EmailAddress() scalars.EmailAddress {
	return scalars.EmailAddress{EmailAddress: m.Member.EmailAddress}
}

func (m MemberEntityExtended) Metadata() MetadataEntityExtemded {
	return MetadataEntityExtemded{Metadata: m.Member.Metadata}
}

func (m MemberEntityExtended) ProfileImage() scalars.ProfilePicture {
	return scalars.ProfilePicture{Base64String: *m.Member.Metadata.ProfileImage}
}

func (m MemberEntityExtended) MemberAccessPublicKey() *scalars.Ed25519PublicKey {
	return &scalars.Ed25519PublicKey{PublicKey: m.Member.MemberAccessPublicKey.Key()}
}

func (m MemberEntityExtended) AccessTokenID() *scalars.UUIDV4 {
	if m.Member.AccessTokenID == nil {
		return &scalars.UUIDV4{UUID: uuid.UUID{}}
	}
	return &scalars.UUIDV4{UUID: *m.Member.AccessTokenID}
}

func (m MemberEntityExtended) Admin() bool {
	return m.Member.Admin
}

func (m MemberEntityExtended) Verified() bool {
	return m.Member.Verified
}
