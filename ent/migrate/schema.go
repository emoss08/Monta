// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountingControlsColumns holds the columns for the "accounting_controls" table.
	AccountingControlsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "rec_threshold", Type: field.TypeInt64, Default: 50},
		{Name: "rec_threshold_action", Type: field.TypeEnum, Enums: []string{"Halt", "Warn"}, Default: "Halt"},
		{Name: "auto_create_journal_entries", Type: field.TypeBool, Default: false},
		{Name: "restrict_manual_journal_entries", Type: field.TypeBool, Default: false},
		{Name: "require_journal_entry_approval", Type: field.TypeBool, Default: false},
		{Name: "enable_rec_notifications", Type: field.TypeBool, Default: true},
		{Name: "halt_on_pending_rec", Type: field.TypeBool, Default: false},
		{Name: "critical_processes", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "business_unit_id", Type: field.TypeUUID},
		{Name: "default_rev_account_id", Type: field.TypeUUID, Nullable: true},
		{Name: "default_exp_account_id", Type: field.TypeUUID, Nullable: true},
		{Name: "organization_id", Type: field.TypeUUID, Unique: true},
	}
	// AccountingControlsTable holds the schema information for the "accounting_controls" table.
	AccountingControlsTable = &schema.Table{
		Name:       "accounting_controls",
		Columns:    AccountingControlsColumns,
		PrimaryKey: []*schema.Column{AccountingControlsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "accounting_controls_business_units_business_unit",
				Columns:    []*schema.Column{AccountingControlsColumns[11]},
				RefColumns: []*schema.Column{BusinessUnitsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "accounting_controls_general_ledger_accounts_default_rev_account",
				Columns:    []*schema.Column{AccountingControlsColumns[12]},
				RefColumns: []*schema.Column{GeneralLedgerAccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "accounting_controls_general_ledger_accounts_default_exp_account",
				Columns:    []*schema.Column{AccountingControlsColumns[13]},
				RefColumns: []*schema.Column{GeneralLedgerAccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "accounting_controls_organizations_accounting_control",
				Columns:    []*schema.Column{AccountingControlsColumns[14]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// BillingControlsColumns holds the columns for the "billing_controls" table.
	BillingControlsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "remove_billing_history", Type: field.TypeBool, Default: false},
		{Name: "auto_bill_shipment", Type: field.TypeBool, Default: false},
		{Name: "auto_mark_ready_to_bill", Type: field.TypeBool, Default: false},
		{Name: "validate_customer_rates", Type: field.TypeBool, Default: false},
		{Name: "auto_bill_criteria", Type: field.TypeEnum, Enums: []string{"Delivered", "TransferredToBilling", "MarkedReadyToBill"}, Default: "MarkedReadyToBill"},
		{Name: "shipment_transfer_criteria", Type: field.TypeEnum, Enums: []string{"ReadyAndCompleted", "Completed", "ReadyToBill"}, Default: "ReadyToBill"},
		{Name: "business_unit_id", Type: field.TypeUUID},
		{Name: "organization_id", Type: field.TypeUUID, Unique: true},
	}
	// BillingControlsTable holds the schema information for the "billing_controls" table.
	BillingControlsTable = &schema.Table{
		Name:       "billing_controls",
		Columns:    BillingControlsColumns,
		PrimaryKey: []*schema.Column{BillingControlsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "billing_controls_business_units_business_unit",
				Columns:    []*schema.Column{BillingControlsColumns[9]},
				RefColumns: []*schema.Column{BusinessUnitsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "billing_controls_organizations_billing_control",
				Columns:    []*schema.Column{BillingControlsColumns[10]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// BusinessUnitsColumns holds the columns for the "business_units" table.
	BusinessUnitsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"A", "I"}, Default: "A"},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "entity_key", Type: field.TypeString, Size: 10},
		{Name: "phone_number", Type: field.TypeString, Size: 15},
		{Name: "address", Type: field.TypeString, Nullable: true},
		{Name: "city", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "state", Type: field.TypeString, Nullable: true, Size: 2},
		{Name: "country", Type: field.TypeString, Nullable: true, Size: 2},
		{Name: "postal_code", Type: field.TypeString, Nullable: true, Size: 10},
		{Name: "tax_id", Type: field.TypeString, Nullable: true, Size: 20},
		{Name: "subscription_plan", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "legal_name", Type: field.TypeString, Nullable: true},
		{Name: "contact_name", Type: field.TypeString, Nullable: true},
		{Name: "contact_email", Type: field.TypeString, Nullable: true},
		{Name: "paid_until", Type: field.TypeTime, Nullable: true},
		{Name: "settings", Type: field.TypeJSON, Nullable: true},
		{Name: "free_trial", Type: field.TypeBool, Default: false},
		{Name: "parent_id", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// BusinessUnitsTable holds the schema information for the "business_units" table.
	BusinessUnitsTable = &schema.Table{
		Name:       "business_units",
		Columns:    BusinessUnitsColumns,
		PrimaryKey: []*schema.Column{BusinessUnitsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "business_units_business_units_next",
				Columns:    []*schema.Column{BusinessUnitsColumns[21]},
				RefColumns: []*schema.Column{BusinessUnitsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "businessunit_name",
				Unique:  true,
				Columns: []*schema.Column{BusinessUnitsColumns[4]},
			},
			{
				Name:    "businessunit_entity_key",
				Unique:  true,
				Columns: []*schema.Column{BusinessUnitsColumns[5]},
			},
		},
	}
	// CommoditiesColumns holds the columns for the "commodities" table.
	CommoditiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"A", "I"}, Default: "A"},
		{Name: "name", Type: field.TypeString, Size: 100},
		{Name: "is_hazmat", Type: field.TypeBool, Default: false},
		{Name: "unit_of_measure", Type: field.TypeEnum, Nullable: true, Enums: []string{"P", "T", "D", "C", "A", "B", "O", "L", "I", "S"}},
		{Name: "min_temp", Type: field.TypeFloat64, Nullable: true},
		{Name: "max_temp", Type: field.TypeFloat64, Nullable: true},
		{Name: "set_point_temp", Type: field.TypeFloat64, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "business_unit_id", Type: field.TypeUUID},
		{Name: "organization_id", Type: field.TypeUUID},
		{Name: "hazardous_material_id", Type: field.TypeUUID},
	}
	// CommoditiesTable holds the schema information for the "commodities" table.
	CommoditiesTable = &schema.Table{
		Name:       "commodities",
		Columns:    CommoditiesColumns,
		PrimaryKey: []*schema.Column{CommoditiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "commodities_business_units_business_unit",
				Columns:    []*schema.Column{CommoditiesColumns[11]},
				RefColumns: []*schema.Column{BusinessUnitsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "commodities_organizations_organization",
				Columns:    []*schema.Column{CommoditiesColumns[12]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "commodities_hazardous_materials_commodities",
				Columns:    []*schema.Column{CommoditiesColumns[13]},
				RefColumns: []*schema.Column{HazardousMaterialsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// DispatchControlsColumns holds the columns for the "dispatch_controls" table.
	DispatchControlsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "record_service_incident", Type: field.TypeEnum, Enums: []string{"Never", "Pickup", "Delivery", "PickupAndDelivery", "AllExceptShipper"}, Default: "Never"},
		{Name: "deadhead_target", Type: field.TypeFloat64, Default: 0},
		{Name: "max_shipment_weight_limit", Type: field.TypeInt, Default: 80000},
		{Name: "grace_period", Type: field.TypeUint8, Default: 0},
		{Name: "enforce_worker_assign", Type: field.TypeBool, Default: true},
		{Name: "trailer_continuity", Type: field.TypeBool, Default: false},
		{Name: "dupe_trailer_check", Type: field.TypeBool, Default: false},
		{Name: "maintenance_compliance", Type: field.TypeBool, Default: true},
		{Name: "regulatory_check", Type: field.TypeBool, Default: false},
		{Name: "prev_shipment_on_hold", Type: field.TypeBool, Default: false},
		{Name: "worker_time_away_restriction", Type: field.TypeBool, Default: true},
		{Name: "tractor_worker_fleet_constraint", Type: field.TypeBool, Default: false},
		{Name: "business_unit_id", Type: field.TypeUUID},
		{Name: "organization_id", Type: field.TypeUUID, Unique: true},
	}
	// DispatchControlsTable holds the schema information for the "dispatch_controls" table.
	DispatchControlsTable = &schema.Table{
		Name:       "dispatch_controls",
		Columns:    DispatchControlsColumns,
		PrimaryKey: []*schema.Column{DispatchControlsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "dispatch_controls_business_units_business_unit",
				Columns:    []*schema.Column{DispatchControlsColumns[15]},
				RefColumns: []*schema.Column{BusinessUnitsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "dispatch_controls_organizations_dispatch_control",
				Columns:    []*schema.Column{DispatchControlsColumns[16]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// FeasibilityToolControlsColumns holds the columns for the "feasibility_tool_controls" table.
	FeasibilityToolControlsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "otp_operator", Type: field.TypeEnum, Enums: []string{"Eq", "Ne", "Gt", "Gte", "Lt", "Lte"}, Default: "Eq"},
		{Name: "otp_value", Type: field.TypeFloat64, Default: 100},
		{Name: "mpw_operator", Type: field.TypeEnum, Enums: []string{"Eq", "Ne", "Gt", "Gte", "Lt", "Lte"}, Default: "Eq"},
		{Name: "mpw_value", Type: field.TypeFloat64, Default: 100},
		{Name: "mpd_operator", Type: field.TypeEnum, Enums: []string{"Eq", "Ne", "Gt", "Gte", "Lt", "Lte"}, Default: "Eq"},
		{Name: "mpd_value", Type: field.TypeFloat64, Default: 100},
		{Name: "mpg_operator", Type: field.TypeEnum, Enums: []string{"Eq", "Ne", "Gt", "Gte", "Lt", "Lte"}, Default: "Eq"},
		{Name: "mpg_value", Type: field.TypeFloat64, Default: 100},
		{Name: "business_unit_id", Type: field.TypeUUID},
		{Name: "organization_id", Type: field.TypeUUID, Unique: true},
	}
	// FeasibilityToolControlsTable holds the schema information for the "feasibility_tool_controls" table.
	FeasibilityToolControlsTable = &schema.Table{
		Name:       "feasibility_tool_controls",
		Columns:    FeasibilityToolControlsColumns,
		PrimaryKey: []*schema.Column{FeasibilityToolControlsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "feasibility_tool_controls_business_units_business_unit",
				Columns:    []*schema.Column{FeasibilityToolControlsColumns[11]},
				RefColumns: []*schema.Column{BusinessUnitsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "feasibility_tool_controls_organizations_feasibility_tool_control",
				Columns:    []*schema.Column{FeasibilityToolControlsColumns[12]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// GeneralLedgerAccountsColumns holds the columns for the "general_ledger_accounts" table.
	GeneralLedgerAccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"A", "I"}, Default: "A"},
		{Name: "account_number", Type: field.TypeString, Size: 7},
		{Name: "account_type", Type: field.TypeEnum, Enums: []string{"AccountTypeAsset", "AccountTypeLiability", "AccountTypeEquity", "AccountTypeRevenue", "AccountTypeExpense"}},
		{Name: "cash_flow_type", Type: field.TypeEnum, Nullable: true, Enums: []string{"CashFlowOperating", "CashFlowInvesting", "CashFlowFinancing"}},
		{Name: "account_sub_type", Type: field.TypeEnum, Nullable: true, Enums: []string{"AccountSubTypeCurrentAsset", "AccountSubTypeFixedAsset", "AccountSubTypeOtherAsset", "AccountSubTypeCurrentLiability", "AccountSubTypeLongTermLiability", "AccountSubTypeEquity", "AccountSubTypeRevenue", "AccountSubTypeCostOfGoodsSold", "AccountSubTypeExpense", "AccountSubTypeOtherIncome", "AccountSubTypeOtherExpense"}},
		{Name: "account_class", Type: field.TypeEnum, Nullable: true, Enums: []string{"AccountClassificationBank", "AccountClassificationCash", "AccountClassificationAR", "AccountClassificationAP", "AccountClassificationINV", "AccountClassificationOCA", "AccountClassificationFA"}},
		{Name: "balance", Type: field.TypeFloat64, Nullable: true},
		{Name: "interest_rate", Type: field.TypeFloat64, Nullable: true},
		{Name: "date_opened", Type: field.TypeTime},
		{Name: "date_closed", Type: field.TypeTime, Nullable: true},
		{Name: "notes", Type: field.TypeString, Nullable: true},
		{Name: "is_tax_relevant", Type: field.TypeBool, Default: false},
		{Name: "is_reconciled", Type: field.TypeBool, Default: false},
		{Name: "business_unit_id", Type: field.TypeUUID},
		{Name: "organization_id", Type: field.TypeUUID},
	}
	// GeneralLedgerAccountsTable holds the schema information for the "general_ledger_accounts" table.
	GeneralLedgerAccountsTable = &schema.Table{
		Name:       "general_ledger_accounts",
		Columns:    GeneralLedgerAccountsColumns,
		PrimaryKey: []*schema.Column{GeneralLedgerAccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "general_ledger_accounts_business_units_business_unit",
				Columns:    []*schema.Column{GeneralLedgerAccountsColumns[16]},
				RefColumns: []*schema.Column{BusinessUnitsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "general_ledger_accounts_organizations_organization",
				Columns:    []*schema.Column{GeneralLedgerAccountsColumns[17]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "generalledgeraccount_account_number_organization_id",
				Unique:  true,
				Columns: []*schema.Column{GeneralLedgerAccountsColumns[4], GeneralLedgerAccountsColumns[17]},
			},
		},
	}
	// HazardousMaterialsColumns holds the columns for the "hazardous_materials" table.
	HazardousMaterialsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 100},
		{Name: "hazard_class", Type: field.TypeEnum, Enums: []string{"HazardClass1And1", "HazardClass1And2", "HazardClass1And3", "HazardClass1And4", "HazardClass1And5", "HazardClass1And6", "HazardClass2And1", "HazardClass2And2", "HazardClass2And3", "HazardClass3", "HazardClass4And1", "HazardClass4And2", "HazardClass4And3", "HazardClass5And1", "HazardClass5And2", "HazardClass6And1", "HazardClass6And2", "HazardClass7", "HazardClass8", "HazardClass9"}, Default: "HazardClass1And1"},
		{Name: "erg_number", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "packing_group", Type: field.TypeEnum, Nullable: true, Enums: []string{"PackingGroupI", "PackingGroupII", "PackingGroupIII"}},
		{Name: "proper_shipping_name", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "business_unit_id", Type: field.TypeUUID},
		{Name: "organization_id", Type: field.TypeUUID},
	}
	// HazardousMaterialsTable holds the schema information for the "hazardous_materials" table.
	HazardousMaterialsTable = &schema.Table{
		Name:       "hazardous_materials",
		Columns:    HazardousMaterialsColumns,
		PrimaryKey: []*schema.Column{HazardousMaterialsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "hazardous_materials_business_units_business_unit",
				Columns:    []*schema.Column{HazardousMaterialsColumns[9]},
				RefColumns: []*schema.Column{BusinessUnitsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "hazardous_materials_organizations_organization",
				Columns:    []*schema.Column{HazardousMaterialsColumns[10]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// InvoiceControlsColumns holds the columns for the "invoice_controls" table.
	InvoiceControlsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "invoice_number_prefix", Type: field.TypeString, Size: 10, Default: "INV-"},
		{Name: "credit_memo_number_prefix", Type: field.TypeString, Size: 10, Default: "CM-"},
		{Name: "invoice_terms", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "invoice_footer", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "invoice_logo_url", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "invoice_date_format", Type: field.TypeEnum, Enums: []string{"InvoiceDateFormatMDY", "InvoiceDateFormatDMY", "InvoiceDateFormatYMD", "InvoiceDateFormatYDM"}, Default: "InvoiceDateFormatMDY"},
		{Name: "invoice_due_after_days", Type: field.TypeUint8, Default: 30},
		{Name: "invoice_logo_width", Type: field.TypeUint16, Default: 100},
		{Name: "show_amount_due", Type: field.TypeBool, Default: true},
		{Name: "attach_pdf", Type: field.TypeBool, Default: true},
		{Name: "show_invoice_due_date", Type: field.TypeBool, Default: true},
		{Name: "business_unit_id", Type: field.TypeUUID},
		{Name: "organization_id", Type: field.TypeUUID, Unique: true},
	}
	// InvoiceControlsTable holds the schema information for the "invoice_controls" table.
	InvoiceControlsTable = &schema.Table{
		Name:       "invoice_controls",
		Columns:    InvoiceControlsColumns,
		PrimaryKey: []*schema.Column{InvoiceControlsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "invoice_controls_business_units_business_unit",
				Columns:    []*schema.Column{InvoiceControlsColumns[14]},
				RefColumns: []*schema.Column{BusinessUnitsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "invoice_controls_organizations_invoice_control",
				Columns:    []*schema.Column{InvoiceControlsColumns[15]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// OrganizationsColumns holds the columns for the "organizations" table.
	OrganizationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 100},
		{Name: "scac_code", Type: field.TypeString, Size: 4},
		{Name: "dot_number", Type: field.TypeString, Size: 12},
		{Name: "logo_url", Type: field.TypeString, Nullable: true},
		{Name: "org_type", Type: field.TypeEnum, Enums: []string{"A", "B", "X"}, Default: "A"},
		{Name: "timezone", Type: field.TypeEnum, Enums: []string{"AmericaLosAngeles", "AmericaDenver", "AmericaChicago", "AmericaNewYork"}, Default: "AmericaLosAngeles"},
		{Name: "business_unit_id", Type: field.TypeUUID},
	}
	// OrganizationsTable holds the schema information for the "organizations" table.
	OrganizationsTable = &schema.Table{
		Name:       "organizations",
		Columns:    OrganizationsColumns,
		PrimaryKey: []*schema.Column{OrganizationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "organizations_business_units_organizations",
				Columns:    []*schema.Column{OrganizationsColumns[9]},
				RefColumns: []*schema.Column{BusinessUnitsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// RouteControlsColumns holds the columns for the "route_controls" table.
	RouteControlsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "distance_method", Type: field.TypeEnum, Enums: []string{"T", "G"}, Default: "T"},
		{Name: "mileage_unit", Type: field.TypeEnum, Enums: []string{"M", "I"}, Default: "M"},
		{Name: "generate_routes", Type: field.TypeBool, Default: false},
		{Name: "organization_id", Type: field.TypeUUID, Unique: true},
		{Name: "business_unit_id", Type: field.TypeUUID},
	}
	// RouteControlsTable holds the schema information for the "route_controls" table.
	RouteControlsTable = &schema.Table{
		Name:       "route_controls",
		Columns:    RouteControlsColumns,
		PrimaryKey: []*schema.Column{RouteControlsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "route_controls_organizations_route_control",
				Columns:    []*schema.Column{RouteControlsColumns[6]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "route_controls_business_units_business_unit",
				Columns:    []*schema.Column{RouteControlsColumns[7]},
				RefColumns: []*schema.Column{BusinessUnitsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "data", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "expires_at", Type: field.TypeTime},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:       "sessions",
		Columns:    SessionsColumns,
		PrimaryKey: []*schema.Column{SessionsColumns[0]},
	}
	// ShipmentControlsColumns holds the columns for the "shipment_controls" table.
	ShipmentControlsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "auto_rate_shipment", Type: field.TypeBool, Default: true},
		{Name: "calculate_distance", Type: field.TypeBool, Default: true},
		{Name: "enforce_rev_code", Type: field.TypeBool, Default: false},
		{Name: "enforce_voided_comm", Type: field.TypeBool, Default: false},
		{Name: "generate_routes", Type: field.TypeBool, Default: false},
		{Name: "enforce_commodity", Type: field.TypeBool, Default: false},
		{Name: "auto_sequence_stops", Type: field.TypeBool, Default: true},
		{Name: "auto_shipment_total", Type: field.TypeBool, Default: true},
		{Name: "enforce_origin_destination", Type: field.TypeBool, Default: false},
		{Name: "check_for_duplicate_bol", Type: field.TypeBool, Default: false},
		{Name: "send_placard_info", Type: field.TypeBool, Default: false},
		{Name: "enforce_hazmat_seg_rules", Type: field.TypeBool, Default: true},
		{Name: "organization_id", Type: field.TypeUUID, Unique: true},
		{Name: "business_unit_id", Type: field.TypeUUID},
	}
	// ShipmentControlsTable holds the schema information for the "shipment_controls" table.
	ShipmentControlsTable = &schema.Table{
		Name:       "shipment_controls",
		Columns:    ShipmentControlsColumns,
		PrimaryKey: []*schema.Column{ShipmentControlsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "shipment_controls_organizations_shipment_control",
				Columns:    []*schema.Column{ShipmentControlsColumns[15]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "shipment_controls_business_units_business_unit",
				Columns:    []*schema.Column{ShipmentControlsColumns[16]},
				RefColumns: []*schema.Column{BusinessUnitsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TableChangeAlertsColumns holds the columns for the "table_change_alerts" table.
	TableChangeAlertsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"A", "I"}, Default: "A"},
		{Name: "name", Type: field.TypeString, Size: 50},
		{Name: "database_action", Type: field.TypeEnum, Enums: []string{"Insert", "Update", "Delete", "All"}},
		{Name: "source", Type: field.TypeEnum, Enums: []string{"Kafka", "Db"}},
		{Name: "table_name", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "topic", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "custom_subject", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "function_name", Type: field.TypeString, Nullable: true, Size: 50},
		{Name: "trigger_name", Type: field.TypeString, Nullable: true, Size: 50},
		{Name: "listener_name", Type: field.TypeString, Nullable: true, Size: 50},
		{Name: "email_recipients", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "effective_date", Type: field.TypeTime, Nullable: true},
		{Name: "expiration_date", Type: field.TypeTime, Nullable: true},
		{Name: "business_unit_id", Type: field.TypeUUID},
		{Name: "organization_id", Type: field.TypeUUID},
	}
	// TableChangeAlertsTable holds the schema information for the "table_change_alerts" table.
	TableChangeAlertsTable = &schema.Table{
		Name:       "table_change_alerts",
		Columns:    TableChangeAlertsColumns,
		PrimaryKey: []*schema.Column{TableChangeAlertsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "table_change_alerts_business_units_business_unit",
				Columns:    []*schema.Column{TableChangeAlertsColumns[17]},
				RefColumns: []*schema.Column{BusinessUnitsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "table_change_alerts_organizations_organization",
				Columns:    []*schema.Column{TableChangeAlertsColumns[18]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 50},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "general_ledger_account_tags", Type: field.TypeUUID, Nullable: true},
		{Name: "business_unit_id", Type: field.TypeUUID},
		{Name: "organization_id", Type: field.TypeUUID},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tags_general_ledger_accounts_tags",
				Columns:    []*schema.Column{TagsColumns[5]},
				RefColumns: []*schema.Column{GeneralLedgerAccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "tags_business_units_business_unit",
				Columns:    []*schema.Column{TagsColumns[6]},
				RefColumns: []*schema.Column{BusinessUnitsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "tags_organizations_organization",
				Columns:    []*schema.Column{TagsColumns[7]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "tag_name_organization_id",
				Unique:  true,
				Columns: []*schema.Column{TagsColumns[3], TagsColumns[7]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"A", "I"}, Default: "A"},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "username", Type: field.TypeString, Size: 30},
		{Name: "password", Type: field.TypeString, Size: 100},
		{Name: "email", Type: field.TypeString, Size: 255},
		{Name: "timezone", Type: field.TypeEnum, Enums: []string{"AmericaLosAngeles", "AmericaDenver", "AmericaChicago", "AmericaNewYork"}, Default: "AmericaLosAngeles"},
		{Name: "profile_pic_url", Type: field.TypeString, Nullable: true},
		{Name: "thumbnail_url", Type: field.TypeString, Nullable: true},
		{Name: "phone_number", Type: field.TypeString, Nullable: true},
		{Name: "is_admin", Type: field.TypeBool, Default: false},
		{Name: "is_super_admin", Type: field.TypeBool, Default: false},
		{Name: "last_login", Type: field.TypeTime, Nullable: true},
		{Name: "business_unit_id", Type: field.TypeUUID},
		{Name: "organization_id", Type: field.TypeUUID},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_business_units_business_unit",
				Columns:    []*schema.Column{UsersColumns[15]},
				RefColumns: []*schema.Column{BusinessUnitsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "users_organizations_organization",
				Columns:    []*schema.Column{UsersColumns[16]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "user_username_email",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[5], UsersColumns[7]},
			},
		},
	}
	// UserFavoritesColumns holds the columns for the "user_favorites" table.
	UserFavoritesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "page_link", Type: field.TypeString, Unique: true, Size: 255},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "business_unit_id", Type: field.TypeUUID},
		{Name: "organization_id", Type: field.TypeUUID},
	}
	// UserFavoritesTable holds the schema information for the "user_favorites" table.
	UserFavoritesTable = &schema.Table{
		Name:       "user_favorites",
		Columns:    UserFavoritesColumns,
		PrimaryKey: []*schema.Column{UserFavoritesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_favorites_users_user_favorites",
				Columns:    []*schema.Column{UserFavoritesColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "user_favorites_business_units_business_unit",
				Columns:    []*schema.Column{UserFavoritesColumns[5]},
				RefColumns: []*schema.Column{BusinessUnitsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_favorites_organizations_organization",
				Columns:    []*schema.Column{UserFavoritesColumns[6]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountingControlsTable,
		BillingControlsTable,
		BusinessUnitsTable,
		CommoditiesTable,
		DispatchControlsTable,
		FeasibilityToolControlsTable,
		GeneralLedgerAccountsTable,
		HazardousMaterialsTable,
		InvoiceControlsTable,
		OrganizationsTable,
		RouteControlsTable,
		SessionsTable,
		ShipmentControlsTable,
		TableChangeAlertsTable,
		TagsTable,
		UsersTable,
		UserFavoritesTable,
	}
)

func init() {
	AccountingControlsTable.ForeignKeys[0].RefTable = BusinessUnitsTable
	AccountingControlsTable.ForeignKeys[1].RefTable = GeneralLedgerAccountsTable
	AccountingControlsTable.ForeignKeys[2].RefTable = GeneralLedgerAccountsTable
	AccountingControlsTable.ForeignKeys[3].RefTable = OrganizationsTable
	BillingControlsTable.ForeignKeys[0].RefTable = BusinessUnitsTable
	BillingControlsTable.ForeignKeys[1].RefTable = OrganizationsTable
	BusinessUnitsTable.ForeignKeys[0].RefTable = BusinessUnitsTable
	CommoditiesTable.ForeignKeys[0].RefTable = BusinessUnitsTable
	CommoditiesTable.ForeignKeys[1].RefTable = OrganizationsTable
	CommoditiesTable.ForeignKeys[2].RefTable = HazardousMaterialsTable
	DispatchControlsTable.ForeignKeys[0].RefTable = BusinessUnitsTable
	DispatchControlsTable.ForeignKeys[1].RefTable = OrganizationsTable
	FeasibilityToolControlsTable.ForeignKeys[0].RefTable = BusinessUnitsTable
	FeasibilityToolControlsTable.ForeignKeys[1].RefTable = OrganizationsTable
	GeneralLedgerAccountsTable.ForeignKeys[0].RefTable = BusinessUnitsTable
	GeneralLedgerAccountsTable.ForeignKeys[1].RefTable = OrganizationsTable
	HazardousMaterialsTable.ForeignKeys[0].RefTable = BusinessUnitsTable
	HazardousMaterialsTable.ForeignKeys[1].RefTable = OrganizationsTable
	InvoiceControlsTable.ForeignKeys[0].RefTable = BusinessUnitsTable
	InvoiceControlsTable.ForeignKeys[1].RefTable = OrganizationsTable
	OrganizationsTable.ForeignKeys[0].RefTable = BusinessUnitsTable
	RouteControlsTable.ForeignKeys[0].RefTable = OrganizationsTable
	RouteControlsTable.ForeignKeys[1].RefTable = BusinessUnitsTable
	ShipmentControlsTable.ForeignKeys[0].RefTable = OrganizationsTable
	ShipmentControlsTable.ForeignKeys[1].RefTable = BusinessUnitsTable
	TableChangeAlertsTable.ForeignKeys[0].RefTable = BusinessUnitsTable
	TableChangeAlertsTable.ForeignKeys[1].RefTable = OrganizationsTable
	TagsTable.ForeignKeys[0].RefTable = GeneralLedgerAccountsTable
	TagsTable.ForeignKeys[1].RefTable = BusinessUnitsTable
	TagsTable.ForeignKeys[2].RefTable = OrganizationsTable
	UsersTable.ForeignKeys[0].RefTable = BusinessUnitsTable
	UsersTable.ForeignKeys[1].RefTable = OrganizationsTable
	UserFavoritesTable.ForeignKeys[0].RefTable = UsersTable
	UserFavoritesTable.ForeignKeys[1].RefTable = BusinessUnitsTable
	UserFavoritesTable.ForeignKeys[2].RefTable = OrganizationsTable
}
