package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// InvoiceControl holds the schema definition for the InvoiceControl entity.
type InvoiceControl struct {
	ent.Schema
}

// Fields of the InvoiceControl.
func (InvoiceControl) Fields() []ent.Field {
	return []ent.Field{
		field.String("invoice_number_prefix").
			Default("INV-").
			MaxLen(10).
			StructTag(`json:"invoiceNumberPrefix"`),
		field.String("credit_memo_number_prefix").
			Default("CM-").
			MaxLen(10).
			StructTag(`json:"creditMemoNumberPrefix"`),
		field.Text("invoice_terms").
			Optional().
			Nillable().
			StructTag(`json:"invoiceTerms"`),
		field.Text("invoice_footer").
			Optional().
			Nillable().
			StructTag(`json:"invoiceFooter"`),
		field.String("invoice_logo_url").
			Optional().
			MaxLen(255).
			Nillable().
			StructTag(`json:"invoiceLogoUrl"`),
		field.Enum("invoice_date_format").
			Values("InvoiceDateFormatMDY", "InvoiceDateFormatDMY", "InvoiceDateFormatYMD", "InvoiceDateFormatYDM").
			Default("InvoiceDateFormatMDY").
			StructTag(`json:"invoiceDateFormat"`),
		field.Uint8("invoice_due_after_days").
			Default(30).
			Positive().
			StructTag(`json:"invoiceDueAfterDays"`),
		field.Uint16("invoice_logo_width").
			Default(100).
			Positive().
			StructTag(`json:"invoiceLogoWidth"`),
		field.Bool("show_amount_due").
			Default(true).
			StructTag(`json:"showAmountDue"`),
		field.Bool("attach_pdf").
			Default(true).
			StructTag(`json:"attachPdf"`),
		field.Bool("show_invoice_due_date").
			Default(true).
			StructTag(`json:"showInvoiceDueDate"`),
	}
}

// Mixin for the InvoiceControl.
func (InvoiceControl) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DefaultMixin{},
	}
}

func (InvoiceControl) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}

// Edges of the InvoiceControl.
func (InvoiceControl) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organization", Organization.Type).
			Ref("invoice_control").
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Required().
			Unique(),
		edge.To("business_unit", BusinessUnit.Type).
			StorageKey(edge.Column("business_unit_id")).
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Required().
			Unique(),
	}
}
