// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/emoss08/trenova/ent/accountingcontrol"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/commodity"
	"github.com/emoss08/trenova/ent/dispatchcontrol"
	"github.com/emoss08/trenova/ent/feasibilitytoolcontrol"
	"github.com/emoss08/trenova/ent/generalledgeraccount"
	"github.com/emoss08/trenova/ent/hazardousmaterial"
	"github.com/emoss08/trenova/ent/invoicecontrol"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/routecontrol"
	"github.com/emoss08/trenova/ent/tablechangealert"
	"github.com/emoss08/trenova/ent/user"
	"github.com/google/uuid"
)

// CreateAccountingControlInput represents a mutation input for creating accountingcontrols.
type CreateAccountingControlInput struct {
	CreatedAt                    *time.Time
	UpdatedAt                    *time.Time
	RecThreshold                 *int64
	RecThresholdAction           *accountingcontrol.RecThresholdAction
	AutoCreateJournalEntries     *bool
	RestrictManualJournalEntries *bool
	RequireJournalEntryApproval  *bool
	EnableRecNotifications       *bool
	HaltOnPendingRec             *bool
	CriticalProcesses            *string
	OrganizationID               uuid.UUID
	BusinessUnitID               uuid.UUID
	DefaultRevAccountID          *uuid.UUID
	DefaultExpAccountID          *uuid.UUID
}

// Mutate applies the CreateAccountingControlInput on the AccountingControlMutation builder.
func (i *CreateAccountingControlInput) Mutate(m *AccountingControlMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.RecThreshold; v != nil {
		m.SetRecThreshold(*v)
	}
	if v := i.RecThresholdAction; v != nil {
		m.SetRecThresholdAction(*v)
	}
	if v := i.AutoCreateJournalEntries; v != nil {
		m.SetAutoCreateJournalEntries(*v)
	}
	if v := i.RestrictManualJournalEntries; v != nil {
		m.SetRestrictManualJournalEntries(*v)
	}
	if v := i.RequireJournalEntryApproval; v != nil {
		m.SetRequireJournalEntryApproval(*v)
	}
	if v := i.EnableRecNotifications; v != nil {
		m.SetEnableRecNotifications(*v)
	}
	if v := i.HaltOnPendingRec; v != nil {
		m.SetHaltOnPendingRec(*v)
	}
	if v := i.CriticalProcesses; v != nil {
		m.SetCriticalProcesses(*v)
	}
	m.SetOrganizationID(i.OrganizationID)
	m.SetBusinessUnitID(i.BusinessUnitID)
	if v := i.DefaultRevAccountID; v != nil {
		m.SetDefaultRevAccountID(*v)
	}
	if v := i.DefaultExpAccountID; v != nil {
		m.SetDefaultExpAccountID(*v)
	}
}

// SetInput applies the change-set in the CreateAccountingControlInput on the AccountingControlCreate builder.
func (c *AccountingControlCreate) SetInput(i CreateAccountingControlInput) *AccountingControlCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateBusinessUnitInput represents a mutation input for creating businessunits.
type CreateBusinessUnitInput struct {
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
	Status           *businessunit.Status
	Name             string
	EntityKey        string
	PhoneNumber      string
	Address          *string
	City             *string
	State            *string
	Country          *string
	PostalCode       *string
	TaxID            *string
	SubscriptionPlan *string
	Description      *string
	LegalName        *string
	ContactName      *string
	ContactEmail     *string
	PaidUntil        *time.Time
	Settings         map[string]interface{}
	FreeTrial        *bool
	PrevID           *uuid.UUID
	NextID           *uuid.UUID
	OrganizationIDs  []uuid.UUID
}

// Mutate applies the CreateBusinessUnitInput on the BusinessUnitMutation builder.
func (i *CreateBusinessUnitInput) Mutate(m *BusinessUnitMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	m.SetName(i.Name)
	m.SetEntityKey(i.EntityKey)
	m.SetPhoneNumber(i.PhoneNumber)
	if v := i.Address; v != nil {
		m.SetAddress(*v)
	}
	if v := i.City; v != nil {
		m.SetCity(*v)
	}
	if v := i.State; v != nil {
		m.SetState(*v)
	}
	if v := i.Country; v != nil {
		m.SetCountry(*v)
	}
	if v := i.PostalCode; v != nil {
		m.SetPostalCode(*v)
	}
	if v := i.TaxID; v != nil {
		m.SetTaxID(*v)
	}
	if v := i.SubscriptionPlan; v != nil {
		m.SetSubscriptionPlan(*v)
	}
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	if v := i.LegalName; v != nil {
		m.SetLegalName(*v)
	}
	if v := i.ContactName; v != nil {
		m.SetContactName(*v)
	}
	if v := i.ContactEmail; v != nil {
		m.SetContactEmail(*v)
	}
	if v := i.PaidUntil; v != nil {
		m.SetPaidUntil(*v)
	}
	if v := i.Settings; v != nil {
		m.SetSettings(v)
	}
	if v := i.FreeTrial; v != nil {
		m.SetFreeTrial(*v)
	}
	if v := i.PrevID; v != nil {
		m.SetPrevID(*v)
	}
	if v := i.NextID; v != nil {
		m.SetNextID(*v)
	}
	if v := i.OrganizationIDs; len(v) > 0 {
		m.AddOrganizationIDs(v...)
	}
}

// SetInput applies the change-set in the CreateBusinessUnitInput on the BusinessUnitCreate builder.
func (c *BusinessUnitCreate) SetInput(i CreateBusinessUnitInput) *BusinessUnitCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateCommodityInput represents a mutation input for creating commodities.
type CreateCommodityInput struct {
	CreatedAt           *time.Time
	UpdatedAt           *time.Time
	Status              *commodity.Status
	Name                string
	IsHazmat            *bool
	UnitOfMeasure       *commodity.UnitOfMeasure
	MinTemp             *float64
	MaxTemp             *float64
	SetPointTemp        *float64
	Description         *string
	BusinessUnitID      uuid.UUID
	OrganizationID      uuid.UUID
	HazardousMaterialID uuid.UUID
}

// Mutate applies the CreateCommodityInput on the CommodityMutation builder.
func (i *CreateCommodityInput) Mutate(m *CommodityMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	m.SetName(i.Name)
	if v := i.IsHazmat; v != nil {
		m.SetIsHazmat(*v)
	}
	if v := i.UnitOfMeasure; v != nil {
		m.SetUnitOfMeasure(*v)
	}
	if v := i.MinTemp; v != nil {
		m.SetMinTemp(*v)
	}
	if v := i.MaxTemp; v != nil {
		m.SetMaxTemp(*v)
	}
	if v := i.SetPointTemp; v != nil {
		m.SetSetPointTemp(*v)
	}
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	m.SetBusinessUnitID(i.BusinessUnitID)
	m.SetOrganizationID(i.OrganizationID)
	m.SetHazardousMaterialID(i.HazardousMaterialID)
}

// SetInput applies the change-set in the CreateCommodityInput on the CommodityCreate builder.
func (c *CommodityCreate) SetInput(i CreateCommodityInput) *CommodityCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateDispatchControlInput represents a mutation input for creating dispatchcontrols.
type CreateDispatchControlInput struct {
	CreatedAt                    *time.Time
	UpdatedAt                    *time.Time
	RecordServiceIncident        *dispatchcontrol.RecordServiceIncident
	DeadheadTarget               *float64
	MaxShipmentWeightLimit       *int
	GracePeriod                  *uint8
	EnforceWorkerAssign          *bool
	TrailerContinuity            *bool
	DupeTrailerCheck             *bool
	MaintenanceCompliance        *bool
	RegulatoryCheck              *bool
	PrevShipmentOnHold           *bool
	WorkerTimeAwayRestriction    *bool
	TractorWorkerFleetConstraint *bool
	OrganizationID               uuid.UUID
	BusinessUnitID               uuid.UUID
}

// Mutate applies the CreateDispatchControlInput on the DispatchControlMutation builder.
func (i *CreateDispatchControlInput) Mutate(m *DispatchControlMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.RecordServiceIncident; v != nil {
		m.SetRecordServiceIncident(*v)
	}
	if v := i.DeadheadTarget; v != nil {
		m.SetDeadheadTarget(*v)
	}
	if v := i.MaxShipmentWeightLimit; v != nil {
		m.SetMaxShipmentWeightLimit(*v)
	}
	if v := i.GracePeriod; v != nil {
		m.SetGracePeriod(*v)
	}
	if v := i.EnforceWorkerAssign; v != nil {
		m.SetEnforceWorkerAssign(*v)
	}
	if v := i.TrailerContinuity; v != nil {
		m.SetTrailerContinuity(*v)
	}
	if v := i.DupeTrailerCheck; v != nil {
		m.SetDupeTrailerCheck(*v)
	}
	if v := i.MaintenanceCompliance; v != nil {
		m.SetMaintenanceCompliance(*v)
	}
	if v := i.RegulatoryCheck; v != nil {
		m.SetRegulatoryCheck(*v)
	}
	if v := i.PrevShipmentOnHold; v != nil {
		m.SetPrevShipmentOnHold(*v)
	}
	if v := i.WorkerTimeAwayRestriction; v != nil {
		m.SetWorkerTimeAwayRestriction(*v)
	}
	if v := i.TractorWorkerFleetConstraint; v != nil {
		m.SetTractorWorkerFleetConstraint(*v)
	}
	m.SetOrganizationID(i.OrganizationID)
	m.SetBusinessUnitID(i.BusinessUnitID)
}

// SetInput applies the change-set in the CreateDispatchControlInput on the DispatchControlCreate builder.
func (c *DispatchControlCreate) SetInput(i CreateDispatchControlInput) *DispatchControlCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateFeasibilityToolControlInput represents a mutation input for creating feasibilitytoolcontrols.
type CreateFeasibilityToolControlInput struct {
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
	OtpOperator    *feasibilitytoolcontrol.OtpOperator
	OtpValue       *float64
	MpwOperator    *feasibilitytoolcontrol.MpwOperator
	MpwValue       *float64
	MpdOperator    *feasibilitytoolcontrol.MpdOperator
	MpdValue       *float64
	MpgOperator    *feasibilitytoolcontrol.MpgOperator
	MpgValue       *float64
	OrganizationID uuid.UUID
	BusinessUnitID uuid.UUID
}

// Mutate applies the CreateFeasibilityToolControlInput on the FeasibilityToolControlMutation builder.
func (i *CreateFeasibilityToolControlInput) Mutate(m *FeasibilityToolControlMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.OtpOperator; v != nil {
		m.SetOtpOperator(*v)
	}
	if v := i.OtpValue; v != nil {
		m.SetOtpValue(*v)
	}
	if v := i.MpwOperator; v != nil {
		m.SetMpwOperator(*v)
	}
	if v := i.MpwValue; v != nil {
		m.SetMpwValue(*v)
	}
	if v := i.MpdOperator; v != nil {
		m.SetMpdOperator(*v)
	}
	if v := i.MpdValue; v != nil {
		m.SetMpdValue(*v)
	}
	if v := i.MpgOperator; v != nil {
		m.SetMpgOperator(*v)
	}
	if v := i.MpgValue; v != nil {
		m.SetMpgValue(*v)
	}
	m.SetOrganizationID(i.OrganizationID)
	m.SetBusinessUnitID(i.BusinessUnitID)
}

// SetInput applies the change-set in the CreateFeasibilityToolControlInput on the FeasibilityToolControlCreate builder.
func (c *FeasibilityToolControlCreate) SetInput(i CreateFeasibilityToolControlInput) *FeasibilityToolControlCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateGeneralLedgerAccountInput represents a mutation input for creating generalledgeraccounts.
type CreateGeneralLedgerAccountInput struct {
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
	Status         *generalledgeraccount.Status
	AccountNumber  string
	AccountType    generalledgeraccount.AccountType
	CashFlowType   *generalledgeraccount.CashFlowType
	AccountSubType *generalledgeraccount.AccountSubType
	AccountClass   *generalledgeraccount.AccountClass
	Balance        *float64
	InterestRate   *float64
	DateOpened     *time.Time
	DateClosed     *time.Time
	Notes          *string
	IsTaxRelevant  *bool
	IsReconciled   *bool
	BusinessUnitID uuid.UUID
	OrganizationID uuid.UUID
	TagIDs         []uuid.UUID
}

// Mutate applies the CreateGeneralLedgerAccountInput on the GeneralLedgerAccountMutation builder.
func (i *CreateGeneralLedgerAccountInput) Mutate(m *GeneralLedgerAccountMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	m.SetAccountNumber(i.AccountNumber)
	m.SetAccountType(i.AccountType)
	if v := i.CashFlowType; v != nil {
		m.SetCashFlowType(*v)
	}
	if v := i.AccountSubType; v != nil {
		m.SetAccountSubType(*v)
	}
	if v := i.AccountClass; v != nil {
		m.SetAccountClass(*v)
	}
	if v := i.Balance; v != nil {
		m.SetBalance(*v)
	}
	if v := i.InterestRate; v != nil {
		m.SetInterestRate(*v)
	}
	if v := i.DateOpened; v != nil {
		m.SetDateOpened(*v)
	}
	if v := i.DateClosed; v != nil {
		m.SetDateClosed(*v)
	}
	if v := i.Notes; v != nil {
		m.SetNotes(*v)
	}
	if v := i.IsTaxRelevant; v != nil {
		m.SetIsTaxRelevant(*v)
	}
	if v := i.IsReconciled; v != nil {
		m.SetIsReconciled(*v)
	}
	m.SetBusinessUnitID(i.BusinessUnitID)
	m.SetOrganizationID(i.OrganizationID)
	if v := i.TagIDs; len(v) > 0 {
		m.AddTagIDs(v...)
	}
}

// SetInput applies the change-set in the CreateGeneralLedgerAccountInput on the GeneralLedgerAccountCreate builder.
func (c *GeneralLedgerAccountCreate) SetInput(i CreateGeneralLedgerAccountInput) *GeneralLedgerAccountCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateHazardousMaterialInput represents a mutation input for creating hazardousmaterials.
type CreateHazardousMaterialInput struct {
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
	Name               string
	HazardClass        *hazardousmaterial.HazardClass
	ErgNumber          *string
	Description        *string
	PackingGroup       *hazardousmaterial.PackingGroup
	ProperShippingName *string
	BusinessUnitID     uuid.UUID
	OrganizationID     uuid.UUID
	CommodityIDs       []uuid.UUID
}

// Mutate applies the CreateHazardousMaterialInput on the HazardousMaterialMutation builder.
func (i *CreateHazardousMaterialInput) Mutate(m *HazardousMaterialMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetName(i.Name)
	if v := i.HazardClass; v != nil {
		m.SetHazardClass(*v)
	}
	if v := i.ErgNumber; v != nil {
		m.SetErgNumber(*v)
	}
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	if v := i.PackingGroup; v != nil {
		m.SetPackingGroup(*v)
	}
	if v := i.ProperShippingName; v != nil {
		m.SetProperShippingName(*v)
	}
	m.SetBusinessUnitID(i.BusinessUnitID)
	m.SetOrganizationID(i.OrganizationID)
	if v := i.CommodityIDs; len(v) > 0 {
		m.AddCommodityIDs(v...)
	}
}

// SetInput applies the change-set in the CreateHazardousMaterialInput on the HazardousMaterialCreate builder.
func (c *HazardousMaterialCreate) SetInput(i CreateHazardousMaterialInput) *HazardousMaterialCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateInvoiceControlInput represents a mutation input for creating invoicecontrols.
type CreateInvoiceControlInput struct {
	CreatedAt              *time.Time
	UpdatedAt              *time.Time
	InvoiceNumberPrefix    *string
	CreditMemoNumberPrefix *string
	InvoiceTerms           *string
	InvoiceFooter          *string
	InvoiceLogoURL         *string
	InvoiceDateFormat      *invoicecontrol.InvoiceDateFormat
	InvoiceDueAfterDays    *uint8
	InvoiceLogoWidth       *uint16
	ShowAmountDue          *bool
	AttachPdf              *bool
	ShowInvoiceDueDate     *bool
	OrganizationID         uuid.UUID
	BusinessUnitID         uuid.UUID
}

// Mutate applies the CreateInvoiceControlInput on the InvoiceControlMutation builder.
func (i *CreateInvoiceControlInput) Mutate(m *InvoiceControlMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.InvoiceNumberPrefix; v != nil {
		m.SetInvoiceNumberPrefix(*v)
	}
	if v := i.CreditMemoNumberPrefix; v != nil {
		m.SetCreditMemoNumberPrefix(*v)
	}
	if v := i.InvoiceTerms; v != nil {
		m.SetInvoiceTerms(*v)
	}
	if v := i.InvoiceFooter; v != nil {
		m.SetInvoiceFooter(*v)
	}
	if v := i.InvoiceLogoURL; v != nil {
		m.SetInvoiceLogoURL(*v)
	}
	if v := i.InvoiceDateFormat; v != nil {
		m.SetInvoiceDateFormat(*v)
	}
	if v := i.InvoiceDueAfterDays; v != nil {
		m.SetInvoiceDueAfterDays(*v)
	}
	if v := i.InvoiceLogoWidth; v != nil {
		m.SetInvoiceLogoWidth(*v)
	}
	if v := i.ShowAmountDue; v != nil {
		m.SetShowAmountDue(*v)
	}
	if v := i.AttachPdf; v != nil {
		m.SetAttachPdf(*v)
	}
	if v := i.ShowInvoiceDueDate; v != nil {
		m.SetShowInvoiceDueDate(*v)
	}
	m.SetOrganizationID(i.OrganizationID)
	m.SetBusinessUnitID(i.BusinessUnitID)
}

// SetInput applies the change-set in the CreateInvoiceControlInput on the InvoiceControlCreate builder.
func (c *InvoiceControlCreate) SetInput(i CreateInvoiceControlInput) *InvoiceControlCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateOrganizationInput represents a mutation input for creating organizations.
type CreateOrganizationInput struct {
	CreatedAt                *time.Time
	UpdatedAt                *time.Time
	Name                     string
	ScacCode                 string
	DotNumber                string
	LogoURL                  *string
	OrgType                  *organization.OrgType
	Timezone                 *organization.Timezone
	BusinessUnitID           uuid.UUID
	AccountingControlID      *uuid.UUID
	BillingControlID         *uuid.UUID
	DispatchControlID        *uuid.UUID
	FeasibilityToolControlID *uuid.UUID
	InvoiceControlID         *uuid.UUID
	RouteControlID           *uuid.UUID
	ShipmentControlID        *uuid.UUID
}

// Mutate applies the CreateOrganizationInput on the OrganizationMutation builder.
func (i *CreateOrganizationInput) Mutate(m *OrganizationMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetName(i.Name)
	m.SetScacCode(i.ScacCode)
	m.SetDotNumber(i.DotNumber)
	if v := i.LogoURL; v != nil {
		m.SetLogoURL(*v)
	}
	if v := i.OrgType; v != nil {
		m.SetOrgType(*v)
	}
	if v := i.Timezone; v != nil {
		m.SetTimezone(*v)
	}
	m.SetBusinessUnitID(i.BusinessUnitID)
	if v := i.AccountingControlID; v != nil {
		m.SetAccountingControlID(*v)
	}
	if v := i.BillingControlID; v != nil {
		m.SetBillingControlID(*v)
	}
	if v := i.DispatchControlID; v != nil {
		m.SetDispatchControlID(*v)
	}
	if v := i.FeasibilityToolControlID; v != nil {
		m.SetFeasibilityToolControlID(*v)
	}
	if v := i.InvoiceControlID; v != nil {
		m.SetInvoiceControlID(*v)
	}
	if v := i.RouteControlID; v != nil {
		m.SetRouteControlID(*v)
	}
	if v := i.ShipmentControlID; v != nil {
		m.SetShipmentControlID(*v)
	}
}

// SetInput applies the change-set in the CreateOrganizationInput on the OrganizationCreate builder.
func (c *OrganizationCreate) SetInput(i CreateOrganizationInput) *OrganizationCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateRouteControlInput represents a mutation input for creating routecontrols.
type CreateRouteControlInput struct {
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
	DistanceMethod *routecontrol.DistanceMethod
	MileageUnit    *routecontrol.MileageUnit
	GenerateRoutes *bool
	OrganizationID uuid.UUID
	BusinessUnitID uuid.UUID
}

// Mutate applies the CreateRouteControlInput on the RouteControlMutation builder.
func (i *CreateRouteControlInput) Mutate(m *RouteControlMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DistanceMethod; v != nil {
		m.SetDistanceMethod(*v)
	}
	if v := i.MileageUnit; v != nil {
		m.SetMileageUnit(*v)
	}
	if v := i.GenerateRoutes; v != nil {
		m.SetGenerateRoutes(*v)
	}
	m.SetOrganizationID(i.OrganizationID)
	m.SetBusinessUnitID(i.BusinessUnitID)
}

// SetInput applies the change-set in the CreateRouteControlInput on the RouteControlCreate builder.
func (c *RouteControlCreate) SetInput(i CreateRouteControlInput) *RouteControlCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateShipmentControlInput represents a mutation input for creating shipmentcontrols.
type CreateShipmentControlInput struct {
	CreatedAt                *time.Time
	UpdatedAt                *time.Time
	AutoRateShipment         *bool
	CalculateDistance        *bool
	EnforceRevCode           *bool
	EnforceVoidedComm        *bool
	GenerateRoutes           *bool
	EnforceCommodity         *bool
	AutoSequenceStops        *bool
	AutoShipmentTotal        *bool
	EnforceOriginDestination *bool
	CheckForDuplicateBol     *bool
	SendPlacardInfo          *bool
	EnforceHazmatSegRules    *bool
	OrganizationID           uuid.UUID
	BusinessUnitID           uuid.UUID
}

// Mutate applies the CreateShipmentControlInput on the ShipmentControlMutation builder.
func (i *CreateShipmentControlInput) Mutate(m *ShipmentControlMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.AutoRateShipment; v != nil {
		m.SetAutoRateShipment(*v)
	}
	if v := i.CalculateDistance; v != nil {
		m.SetCalculateDistance(*v)
	}
	if v := i.EnforceRevCode; v != nil {
		m.SetEnforceRevCode(*v)
	}
	if v := i.EnforceVoidedComm; v != nil {
		m.SetEnforceVoidedComm(*v)
	}
	if v := i.GenerateRoutes; v != nil {
		m.SetGenerateRoutes(*v)
	}
	if v := i.EnforceCommodity; v != nil {
		m.SetEnforceCommodity(*v)
	}
	if v := i.AutoSequenceStops; v != nil {
		m.SetAutoSequenceStops(*v)
	}
	if v := i.AutoShipmentTotal; v != nil {
		m.SetAutoShipmentTotal(*v)
	}
	if v := i.EnforceOriginDestination; v != nil {
		m.SetEnforceOriginDestination(*v)
	}
	if v := i.CheckForDuplicateBol; v != nil {
		m.SetCheckForDuplicateBol(*v)
	}
	if v := i.SendPlacardInfo; v != nil {
		m.SetSendPlacardInfo(*v)
	}
	if v := i.EnforceHazmatSegRules; v != nil {
		m.SetEnforceHazmatSegRules(*v)
	}
	m.SetOrganizationID(i.OrganizationID)
	m.SetBusinessUnitID(i.BusinessUnitID)
}

// SetInput applies the change-set in the CreateShipmentControlInput on the ShipmentControlCreate builder.
func (c *ShipmentControlCreate) SetInput(i CreateShipmentControlInput) *ShipmentControlCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateTableChangeAlertInput represents a mutation input for creating tablechangealerts.
type CreateTableChangeAlertInput struct {
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
	Status          *tablechangealert.Status
	Name            string
	DatabaseAction  tablechangealert.DatabaseAction
	Source          tablechangealert.Source
	TableName       *string
	Topic           *string
	Description     *string
	CustomSubject   *string
	FunctionName    *string
	TriggerName     *string
	ListenerName    *string
	EmailRecipients *string
	EffectiveDate   *time.Time
	ExpirationDate  *time.Time
	BusinessUnitID  uuid.UUID
	OrganizationID  uuid.UUID
}

// Mutate applies the CreateTableChangeAlertInput on the TableChangeAlertMutation builder.
func (i *CreateTableChangeAlertInput) Mutate(m *TableChangeAlertMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	m.SetName(i.Name)
	m.SetDatabaseAction(i.DatabaseAction)
	m.SetSource(i.Source)
	if v := i.TableName; v != nil {
		m.SetTableName(*v)
	}
	if v := i.Topic; v != nil {
		m.SetTopic(*v)
	}
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	if v := i.CustomSubject; v != nil {
		m.SetCustomSubject(*v)
	}
	if v := i.FunctionName; v != nil {
		m.SetFunctionName(*v)
	}
	if v := i.TriggerName; v != nil {
		m.SetTriggerName(*v)
	}
	if v := i.ListenerName; v != nil {
		m.SetListenerName(*v)
	}
	if v := i.EmailRecipients; v != nil {
		m.SetEmailRecipients(*v)
	}
	if v := i.EffectiveDate; v != nil {
		m.SetEffectiveDate(*v)
	}
	if v := i.ExpirationDate; v != nil {
		m.SetExpirationDate(*v)
	}
	m.SetBusinessUnitID(i.BusinessUnitID)
	m.SetOrganizationID(i.OrganizationID)
}

// SetInput applies the change-set in the CreateTableChangeAlertInput on the TableChangeAlertCreate builder.
func (c *TableChangeAlertCreate) SetInput(i CreateTableChangeAlertInput) *TableChangeAlertCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateTagInput represents a mutation input for creating tags.
type CreateTagInput struct {
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
	Name           string
	Description    *string
	BusinessUnitID uuid.UUID
	OrganizationID uuid.UUID
}

// Mutate applies the CreateTagInput on the TagMutation builder.
func (i *CreateTagInput) Mutate(m *TagMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetName(i.Name)
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	m.SetBusinessUnitID(i.BusinessUnitID)
	m.SetOrganizationID(i.OrganizationID)
}

// SetInput applies the change-set in the CreateTagInput on the TagCreate builder.
func (c *TagCreate) SetInput(i CreateTagInput) *TagCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
	Status         *user.Status
	Name           string
	Username       string
	Password       string
	Email          string
	DateJoined     *string
	Timezone       *user.Timezone
	ProfilePicURL  *string
	ThumbnailURL   *string
	PhoneNumber    *string
	IsAdmin        *bool
	IsSuperAdmin   *bool
	BusinessUnitID uuid.UUID
	OrganizationID uuid.UUID
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	m.SetName(i.Name)
	m.SetUsername(i.Username)
	m.SetPassword(i.Password)
	m.SetEmail(i.Email)
	if v := i.DateJoined; v != nil {
		m.SetDateJoined(*v)
	}
	if v := i.Timezone; v != nil {
		m.SetTimezone(*v)
	}
	if v := i.ProfilePicURL; v != nil {
		m.SetProfilePicURL(*v)
	}
	if v := i.ThumbnailURL; v != nil {
		m.SetThumbnailURL(*v)
	}
	if v := i.PhoneNumber; v != nil {
		m.SetPhoneNumber(*v)
	}
	if v := i.IsAdmin; v != nil {
		m.SetIsAdmin(*v)
	}
	if v := i.IsSuperAdmin; v != nil {
		m.SetIsSuperAdmin(*v)
	}
	m.SetBusinessUnitID(i.BusinessUnitID)
	m.SetOrganizationID(i.OrganizationID)
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}
