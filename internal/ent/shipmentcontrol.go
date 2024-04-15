// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/emoss08/trenova/internal/ent/businessunit"
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/emoss08/trenova/internal/ent/shipmentcontrol"
	"github.com/google/uuid"
)

// ShipmentControl is the model entity for the ShipmentControl schema.
type ShipmentControl struct {
	config `json:"-" validate:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt" validate:"omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt" validate:"omitempty"`
	// AutoRateShipment holds the value of the "auto_rate_shipment" field.
	AutoRateShipment bool `json:"autoRateShipment"`
	// CalculateDistance holds the value of the "calculate_distance" field.
	CalculateDistance bool `json:"calculateDistance"`
	// EnforceRevCode holds the value of the "enforce_rev_code" field.
	EnforceRevCode bool `json:"enforceRevCode"`
	// EnforceVoidedComm holds the value of the "enforce_voided_comm" field.
	EnforceVoidedComm bool `json:"enforceVoidedComm"`
	// GenerateRoutes holds the value of the "generate_routes" field.
	GenerateRoutes bool `json:"generateRoutes"`
	// EnforceCommodity holds the value of the "enforce_commodity" field.
	EnforceCommodity bool `json:"enforceCommodity"`
	// AutoSequenceStops holds the value of the "auto_sequence_stops" field.
	AutoSequenceStops bool `json:"autoSequenceStops"`
	// AutoShipmentTotal holds the value of the "auto_shipment_total" field.
	AutoShipmentTotal bool `json:"autoShipmentTotal"`
	// EnforceOriginDestination holds the value of the "enforce_origin_destination" field.
	EnforceOriginDestination bool `json:"enforceOriginDestination"`
	// CheckForDuplicateBol holds the value of the "check_for_duplicate_bol" field.
	CheckForDuplicateBol bool `json:"checkForDuplicateBol"`
	// SendPlacardInfo holds the value of the "send_placard_info" field.
	SendPlacardInfo bool `json:"sendPlacardInfo"`
	// EnforceHazmatSegRules holds the value of the "enforce_hazmat_seg_rules" field.
	EnforceHazmatSegRules bool `json:"enforceHazmatSegRules"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ShipmentControlQuery when eager-loading is set.
	Edges            ShipmentControlEdges `json:"edges"`
	organization_id  *uuid.UUID
	business_unit_id *uuid.UUID
	selectValues     sql.SelectValues
}

// ShipmentControlEdges holds the relations/edges for other nodes in the graph.
type ShipmentControlEdges struct {
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// BusinessUnit holds the value of the business_unit edge.
	BusinessUnit *BusinessUnit `json:"business_unit,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ShipmentControlEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// BusinessUnitOrErr returns the BusinessUnit value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ShipmentControlEdges) BusinessUnitOrErr() (*BusinessUnit, error) {
	if e.BusinessUnit != nil {
		return e.BusinessUnit, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: businessunit.Label}
	}
	return nil, &NotLoadedError{edge: "business_unit"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ShipmentControl) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case shipmentcontrol.FieldAutoRateShipment, shipmentcontrol.FieldCalculateDistance, shipmentcontrol.FieldEnforceRevCode, shipmentcontrol.FieldEnforceVoidedComm, shipmentcontrol.FieldGenerateRoutes, shipmentcontrol.FieldEnforceCommodity, shipmentcontrol.FieldAutoSequenceStops, shipmentcontrol.FieldAutoShipmentTotal, shipmentcontrol.FieldEnforceOriginDestination, shipmentcontrol.FieldCheckForDuplicateBol, shipmentcontrol.FieldSendPlacardInfo, shipmentcontrol.FieldEnforceHazmatSegRules:
			values[i] = new(sql.NullBool)
		case shipmentcontrol.FieldCreatedAt, shipmentcontrol.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case shipmentcontrol.FieldID:
			values[i] = new(uuid.UUID)
		case shipmentcontrol.ForeignKeys[0]: // organization_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case shipmentcontrol.ForeignKeys[1]: // business_unit_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ShipmentControl fields.
func (sc *ShipmentControl) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case shipmentcontrol.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				sc.ID = *value
			}
		case shipmentcontrol.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sc.CreatedAt = value.Time
			}
		case shipmentcontrol.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sc.UpdatedAt = value.Time
			}
		case shipmentcontrol.FieldAutoRateShipment:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field auto_rate_shipment", values[i])
			} else if value.Valid {
				sc.AutoRateShipment = value.Bool
			}
		case shipmentcontrol.FieldCalculateDistance:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field calculate_distance", values[i])
			} else if value.Valid {
				sc.CalculateDistance = value.Bool
			}
		case shipmentcontrol.FieldEnforceRevCode:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enforce_rev_code", values[i])
			} else if value.Valid {
				sc.EnforceRevCode = value.Bool
			}
		case shipmentcontrol.FieldEnforceVoidedComm:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enforce_voided_comm", values[i])
			} else if value.Valid {
				sc.EnforceVoidedComm = value.Bool
			}
		case shipmentcontrol.FieldGenerateRoutes:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field generate_routes", values[i])
			} else if value.Valid {
				sc.GenerateRoutes = value.Bool
			}
		case shipmentcontrol.FieldEnforceCommodity:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enforce_commodity", values[i])
			} else if value.Valid {
				sc.EnforceCommodity = value.Bool
			}
		case shipmentcontrol.FieldAutoSequenceStops:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field auto_sequence_stops", values[i])
			} else if value.Valid {
				sc.AutoSequenceStops = value.Bool
			}
		case shipmentcontrol.FieldAutoShipmentTotal:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field auto_shipment_total", values[i])
			} else if value.Valid {
				sc.AutoShipmentTotal = value.Bool
			}
		case shipmentcontrol.FieldEnforceOriginDestination:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enforce_origin_destination", values[i])
			} else if value.Valid {
				sc.EnforceOriginDestination = value.Bool
			}
		case shipmentcontrol.FieldCheckForDuplicateBol:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field check_for_duplicate_bol", values[i])
			} else if value.Valid {
				sc.CheckForDuplicateBol = value.Bool
			}
		case shipmentcontrol.FieldSendPlacardInfo:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field send_placard_info", values[i])
			} else if value.Valid {
				sc.SendPlacardInfo = value.Bool
			}
		case shipmentcontrol.FieldEnforceHazmatSegRules:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enforce_hazmat_seg_rules", values[i])
			} else if value.Valid {
				sc.EnforceHazmatSegRules = value.Bool
			}
		case shipmentcontrol.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value.Valid {
				sc.organization_id = new(uuid.UUID)
				*sc.organization_id = *value.S.(*uuid.UUID)
			}
		case shipmentcontrol.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field business_unit_id", values[i])
			} else if value.Valid {
				sc.business_unit_id = new(uuid.UUID)
				*sc.business_unit_id = *value.S.(*uuid.UUID)
			}
		default:
			sc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ShipmentControl.
// This includes values selected through modifiers, order, etc.
func (sc *ShipmentControl) Value(name string) (ent.Value, error) {
	return sc.selectValues.Get(name)
}

// QueryOrganization queries the "organization" edge of the ShipmentControl entity.
func (sc *ShipmentControl) QueryOrganization() *OrganizationQuery {
	return NewShipmentControlClient(sc.config).QueryOrganization(sc)
}

// QueryBusinessUnit queries the "business_unit" edge of the ShipmentControl entity.
func (sc *ShipmentControl) QueryBusinessUnit() *BusinessUnitQuery {
	return NewShipmentControlClient(sc.config).QueryBusinessUnit(sc)
}

// Update returns a builder for updating this ShipmentControl.
// Note that you need to call ShipmentControl.Unwrap() before calling this method if this ShipmentControl
// was returned from a transaction, and the transaction was committed or rolled back.
func (sc *ShipmentControl) Update() *ShipmentControlUpdateOne {
	return NewShipmentControlClient(sc.config).UpdateOne(sc)
}

// Unwrap unwraps the ShipmentControl entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sc *ShipmentControl) Unwrap() *ShipmentControl {
	_tx, ok := sc.config.driver.(*txDriver)
	if !ok {
		panic("ent: ShipmentControl is not a transactional entity")
	}
	sc.config.driver = _tx.drv
	return sc
}

// String implements the fmt.Stringer.
func (sc *ShipmentControl) String() string {
	var builder strings.Builder
	builder.WriteString("ShipmentControl(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(sc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("auto_rate_shipment=")
	builder.WriteString(fmt.Sprintf("%v", sc.AutoRateShipment))
	builder.WriteString(", ")
	builder.WriteString("calculate_distance=")
	builder.WriteString(fmt.Sprintf("%v", sc.CalculateDistance))
	builder.WriteString(", ")
	builder.WriteString("enforce_rev_code=")
	builder.WriteString(fmt.Sprintf("%v", sc.EnforceRevCode))
	builder.WriteString(", ")
	builder.WriteString("enforce_voided_comm=")
	builder.WriteString(fmt.Sprintf("%v", sc.EnforceVoidedComm))
	builder.WriteString(", ")
	builder.WriteString("generate_routes=")
	builder.WriteString(fmt.Sprintf("%v", sc.GenerateRoutes))
	builder.WriteString(", ")
	builder.WriteString("enforce_commodity=")
	builder.WriteString(fmt.Sprintf("%v", sc.EnforceCommodity))
	builder.WriteString(", ")
	builder.WriteString("auto_sequence_stops=")
	builder.WriteString(fmt.Sprintf("%v", sc.AutoSequenceStops))
	builder.WriteString(", ")
	builder.WriteString("auto_shipment_total=")
	builder.WriteString(fmt.Sprintf("%v", sc.AutoShipmentTotal))
	builder.WriteString(", ")
	builder.WriteString("enforce_origin_destination=")
	builder.WriteString(fmt.Sprintf("%v", sc.EnforceOriginDestination))
	builder.WriteString(", ")
	builder.WriteString("check_for_duplicate_bol=")
	builder.WriteString(fmt.Sprintf("%v", sc.CheckForDuplicateBol))
	builder.WriteString(", ")
	builder.WriteString("send_placard_info=")
	builder.WriteString(fmt.Sprintf("%v", sc.SendPlacardInfo))
	builder.WriteString(", ")
	builder.WriteString("enforce_hazmat_seg_rules=")
	builder.WriteString(fmt.Sprintf("%v", sc.EnforceHazmatSegRules))
	builder.WriteByte(')')
	return builder.String()
}

// ShipmentControls is a parsable slice of ShipmentControl.
type ShipmentControls []*ShipmentControl