// Code generated by ent, DO NOT EDIT.

package invoicecontrol

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldUpdatedAt, v))
}

// InvoiceNumberPrefix applies equality check predicate on the "invoice_number_prefix" field. It's identical to InvoiceNumberPrefixEQ.
func InvoiceNumberPrefix(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldInvoiceNumberPrefix, v))
}

// CreditMemoNumberPrefix applies equality check predicate on the "credit_memo_number_prefix" field. It's identical to CreditMemoNumberPrefixEQ.
func CreditMemoNumberPrefix(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldCreditMemoNumberPrefix, v))
}

// InvoiceTerms applies equality check predicate on the "invoice_terms" field. It's identical to InvoiceTermsEQ.
func InvoiceTerms(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldInvoiceTerms, v))
}

// InvoiceFooter applies equality check predicate on the "invoice_footer" field. It's identical to InvoiceFooterEQ.
func InvoiceFooter(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldInvoiceFooter, v))
}

// InvoiceLogoURL applies equality check predicate on the "invoice_logo_url" field. It's identical to InvoiceLogoURLEQ.
func InvoiceLogoURL(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldInvoiceLogoURL, v))
}

// InvoiceDueAfterDays applies equality check predicate on the "invoice_due_after_days" field. It's identical to InvoiceDueAfterDaysEQ.
func InvoiceDueAfterDays(v uint8) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldInvoiceDueAfterDays, v))
}

// InvoiceLogoWidth applies equality check predicate on the "invoice_logo_width" field. It's identical to InvoiceLogoWidthEQ.
func InvoiceLogoWidth(v uint16) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldInvoiceLogoWidth, v))
}

// ShowAmountDue applies equality check predicate on the "show_amount_due" field. It's identical to ShowAmountDueEQ.
func ShowAmountDue(v bool) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldShowAmountDue, v))
}

// AttachPdf applies equality check predicate on the "attach_pdf" field. It's identical to AttachPdfEQ.
func AttachPdf(v bool) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldAttachPdf, v))
}

// ShowInvoiceDueDate applies equality check predicate on the "show_invoice_due_date" field. It's identical to ShowInvoiceDueDateEQ.
func ShowInvoiceDueDate(v bool) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldShowInvoiceDueDate, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldLTE(FieldUpdatedAt, v))
}

// InvoiceNumberPrefixEQ applies the EQ predicate on the "invoice_number_prefix" field.
func InvoiceNumberPrefixEQ(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldInvoiceNumberPrefix, v))
}

// InvoiceNumberPrefixNEQ applies the NEQ predicate on the "invoice_number_prefix" field.
func InvoiceNumberPrefixNEQ(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNEQ(FieldInvoiceNumberPrefix, v))
}

// InvoiceNumberPrefixIn applies the In predicate on the "invoice_number_prefix" field.
func InvoiceNumberPrefixIn(vs ...string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldIn(FieldInvoiceNumberPrefix, vs...))
}

// InvoiceNumberPrefixNotIn applies the NotIn predicate on the "invoice_number_prefix" field.
func InvoiceNumberPrefixNotIn(vs ...string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNotIn(FieldInvoiceNumberPrefix, vs...))
}

// InvoiceNumberPrefixGT applies the GT predicate on the "invoice_number_prefix" field.
func InvoiceNumberPrefixGT(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldGT(FieldInvoiceNumberPrefix, v))
}

// InvoiceNumberPrefixGTE applies the GTE predicate on the "invoice_number_prefix" field.
func InvoiceNumberPrefixGTE(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldGTE(FieldInvoiceNumberPrefix, v))
}

// InvoiceNumberPrefixLT applies the LT predicate on the "invoice_number_prefix" field.
func InvoiceNumberPrefixLT(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldLT(FieldInvoiceNumberPrefix, v))
}

// InvoiceNumberPrefixLTE applies the LTE predicate on the "invoice_number_prefix" field.
func InvoiceNumberPrefixLTE(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldLTE(FieldInvoiceNumberPrefix, v))
}

// InvoiceNumberPrefixContains applies the Contains predicate on the "invoice_number_prefix" field.
func InvoiceNumberPrefixContains(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldContains(FieldInvoiceNumberPrefix, v))
}

// InvoiceNumberPrefixHasPrefix applies the HasPrefix predicate on the "invoice_number_prefix" field.
func InvoiceNumberPrefixHasPrefix(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldHasPrefix(FieldInvoiceNumberPrefix, v))
}

// InvoiceNumberPrefixHasSuffix applies the HasSuffix predicate on the "invoice_number_prefix" field.
func InvoiceNumberPrefixHasSuffix(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldHasSuffix(FieldInvoiceNumberPrefix, v))
}

// InvoiceNumberPrefixEqualFold applies the EqualFold predicate on the "invoice_number_prefix" field.
func InvoiceNumberPrefixEqualFold(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEqualFold(FieldInvoiceNumberPrefix, v))
}

// InvoiceNumberPrefixContainsFold applies the ContainsFold predicate on the "invoice_number_prefix" field.
func InvoiceNumberPrefixContainsFold(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldContainsFold(FieldInvoiceNumberPrefix, v))
}

// CreditMemoNumberPrefixEQ applies the EQ predicate on the "credit_memo_number_prefix" field.
func CreditMemoNumberPrefixEQ(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldCreditMemoNumberPrefix, v))
}

// CreditMemoNumberPrefixNEQ applies the NEQ predicate on the "credit_memo_number_prefix" field.
func CreditMemoNumberPrefixNEQ(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNEQ(FieldCreditMemoNumberPrefix, v))
}

// CreditMemoNumberPrefixIn applies the In predicate on the "credit_memo_number_prefix" field.
func CreditMemoNumberPrefixIn(vs ...string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldIn(FieldCreditMemoNumberPrefix, vs...))
}

// CreditMemoNumberPrefixNotIn applies the NotIn predicate on the "credit_memo_number_prefix" field.
func CreditMemoNumberPrefixNotIn(vs ...string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNotIn(FieldCreditMemoNumberPrefix, vs...))
}

// CreditMemoNumberPrefixGT applies the GT predicate on the "credit_memo_number_prefix" field.
func CreditMemoNumberPrefixGT(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldGT(FieldCreditMemoNumberPrefix, v))
}

// CreditMemoNumberPrefixGTE applies the GTE predicate on the "credit_memo_number_prefix" field.
func CreditMemoNumberPrefixGTE(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldGTE(FieldCreditMemoNumberPrefix, v))
}

// CreditMemoNumberPrefixLT applies the LT predicate on the "credit_memo_number_prefix" field.
func CreditMemoNumberPrefixLT(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldLT(FieldCreditMemoNumberPrefix, v))
}

// CreditMemoNumberPrefixLTE applies the LTE predicate on the "credit_memo_number_prefix" field.
func CreditMemoNumberPrefixLTE(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldLTE(FieldCreditMemoNumberPrefix, v))
}

// CreditMemoNumberPrefixContains applies the Contains predicate on the "credit_memo_number_prefix" field.
func CreditMemoNumberPrefixContains(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldContains(FieldCreditMemoNumberPrefix, v))
}

// CreditMemoNumberPrefixHasPrefix applies the HasPrefix predicate on the "credit_memo_number_prefix" field.
func CreditMemoNumberPrefixHasPrefix(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldHasPrefix(FieldCreditMemoNumberPrefix, v))
}

// CreditMemoNumberPrefixHasSuffix applies the HasSuffix predicate on the "credit_memo_number_prefix" field.
func CreditMemoNumberPrefixHasSuffix(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldHasSuffix(FieldCreditMemoNumberPrefix, v))
}

// CreditMemoNumberPrefixEqualFold applies the EqualFold predicate on the "credit_memo_number_prefix" field.
func CreditMemoNumberPrefixEqualFold(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEqualFold(FieldCreditMemoNumberPrefix, v))
}

// CreditMemoNumberPrefixContainsFold applies the ContainsFold predicate on the "credit_memo_number_prefix" field.
func CreditMemoNumberPrefixContainsFold(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldContainsFold(FieldCreditMemoNumberPrefix, v))
}

// InvoiceTermsEQ applies the EQ predicate on the "invoice_terms" field.
func InvoiceTermsEQ(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldInvoiceTerms, v))
}

// InvoiceTermsNEQ applies the NEQ predicate on the "invoice_terms" field.
func InvoiceTermsNEQ(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNEQ(FieldInvoiceTerms, v))
}

// InvoiceTermsIn applies the In predicate on the "invoice_terms" field.
func InvoiceTermsIn(vs ...string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldIn(FieldInvoiceTerms, vs...))
}

// InvoiceTermsNotIn applies the NotIn predicate on the "invoice_terms" field.
func InvoiceTermsNotIn(vs ...string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNotIn(FieldInvoiceTerms, vs...))
}

// InvoiceTermsGT applies the GT predicate on the "invoice_terms" field.
func InvoiceTermsGT(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldGT(FieldInvoiceTerms, v))
}

// InvoiceTermsGTE applies the GTE predicate on the "invoice_terms" field.
func InvoiceTermsGTE(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldGTE(FieldInvoiceTerms, v))
}

// InvoiceTermsLT applies the LT predicate on the "invoice_terms" field.
func InvoiceTermsLT(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldLT(FieldInvoiceTerms, v))
}

// InvoiceTermsLTE applies the LTE predicate on the "invoice_terms" field.
func InvoiceTermsLTE(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldLTE(FieldInvoiceTerms, v))
}

// InvoiceTermsContains applies the Contains predicate on the "invoice_terms" field.
func InvoiceTermsContains(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldContains(FieldInvoiceTerms, v))
}

// InvoiceTermsHasPrefix applies the HasPrefix predicate on the "invoice_terms" field.
func InvoiceTermsHasPrefix(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldHasPrefix(FieldInvoiceTerms, v))
}

// InvoiceTermsHasSuffix applies the HasSuffix predicate on the "invoice_terms" field.
func InvoiceTermsHasSuffix(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldHasSuffix(FieldInvoiceTerms, v))
}

// InvoiceTermsIsNil applies the IsNil predicate on the "invoice_terms" field.
func InvoiceTermsIsNil() predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldIsNull(FieldInvoiceTerms))
}

// InvoiceTermsNotNil applies the NotNil predicate on the "invoice_terms" field.
func InvoiceTermsNotNil() predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNotNull(FieldInvoiceTerms))
}

// InvoiceTermsEqualFold applies the EqualFold predicate on the "invoice_terms" field.
func InvoiceTermsEqualFold(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEqualFold(FieldInvoiceTerms, v))
}

// InvoiceTermsContainsFold applies the ContainsFold predicate on the "invoice_terms" field.
func InvoiceTermsContainsFold(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldContainsFold(FieldInvoiceTerms, v))
}

// InvoiceFooterEQ applies the EQ predicate on the "invoice_footer" field.
func InvoiceFooterEQ(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldInvoiceFooter, v))
}

// InvoiceFooterNEQ applies the NEQ predicate on the "invoice_footer" field.
func InvoiceFooterNEQ(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNEQ(FieldInvoiceFooter, v))
}

// InvoiceFooterIn applies the In predicate on the "invoice_footer" field.
func InvoiceFooterIn(vs ...string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldIn(FieldInvoiceFooter, vs...))
}

// InvoiceFooterNotIn applies the NotIn predicate on the "invoice_footer" field.
func InvoiceFooterNotIn(vs ...string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNotIn(FieldInvoiceFooter, vs...))
}

// InvoiceFooterGT applies the GT predicate on the "invoice_footer" field.
func InvoiceFooterGT(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldGT(FieldInvoiceFooter, v))
}

// InvoiceFooterGTE applies the GTE predicate on the "invoice_footer" field.
func InvoiceFooterGTE(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldGTE(FieldInvoiceFooter, v))
}

// InvoiceFooterLT applies the LT predicate on the "invoice_footer" field.
func InvoiceFooterLT(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldLT(FieldInvoiceFooter, v))
}

// InvoiceFooterLTE applies the LTE predicate on the "invoice_footer" field.
func InvoiceFooterLTE(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldLTE(FieldInvoiceFooter, v))
}

// InvoiceFooterContains applies the Contains predicate on the "invoice_footer" field.
func InvoiceFooterContains(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldContains(FieldInvoiceFooter, v))
}

// InvoiceFooterHasPrefix applies the HasPrefix predicate on the "invoice_footer" field.
func InvoiceFooterHasPrefix(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldHasPrefix(FieldInvoiceFooter, v))
}

// InvoiceFooterHasSuffix applies the HasSuffix predicate on the "invoice_footer" field.
func InvoiceFooterHasSuffix(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldHasSuffix(FieldInvoiceFooter, v))
}

// InvoiceFooterIsNil applies the IsNil predicate on the "invoice_footer" field.
func InvoiceFooterIsNil() predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldIsNull(FieldInvoiceFooter))
}

// InvoiceFooterNotNil applies the NotNil predicate on the "invoice_footer" field.
func InvoiceFooterNotNil() predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNotNull(FieldInvoiceFooter))
}

// InvoiceFooterEqualFold applies the EqualFold predicate on the "invoice_footer" field.
func InvoiceFooterEqualFold(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEqualFold(FieldInvoiceFooter, v))
}

// InvoiceFooterContainsFold applies the ContainsFold predicate on the "invoice_footer" field.
func InvoiceFooterContainsFold(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldContainsFold(FieldInvoiceFooter, v))
}

// InvoiceLogoURLEQ applies the EQ predicate on the "invoice_logo_url" field.
func InvoiceLogoURLEQ(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldInvoiceLogoURL, v))
}

// InvoiceLogoURLNEQ applies the NEQ predicate on the "invoice_logo_url" field.
func InvoiceLogoURLNEQ(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNEQ(FieldInvoiceLogoURL, v))
}

// InvoiceLogoURLIn applies the In predicate on the "invoice_logo_url" field.
func InvoiceLogoURLIn(vs ...string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldIn(FieldInvoiceLogoURL, vs...))
}

// InvoiceLogoURLNotIn applies the NotIn predicate on the "invoice_logo_url" field.
func InvoiceLogoURLNotIn(vs ...string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNotIn(FieldInvoiceLogoURL, vs...))
}

// InvoiceLogoURLGT applies the GT predicate on the "invoice_logo_url" field.
func InvoiceLogoURLGT(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldGT(FieldInvoiceLogoURL, v))
}

// InvoiceLogoURLGTE applies the GTE predicate on the "invoice_logo_url" field.
func InvoiceLogoURLGTE(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldGTE(FieldInvoiceLogoURL, v))
}

// InvoiceLogoURLLT applies the LT predicate on the "invoice_logo_url" field.
func InvoiceLogoURLLT(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldLT(FieldInvoiceLogoURL, v))
}

// InvoiceLogoURLLTE applies the LTE predicate on the "invoice_logo_url" field.
func InvoiceLogoURLLTE(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldLTE(FieldInvoiceLogoURL, v))
}

// InvoiceLogoURLContains applies the Contains predicate on the "invoice_logo_url" field.
func InvoiceLogoURLContains(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldContains(FieldInvoiceLogoURL, v))
}

// InvoiceLogoURLHasPrefix applies the HasPrefix predicate on the "invoice_logo_url" field.
func InvoiceLogoURLHasPrefix(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldHasPrefix(FieldInvoiceLogoURL, v))
}

// InvoiceLogoURLHasSuffix applies the HasSuffix predicate on the "invoice_logo_url" field.
func InvoiceLogoURLHasSuffix(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldHasSuffix(FieldInvoiceLogoURL, v))
}

// InvoiceLogoURLIsNil applies the IsNil predicate on the "invoice_logo_url" field.
func InvoiceLogoURLIsNil() predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldIsNull(FieldInvoiceLogoURL))
}

// InvoiceLogoURLNotNil applies the NotNil predicate on the "invoice_logo_url" field.
func InvoiceLogoURLNotNil() predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNotNull(FieldInvoiceLogoURL))
}

// InvoiceLogoURLEqualFold applies the EqualFold predicate on the "invoice_logo_url" field.
func InvoiceLogoURLEqualFold(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEqualFold(FieldInvoiceLogoURL, v))
}

// InvoiceLogoURLContainsFold applies the ContainsFold predicate on the "invoice_logo_url" field.
func InvoiceLogoURLContainsFold(v string) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldContainsFold(FieldInvoiceLogoURL, v))
}

// InvoiceDateFormatEQ applies the EQ predicate on the "invoice_date_format" field.
func InvoiceDateFormatEQ(v InvoiceDateFormat) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldInvoiceDateFormat, v))
}

// InvoiceDateFormatNEQ applies the NEQ predicate on the "invoice_date_format" field.
func InvoiceDateFormatNEQ(v InvoiceDateFormat) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNEQ(FieldInvoiceDateFormat, v))
}

// InvoiceDateFormatIn applies the In predicate on the "invoice_date_format" field.
func InvoiceDateFormatIn(vs ...InvoiceDateFormat) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldIn(FieldInvoiceDateFormat, vs...))
}

// InvoiceDateFormatNotIn applies the NotIn predicate on the "invoice_date_format" field.
func InvoiceDateFormatNotIn(vs ...InvoiceDateFormat) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNotIn(FieldInvoiceDateFormat, vs...))
}

// InvoiceDueAfterDaysEQ applies the EQ predicate on the "invoice_due_after_days" field.
func InvoiceDueAfterDaysEQ(v uint8) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldInvoiceDueAfterDays, v))
}

// InvoiceDueAfterDaysNEQ applies the NEQ predicate on the "invoice_due_after_days" field.
func InvoiceDueAfterDaysNEQ(v uint8) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNEQ(FieldInvoiceDueAfterDays, v))
}

// InvoiceDueAfterDaysIn applies the In predicate on the "invoice_due_after_days" field.
func InvoiceDueAfterDaysIn(vs ...uint8) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldIn(FieldInvoiceDueAfterDays, vs...))
}

// InvoiceDueAfterDaysNotIn applies the NotIn predicate on the "invoice_due_after_days" field.
func InvoiceDueAfterDaysNotIn(vs ...uint8) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNotIn(FieldInvoiceDueAfterDays, vs...))
}

// InvoiceDueAfterDaysGT applies the GT predicate on the "invoice_due_after_days" field.
func InvoiceDueAfterDaysGT(v uint8) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldGT(FieldInvoiceDueAfterDays, v))
}

// InvoiceDueAfterDaysGTE applies the GTE predicate on the "invoice_due_after_days" field.
func InvoiceDueAfterDaysGTE(v uint8) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldGTE(FieldInvoiceDueAfterDays, v))
}

// InvoiceDueAfterDaysLT applies the LT predicate on the "invoice_due_after_days" field.
func InvoiceDueAfterDaysLT(v uint8) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldLT(FieldInvoiceDueAfterDays, v))
}

// InvoiceDueAfterDaysLTE applies the LTE predicate on the "invoice_due_after_days" field.
func InvoiceDueAfterDaysLTE(v uint8) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldLTE(FieldInvoiceDueAfterDays, v))
}

// InvoiceLogoWidthEQ applies the EQ predicate on the "invoice_logo_width" field.
func InvoiceLogoWidthEQ(v uint16) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldInvoiceLogoWidth, v))
}

// InvoiceLogoWidthNEQ applies the NEQ predicate on the "invoice_logo_width" field.
func InvoiceLogoWidthNEQ(v uint16) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNEQ(FieldInvoiceLogoWidth, v))
}

// InvoiceLogoWidthIn applies the In predicate on the "invoice_logo_width" field.
func InvoiceLogoWidthIn(vs ...uint16) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldIn(FieldInvoiceLogoWidth, vs...))
}

// InvoiceLogoWidthNotIn applies the NotIn predicate on the "invoice_logo_width" field.
func InvoiceLogoWidthNotIn(vs ...uint16) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNotIn(FieldInvoiceLogoWidth, vs...))
}

// InvoiceLogoWidthGT applies the GT predicate on the "invoice_logo_width" field.
func InvoiceLogoWidthGT(v uint16) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldGT(FieldInvoiceLogoWidth, v))
}

// InvoiceLogoWidthGTE applies the GTE predicate on the "invoice_logo_width" field.
func InvoiceLogoWidthGTE(v uint16) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldGTE(FieldInvoiceLogoWidth, v))
}

// InvoiceLogoWidthLT applies the LT predicate on the "invoice_logo_width" field.
func InvoiceLogoWidthLT(v uint16) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldLT(FieldInvoiceLogoWidth, v))
}

// InvoiceLogoWidthLTE applies the LTE predicate on the "invoice_logo_width" field.
func InvoiceLogoWidthLTE(v uint16) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldLTE(FieldInvoiceLogoWidth, v))
}

// ShowAmountDueEQ applies the EQ predicate on the "show_amount_due" field.
func ShowAmountDueEQ(v bool) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldShowAmountDue, v))
}

// ShowAmountDueNEQ applies the NEQ predicate on the "show_amount_due" field.
func ShowAmountDueNEQ(v bool) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNEQ(FieldShowAmountDue, v))
}

// AttachPdfEQ applies the EQ predicate on the "attach_pdf" field.
func AttachPdfEQ(v bool) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldAttachPdf, v))
}

// AttachPdfNEQ applies the NEQ predicate on the "attach_pdf" field.
func AttachPdfNEQ(v bool) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNEQ(FieldAttachPdf, v))
}

// ShowInvoiceDueDateEQ applies the EQ predicate on the "show_invoice_due_date" field.
func ShowInvoiceDueDateEQ(v bool) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldEQ(FieldShowInvoiceDueDate, v))
}

// ShowInvoiceDueDateNEQ applies the NEQ predicate on the "show_invoice_due_date" field.
func ShowInvoiceDueDateNEQ(v bool) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.FieldNEQ(FieldShowInvoiceDueDate, v))
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.InvoiceControl {
	return predicate.InvoiceControl(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.InvoiceControl {
	return predicate.InvoiceControl(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBusinessUnit applies the HasEdge predicate on the "business_unit" edge.
func HasBusinessUnit() predicate.InvoiceControl {
	return predicate.InvoiceControl(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessUnitWith applies the HasEdge predicate on the "business_unit" edge with a given conditions (other predicates).
func HasBusinessUnitWith(preds ...predicate.BusinessUnit) predicate.InvoiceControl {
	return predicate.InvoiceControl(func(s *sql.Selector) {
		step := newBusinessUnitStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.InvoiceControl) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.InvoiceControl) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.InvoiceControl) predicate.InvoiceControl {
	return predicate.InvoiceControl(sql.NotPredicates(p))
}