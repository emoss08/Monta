// Code generated by entc, DO NOT EDIT.

package customeremailprofile

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/emoss08/trenova/internal/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldLTE(FieldID, id))
}

// BusinessUnitID applies equality check predicate on the "business_unit_id" field. It's identical to BusinessUnitIDEQ.
func BusinessUnitID(v uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEQ(FieldBusinessUnitID, v))
}

// OrganizationID applies equality check predicate on the "organization_id" field. It's identical to OrganizationIDEQ.
func OrganizationID(v uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEQ(FieldOrganizationID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEQ(FieldUpdatedAt, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEQ(FieldVersion, v))
}

// CustomerID applies equality check predicate on the "customer_id" field. It's identical to CustomerIDEQ.
func CustomerID(v uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEQ(FieldCustomerID, v))
}

// Subject applies equality check predicate on the "subject" field. It's identical to SubjectEQ.
func Subject(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEQ(FieldSubject, v))
}

// EmailProfileID applies equality check predicate on the "email_profile_id" field. It's identical to EmailProfileIDEQ.
func EmailProfileID(v uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEQ(FieldEmailProfileID, v))
}

// EmailRecipients applies equality check predicate on the "email_recipients" field. It's identical to EmailRecipientsEQ.
func EmailRecipients(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEQ(FieldEmailRecipients, v))
}

// EmailCcRecipients applies equality check predicate on the "email_cc_recipients" field. It's identical to EmailCcRecipientsEQ.
func EmailCcRecipients(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEQ(FieldEmailCcRecipients, v))
}

// AttachmentName applies equality check predicate on the "attachment_name" field. It's identical to AttachmentNameEQ.
func AttachmentName(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEQ(FieldAttachmentName, v))
}

// BusinessUnitIDEQ applies the EQ predicate on the "business_unit_id" field.
func BusinessUnitIDEQ(v uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDNEQ applies the NEQ predicate on the "business_unit_id" field.
func BusinessUnitIDNEQ(v uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDIn applies the In predicate on the "business_unit_id" field.
func BusinessUnitIDIn(vs ...uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldIn(FieldBusinessUnitID, vs...))
}

// BusinessUnitIDNotIn applies the NotIn predicate on the "business_unit_id" field.
func BusinessUnitIDNotIn(vs ...uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNotIn(FieldBusinessUnitID, vs...))
}

// OrganizationIDEQ applies the EQ predicate on the "organization_id" field.
func OrganizationIDEQ(v uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEQ(FieldOrganizationID, v))
}

// OrganizationIDNEQ applies the NEQ predicate on the "organization_id" field.
func OrganizationIDNEQ(v uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNEQ(FieldOrganizationID, v))
}

// OrganizationIDIn applies the In predicate on the "organization_id" field.
func OrganizationIDIn(vs ...uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldIn(FieldOrganizationID, vs...))
}

// OrganizationIDNotIn applies the NotIn predicate on the "organization_id" field.
func OrganizationIDNotIn(vs ...uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNotIn(FieldOrganizationID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldLTE(FieldUpdatedAt, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldLTE(FieldVersion, v))
}

// CustomerIDEQ applies the EQ predicate on the "customer_id" field.
func CustomerIDEQ(v uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEQ(FieldCustomerID, v))
}

// CustomerIDNEQ applies the NEQ predicate on the "customer_id" field.
func CustomerIDNEQ(v uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNEQ(FieldCustomerID, v))
}

// CustomerIDIn applies the In predicate on the "customer_id" field.
func CustomerIDIn(vs ...uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldIn(FieldCustomerID, vs...))
}

// CustomerIDNotIn applies the NotIn predicate on the "customer_id" field.
func CustomerIDNotIn(vs ...uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNotIn(FieldCustomerID, vs...))
}

// SubjectEQ applies the EQ predicate on the "subject" field.
func SubjectEQ(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEQ(FieldSubject, v))
}

// SubjectNEQ applies the NEQ predicate on the "subject" field.
func SubjectNEQ(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNEQ(FieldSubject, v))
}

// SubjectIn applies the In predicate on the "subject" field.
func SubjectIn(vs ...string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldIn(FieldSubject, vs...))
}

// SubjectNotIn applies the NotIn predicate on the "subject" field.
func SubjectNotIn(vs ...string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNotIn(FieldSubject, vs...))
}

// SubjectGT applies the GT predicate on the "subject" field.
func SubjectGT(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldGT(FieldSubject, v))
}

// SubjectGTE applies the GTE predicate on the "subject" field.
func SubjectGTE(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldGTE(FieldSubject, v))
}

// SubjectLT applies the LT predicate on the "subject" field.
func SubjectLT(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldLT(FieldSubject, v))
}

// SubjectLTE applies the LTE predicate on the "subject" field.
func SubjectLTE(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldLTE(FieldSubject, v))
}

// SubjectContains applies the Contains predicate on the "subject" field.
func SubjectContains(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldContains(FieldSubject, v))
}

// SubjectHasPrefix applies the HasPrefix predicate on the "subject" field.
func SubjectHasPrefix(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldHasPrefix(FieldSubject, v))
}

// SubjectHasSuffix applies the HasSuffix predicate on the "subject" field.
func SubjectHasSuffix(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldHasSuffix(FieldSubject, v))
}

// SubjectIsNil applies the IsNil predicate on the "subject" field.
func SubjectIsNil() predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldIsNull(FieldSubject))
}

// SubjectNotNil applies the NotNil predicate on the "subject" field.
func SubjectNotNil() predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNotNull(FieldSubject))
}

// SubjectEqualFold applies the EqualFold predicate on the "subject" field.
func SubjectEqualFold(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEqualFold(FieldSubject, v))
}

// SubjectContainsFold applies the ContainsFold predicate on the "subject" field.
func SubjectContainsFold(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldContainsFold(FieldSubject, v))
}

// EmailProfileIDEQ applies the EQ predicate on the "email_profile_id" field.
func EmailProfileIDEQ(v uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEQ(FieldEmailProfileID, v))
}

// EmailProfileIDNEQ applies the NEQ predicate on the "email_profile_id" field.
func EmailProfileIDNEQ(v uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNEQ(FieldEmailProfileID, v))
}

// EmailProfileIDIn applies the In predicate on the "email_profile_id" field.
func EmailProfileIDIn(vs ...uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldIn(FieldEmailProfileID, vs...))
}

// EmailProfileIDNotIn applies the NotIn predicate on the "email_profile_id" field.
func EmailProfileIDNotIn(vs ...uuid.UUID) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNotIn(FieldEmailProfileID, vs...))
}

// EmailProfileIDIsNil applies the IsNil predicate on the "email_profile_id" field.
func EmailProfileIDIsNil() predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldIsNull(FieldEmailProfileID))
}

// EmailProfileIDNotNil applies the NotNil predicate on the "email_profile_id" field.
func EmailProfileIDNotNil() predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNotNull(FieldEmailProfileID))
}

// EmailRecipientsEQ applies the EQ predicate on the "email_recipients" field.
func EmailRecipientsEQ(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEQ(FieldEmailRecipients, v))
}

// EmailRecipientsNEQ applies the NEQ predicate on the "email_recipients" field.
func EmailRecipientsNEQ(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNEQ(FieldEmailRecipients, v))
}

// EmailRecipientsIn applies the In predicate on the "email_recipients" field.
func EmailRecipientsIn(vs ...string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldIn(FieldEmailRecipients, vs...))
}

// EmailRecipientsNotIn applies the NotIn predicate on the "email_recipients" field.
func EmailRecipientsNotIn(vs ...string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNotIn(FieldEmailRecipients, vs...))
}

// EmailRecipientsGT applies the GT predicate on the "email_recipients" field.
func EmailRecipientsGT(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldGT(FieldEmailRecipients, v))
}

// EmailRecipientsGTE applies the GTE predicate on the "email_recipients" field.
func EmailRecipientsGTE(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldGTE(FieldEmailRecipients, v))
}

// EmailRecipientsLT applies the LT predicate on the "email_recipients" field.
func EmailRecipientsLT(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldLT(FieldEmailRecipients, v))
}

// EmailRecipientsLTE applies the LTE predicate on the "email_recipients" field.
func EmailRecipientsLTE(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldLTE(FieldEmailRecipients, v))
}

// EmailRecipientsContains applies the Contains predicate on the "email_recipients" field.
func EmailRecipientsContains(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldContains(FieldEmailRecipients, v))
}

// EmailRecipientsHasPrefix applies the HasPrefix predicate on the "email_recipients" field.
func EmailRecipientsHasPrefix(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldHasPrefix(FieldEmailRecipients, v))
}

// EmailRecipientsHasSuffix applies the HasSuffix predicate on the "email_recipients" field.
func EmailRecipientsHasSuffix(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldHasSuffix(FieldEmailRecipients, v))
}

// EmailRecipientsEqualFold applies the EqualFold predicate on the "email_recipients" field.
func EmailRecipientsEqualFold(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEqualFold(FieldEmailRecipients, v))
}

// EmailRecipientsContainsFold applies the ContainsFold predicate on the "email_recipients" field.
func EmailRecipientsContainsFold(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldContainsFold(FieldEmailRecipients, v))
}

// EmailCcRecipientsEQ applies the EQ predicate on the "email_cc_recipients" field.
func EmailCcRecipientsEQ(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEQ(FieldEmailCcRecipients, v))
}

// EmailCcRecipientsNEQ applies the NEQ predicate on the "email_cc_recipients" field.
func EmailCcRecipientsNEQ(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNEQ(FieldEmailCcRecipients, v))
}

// EmailCcRecipientsIn applies the In predicate on the "email_cc_recipients" field.
func EmailCcRecipientsIn(vs ...string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldIn(FieldEmailCcRecipients, vs...))
}

// EmailCcRecipientsNotIn applies the NotIn predicate on the "email_cc_recipients" field.
func EmailCcRecipientsNotIn(vs ...string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNotIn(FieldEmailCcRecipients, vs...))
}

// EmailCcRecipientsGT applies the GT predicate on the "email_cc_recipients" field.
func EmailCcRecipientsGT(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldGT(FieldEmailCcRecipients, v))
}

// EmailCcRecipientsGTE applies the GTE predicate on the "email_cc_recipients" field.
func EmailCcRecipientsGTE(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldGTE(FieldEmailCcRecipients, v))
}

// EmailCcRecipientsLT applies the LT predicate on the "email_cc_recipients" field.
func EmailCcRecipientsLT(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldLT(FieldEmailCcRecipients, v))
}

// EmailCcRecipientsLTE applies the LTE predicate on the "email_cc_recipients" field.
func EmailCcRecipientsLTE(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldLTE(FieldEmailCcRecipients, v))
}

// EmailCcRecipientsContains applies the Contains predicate on the "email_cc_recipients" field.
func EmailCcRecipientsContains(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldContains(FieldEmailCcRecipients, v))
}

// EmailCcRecipientsHasPrefix applies the HasPrefix predicate on the "email_cc_recipients" field.
func EmailCcRecipientsHasPrefix(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldHasPrefix(FieldEmailCcRecipients, v))
}

// EmailCcRecipientsHasSuffix applies the HasSuffix predicate on the "email_cc_recipients" field.
func EmailCcRecipientsHasSuffix(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldHasSuffix(FieldEmailCcRecipients, v))
}

// EmailCcRecipientsIsNil applies the IsNil predicate on the "email_cc_recipients" field.
func EmailCcRecipientsIsNil() predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldIsNull(FieldEmailCcRecipients))
}

// EmailCcRecipientsNotNil applies the NotNil predicate on the "email_cc_recipients" field.
func EmailCcRecipientsNotNil() predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNotNull(FieldEmailCcRecipients))
}

// EmailCcRecipientsEqualFold applies the EqualFold predicate on the "email_cc_recipients" field.
func EmailCcRecipientsEqualFold(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEqualFold(FieldEmailCcRecipients, v))
}

// EmailCcRecipientsContainsFold applies the ContainsFold predicate on the "email_cc_recipients" field.
func EmailCcRecipientsContainsFold(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldContainsFold(FieldEmailCcRecipients, v))
}

// AttachmentNameEQ applies the EQ predicate on the "attachment_name" field.
func AttachmentNameEQ(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEQ(FieldAttachmentName, v))
}

// AttachmentNameNEQ applies the NEQ predicate on the "attachment_name" field.
func AttachmentNameNEQ(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNEQ(FieldAttachmentName, v))
}

// AttachmentNameIn applies the In predicate on the "attachment_name" field.
func AttachmentNameIn(vs ...string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldIn(FieldAttachmentName, vs...))
}

// AttachmentNameNotIn applies the NotIn predicate on the "attachment_name" field.
func AttachmentNameNotIn(vs ...string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNotIn(FieldAttachmentName, vs...))
}

// AttachmentNameGT applies the GT predicate on the "attachment_name" field.
func AttachmentNameGT(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldGT(FieldAttachmentName, v))
}

// AttachmentNameGTE applies the GTE predicate on the "attachment_name" field.
func AttachmentNameGTE(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldGTE(FieldAttachmentName, v))
}

// AttachmentNameLT applies the LT predicate on the "attachment_name" field.
func AttachmentNameLT(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldLT(FieldAttachmentName, v))
}

// AttachmentNameLTE applies the LTE predicate on the "attachment_name" field.
func AttachmentNameLTE(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldLTE(FieldAttachmentName, v))
}

// AttachmentNameContains applies the Contains predicate on the "attachment_name" field.
func AttachmentNameContains(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldContains(FieldAttachmentName, v))
}

// AttachmentNameHasPrefix applies the HasPrefix predicate on the "attachment_name" field.
func AttachmentNameHasPrefix(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldHasPrefix(FieldAttachmentName, v))
}

// AttachmentNameHasSuffix applies the HasSuffix predicate on the "attachment_name" field.
func AttachmentNameHasSuffix(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldHasSuffix(FieldAttachmentName, v))
}

// AttachmentNameIsNil applies the IsNil predicate on the "attachment_name" field.
func AttachmentNameIsNil() predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldIsNull(FieldAttachmentName))
}

// AttachmentNameNotNil applies the NotNil predicate on the "attachment_name" field.
func AttachmentNameNotNil() predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNotNull(FieldAttachmentName))
}

// AttachmentNameEqualFold applies the EqualFold predicate on the "attachment_name" field.
func AttachmentNameEqualFold(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEqualFold(FieldAttachmentName, v))
}

// AttachmentNameContainsFold applies the ContainsFold predicate on the "attachment_name" field.
func AttachmentNameContainsFold(v string) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldContainsFold(FieldAttachmentName, v))
}

// EmailFormatEQ applies the EQ predicate on the "email_format" field.
func EmailFormatEQ(v EmailFormat) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldEQ(FieldEmailFormat, v))
}

// EmailFormatNEQ applies the NEQ predicate on the "email_format" field.
func EmailFormatNEQ(v EmailFormat) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNEQ(FieldEmailFormat, v))
}

// EmailFormatIn applies the In predicate on the "email_format" field.
func EmailFormatIn(vs ...EmailFormat) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldIn(FieldEmailFormat, vs...))
}

// EmailFormatNotIn applies the NotIn predicate on the "email_format" field.
func EmailFormatNotIn(vs ...EmailFormat) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.FieldNotIn(FieldEmailFormat, vs...))
}

// HasBusinessUnit applies the HasEdge predicate on the "business_unit" edge.
func HasBusinessUnit() predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessUnitWith applies the HasEdge predicate on the "business_unit" edge with a given conditions (other predicates).
func HasBusinessUnitWith(preds ...predicate.BusinessUnit) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(func(s *sql.Selector) {
		step := newBusinessUnitStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCustomer applies the HasEdge predicate on the "customer" edge.
func HasCustomer() predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, CustomerTable, CustomerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCustomerWith applies the HasEdge predicate on the "customer" edge with a given conditions (other predicates).
func HasCustomerWith(preds ...predicate.Customer) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(func(s *sql.Selector) {
		step := newCustomerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEmailProfile applies the HasEdge predicate on the "email_profile" edge.
func HasEmailProfile() predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, EmailProfileTable, EmailProfileColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEmailProfileWith applies the HasEdge predicate on the "email_profile" edge with a given conditions (other predicates).
func HasEmailProfileWith(preds ...predicate.EmailProfile) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(func(s *sql.Selector) {
		step := newEmailProfileStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CustomerEmailProfile) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CustomerEmailProfile) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CustomerEmailProfile) predicate.CustomerEmailProfile {
	return predicate.CustomerEmailProfile(sql.NotPredicates(p))
}