// Code generated by ent, DO NOT EDIT.

package tractor

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldLTE(FieldID, id))
}

// BusinessUnitID applies equality check predicate on the "business_unit_id" field. It's identical to BusinessUnitIDEQ.
func BusinessUnitID(v uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldBusinessUnitID, v))
}

// OrganizationID applies equality check predicate on the "organization_id" field. It's identical to OrganizationIDEQ.
func OrganizationID(v uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldOrganizationID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldUpdatedAt, v))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldCode, v))
}

// EquipmentTypeID applies equality check predicate on the "equipment_type_id" field. It's identical to EquipmentTypeIDEQ.
func EquipmentTypeID(v uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldEquipmentTypeID, v))
}

// LicensePlateNumber applies equality check predicate on the "license_plate_number" field. It's identical to LicensePlateNumberEQ.
func LicensePlateNumber(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldLicensePlateNumber, v))
}

// Vin applies equality check predicate on the "vin" field. It's identical to VinEQ.
func Vin(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldVin, v))
}

// EquipmentManufacturerID applies equality check predicate on the "equipment_manufacturer_id" field. It's identical to EquipmentManufacturerIDEQ.
func EquipmentManufacturerID(v uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldEquipmentManufacturerID, v))
}

// Model applies equality check predicate on the "model" field. It's identical to ModelEQ.
func Model(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldModel, v))
}

// Year applies equality check predicate on the "year" field. It's identical to YearEQ.
func Year(v int) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldYear, v))
}

// StateID applies equality check predicate on the "state_id" field. It's identical to StateIDEQ.
func StateID(v uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldStateID, v))
}

// Leased applies equality check predicate on the "leased" field. It's identical to LeasedEQ.
func Leased(v bool) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldLeased, v))
}

// LeasedDate applies equality check predicate on the "leased_date" field. It's identical to LeasedDateEQ.
func LeasedDate(v time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldLeasedDate, v))
}

// PrimaryWorkerID applies equality check predicate on the "primary_worker_id" field. It's identical to PrimaryWorkerIDEQ.
func PrimaryWorkerID(v uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldPrimaryWorkerID, v))
}

// SecondaryWorkerID applies equality check predicate on the "secondary_worker_id" field. It's identical to SecondaryWorkerIDEQ.
func SecondaryWorkerID(v uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldSecondaryWorkerID, v))
}

// FleetCodeID applies equality check predicate on the "fleet_code_id" field. It's identical to FleetCodeIDEQ.
func FleetCodeID(v uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldFleetCodeID, v))
}

// BusinessUnitIDEQ applies the EQ predicate on the "business_unit_id" field.
func BusinessUnitIDEQ(v uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDNEQ applies the NEQ predicate on the "business_unit_id" field.
func BusinessUnitIDNEQ(v uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldNEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDIn applies the In predicate on the "business_unit_id" field.
func BusinessUnitIDIn(vs ...uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldIn(FieldBusinessUnitID, vs...))
}

// BusinessUnitIDNotIn applies the NotIn predicate on the "business_unit_id" field.
func BusinessUnitIDNotIn(vs ...uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldNotIn(FieldBusinessUnitID, vs...))
}

// OrganizationIDEQ applies the EQ predicate on the "organization_id" field.
func OrganizationIDEQ(v uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldOrganizationID, v))
}

// OrganizationIDNEQ applies the NEQ predicate on the "organization_id" field.
func OrganizationIDNEQ(v uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldNEQ(FieldOrganizationID, v))
}

// OrganizationIDIn applies the In predicate on the "organization_id" field.
func OrganizationIDIn(vs ...uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldIn(FieldOrganizationID, vs...))
}

// OrganizationIDNotIn applies the NotIn predicate on the "organization_id" field.
func OrganizationIDNotIn(vs ...uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldNotIn(FieldOrganizationID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldLTE(FieldUpdatedAt, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.Tractor {
	return predicate.Tractor(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.Tractor {
	return predicate.Tractor(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldContainsFold(FieldCode, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Tractor {
	return predicate.Tractor(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Tractor {
	return predicate.Tractor(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Tractor {
	return predicate.Tractor(sql.FieldNotIn(FieldStatus, vs...))
}

// EquipmentTypeIDEQ applies the EQ predicate on the "equipment_type_id" field.
func EquipmentTypeIDEQ(v uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldEquipmentTypeID, v))
}

// EquipmentTypeIDNEQ applies the NEQ predicate on the "equipment_type_id" field.
func EquipmentTypeIDNEQ(v uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldNEQ(FieldEquipmentTypeID, v))
}

// EquipmentTypeIDIn applies the In predicate on the "equipment_type_id" field.
func EquipmentTypeIDIn(vs ...uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldIn(FieldEquipmentTypeID, vs...))
}

// EquipmentTypeIDNotIn applies the NotIn predicate on the "equipment_type_id" field.
func EquipmentTypeIDNotIn(vs ...uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldNotIn(FieldEquipmentTypeID, vs...))
}

// EquipmentTypeIDIsNil applies the IsNil predicate on the "equipment_type_id" field.
func EquipmentTypeIDIsNil() predicate.Tractor {
	return predicate.Tractor(sql.FieldIsNull(FieldEquipmentTypeID))
}

// EquipmentTypeIDNotNil applies the NotNil predicate on the "equipment_type_id" field.
func EquipmentTypeIDNotNil() predicate.Tractor {
	return predicate.Tractor(sql.FieldNotNull(FieldEquipmentTypeID))
}

// LicensePlateNumberEQ applies the EQ predicate on the "license_plate_number" field.
func LicensePlateNumberEQ(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldLicensePlateNumber, v))
}

// LicensePlateNumberNEQ applies the NEQ predicate on the "license_plate_number" field.
func LicensePlateNumberNEQ(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldNEQ(FieldLicensePlateNumber, v))
}

// LicensePlateNumberIn applies the In predicate on the "license_plate_number" field.
func LicensePlateNumberIn(vs ...string) predicate.Tractor {
	return predicate.Tractor(sql.FieldIn(FieldLicensePlateNumber, vs...))
}

// LicensePlateNumberNotIn applies the NotIn predicate on the "license_plate_number" field.
func LicensePlateNumberNotIn(vs ...string) predicate.Tractor {
	return predicate.Tractor(sql.FieldNotIn(FieldLicensePlateNumber, vs...))
}

// LicensePlateNumberGT applies the GT predicate on the "license_plate_number" field.
func LicensePlateNumberGT(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldGT(FieldLicensePlateNumber, v))
}

// LicensePlateNumberGTE applies the GTE predicate on the "license_plate_number" field.
func LicensePlateNumberGTE(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldGTE(FieldLicensePlateNumber, v))
}

// LicensePlateNumberLT applies the LT predicate on the "license_plate_number" field.
func LicensePlateNumberLT(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldLT(FieldLicensePlateNumber, v))
}

// LicensePlateNumberLTE applies the LTE predicate on the "license_plate_number" field.
func LicensePlateNumberLTE(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldLTE(FieldLicensePlateNumber, v))
}

// LicensePlateNumberContains applies the Contains predicate on the "license_plate_number" field.
func LicensePlateNumberContains(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldContains(FieldLicensePlateNumber, v))
}

// LicensePlateNumberHasPrefix applies the HasPrefix predicate on the "license_plate_number" field.
func LicensePlateNumberHasPrefix(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldHasPrefix(FieldLicensePlateNumber, v))
}

// LicensePlateNumberHasSuffix applies the HasSuffix predicate on the "license_plate_number" field.
func LicensePlateNumberHasSuffix(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldHasSuffix(FieldLicensePlateNumber, v))
}

// LicensePlateNumberIsNil applies the IsNil predicate on the "license_plate_number" field.
func LicensePlateNumberIsNil() predicate.Tractor {
	return predicate.Tractor(sql.FieldIsNull(FieldLicensePlateNumber))
}

// LicensePlateNumberNotNil applies the NotNil predicate on the "license_plate_number" field.
func LicensePlateNumberNotNil() predicate.Tractor {
	return predicate.Tractor(sql.FieldNotNull(FieldLicensePlateNumber))
}

// LicensePlateNumberEqualFold applies the EqualFold predicate on the "license_plate_number" field.
func LicensePlateNumberEqualFold(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldEqualFold(FieldLicensePlateNumber, v))
}

// LicensePlateNumberContainsFold applies the ContainsFold predicate on the "license_plate_number" field.
func LicensePlateNumberContainsFold(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldContainsFold(FieldLicensePlateNumber, v))
}

// VinEQ applies the EQ predicate on the "vin" field.
func VinEQ(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldVin, v))
}

// VinNEQ applies the NEQ predicate on the "vin" field.
func VinNEQ(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldNEQ(FieldVin, v))
}

// VinIn applies the In predicate on the "vin" field.
func VinIn(vs ...string) predicate.Tractor {
	return predicate.Tractor(sql.FieldIn(FieldVin, vs...))
}

// VinNotIn applies the NotIn predicate on the "vin" field.
func VinNotIn(vs ...string) predicate.Tractor {
	return predicate.Tractor(sql.FieldNotIn(FieldVin, vs...))
}

// VinGT applies the GT predicate on the "vin" field.
func VinGT(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldGT(FieldVin, v))
}

// VinGTE applies the GTE predicate on the "vin" field.
func VinGTE(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldGTE(FieldVin, v))
}

// VinLT applies the LT predicate on the "vin" field.
func VinLT(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldLT(FieldVin, v))
}

// VinLTE applies the LTE predicate on the "vin" field.
func VinLTE(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldLTE(FieldVin, v))
}

// VinContains applies the Contains predicate on the "vin" field.
func VinContains(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldContains(FieldVin, v))
}

// VinHasPrefix applies the HasPrefix predicate on the "vin" field.
func VinHasPrefix(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldHasPrefix(FieldVin, v))
}

// VinHasSuffix applies the HasSuffix predicate on the "vin" field.
func VinHasSuffix(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldHasSuffix(FieldVin, v))
}

// VinIsNil applies the IsNil predicate on the "vin" field.
func VinIsNil() predicate.Tractor {
	return predicate.Tractor(sql.FieldIsNull(FieldVin))
}

// VinNotNil applies the NotNil predicate on the "vin" field.
func VinNotNil() predicate.Tractor {
	return predicate.Tractor(sql.FieldNotNull(FieldVin))
}

// VinEqualFold applies the EqualFold predicate on the "vin" field.
func VinEqualFold(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldEqualFold(FieldVin, v))
}

// VinContainsFold applies the ContainsFold predicate on the "vin" field.
func VinContainsFold(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldContainsFold(FieldVin, v))
}

// EquipmentManufacturerIDEQ applies the EQ predicate on the "equipment_manufacturer_id" field.
func EquipmentManufacturerIDEQ(v uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldEquipmentManufacturerID, v))
}

// EquipmentManufacturerIDNEQ applies the NEQ predicate on the "equipment_manufacturer_id" field.
func EquipmentManufacturerIDNEQ(v uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldNEQ(FieldEquipmentManufacturerID, v))
}

// EquipmentManufacturerIDIn applies the In predicate on the "equipment_manufacturer_id" field.
func EquipmentManufacturerIDIn(vs ...uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldIn(FieldEquipmentManufacturerID, vs...))
}

// EquipmentManufacturerIDNotIn applies the NotIn predicate on the "equipment_manufacturer_id" field.
func EquipmentManufacturerIDNotIn(vs ...uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldNotIn(FieldEquipmentManufacturerID, vs...))
}

// EquipmentManufacturerIDIsNil applies the IsNil predicate on the "equipment_manufacturer_id" field.
func EquipmentManufacturerIDIsNil() predicate.Tractor {
	return predicate.Tractor(sql.FieldIsNull(FieldEquipmentManufacturerID))
}

// EquipmentManufacturerIDNotNil applies the NotNil predicate on the "equipment_manufacturer_id" field.
func EquipmentManufacturerIDNotNil() predicate.Tractor {
	return predicate.Tractor(sql.FieldNotNull(FieldEquipmentManufacturerID))
}

// ModelEQ applies the EQ predicate on the "model" field.
func ModelEQ(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldModel, v))
}

// ModelNEQ applies the NEQ predicate on the "model" field.
func ModelNEQ(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldNEQ(FieldModel, v))
}

// ModelIn applies the In predicate on the "model" field.
func ModelIn(vs ...string) predicate.Tractor {
	return predicate.Tractor(sql.FieldIn(FieldModel, vs...))
}

// ModelNotIn applies the NotIn predicate on the "model" field.
func ModelNotIn(vs ...string) predicate.Tractor {
	return predicate.Tractor(sql.FieldNotIn(FieldModel, vs...))
}

// ModelGT applies the GT predicate on the "model" field.
func ModelGT(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldGT(FieldModel, v))
}

// ModelGTE applies the GTE predicate on the "model" field.
func ModelGTE(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldGTE(FieldModel, v))
}

// ModelLT applies the LT predicate on the "model" field.
func ModelLT(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldLT(FieldModel, v))
}

// ModelLTE applies the LTE predicate on the "model" field.
func ModelLTE(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldLTE(FieldModel, v))
}

// ModelContains applies the Contains predicate on the "model" field.
func ModelContains(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldContains(FieldModel, v))
}

// ModelHasPrefix applies the HasPrefix predicate on the "model" field.
func ModelHasPrefix(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldHasPrefix(FieldModel, v))
}

// ModelHasSuffix applies the HasSuffix predicate on the "model" field.
func ModelHasSuffix(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldHasSuffix(FieldModel, v))
}

// ModelIsNil applies the IsNil predicate on the "model" field.
func ModelIsNil() predicate.Tractor {
	return predicate.Tractor(sql.FieldIsNull(FieldModel))
}

// ModelNotNil applies the NotNil predicate on the "model" field.
func ModelNotNil() predicate.Tractor {
	return predicate.Tractor(sql.FieldNotNull(FieldModel))
}

// ModelEqualFold applies the EqualFold predicate on the "model" field.
func ModelEqualFold(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldEqualFold(FieldModel, v))
}

// ModelContainsFold applies the ContainsFold predicate on the "model" field.
func ModelContainsFold(v string) predicate.Tractor {
	return predicate.Tractor(sql.FieldContainsFold(FieldModel, v))
}

// YearEQ applies the EQ predicate on the "year" field.
func YearEQ(v int) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldYear, v))
}

// YearNEQ applies the NEQ predicate on the "year" field.
func YearNEQ(v int) predicate.Tractor {
	return predicate.Tractor(sql.FieldNEQ(FieldYear, v))
}

// YearIn applies the In predicate on the "year" field.
func YearIn(vs ...int) predicate.Tractor {
	return predicate.Tractor(sql.FieldIn(FieldYear, vs...))
}

// YearNotIn applies the NotIn predicate on the "year" field.
func YearNotIn(vs ...int) predicate.Tractor {
	return predicate.Tractor(sql.FieldNotIn(FieldYear, vs...))
}

// YearGT applies the GT predicate on the "year" field.
func YearGT(v int) predicate.Tractor {
	return predicate.Tractor(sql.FieldGT(FieldYear, v))
}

// YearGTE applies the GTE predicate on the "year" field.
func YearGTE(v int) predicate.Tractor {
	return predicate.Tractor(sql.FieldGTE(FieldYear, v))
}

// YearLT applies the LT predicate on the "year" field.
func YearLT(v int) predicate.Tractor {
	return predicate.Tractor(sql.FieldLT(FieldYear, v))
}

// YearLTE applies the LTE predicate on the "year" field.
func YearLTE(v int) predicate.Tractor {
	return predicate.Tractor(sql.FieldLTE(FieldYear, v))
}

// YearIsNil applies the IsNil predicate on the "year" field.
func YearIsNil() predicate.Tractor {
	return predicate.Tractor(sql.FieldIsNull(FieldYear))
}

// YearNotNil applies the NotNil predicate on the "year" field.
func YearNotNil() predicate.Tractor {
	return predicate.Tractor(sql.FieldNotNull(FieldYear))
}

// StateIDEQ applies the EQ predicate on the "state_id" field.
func StateIDEQ(v uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldStateID, v))
}

// StateIDNEQ applies the NEQ predicate on the "state_id" field.
func StateIDNEQ(v uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldNEQ(FieldStateID, v))
}

// StateIDIn applies the In predicate on the "state_id" field.
func StateIDIn(vs ...uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldIn(FieldStateID, vs...))
}

// StateIDNotIn applies the NotIn predicate on the "state_id" field.
func StateIDNotIn(vs ...uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldNotIn(FieldStateID, vs...))
}

// StateIDIsNil applies the IsNil predicate on the "state_id" field.
func StateIDIsNil() predicate.Tractor {
	return predicate.Tractor(sql.FieldIsNull(FieldStateID))
}

// StateIDNotNil applies the NotNil predicate on the "state_id" field.
func StateIDNotNil() predicate.Tractor {
	return predicate.Tractor(sql.FieldNotNull(FieldStateID))
}

// LeasedEQ applies the EQ predicate on the "leased" field.
func LeasedEQ(v bool) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldLeased, v))
}

// LeasedNEQ applies the NEQ predicate on the "leased" field.
func LeasedNEQ(v bool) predicate.Tractor {
	return predicate.Tractor(sql.FieldNEQ(FieldLeased, v))
}

// LeasedDateEQ applies the EQ predicate on the "leased_date" field.
func LeasedDateEQ(v time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldLeasedDate, v))
}

// LeasedDateNEQ applies the NEQ predicate on the "leased_date" field.
func LeasedDateNEQ(v time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldNEQ(FieldLeasedDate, v))
}

// LeasedDateIn applies the In predicate on the "leased_date" field.
func LeasedDateIn(vs ...time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldIn(FieldLeasedDate, vs...))
}

// LeasedDateNotIn applies the NotIn predicate on the "leased_date" field.
func LeasedDateNotIn(vs ...time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldNotIn(FieldLeasedDate, vs...))
}

// LeasedDateGT applies the GT predicate on the "leased_date" field.
func LeasedDateGT(v time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldGT(FieldLeasedDate, v))
}

// LeasedDateGTE applies the GTE predicate on the "leased_date" field.
func LeasedDateGTE(v time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldGTE(FieldLeasedDate, v))
}

// LeasedDateLT applies the LT predicate on the "leased_date" field.
func LeasedDateLT(v time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldLT(FieldLeasedDate, v))
}

// LeasedDateLTE applies the LTE predicate on the "leased_date" field.
func LeasedDateLTE(v time.Time) predicate.Tractor {
	return predicate.Tractor(sql.FieldLTE(FieldLeasedDate, v))
}

// LeasedDateIsNil applies the IsNil predicate on the "leased_date" field.
func LeasedDateIsNil() predicate.Tractor {
	return predicate.Tractor(sql.FieldIsNull(FieldLeasedDate))
}

// LeasedDateNotNil applies the NotNil predicate on the "leased_date" field.
func LeasedDateNotNil() predicate.Tractor {
	return predicate.Tractor(sql.FieldNotNull(FieldLeasedDate))
}

// PrimaryWorkerIDEQ applies the EQ predicate on the "primary_worker_id" field.
func PrimaryWorkerIDEQ(v uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldPrimaryWorkerID, v))
}

// PrimaryWorkerIDNEQ applies the NEQ predicate on the "primary_worker_id" field.
func PrimaryWorkerIDNEQ(v uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldNEQ(FieldPrimaryWorkerID, v))
}

// PrimaryWorkerIDIn applies the In predicate on the "primary_worker_id" field.
func PrimaryWorkerIDIn(vs ...uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldIn(FieldPrimaryWorkerID, vs...))
}

// PrimaryWorkerIDNotIn applies the NotIn predicate on the "primary_worker_id" field.
func PrimaryWorkerIDNotIn(vs ...uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldNotIn(FieldPrimaryWorkerID, vs...))
}

// SecondaryWorkerIDEQ applies the EQ predicate on the "secondary_worker_id" field.
func SecondaryWorkerIDEQ(v uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldSecondaryWorkerID, v))
}

// SecondaryWorkerIDNEQ applies the NEQ predicate on the "secondary_worker_id" field.
func SecondaryWorkerIDNEQ(v uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldNEQ(FieldSecondaryWorkerID, v))
}

// SecondaryWorkerIDIn applies the In predicate on the "secondary_worker_id" field.
func SecondaryWorkerIDIn(vs ...uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldIn(FieldSecondaryWorkerID, vs...))
}

// SecondaryWorkerIDNotIn applies the NotIn predicate on the "secondary_worker_id" field.
func SecondaryWorkerIDNotIn(vs ...uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldNotIn(FieldSecondaryWorkerID, vs...))
}

// SecondaryWorkerIDIsNil applies the IsNil predicate on the "secondary_worker_id" field.
func SecondaryWorkerIDIsNil() predicate.Tractor {
	return predicate.Tractor(sql.FieldIsNull(FieldSecondaryWorkerID))
}

// SecondaryWorkerIDNotNil applies the NotNil predicate on the "secondary_worker_id" field.
func SecondaryWorkerIDNotNil() predicate.Tractor {
	return predicate.Tractor(sql.FieldNotNull(FieldSecondaryWorkerID))
}

// FleetCodeIDEQ applies the EQ predicate on the "fleet_code_id" field.
func FleetCodeIDEQ(v uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldEQ(FieldFleetCodeID, v))
}

// FleetCodeIDNEQ applies the NEQ predicate on the "fleet_code_id" field.
func FleetCodeIDNEQ(v uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldNEQ(FieldFleetCodeID, v))
}

// FleetCodeIDIn applies the In predicate on the "fleet_code_id" field.
func FleetCodeIDIn(vs ...uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldIn(FieldFleetCodeID, vs...))
}

// FleetCodeIDNotIn applies the NotIn predicate on the "fleet_code_id" field.
func FleetCodeIDNotIn(vs ...uuid.UUID) predicate.Tractor {
	return predicate.Tractor(sql.FieldNotIn(FieldFleetCodeID, vs...))
}

// HasBusinessUnit applies the HasEdge predicate on the "business_unit" edge.
func HasBusinessUnit() predicate.Tractor {
	return predicate.Tractor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessUnitWith applies the HasEdge predicate on the "business_unit" edge with a given conditions (other predicates).
func HasBusinessUnitWith(preds ...predicate.BusinessUnit) predicate.Tractor {
	return predicate.Tractor(func(s *sql.Selector) {
		step := newBusinessUnitStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.Tractor {
	return predicate.Tractor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.Tractor {
	return predicate.Tractor(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEquipmentType applies the HasEdge predicate on the "equipment_type" edge.
func HasEquipmentType() predicate.Tractor {
	return predicate.Tractor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, EquipmentTypeTable, EquipmentTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEquipmentTypeWith applies the HasEdge predicate on the "equipment_type" edge with a given conditions (other predicates).
func HasEquipmentTypeWith(preds ...predicate.EquipmentType) predicate.Tractor {
	return predicate.Tractor(func(s *sql.Selector) {
		step := newEquipmentTypeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEquipmentManufacturer applies the HasEdge predicate on the "equipment_manufacturer" edge.
func HasEquipmentManufacturer() predicate.Tractor {
	return predicate.Tractor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, EquipmentManufacturerTable, EquipmentManufacturerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEquipmentManufacturerWith applies the HasEdge predicate on the "equipment_manufacturer" edge with a given conditions (other predicates).
func HasEquipmentManufacturerWith(preds ...predicate.EquipmentManufactuer) predicate.Tractor {
	return predicate.Tractor(func(s *sql.Selector) {
		step := newEquipmentManufacturerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasState applies the HasEdge predicate on the "state" edge.
func HasState() predicate.Tractor {
	return predicate.Tractor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, StateTable, StateColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStateWith applies the HasEdge predicate on the "state" edge with a given conditions (other predicates).
func HasStateWith(preds ...predicate.UsState) predicate.Tractor {
	return predicate.Tractor(func(s *sql.Selector) {
		step := newStateStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPrimaryWorker applies the HasEdge predicate on the "primary_worker" edge.
func HasPrimaryWorker() predicate.Tractor {
	return predicate.Tractor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, PrimaryWorkerTable, PrimaryWorkerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPrimaryWorkerWith applies the HasEdge predicate on the "primary_worker" edge with a given conditions (other predicates).
func HasPrimaryWorkerWith(preds ...predicate.Worker) predicate.Tractor {
	return predicate.Tractor(func(s *sql.Selector) {
		step := newPrimaryWorkerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSecondaryWorker applies the HasEdge predicate on the "secondary_worker" edge.
func HasSecondaryWorker() predicate.Tractor {
	return predicate.Tractor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, SecondaryWorkerTable, SecondaryWorkerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSecondaryWorkerWith applies the HasEdge predicate on the "secondary_worker" edge with a given conditions (other predicates).
func HasSecondaryWorkerWith(preds ...predicate.Worker) predicate.Tractor {
	return predicate.Tractor(func(s *sql.Selector) {
		step := newSecondaryWorkerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFleetCode applies the HasEdge predicate on the "fleet_code" edge.
func HasFleetCode() predicate.Tractor {
	return predicate.Tractor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, FleetCodeTable, FleetCodeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFleetCodeWith applies the HasEdge predicate on the "fleet_code" edge with a given conditions (other predicates).
func HasFleetCodeWith(preds ...predicate.FleetCode) predicate.Tractor {
	return predicate.Tractor(func(s *sql.Selector) {
		step := newFleetCodeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Tractor) predicate.Tractor {
	return predicate.Tractor(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Tractor) predicate.Tractor {
	return predicate.Tractor(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Tractor) predicate.Tractor {
	return predicate.Tractor(sql.NotPredicates(p))
}
