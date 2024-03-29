// Code generated by ent, DO NOT EDIT.

package hazardousmaterialsegregation

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldLTE(FieldID, id))
}

// BusinessUnitID applies equality check predicate on the "business_unit_id" field. It's identical to BusinessUnitIDEQ.
func BusinessUnitID(v uuid.UUID) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldEQ(FieldBusinessUnitID, v))
}

// OrganizationID applies equality check predicate on the "organization_id" field. It's identical to OrganizationIDEQ.
func OrganizationID(v uuid.UUID) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldEQ(FieldOrganizationID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldEQ(FieldUpdatedAt, v))
}

// BusinessUnitIDEQ applies the EQ predicate on the "business_unit_id" field.
func BusinessUnitIDEQ(v uuid.UUID) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDNEQ applies the NEQ predicate on the "business_unit_id" field.
func BusinessUnitIDNEQ(v uuid.UUID) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldNEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDIn applies the In predicate on the "business_unit_id" field.
func BusinessUnitIDIn(vs ...uuid.UUID) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldIn(FieldBusinessUnitID, vs...))
}

// BusinessUnitIDNotIn applies the NotIn predicate on the "business_unit_id" field.
func BusinessUnitIDNotIn(vs ...uuid.UUID) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldNotIn(FieldBusinessUnitID, vs...))
}

// OrganizationIDEQ applies the EQ predicate on the "organization_id" field.
func OrganizationIDEQ(v uuid.UUID) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldEQ(FieldOrganizationID, v))
}

// OrganizationIDNEQ applies the NEQ predicate on the "organization_id" field.
func OrganizationIDNEQ(v uuid.UUID) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldNEQ(FieldOrganizationID, v))
}

// OrganizationIDIn applies the In predicate on the "organization_id" field.
func OrganizationIDIn(vs ...uuid.UUID) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldIn(FieldOrganizationID, vs...))
}

// OrganizationIDNotIn applies the NotIn predicate on the "organization_id" field.
func OrganizationIDNotIn(vs ...uuid.UUID) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldNotIn(FieldOrganizationID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldLTE(FieldUpdatedAt, v))
}

// ClassAEQ applies the EQ predicate on the "class_a" field.
func ClassAEQ(v ClassA) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldEQ(FieldClassA, v))
}

// ClassANEQ applies the NEQ predicate on the "class_a" field.
func ClassANEQ(v ClassA) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldNEQ(FieldClassA, v))
}

// ClassAIn applies the In predicate on the "class_a" field.
func ClassAIn(vs ...ClassA) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldIn(FieldClassA, vs...))
}

// ClassANotIn applies the NotIn predicate on the "class_a" field.
func ClassANotIn(vs ...ClassA) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldNotIn(FieldClassA, vs...))
}

// ClassBEQ applies the EQ predicate on the "class_b" field.
func ClassBEQ(v ClassB) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldEQ(FieldClassB, v))
}

// ClassBNEQ applies the NEQ predicate on the "class_b" field.
func ClassBNEQ(v ClassB) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldNEQ(FieldClassB, v))
}

// ClassBIn applies the In predicate on the "class_b" field.
func ClassBIn(vs ...ClassB) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldIn(FieldClassB, vs...))
}

// ClassBNotIn applies the NotIn predicate on the "class_b" field.
func ClassBNotIn(vs ...ClassB) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldNotIn(FieldClassB, vs...))
}

// SegregationTypeEQ applies the EQ predicate on the "segregation_type" field.
func SegregationTypeEQ(v SegregationType) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldEQ(FieldSegregationType, v))
}

// SegregationTypeNEQ applies the NEQ predicate on the "segregation_type" field.
func SegregationTypeNEQ(v SegregationType) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldNEQ(FieldSegregationType, v))
}

// SegregationTypeIn applies the In predicate on the "segregation_type" field.
func SegregationTypeIn(vs ...SegregationType) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldIn(FieldSegregationType, vs...))
}

// SegregationTypeNotIn applies the NotIn predicate on the "segregation_type" field.
func SegregationTypeNotIn(vs ...SegregationType) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.FieldNotIn(FieldSegregationType, vs...))
}

// HasBusinessUnit applies the HasEdge predicate on the "business_unit" edge.
func HasBusinessUnit() predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessUnitWith applies the HasEdge predicate on the "business_unit" edge with a given conditions (other predicates).
func HasBusinessUnitWith(preds ...predicate.BusinessUnit) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(func(s *sql.Selector) {
		step := newBusinessUnitStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.HazardousMaterialSegregation) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.HazardousMaterialSegregation) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.HazardousMaterialSegregation) predicate.HazardousMaterialSegregation {
	return predicate.HazardousMaterialSegregation(sql.NotPredicates(p))
}
