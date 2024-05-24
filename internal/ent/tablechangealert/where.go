// Code generated by entc, DO NOT EDIT.

package tablechangealert

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/emoss08/trenova/internal/ent/predicate"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldLTE(FieldID, id))
}

// BusinessUnitID applies equality check predicate on the "business_unit_id" field. It's identical to BusinessUnitIDEQ.
func BusinessUnitID(v uuid.UUID) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldBusinessUnitID, v))
}

// OrganizationID applies equality check predicate on the "organization_id" field. It's identical to OrganizationIDEQ.
func OrganizationID(v uuid.UUID) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldOrganizationID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldUpdatedAt, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldVersion, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldName, v))
}

// TopicName applies equality check predicate on the "topic_name" field. It's identical to TopicNameEQ.
func TopicName(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldTopicName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldDescription, v))
}

// CustomSubject applies equality check predicate on the "custom_subject" field. It's identical to CustomSubjectEQ.
func CustomSubject(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldCustomSubject, v))
}

// EmailRecipients applies equality check predicate on the "email_recipients" field. It's identical to EmailRecipientsEQ.
func EmailRecipients(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldEmailRecipients, v))
}

// EffectiveDate applies equality check predicate on the "effective_date" field. It's identical to EffectiveDateEQ.
func EffectiveDate(v *pgtype.Date) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldEffectiveDate, v))
}

// ExpirationDate applies equality check predicate on the "expiration_date" field. It's identical to ExpirationDateEQ.
func ExpirationDate(v *pgtype.Date) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldExpirationDate, v))
}

// BusinessUnitIDEQ applies the EQ predicate on the "business_unit_id" field.
func BusinessUnitIDEQ(v uuid.UUID) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDNEQ applies the NEQ predicate on the "business_unit_id" field.
func BusinessUnitIDNEQ(v uuid.UUID) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDIn applies the In predicate on the "business_unit_id" field.
func BusinessUnitIDIn(vs ...uuid.UUID) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldIn(FieldBusinessUnitID, vs...))
}

// BusinessUnitIDNotIn applies the NotIn predicate on the "business_unit_id" field.
func BusinessUnitIDNotIn(vs ...uuid.UUID) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNotIn(FieldBusinessUnitID, vs...))
}

// OrganizationIDEQ applies the EQ predicate on the "organization_id" field.
func OrganizationIDEQ(v uuid.UUID) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldOrganizationID, v))
}

// OrganizationIDNEQ applies the NEQ predicate on the "organization_id" field.
func OrganizationIDNEQ(v uuid.UUID) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNEQ(FieldOrganizationID, v))
}

// OrganizationIDIn applies the In predicate on the "organization_id" field.
func OrganizationIDIn(vs ...uuid.UUID) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldIn(FieldOrganizationID, vs...))
}

// OrganizationIDNotIn applies the NotIn predicate on the "organization_id" field.
func OrganizationIDNotIn(vs ...uuid.UUID) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNotIn(FieldOrganizationID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldLTE(FieldUpdatedAt, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldLTE(FieldVersion, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNotIn(FieldStatus, vs...))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldContainsFold(FieldName, v))
}

// DatabaseActionEQ applies the EQ predicate on the "database_action" field.
func DatabaseActionEQ(v DatabaseAction) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldDatabaseAction, v))
}

// DatabaseActionNEQ applies the NEQ predicate on the "database_action" field.
func DatabaseActionNEQ(v DatabaseAction) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNEQ(FieldDatabaseAction, v))
}

// DatabaseActionIn applies the In predicate on the "database_action" field.
func DatabaseActionIn(vs ...DatabaseAction) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldIn(FieldDatabaseAction, vs...))
}

// DatabaseActionNotIn applies the NotIn predicate on the "database_action" field.
func DatabaseActionNotIn(vs ...DatabaseAction) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNotIn(FieldDatabaseAction, vs...))
}

// TopicNameEQ applies the EQ predicate on the "topic_name" field.
func TopicNameEQ(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldTopicName, v))
}

// TopicNameNEQ applies the NEQ predicate on the "topic_name" field.
func TopicNameNEQ(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNEQ(FieldTopicName, v))
}

// TopicNameIn applies the In predicate on the "topic_name" field.
func TopicNameIn(vs ...string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldIn(FieldTopicName, vs...))
}

// TopicNameNotIn applies the NotIn predicate on the "topic_name" field.
func TopicNameNotIn(vs ...string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNotIn(FieldTopicName, vs...))
}

// TopicNameGT applies the GT predicate on the "topic_name" field.
func TopicNameGT(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldGT(FieldTopicName, v))
}

// TopicNameGTE applies the GTE predicate on the "topic_name" field.
func TopicNameGTE(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldGTE(FieldTopicName, v))
}

// TopicNameLT applies the LT predicate on the "topic_name" field.
func TopicNameLT(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldLT(FieldTopicName, v))
}

// TopicNameLTE applies the LTE predicate on the "topic_name" field.
func TopicNameLTE(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldLTE(FieldTopicName, v))
}

// TopicNameContains applies the Contains predicate on the "topic_name" field.
func TopicNameContains(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldContains(FieldTopicName, v))
}

// TopicNameHasPrefix applies the HasPrefix predicate on the "topic_name" field.
func TopicNameHasPrefix(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldHasPrefix(FieldTopicName, v))
}

// TopicNameHasSuffix applies the HasSuffix predicate on the "topic_name" field.
func TopicNameHasSuffix(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldHasSuffix(FieldTopicName, v))
}

// TopicNameIsNil applies the IsNil predicate on the "topic_name" field.
func TopicNameIsNil() predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldIsNull(FieldTopicName))
}

// TopicNameNotNil applies the NotNil predicate on the "topic_name" field.
func TopicNameNotNil() predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNotNull(FieldTopicName))
}

// TopicNameEqualFold applies the EqualFold predicate on the "topic_name" field.
func TopicNameEqualFold(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEqualFold(FieldTopicName, v))
}

// TopicNameContainsFold applies the ContainsFold predicate on the "topic_name" field.
func TopicNameContainsFold(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldContainsFold(FieldTopicName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldContainsFold(FieldDescription, v))
}

// CustomSubjectEQ applies the EQ predicate on the "custom_subject" field.
func CustomSubjectEQ(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldCustomSubject, v))
}

// CustomSubjectNEQ applies the NEQ predicate on the "custom_subject" field.
func CustomSubjectNEQ(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNEQ(FieldCustomSubject, v))
}

// CustomSubjectIn applies the In predicate on the "custom_subject" field.
func CustomSubjectIn(vs ...string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldIn(FieldCustomSubject, vs...))
}

// CustomSubjectNotIn applies the NotIn predicate on the "custom_subject" field.
func CustomSubjectNotIn(vs ...string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNotIn(FieldCustomSubject, vs...))
}

// CustomSubjectGT applies the GT predicate on the "custom_subject" field.
func CustomSubjectGT(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldGT(FieldCustomSubject, v))
}

// CustomSubjectGTE applies the GTE predicate on the "custom_subject" field.
func CustomSubjectGTE(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldGTE(FieldCustomSubject, v))
}

// CustomSubjectLT applies the LT predicate on the "custom_subject" field.
func CustomSubjectLT(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldLT(FieldCustomSubject, v))
}

// CustomSubjectLTE applies the LTE predicate on the "custom_subject" field.
func CustomSubjectLTE(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldLTE(FieldCustomSubject, v))
}

// CustomSubjectContains applies the Contains predicate on the "custom_subject" field.
func CustomSubjectContains(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldContains(FieldCustomSubject, v))
}

// CustomSubjectHasPrefix applies the HasPrefix predicate on the "custom_subject" field.
func CustomSubjectHasPrefix(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldHasPrefix(FieldCustomSubject, v))
}

// CustomSubjectHasSuffix applies the HasSuffix predicate on the "custom_subject" field.
func CustomSubjectHasSuffix(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldHasSuffix(FieldCustomSubject, v))
}

// CustomSubjectIsNil applies the IsNil predicate on the "custom_subject" field.
func CustomSubjectIsNil() predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldIsNull(FieldCustomSubject))
}

// CustomSubjectNotNil applies the NotNil predicate on the "custom_subject" field.
func CustomSubjectNotNil() predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNotNull(FieldCustomSubject))
}

// CustomSubjectEqualFold applies the EqualFold predicate on the "custom_subject" field.
func CustomSubjectEqualFold(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEqualFold(FieldCustomSubject, v))
}

// CustomSubjectContainsFold applies the ContainsFold predicate on the "custom_subject" field.
func CustomSubjectContainsFold(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldContainsFold(FieldCustomSubject, v))
}

// DeliveryMethodEQ applies the EQ predicate on the "delivery_method" field.
func DeliveryMethodEQ(v DeliveryMethod) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldDeliveryMethod, v))
}

// DeliveryMethodNEQ applies the NEQ predicate on the "delivery_method" field.
func DeliveryMethodNEQ(v DeliveryMethod) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNEQ(FieldDeliveryMethod, v))
}

// DeliveryMethodIn applies the In predicate on the "delivery_method" field.
func DeliveryMethodIn(vs ...DeliveryMethod) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldIn(FieldDeliveryMethod, vs...))
}

// DeliveryMethodNotIn applies the NotIn predicate on the "delivery_method" field.
func DeliveryMethodNotIn(vs ...DeliveryMethod) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNotIn(FieldDeliveryMethod, vs...))
}

// EmailRecipientsEQ applies the EQ predicate on the "email_recipients" field.
func EmailRecipientsEQ(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldEmailRecipients, v))
}

// EmailRecipientsNEQ applies the NEQ predicate on the "email_recipients" field.
func EmailRecipientsNEQ(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNEQ(FieldEmailRecipients, v))
}

// EmailRecipientsIn applies the In predicate on the "email_recipients" field.
func EmailRecipientsIn(vs ...string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldIn(FieldEmailRecipients, vs...))
}

// EmailRecipientsNotIn applies the NotIn predicate on the "email_recipients" field.
func EmailRecipientsNotIn(vs ...string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNotIn(FieldEmailRecipients, vs...))
}

// EmailRecipientsGT applies the GT predicate on the "email_recipients" field.
func EmailRecipientsGT(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldGT(FieldEmailRecipients, v))
}

// EmailRecipientsGTE applies the GTE predicate on the "email_recipients" field.
func EmailRecipientsGTE(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldGTE(FieldEmailRecipients, v))
}

// EmailRecipientsLT applies the LT predicate on the "email_recipients" field.
func EmailRecipientsLT(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldLT(FieldEmailRecipients, v))
}

// EmailRecipientsLTE applies the LTE predicate on the "email_recipients" field.
func EmailRecipientsLTE(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldLTE(FieldEmailRecipients, v))
}

// EmailRecipientsContains applies the Contains predicate on the "email_recipients" field.
func EmailRecipientsContains(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldContains(FieldEmailRecipients, v))
}

// EmailRecipientsHasPrefix applies the HasPrefix predicate on the "email_recipients" field.
func EmailRecipientsHasPrefix(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldHasPrefix(FieldEmailRecipients, v))
}

// EmailRecipientsHasSuffix applies the HasSuffix predicate on the "email_recipients" field.
func EmailRecipientsHasSuffix(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldHasSuffix(FieldEmailRecipients, v))
}

// EmailRecipientsIsNil applies the IsNil predicate on the "email_recipients" field.
func EmailRecipientsIsNil() predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldIsNull(FieldEmailRecipients))
}

// EmailRecipientsNotNil applies the NotNil predicate on the "email_recipients" field.
func EmailRecipientsNotNil() predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNotNull(FieldEmailRecipients))
}

// EmailRecipientsEqualFold applies the EqualFold predicate on the "email_recipients" field.
func EmailRecipientsEqualFold(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEqualFold(FieldEmailRecipients, v))
}

// EmailRecipientsContainsFold applies the ContainsFold predicate on the "email_recipients" field.
func EmailRecipientsContainsFold(v string) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldContainsFold(FieldEmailRecipients, v))
}

// EffectiveDateEQ applies the EQ predicate on the "effective_date" field.
func EffectiveDateEQ(v *pgtype.Date) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldEffectiveDate, v))
}

// EffectiveDateNEQ applies the NEQ predicate on the "effective_date" field.
func EffectiveDateNEQ(v *pgtype.Date) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNEQ(FieldEffectiveDate, v))
}

// EffectiveDateIn applies the In predicate on the "effective_date" field.
func EffectiveDateIn(vs ...*pgtype.Date) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldIn(FieldEffectiveDate, vs...))
}

// EffectiveDateNotIn applies the NotIn predicate on the "effective_date" field.
func EffectiveDateNotIn(vs ...*pgtype.Date) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNotIn(FieldEffectiveDate, vs...))
}

// EffectiveDateGT applies the GT predicate on the "effective_date" field.
func EffectiveDateGT(v *pgtype.Date) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldGT(FieldEffectiveDate, v))
}

// EffectiveDateGTE applies the GTE predicate on the "effective_date" field.
func EffectiveDateGTE(v *pgtype.Date) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldGTE(FieldEffectiveDate, v))
}

// EffectiveDateLT applies the LT predicate on the "effective_date" field.
func EffectiveDateLT(v *pgtype.Date) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldLT(FieldEffectiveDate, v))
}

// EffectiveDateLTE applies the LTE predicate on the "effective_date" field.
func EffectiveDateLTE(v *pgtype.Date) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldLTE(FieldEffectiveDate, v))
}

// EffectiveDateIsNil applies the IsNil predicate on the "effective_date" field.
func EffectiveDateIsNil() predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldIsNull(FieldEffectiveDate))
}

// EffectiveDateNotNil applies the NotNil predicate on the "effective_date" field.
func EffectiveDateNotNil() predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNotNull(FieldEffectiveDate))
}

// ExpirationDateEQ applies the EQ predicate on the "expiration_date" field.
func ExpirationDateEQ(v *pgtype.Date) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldEQ(FieldExpirationDate, v))
}

// ExpirationDateNEQ applies the NEQ predicate on the "expiration_date" field.
func ExpirationDateNEQ(v *pgtype.Date) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNEQ(FieldExpirationDate, v))
}

// ExpirationDateIn applies the In predicate on the "expiration_date" field.
func ExpirationDateIn(vs ...*pgtype.Date) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldIn(FieldExpirationDate, vs...))
}

// ExpirationDateNotIn applies the NotIn predicate on the "expiration_date" field.
func ExpirationDateNotIn(vs ...*pgtype.Date) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNotIn(FieldExpirationDate, vs...))
}

// ExpirationDateGT applies the GT predicate on the "expiration_date" field.
func ExpirationDateGT(v *pgtype.Date) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldGT(FieldExpirationDate, v))
}

// ExpirationDateGTE applies the GTE predicate on the "expiration_date" field.
func ExpirationDateGTE(v *pgtype.Date) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldGTE(FieldExpirationDate, v))
}

// ExpirationDateLT applies the LT predicate on the "expiration_date" field.
func ExpirationDateLT(v *pgtype.Date) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldLT(FieldExpirationDate, v))
}

// ExpirationDateLTE applies the LTE predicate on the "expiration_date" field.
func ExpirationDateLTE(v *pgtype.Date) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldLTE(FieldExpirationDate, v))
}

// ExpirationDateIsNil applies the IsNil predicate on the "expiration_date" field.
func ExpirationDateIsNil() predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldIsNull(FieldExpirationDate))
}

// ExpirationDateNotNil applies the NotNil predicate on the "expiration_date" field.
func ExpirationDateNotNil() predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNotNull(FieldExpirationDate))
}

// ConditionalLogicIsNil applies the IsNil predicate on the "conditional_logic" field.
func ConditionalLogicIsNil() predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldIsNull(FieldConditionalLogic))
}

// ConditionalLogicNotNil applies the NotNil predicate on the "conditional_logic" field.
func ConditionalLogicNotNil() predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.FieldNotNull(FieldConditionalLogic))
}

// HasBusinessUnit applies the HasEdge predicate on the "business_unit" edge.
func HasBusinessUnit() predicate.TableChangeAlert {
	return predicate.TableChangeAlert(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessUnitWith applies the HasEdge predicate on the "business_unit" edge with a given conditions (other predicates).
func HasBusinessUnitWith(preds ...predicate.BusinessUnit) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(func(s *sql.Selector) {
		step := newBusinessUnitStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.TableChangeAlert {
	return predicate.TableChangeAlert(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TableChangeAlert) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TableChangeAlert) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TableChangeAlert) predicate.TableChangeAlert {
	return predicate.TableChangeAlert(sql.NotPredicates(p))
}
