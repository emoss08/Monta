// Code generated by entc, DO NOT EDIT.

package googleapi

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/emoss08/trenova/internal/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldEQ(FieldUpdatedAt, v))
}

// APIKey applies equality check predicate on the "api_key" field. It's identical to APIKeyEQ.
func APIKey(v string) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldEQ(FieldAPIKey, v))
}

// AddCustomerLocation applies equality check predicate on the "add_customer_location" field. It's identical to AddCustomerLocationEQ.
func AddCustomerLocation(v bool) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldEQ(FieldAddCustomerLocation, v))
}

// AutoGeocode applies equality check predicate on the "auto_geocode" field. It's identical to AutoGeocodeEQ.
func AutoGeocode(v bool) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldEQ(FieldAutoGeocode, v))
}

// AddLocation applies equality check predicate on the "add_location" field. It's identical to AddLocationEQ.
func AddLocation(v bool) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldEQ(FieldAddLocation, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldLTE(FieldUpdatedAt, v))
}

// APIKeyEQ applies the EQ predicate on the "api_key" field.
func APIKeyEQ(v string) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldEQ(FieldAPIKey, v))
}

// APIKeyNEQ applies the NEQ predicate on the "api_key" field.
func APIKeyNEQ(v string) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldNEQ(FieldAPIKey, v))
}

// APIKeyIn applies the In predicate on the "api_key" field.
func APIKeyIn(vs ...string) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldIn(FieldAPIKey, vs...))
}

// APIKeyNotIn applies the NotIn predicate on the "api_key" field.
func APIKeyNotIn(vs ...string) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldNotIn(FieldAPIKey, vs...))
}

// APIKeyGT applies the GT predicate on the "api_key" field.
func APIKeyGT(v string) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldGT(FieldAPIKey, v))
}

// APIKeyGTE applies the GTE predicate on the "api_key" field.
func APIKeyGTE(v string) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldGTE(FieldAPIKey, v))
}

// APIKeyLT applies the LT predicate on the "api_key" field.
func APIKeyLT(v string) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldLT(FieldAPIKey, v))
}

// APIKeyLTE applies the LTE predicate on the "api_key" field.
func APIKeyLTE(v string) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldLTE(FieldAPIKey, v))
}

// APIKeyContains applies the Contains predicate on the "api_key" field.
func APIKeyContains(v string) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldContains(FieldAPIKey, v))
}

// APIKeyHasPrefix applies the HasPrefix predicate on the "api_key" field.
func APIKeyHasPrefix(v string) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldHasPrefix(FieldAPIKey, v))
}

// APIKeyHasSuffix applies the HasSuffix predicate on the "api_key" field.
func APIKeyHasSuffix(v string) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldHasSuffix(FieldAPIKey, v))
}

// APIKeyEqualFold applies the EqualFold predicate on the "api_key" field.
func APIKeyEqualFold(v string) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldEqualFold(FieldAPIKey, v))
}

// APIKeyContainsFold applies the ContainsFold predicate on the "api_key" field.
func APIKeyContainsFold(v string) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldContainsFold(FieldAPIKey, v))
}

// MileageUnitEQ applies the EQ predicate on the "mileage_unit" field.
func MileageUnitEQ(v MileageUnit) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldEQ(FieldMileageUnit, v))
}

// MileageUnitNEQ applies the NEQ predicate on the "mileage_unit" field.
func MileageUnitNEQ(v MileageUnit) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldNEQ(FieldMileageUnit, v))
}

// MileageUnitIn applies the In predicate on the "mileage_unit" field.
func MileageUnitIn(vs ...MileageUnit) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldIn(FieldMileageUnit, vs...))
}

// MileageUnitNotIn applies the NotIn predicate on the "mileage_unit" field.
func MileageUnitNotIn(vs ...MileageUnit) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldNotIn(FieldMileageUnit, vs...))
}

// AddCustomerLocationEQ applies the EQ predicate on the "add_customer_location" field.
func AddCustomerLocationEQ(v bool) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldEQ(FieldAddCustomerLocation, v))
}

// AddCustomerLocationNEQ applies the NEQ predicate on the "add_customer_location" field.
func AddCustomerLocationNEQ(v bool) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldNEQ(FieldAddCustomerLocation, v))
}

// AutoGeocodeEQ applies the EQ predicate on the "auto_geocode" field.
func AutoGeocodeEQ(v bool) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldEQ(FieldAutoGeocode, v))
}

// AutoGeocodeNEQ applies the NEQ predicate on the "auto_geocode" field.
func AutoGeocodeNEQ(v bool) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldNEQ(FieldAutoGeocode, v))
}

// AddLocationEQ applies the EQ predicate on the "add_location" field.
func AddLocationEQ(v bool) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldEQ(FieldAddLocation, v))
}

// AddLocationNEQ applies the NEQ predicate on the "add_location" field.
func AddLocationNEQ(v bool) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldNEQ(FieldAddLocation, v))
}

// TrafficModelEQ applies the EQ predicate on the "traffic_model" field.
func TrafficModelEQ(v TrafficModel) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldEQ(FieldTrafficModel, v))
}

// TrafficModelNEQ applies the NEQ predicate on the "traffic_model" field.
func TrafficModelNEQ(v TrafficModel) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldNEQ(FieldTrafficModel, v))
}

// TrafficModelIn applies the In predicate on the "traffic_model" field.
func TrafficModelIn(vs ...TrafficModel) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldIn(FieldTrafficModel, vs...))
}

// TrafficModelNotIn applies the NotIn predicate on the "traffic_model" field.
func TrafficModelNotIn(vs ...TrafficModel) predicate.GoogleApi {
	return predicate.GoogleApi(sql.FieldNotIn(FieldTrafficModel, vs...))
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.GoogleApi {
	return predicate.GoogleApi(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.GoogleApi {
	return predicate.GoogleApi(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBusinessUnit applies the HasEdge predicate on the "business_unit" edge.
func HasBusinessUnit() predicate.GoogleApi {
	return predicate.GoogleApi(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessUnitWith applies the HasEdge predicate on the "business_unit" edge with a given conditions (other predicates).
func HasBusinessUnitWith(preds ...predicate.BusinessUnit) predicate.GoogleApi {
	return predicate.GoogleApi(func(s *sql.Selector) {
		step := newBusinessUnitStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GoogleApi) predicate.GoogleApi {
	return predicate.GoogleApi(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GoogleApi) predicate.GoogleApi {
	return predicate.GoogleApi(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GoogleApi) predicate.GoogleApi {
	return predicate.GoogleApi(sql.NotPredicates(p))
}