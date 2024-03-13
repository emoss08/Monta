package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Organization holds the schema definition for the Organization entity.
type Organization struct {
	ent.Schema
}

// Fields of the Organization.
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MaxLen(100),
		field.String("scac_code").
			MaxLen(4).
			StructTag(`json:"scacCode"`),
		field.String("dot_number").
			MaxLen(12).
			StructTag(`json:"dotNumber"`),
		field.String("logo_url").
			Optional().
			StructTag(`json:"logoUrl"`),
		field.Enum("org_type").
			Values("A", "B", "X").
			Default("A").
			StructTag(`json:"orgType"`),
		field.Enum("timezone").
			Values("TimezoneAmericaLosAngeles", "TimezoneAmericaDenver", "TimezoneAmericaChicago", "TimezoneAmericaNewYork").
			Default("TimezoneAmericaLosAngeles"),
	}
}

// Edges of the Organization.
func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("business_unit", BusinessUnit.Type).Ref("organizations").
			Required().
			Unique(),
		edge.To("accounting_control", AccountingControl.Type).
			Required().
			Unique(),
		edge.To("billing_control", BillingControl.Type).
			Required().
			Unique(),
		edge.To("dispatch_control", DispatchControl.Type).
			Required().
			Unique(),
	}
}

// Mixin of the Organization.
func (Organization) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DefaultMixin{},
	}
}
