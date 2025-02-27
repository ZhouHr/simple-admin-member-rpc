//go:build tools && tools
// +build tools,tools

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-member-rpc/ent/member"
	"github.com/suyuan32/simple-admin-member-rpc/ent/memberrank"
	"github.com/suyuan32/simple-admin-member-rpc/ent/oauthprovider"
	"github.com/suyuan32/simple-admin-member-rpc/ent/schema"
	"github.com/suyuan32/simple-admin-member-rpc/ent/token"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	memberMixin := schema.Member{}.Mixin()
	memberMixinFields0 := memberMixin[0].Fields()
	_ = memberMixinFields0
	memberMixinFields1 := memberMixin[1].Fields()
	_ = memberMixinFields1
	memberFields := schema.Member{}.Fields()
	_ = memberFields
	// memberDescCreatedAt is the schema descriptor for created_at field.
	memberDescCreatedAt := memberMixinFields0[1].Descriptor()
	// member.DefaultCreatedAt holds the default value on creation for the created_at field.
	member.DefaultCreatedAt = memberDescCreatedAt.Default.(func() time.Time)
	// memberDescUpdatedAt is the schema descriptor for updated_at field.
	memberDescUpdatedAt := memberMixinFields0[2].Descriptor()
	// member.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	member.DefaultUpdatedAt = memberDescUpdatedAt.Default.(func() time.Time)
	// member.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	member.UpdateDefaultUpdatedAt = memberDescUpdatedAt.UpdateDefault.(func() time.Time)
	// memberDescStatus is the schema descriptor for status field.
	memberDescStatus := memberMixinFields1[0].Descriptor()
	// member.DefaultStatus holds the default value on creation for the status field.
	member.DefaultStatus = memberDescStatus.Default.(uint8)
	// memberDescRankID is the schema descriptor for rank_id field.
	memberDescRankID := memberFields[3].Descriptor()
	// member.DefaultRankID holds the default value on creation for the rank_id field.
	member.DefaultRankID = memberDescRankID.Default.(uint64)
	// memberDescAvatar is the schema descriptor for avatar field.
	memberDescAvatar := memberFields[6].Descriptor()
	// member.DefaultAvatar holds the default value on creation for the avatar field.
	member.DefaultAvatar = memberDescAvatar.Default.(string)
	// memberDescExpiredAt is the schema descriptor for expired_at field.
	memberDescExpiredAt := memberFields[8].Descriptor()
	// member.DefaultExpiredAt holds the default value on creation for the expired_at field.
	member.DefaultExpiredAt = memberDescExpiredAt.Default.(time.Time)
	// memberDescID is the schema descriptor for id field.
	memberDescID := memberMixinFields0[0].Descriptor()
	// member.DefaultID holds the default value on creation for the id field.
	member.DefaultID = memberDescID.Default.(func() uuid.UUID)
	memberrankMixin := schema.MemberRank{}.Mixin()
	memberrankMixinFields0 := memberrankMixin[0].Fields()
	_ = memberrankMixinFields0
	memberrankFields := schema.MemberRank{}.Fields()
	_ = memberrankFields
	// memberrankDescCreatedAt is the schema descriptor for created_at field.
	memberrankDescCreatedAt := memberrankMixinFields0[1].Descriptor()
	// memberrank.DefaultCreatedAt holds the default value on creation for the created_at field.
	memberrank.DefaultCreatedAt = memberrankDescCreatedAt.Default.(func() time.Time)
	// memberrankDescUpdatedAt is the schema descriptor for updated_at field.
	memberrankDescUpdatedAt := memberrankMixinFields0[2].Descriptor()
	// memberrank.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	memberrank.DefaultUpdatedAt = memberrankDescUpdatedAt.Default.(func() time.Time)
	// memberrank.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	memberrank.UpdateDefaultUpdatedAt = memberrankDescUpdatedAt.UpdateDefault.(func() time.Time)
	oauthproviderMixin := schema.OauthProvider{}.Mixin()
	oauthproviderMixinFields0 := oauthproviderMixin[0].Fields()
	_ = oauthproviderMixinFields0
	oauthproviderFields := schema.OauthProvider{}.Fields()
	_ = oauthproviderFields
	// oauthproviderDescCreatedAt is the schema descriptor for created_at field.
	oauthproviderDescCreatedAt := oauthproviderMixinFields0[1].Descriptor()
	// oauthprovider.DefaultCreatedAt holds the default value on creation for the created_at field.
	oauthprovider.DefaultCreatedAt = oauthproviderDescCreatedAt.Default.(func() time.Time)
	// oauthproviderDescUpdatedAt is the schema descriptor for updated_at field.
	oauthproviderDescUpdatedAt := oauthproviderMixinFields0[2].Descriptor()
	// oauthprovider.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	oauthprovider.DefaultUpdatedAt = oauthproviderDescUpdatedAt.Default.(func() time.Time)
	// oauthprovider.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	oauthprovider.UpdateDefaultUpdatedAt = oauthproviderDescUpdatedAt.UpdateDefault.(func() time.Time)
	tokenMixin := schema.Token{}.Mixin()
	tokenMixinFields0 := tokenMixin[0].Fields()
	_ = tokenMixinFields0
	tokenMixinFields1 := tokenMixin[1].Fields()
	_ = tokenMixinFields1
	tokenFields := schema.Token{}.Fields()
	_ = tokenFields
	// tokenDescCreatedAt is the schema descriptor for created_at field.
	tokenDescCreatedAt := tokenMixinFields0[1].Descriptor()
	// token.DefaultCreatedAt holds the default value on creation for the created_at field.
	token.DefaultCreatedAt = tokenDescCreatedAt.Default.(func() time.Time)
	// tokenDescUpdatedAt is the schema descriptor for updated_at field.
	tokenDescUpdatedAt := tokenMixinFields0[2].Descriptor()
	// token.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	token.DefaultUpdatedAt = tokenDescUpdatedAt.Default.(func() time.Time)
	// token.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	token.UpdateDefaultUpdatedAt = tokenDescUpdatedAt.UpdateDefault.(func() time.Time)
	// tokenDescStatus is the schema descriptor for status field.
	tokenDescStatus := tokenMixinFields1[0].Descriptor()
	// token.DefaultStatus holds the default value on creation for the status field.
	token.DefaultStatus = tokenDescStatus.Default.(uint8)
	// tokenDescUsername is the schema descriptor for username field.
	tokenDescUsername := tokenFields[2].Descriptor()
	// token.DefaultUsername holds the default value on creation for the username field.
	token.DefaultUsername = tokenDescUsername.Default.(string)
	// tokenDescID is the schema descriptor for id field.
	tokenDescID := tokenMixinFields0[0].Descriptor()
	// token.DefaultID holds the default value on creation for the id field.
	token.DefaultID = tokenDescID.Default.(func() uuid.UUID)
}
