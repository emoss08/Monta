// Code generated by entc, DO NOT EDIT.

package shipmentroute

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/emoss08/trenova/internal/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldLTE(FieldID, id))
}

// BusinessUnitID applies equality check predicate on the "business_unit_id" field. It's identical to BusinessUnitIDEQ.
func BusinessUnitID(v uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldEQ(FieldBusinessUnitID, v))
}

// OrganizationID applies equality check predicate on the "organization_id" field. It's identical to OrganizationIDEQ.
func OrganizationID(v uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldEQ(FieldOrganizationID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldEQ(FieldUpdatedAt, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldEQ(FieldVersion, v))
}

// OriginLocationID applies equality check predicate on the "origin_location_id" field. It's identical to OriginLocationIDEQ.
func OriginLocationID(v uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldEQ(FieldOriginLocationID, v))
}

// DestinationLocationID applies equality check predicate on the "destination_location_id" field. It's identical to DestinationLocationIDEQ.
func DestinationLocationID(v uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldEQ(FieldDestinationLocationID, v))
}

// Mileage applies equality check predicate on the "mileage" field. It's identical to MileageEQ.
func Mileage(v float64) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldEQ(FieldMileage, v))
}

// Duration applies equality check predicate on the "duration" field. It's identical to DurationEQ.
func Duration(v int) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldEQ(FieldDuration, v))
}

// DistanceMethod applies equality check predicate on the "distance_method" field. It's identical to DistanceMethodEQ.
func DistanceMethod(v string) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldEQ(FieldDistanceMethod, v))
}

// AutoGenerated applies equality check predicate on the "auto_generated" field. It's identical to AutoGeneratedEQ.
func AutoGenerated(v bool) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldEQ(FieldAutoGenerated, v))
}

// BusinessUnitIDEQ applies the EQ predicate on the "business_unit_id" field.
func BusinessUnitIDEQ(v uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDNEQ applies the NEQ predicate on the "business_unit_id" field.
func BusinessUnitIDNEQ(v uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldNEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDIn applies the In predicate on the "business_unit_id" field.
func BusinessUnitIDIn(vs ...uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldIn(FieldBusinessUnitID, vs...))
}

// BusinessUnitIDNotIn applies the NotIn predicate on the "business_unit_id" field.
func BusinessUnitIDNotIn(vs ...uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldNotIn(FieldBusinessUnitID, vs...))
}

// OrganizationIDEQ applies the EQ predicate on the "organization_id" field.
func OrganizationIDEQ(v uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldEQ(FieldOrganizationID, v))
}

// OrganizationIDNEQ applies the NEQ predicate on the "organization_id" field.
func OrganizationIDNEQ(v uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldNEQ(FieldOrganizationID, v))
}

// OrganizationIDIn applies the In predicate on the "organization_id" field.
func OrganizationIDIn(vs ...uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldIn(FieldOrganizationID, vs...))
}

// OrganizationIDNotIn applies the NotIn predicate on the "organization_id" field.
func OrganizationIDNotIn(vs ...uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldNotIn(FieldOrganizationID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldLTE(FieldUpdatedAt, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldLTE(FieldVersion, v))
}

// OriginLocationIDEQ applies the EQ predicate on the "origin_location_id" field.
func OriginLocationIDEQ(v uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldEQ(FieldOriginLocationID, v))
}

// OriginLocationIDNEQ applies the NEQ predicate on the "origin_location_id" field.
func OriginLocationIDNEQ(v uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldNEQ(FieldOriginLocationID, v))
}

// OriginLocationIDIn applies the In predicate on the "origin_location_id" field.
func OriginLocationIDIn(vs ...uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldIn(FieldOriginLocationID, vs...))
}

// OriginLocationIDNotIn applies the NotIn predicate on the "origin_location_id" field.
func OriginLocationIDNotIn(vs ...uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldNotIn(FieldOriginLocationID, vs...))
}

// DestinationLocationIDEQ applies the EQ predicate on the "destination_location_id" field.
func DestinationLocationIDEQ(v uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldEQ(FieldDestinationLocationID, v))
}

// DestinationLocationIDNEQ applies the NEQ predicate on the "destination_location_id" field.
func DestinationLocationIDNEQ(v uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldNEQ(FieldDestinationLocationID, v))
}

// DestinationLocationIDIn applies the In predicate on the "destination_location_id" field.
func DestinationLocationIDIn(vs ...uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldIn(FieldDestinationLocationID, vs...))
}

// DestinationLocationIDNotIn applies the NotIn predicate on the "destination_location_id" field.
func DestinationLocationIDNotIn(vs ...uuid.UUID) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldNotIn(FieldDestinationLocationID, vs...))
}

// MileageEQ applies the EQ predicate on the "mileage" field.
func MileageEQ(v float64) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldEQ(FieldMileage, v))
}

// MileageNEQ applies the NEQ predicate on the "mileage" field.
func MileageNEQ(v float64) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldNEQ(FieldMileage, v))
}

// MileageIn applies the In predicate on the "mileage" field.
func MileageIn(vs ...float64) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldIn(FieldMileage, vs...))
}

// MileageNotIn applies the NotIn predicate on the "mileage" field.
func MileageNotIn(vs ...float64) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldNotIn(FieldMileage, vs...))
}

// MileageGT applies the GT predicate on the "mileage" field.
func MileageGT(v float64) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldGT(FieldMileage, v))
}

// MileageGTE applies the GTE predicate on the "mileage" field.
func MileageGTE(v float64) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldGTE(FieldMileage, v))
}

// MileageLT applies the LT predicate on the "mileage" field.
func MileageLT(v float64) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldLT(FieldMileage, v))
}

// MileageLTE applies the LTE predicate on the "mileage" field.
func MileageLTE(v float64) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldLTE(FieldMileage, v))
}

// DurationEQ applies the EQ predicate on the "duration" field.
func DurationEQ(v int) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldEQ(FieldDuration, v))
}

// DurationNEQ applies the NEQ predicate on the "duration" field.
func DurationNEQ(v int) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldNEQ(FieldDuration, v))
}

// DurationIn applies the In predicate on the "duration" field.
func DurationIn(vs ...int) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldIn(FieldDuration, vs...))
}

// DurationNotIn applies the NotIn predicate on the "duration" field.
func DurationNotIn(vs ...int) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldNotIn(FieldDuration, vs...))
}

// DurationGT applies the GT predicate on the "duration" field.
func DurationGT(v int) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldGT(FieldDuration, v))
}

// DurationGTE applies the GTE predicate on the "duration" field.
func DurationGTE(v int) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldGTE(FieldDuration, v))
}

// DurationLT applies the LT predicate on the "duration" field.
func DurationLT(v int) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldLT(FieldDuration, v))
}

// DurationLTE applies the LTE predicate on the "duration" field.
func DurationLTE(v int) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldLTE(FieldDuration, v))
}

// DurationIsNil applies the IsNil predicate on the "duration" field.
func DurationIsNil() predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldIsNull(FieldDuration))
}

// DurationNotNil applies the NotNil predicate on the "duration" field.
func DurationNotNil() predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldNotNull(FieldDuration))
}

// DistanceMethodEQ applies the EQ predicate on the "distance_method" field.
func DistanceMethodEQ(v string) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldEQ(FieldDistanceMethod, v))
}

// DistanceMethodNEQ applies the NEQ predicate on the "distance_method" field.
func DistanceMethodNEQ(v string) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldNEQ(FieldDistanceMethod, v))
}

// DistanceMethodIn applies the In predicate on the "distance_method" field.
func DistanceMethodIn(vs ...string) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldIn(FieldDistanceMethod, vs...))
}

// DistanceMethodNotIn applies the NotIn predicate on the "distance_method" field.
func DistanceMethodNotIn(vs ...string) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldNotIn(FieldDistanceMethod, vs...))
}

// DistanceMethodGT applies the GT predicate on the "distance_method" field.
func DistanceMethodGT(v string) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldGT(FieldDistanceMethod, v))
}

// DistanceMethodGTE applies the GTE predicate on the "distance_method" field.
func DistanceMethodGTE(v string) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldGTE(FieldDistanceMethod, v))
}

// DistanceMethodLT applies the LT predicate on the "distance_method" field.
func DistanceMethodLT(v string) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldLT(FieldDistanceMethod, v))
}

// DistanceMethodLTE applies the LTE predicate on the "distance_method" field.
func DistanceMethodLTE(v string) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldLTE(FieldDistanceMethod, v))
}

// DistanceMethodContains applies the Contains predicate on the "distance_method" field.
func DistanceMethodContains(v string) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldContains(FieldDistanceMethod, v))
}

// DistanceMethodHasPrefix applies the HasPrefix predicate on the "distance_method" field.
func DistanceMethodHasPrefix(v string) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldHasPrefix(FieldDistanceMethod, v))
}

// DistanceMethodHasSuffix applies the HasSuffix predicate on the "distance_method" field.
func DistanceMethodHasSuffix(v string) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldHasSuffix(FieldDistanceMethod, v))
}

// DistanceMethodIsNil applies the IsNil predicate on the "distance_method" field.
func DistanceMethodIsNil() predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldIsNull(FieldDistanceMethod))
}

// DistanceMethodNotNil applies the NotNil predicate on the "distance_method" field.
func DistanceMethodNotNil() predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldNotNull(FieldDistanceMethod))
}

// DistanceMethodEqualFold applies the EqualFold predicate on the "distance_method" field.
func DistanceMethodEqualFold(v string) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldEqualFold(FieldDistanceMethod, v))
}

// DistanceMethodContainsFold applies the ContainsFold predicate on the "distance_method" field.
func DistanceMethodContainsFold(v string) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldContainsFold(FieldDistanceMethod, v))
}

// AutoGeneratedEQ applies the EQ predicate on the "auto_generated" field.
func AutoGeneratedEQ(v bool) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldEQ(FieldAutoGenerated, v))
}

// AutoGeneratedNEQ applies the NEQ predicate on the "auto_generated" field.
func AutoGeneratedNEQ(v bool) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.FieldNEQ(FieldAutoGenerated, v))
}

// HasBusinessUnit applies the HasEdge predicate on the "business_unit" edge.
func HasBusinessUnit() predicate.ShipmentRoute {
	return predicate.ShipmentRoute(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessUnitWith applies the HasEdge predicate on the "business_unit" edge with a given conditions (other predicates).
func HasBusinessUnitWith(preds ...predicate.BusinessUnit) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(func(s *sql.Selector) {
		step := newBusinessUnitStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.ShipmentRoute {
	return predicate.ShipmentRoute(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOriginLocation applies the HasEdge predicate on the "origin_location" edge.
func HasOriginLocation() predicate.ShipmentRoute {
	return predicate.ShipmentRoute(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OriginLocationTable, OriginLocationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOriginLocationWith applies the HasEdge predicate on the "origin_location" edge with a given conditions (other predicates).
func HasOriginLocationWith(preds ...predicate.Location) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(func(s *sql.Selector) {
		step := newOriginLocationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDestinationLocation applies the HasEdge predicate on the "destination_location" edge.
func HasDestinationLocation() predicate.ShipmentRoute {
	return predicate.ShipmentRoute(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DestinationLocationTable, DestinationLocationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDestinationLocationWith applies the HasEdge predicate on the "destination_location" edge with a given conditions (other predicates).
func HasDestinationLocationWith(preds ...predicate.Location) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(func(s *sql.Selector) {
		step := newDestinationLocationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ShipmentRoute) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ShipmentRoute) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ShipmentRoute) predicate.ShipmentRoute {
	return predicate.ShipmentRoute(sql.NotPredicates(p))
}
