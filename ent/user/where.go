// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// BusinessUnitID applies equality check predicate on the "business_unit_id" field. It's identical to BusinessUnitIDEQ.
func BusinessUnitID(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBusinessUnitID, v))
}

// OrganizationID applies equality check predicate on the "organization_id" field. It's identical to OrganizationIDEQ.
func OrganizationID(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldOrganizationID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// ProfilePicURL applies equality check predicate on the "profile_pic_url" field. It's identical to ProfilePicURLEQ.
func ProfilePicURL(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldProfilePicURL, v))
}

// ThumbnailURL applies equality check predicate on the "thumbnail_url" field. It's identical to ThumbnailURLEQ.
func ThumbnailURL(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldThumbnailURL, v))
}

// PhoneNumber applies equality check predicate on the "phone_number" field. It's identical to PhoneNumberEQ.
func PhoneNumber(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPhoneNumber, v))
}

// IsAdmin applies equality check predicate on the "is_admin" field. It's identical to IsAdminEQ.
func IsAdmin(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIsAdmin, v))
}

// IsSuperAdmin applies equality check predicate on the "is_super_admin" field. It's identical to IsSuperAdminEQ.
func IsSuperAdmin(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIsSuperAdmin, v))
}

// LastLogin applies equality check predicate on the "last_login" field. It's identical to LastLoginEQ.
func LastLogin(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLastLogin, v))
}

// BusinessUnitIDEQ applies the EQ predicate on the "business_unit_id" field.
func BusinessUnitIDEQ(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDNEQ applies the NEQ predicate on the "business_unit_id" field.
func BusinessUnitIDNEQ(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDIn applies the In predicate on the "business_unit_id" field.
func BusinessUnitIDIn(vs ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldIn(FieldBusinessUnitID, vs...))
}

// BusinessUnitIDNotIn applies the NotIn predicate on the "business_unit_id" field.
func BusinessUnitIDNotIn(vs ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldBusinessUnitID, vs...))
}

// OrganizationIDEQ applies the EQ predicate on the "organization_id" field.
func OrganizationIDEQ(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldOrganizationID, v))
}

// OrganizationIDNEQ applies the NEQ predicate on the "organization_id" field.
func OrganizationIDNEQ(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldOrganizationID, v))
}

// OrganizationIDIn applies the In predicate on the "organization_id" field.
func OrganizationIDIn(vs ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldIn(FieldOrganizationID, vs...))
}

// OrganizationIDNotIn applies the NotIn predicate on the "organization_id" field.
func OrganizationIDNotIn(vs ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldOrganizationID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdatedAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.User {
	return predicate.User(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.User {
	return predicate.User(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldStatus, vs...))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldName, v))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUsername, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPassword, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmail, v))
}

// TimezoneEQ applies the EQ predicate on the "timezone" field.
func TimezoneEQ(v Timezone) predicate.User {
	return predicate.User(sql.FieldEQ(FieldTimezone, v))
}

// TimezoneNEQ applies the NEQ predicate on the "timezone" field.
func TimezoneNEQ(v Timezone) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldTimezone, v))
}

// TimezoneIn applies the In predicate on the "timezone" field.
func TimezoneIn(vs ...Timezone) predicate.User {
	return predicate.User(sql.FieldIn(FieldTimezone, vs...))
}

// TimezoneNotIn applies the NotIn predicate on the "timezone" field.
func TimezoneNotIn(vs ...Timezone) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldTimezone, vs...))
}

// ProfilePicURLEQ applies the EQ predicate on the "profile_pic_url" field.
func ProfilePicURLEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldProfilePicURL, v))
}

// ProfilePicURLNEQ applies the NEQ predicate on the "profile_pic_url" field.
func ProfilePicURLNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldProfilePicURL, v))
}

// ProfilePicURLIn applies the In predicate on the "profile_pic_url" field.
func ProfilePicURLIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldProfilePicURL, vs...))
}

// ProfilePicURLNotIn applies the NotIn predicate on the "profile_pic_url" field.
func ProfilePicURLNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldProfilePicURL, vs...))
}

// ProfilePicURLGT applies the GT predicate on the "profile_pic_url" field.
func ProfilePicURLGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldProfilePicURL, v))
}

// ProfilePicURLGTE applies the GTE predicate on the "profile_pic_url" field.
func ProfilePicURLGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldProfilePicURL, v))
}

// ProfilePicURLLT applies the LT predicate on the "profile_pic_url" field.
func ProfilePicURLLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldProfilePicURL, v))
}

// ProfilePicURLLTE applies the LTE predicate on the "profile_pic_url" field.
func ProfilePicURLLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldProfilePicURL, v))
}

// ProfilePicURLContains applies the Contains predicate on the "profile_pic_url" field.
func ProfilePicURLContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldProfilePicURL, v))
}

// ProfilePicURLHasPrefix applies the HasPrefix predicate on the "profile_pic_url" field.
func ProfilePicURLHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldProfilePicURL, v))
}

// ProfilePicURLHasSuffix applies the HasSuffix predicate on the "profile_pic_url" field.
func ProfilePicURLHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldProfilePicURL, v))
}

// ProfilePicURLIsNil applies the IsNil predicate on the "profile_pic_url" field.
func ProfilePicURLIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldProfilePicURL))
}

// ProfilePicURLNotNil applies the NotNil predicate on the "profile_pic_url" field.
func ProfilePicURLNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldProfilePicURL))
}

// ProfilePicURLEqualFold applies the EqualFold predicate on the "profile_pic_url" field.
func ProfilePicURLEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldProfilePicURL, v))
}

// ProfilePicURLContainsFold applies the ContainsFold predicate on the "profile_pic_url" field.
func ProfilePicURLContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldProfilePicURL, v))
}

// ThumbnailURLEQ applies the EQ predicate on the "thumbnail_url" field.
func ThumbnailURLEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldThumbnailURL, v))
}

// ThumbnailURLNEQ applies the NEQ predicate on the "thumbnail_url" field.
func ThumbnailURLNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldThumbnailURL, v))
}

// ThumbnailURLIn applies the In predicate on the "thumbnail_url" field.
func ThumbnailURLIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldThumbnailURL, vs...))
}

// ThumbnailURLNotIn applies the NotIn predicate on the "thumbnail_url" field.
func ThumbnailURLNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldThumbnailURL, vs...))
}

// ThumbnailURLGT applies the GT predicate on the "thumbnail_url" field.
func ThumbnailURLGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldThumbnailURL, v))
}

// ThumbnailURLGTE applies the GTE predicate on the "thumbnail_url" field.
func ThumbnailURLGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldThumbnailURL, v))
}

// ThumbnailURLLT applies the LT predicate on the "thumbnail_url" field.
func ThumbnailURLLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldThumbnailURL, v))
}

// ThumbnailURLLTE applies the LTE predicate on the "thumbnail_url" field.
func ThumbnailURLLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldThumbnailURL, v))
}

// ThumbnailURLContains applies the Contains predicate on the "thumbnail_url" field.
func ThumbnailURLContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldThumbnailURL, v))
}

// ThumbnailURLHasPrefix applies the HasPrefix predicate on the "thumbnail_url" field.
func ThumbnailURLHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldThumbnailURL, v))
}

// ThumbnailURLHasSuffix applies the HasSuffix predicate on the "thumbnail_url" field.
func ThumbnailURLHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldThumbnailURL, v))
}

// ThumbnailURLIsNil applies the IsNil predicate on the "thumbnail_url" field.
func ThumbnailURLIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldThumbnailURL))
}

// ThumbnailURLNotNil applies the NotNil predicate on the "thumbnail_url" field.
func ThumbnailURLNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldThumbnailURL))
}

// ThumbnailURLEqualFold applies the EqualFold predicate on the "thumbnail_url" field.
func ThumbnailURLEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldThumbnailURL, v))
}

// ThumbnailURLContainsFold applies the ContainsFold predicate on the "thumbnail_url" field.
func ThumbnailURLContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldThumbnailURL, v))
}

// PhoneNumberEQ applies the EQ predicate on the "phone_number" field.
func PhoneNumberEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPhoneNumber, v))
}

// PhoneNumberNEQ applies the NEQ predicate on the "phone_number" field.
func PhoneNumberNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPhoneNumber, v))
}

// PhoneNumberIn applies the In predicate on the "phone_number" field.
func PhoneNumberIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPhoneNumber, vs...))
}

// PhoneNumberNotIn applies the NotIn predicate on the "phone_number" field.
func PhoneNumberNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPhoneNumber, vs...))
}

// PhoneNumberGT applies the GT predicate on the "phone_number" field.
func PhoneNumberGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPhoneNumber, v))
}

// PhoneNumberGTE applies the GTE predicate on the "phone_number" field.
func PhoneNumberGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPhoneNumber, v))
}

// PhoneNumberLT applies the LT predicate on the "phone_number" field.
func PhoneNumberLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPhoneNumber, v))
}

// PhoneNumberLTE applies the LTE predicate on the "phone_number" field.
func PhoneNumberLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPhoneNumber, v))
}

// PhoneNumberContains applies the Contains predicate on the "phone_number" field.
func PhoneNumberContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPhoneNumber, v))
}

// PhoneNumberHasPrefix applies the HasPrefix predicate on the "phone_number" field.
func PhoneNumberHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPhoneNumber, v))
}

// PhoneNumberHasSuffix applies the HasSuffix predicate on the "phone_number" field.
func PhoneNumberHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPhoneNumber, v))
}

// PhoneNumberIsNil applies the IsNil predicate on the "phone_number" field.
func PhoneNumberIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldPhoneNumber))
}

// PhoneNumberNotNil applies the NotNil predicate on the "phone_number" field.
func PhoneNumberNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldPhoneNumber))
}

// PhoneNumberEqualFold applies the EqualFold predicate on the "phone_number" field.
func PhoneNumberEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPhoneNumber, v))
}

// PhoneNumberContainsFold applies the ContainsFold predicate on the "phone_number" field.
func PhoneNumberContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPhoneNumber, v))
}

// IsAdminEQ applies the EQ predicate on the "is_admin" field.
func IsAdminEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIsAdmin, v))
}

// IsAdminNEQ applies the NEQ predicate on the "is_admin" field.
func IsAdminNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldIsAdmin, v))
}

// IsSuperAdminEQ applies the EQ predicate on the "is_super_admin" field.
func IsSuperAdminEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIsSuperAdmin, v))
}

// IsSuperAdminNEQ applies the NEQ predicate on the "is_super_admin" field.
func IsSuperAdminNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldIsSuperAdmin, v))
}

// LastLoginEQ applies the EQ predicate on the "last_login" field.
func LastLoginEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLastLogin, v))
}

// LastLoginNEQ applies the NEQ predicate on the "last_login" field.
func LastLoginNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldLastLogin, v))
}

// LastLoginIn applies the In predicate on the "last_login" field.
func LastLoginIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldLastLogin, vs...))
}

// LastLoginNotIn applies the NotIn predicate on the "last_login" field.
func LastLoginNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldLastLogin, vs...))
}

// LastLoginGT applies the GT predicate on the "last_login" field.
func LastLoginGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldLastLogin, v))
}

// LastLoginGTE applies the GTE predicate on the "last_login" field.
func LastLoginGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldLastLogin, v))
}

// LastLoginLT applies the LT predicate on the "last_login" field.
func LastLoginLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldLastLogin, v))
}

// LastLoginLTE applies the LTE predicate on the "last_login" field.
func LastLoginLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldLastLogin, v))
}

// LastLoginIsNil applies the IsNil predicate on the "last_login" field.
func LastLoginIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldLastLogin))
}

// LastLoginNotNil applies the NotNil predicate on the "last_login" field.
func LastLoginNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldLastLogin))
}

// HasBusinessUnit applies the HasEdge predicate on the "business_unit" edge.
func HasBusinessUnit() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessUnitWith applies the HasEdge predicate on the "business_unit" edge with a given conditions (other predicates).
func HasBusinessUnitWith(preds ...predicate.BusinessUnit) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newBusinessUnitStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserFavorites applies the HasEdge predicate on the "user_favorites" edge.
func HasUserFavorites() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserFavoritesTable, UserFavoritesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserFavoritesWith applies the HasEdge predicate on the "user_favorites" edge with a given conditions (other predicates).
func HasUserFavoritesWith(preds ...predicate.UserFavorite) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newUserFavoritesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
