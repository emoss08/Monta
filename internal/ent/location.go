// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/emoss08/trenova/internal/ent/businessunit"
	"github.com/emoss08/trenova/internal/ent/location"
	"github.com/emoss08/trenova/internal/ent/locationcategory"
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/emoss08/trenova/internal/ent/usstate"
	"github.com/google/uuid"
)

// Location is the model entity for the Location schema.
type Location struct {
	config `json:"-" validate:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// BusinessUnitID holds the value of the "business_unit_id" field.
	BusinessUnitID uuid.UUID `json:"businessUnitId"`
	// OrganizationID holds the value of the "organization_id" field.
	OrganizationID uuid.UUID `json:"organizationId"`
	// The time that this entity was created.
	CreatedAt time.Time `json:"createdAt" validate:"omitempty"`
	// The last time that this entity was updated.
	UpdatedAt time.Time `json:"updatedAt" validate:"omitempty"`
	// The current version of this entity.
	Version int `json:"version" validate:"omitempty"`
	// Current status of the location.
	Status location.Status `json:"status" validate:"required,oneof=A I"`
	// Unique code for the location.
	Code string `json:"code" validate:"required,max=10"`
	// Location category ID.
	LocationCategoryID *uuid.UUID `json:"locationCategoryId" validate:"omitempty"`
	// Name of the location.
	Name string `json:"name" validate:"required"`
	// Description of the location.
	Description string `json:"description" validate:"omitempty"`
	// Adress Line 1 of the location.
	AddressLine1 string `json:"addressLine1" validate:"required,max=150"`
	// Adress Line 2 of the location.
	AddressLine2 string `json:"addressLine2" validate:"omitempty,max=150"`
	// City of the location.
	City string `json:"city" validate:"required,max=150"`
	// State ID.
	StateID uuid.UUID `json:"stateId" validate:"omitempty,uuid"`
	// Postal code of the location.
	PostalCode string `json:"postalCode" validate:"required,max=10"`
	// Longitude of the location.
	Longitude float64 `json:"longitude" validate:"omitempty"`
	// Latitude of the location.
	Latitude float64 `json:"latitude" validate:"omitempty"`
	// Place ID from Google Maps API.
	PlaceID string `json:"placeId" validate:"omitempty,max=255"`
	// Is the location geocoded?
	IsGeocoded bool `json:"isGeocoded"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LocationQuery when eager-loading is set.
	Edges        LocationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// LocationEdges holds the relations/edges for other nodes in the graph.
type LocationEdges struct {
	// BusinessUnit holds the value of the business_unit edge.
	BusinessUnit *BusinessUnit `json:"business_unit,omitempty"`
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// LocationCategory holds the value of the location_category edge.
	LocationCategory *LocationCategory `json:"locationCategory"`
	// State holds the value of the state edge.
	State *UsState `json:"state"`
	// Comments holds the value of the comments edge.
	Comments []*LocationComment `json:"comments"`
	// Contacts holds the value of the contacts edge.
	Contacts []*LocationContact `json:"contacts"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes   [6]bool
	namedComments map[string][]*LocationComment
	namedContacts map[string][]*LocationContact
}

// BusinessUnitOrErr returns the BusinessUnit value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LocationEdges) BusinessUnitOrErr() (*BusinessUnit, error) {
	if e.BusinessUnit != nil {
		return e.BusinessUnit, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: businessunit.Label}
	}
	return nil, &NotLoadedError{edge: "business_unit"}
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LocationEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// LocationCategoryOrErr returns the LocationCategory value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LocationEdges) LocationCategoryOrErr() (*LocationCategory, error) {
	if e.LocationCategory != nil {
		return e.LocationCategory, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: locationcategory.Label}
	}
	return nil, &NotLoadedError{edge: "location_category"}
}

// StateOrErr returns the State value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LocationEdges) StateOrErr() (*UsState, error) {
	if e.State != nil {
		return e.State, nil
	} else if e.loadedTypes[3] {
		return nil, &NotFoundError{label: usstate.Label}
	}
	return nil, &NotLoadedError{edge: "state"}
}

// CommentsOrErr returns the Comments value or an error if the edge
// was not loaded in eager-loading.
func (e LocationEdges) CommentsOrErr() ([]*LocationComment, error) {
	if e.loadedTypes[4] {
		return e.Comments, nil
	}
	return nil, &NotLoadedError{edge: "comments"}
}

// ContactsOrErr returns the Contacts value or an error if the edge
// was not loaded in eager-loading.
func (e LocationEdges) ContactsOrErr() ([]*LocationContact, error) {
	if e.loadedTypes[5] {
		return e.Contacts, nil
	}
	return nil, &NotLoadedError{edge: "contacts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Location) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case location.FieldLocationCategoryID:
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case location.FieldIsGeocoded:
			values[i] = new(sql.NullBool)
		case location.FieldLongitude, location.FieldLatitude:
			values[i] = new(sql.NullFloat64)
		case location.FieldVersion:
			values[i] = new(sql.NullInt64)
		case location.FieldStatus, location.FieldCode, location.FieldName, location.FieldDescription, location.FieldAddressLine1, location.FieldAddressLine2, location.FieldCity, location.FieldPostalCode, location.FieldPlaceID:
			values[i] = new(sql.NullString)
		case location.FieldCreatedAt, location.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case location.FieldID, location.FieldBusinessUnitID, location.FieldOrganizationID, location.FieldStateID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Location fields.
func (l *Location) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case location.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				l.ID = *value
			}
		case location.FieldBusinessUnitID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field business_unit_id", values[i])
			} else if value != nil {
				l.BusinessUnitID = *value
			}
		case location.FieldOrganizationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value != nil {
				l.OrganizationID = *value
			}
		case location.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				l.CreatedAt = value.Time
			}
		case location.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				l.UpdatedAt = value.Time
			}
		case location.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				l.Version = int(value.Int64)
			}
		case location.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				l.Status = location.Status(value.String)
			}
		case location.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				l.Code = value.String
			}
		case location.FieldLocationCategoryID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field location_category_id", values[i])
			} else if value.Valid {
				l.LocationCategoryID = new(uuid.UUID)
				*l.LocationCategoryID = *value.S.(*uuid.UUID)
			}
		case location.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				l.Name = value.String
			}
		case location.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				l.Description = value.String
			}
		case location.FieldAddressLine1:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address_line_1", values[i])
			} else if value.Valid {
				l.AddressLine1 = value.String
			}
		case location.FieldAddressLine2:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address_line_2", values[i])
			} else if value.Valid {
				l.AddressLine2 = value.String
			}
		case location.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				l.City = value.String
			}
		case location.FieldStateID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field state_id", values[i])
			} else if value != nil {
				l.StateID = *value
			}
		case location.FieldPostalCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field postal_code", values[i])
			} else if value.Valid {
				l.PostalCode = value.String
			}
		case location.FieldLongitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field longitude", values[i])
			} else if value.Valid {
				l.Longitude = value.Float64
			}
		case location.FieldLatitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field latitude", values[i])
			} else if value.Valid {
				l.Latitude = value.Float64
			}
		case location.FieldPlaceID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field place_id", values[i])
			} else if value.Valid {
				l.PlaceID = value.String
			}
		case location.FieldIsGeocoded:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_geocoded", values[i])
			} else if value.Valid {
				l.IsGeocoded = value.Bool
			}
		default:
			l.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Location.
// This includes values selected through modifiers, order, etc.
func (l *Location) Value(name string) (ent.Value, error) {
	return l.selectValues.Get(name)
}

// QueryBusinessUnit queries the "business_unit" edge of the Location entity.
func (l *Location) QueryBusinessUnit() *BusinessUnitQuery {
	return NewLocationClient(l.config).QueryBusinessUnit(l)
}

// QueryOrganization queries the "organization" edge of the Location entity.
func (l *Location) QueryOrganization() *OrganizationQuery {
	return NewLocationClient(l.config).QueryOrganization(l)
}

// QueryLocationCategory queries the "location_category" edge of the Location entity.
func (l *Location) QueryLocationCategory() *LocationCategoryQuery {
	return NewLocationClient(l.config).QueryLocationCategory(l)
}

// QueryState queries the "state" edge of the Location entity.
func (l *Location) QueryState() *UsStateQuery {
	return NewLocationClient(l.config).QueryState(l)
}

// QueryComments queries the "comments" edge of the Location entity.
func (l *Location) QueryComments() *LocationCommentQuery {
	return NewLocationClient(l.config).QueryComments(l)
}

// QueryContacts queries the "contacts" edge of the Location entity.
func (l *Location) QueryContacts() *LocationContactQuery {
	return NewLocationClient(l.config).QueryContacts(l)
}

// Update returns a builder for updating this Location.
// Note that you need to call Location.Unwrap() before calling this method if this Location
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Location) Update() *LocationUpdateOne {
	return NewLocationClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the Location entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Location) Unwrap() *Location {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Location is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Location) String() string {
	var builder strings.Builder
	builder.WriteString("Location(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("business_unit_id=")
	builder.WriteString(fmt.Sprintf("%v", l.BusinessUnitID))
	builder.WriteString(", ")
	builder.WriteString("organization_id=")
	builder.WriteString(fmt.Sprintf("%v", l.OrganizationID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(l.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(l.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", l.Version))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", l.Status))
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(l.Code)
	builder.WriteString(", ")
	if v := l.LocationCategoryID; v != nil {
		builder.WriteString("location_category_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(l.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(l.Description)
	builder.WriteString(", ")
	builder.WriteString("address_line_1=")
	builder.WriteString(l.AddressLine1)
	builder.WriteString(", ")
	builder.WriteString("address_line_2=")
	builder.WriteString(l.AddressLine2)
	builder.WriteString(", ")
	builder.WriteString("city=")
	builder.WriteString(l.City)
	builder.WriteString(", ")
	builder.WriteString("state_id=")
	builder.WriteString(fmt.Sprintf("%v", l.StateID))
	builder.WriteString(", ")
	builder.WriteString("postal_code=")
	builder.WriteString(l.PostalCode)
	builder.WriteString(", ")
	builder.WriteString("longitude=")
	builder.WriteString(fmt.Sprintf("%v", l.Longitude))
	builder.WriteString(", ")
	builder.WriteString("latitude=")
	builder.WriteString(fmt.Sprintf("%v", l.Latitude))
	builder.WriteString(", ")
	builder.WriteString("place_id=")
	builder.WriteString(l.PlaceID)
	builder.WriteString(", ")
	builder.WriteString("is_geocoded=")
	builder.WriteString(fmt.Sprintf("%v", l.IsGeocoded))
	builder.WriteByte(')')
	return builder.String()
}

// NamedComments returns the Comments named value or an error if the edge was not
// loaded in eager-loading with this name.
func (l *Location) NamedComments(name string) ([]*LocationComment, error) {
	if l.Edges.namedComments == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := l.Edges.namedComments[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (l *Location) appendNamedComments(name string, edges ...*LocationComment) {
	if l.Edges.namedComments == nil {
		l.Edges.namedComments = make(map[string][]*LocationComment)
	}
	if len(edges) == 0 {
		l.Edges.namedComments[name] = []*LocationComment{}
	} else {
		l.Edges.namedComments[name] = append(l.Edges.namedComments[name], edges...)
	}
}

// NamedContacts returns the Contacts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (l *Location) NamedContacts(name string) ([]*LocationContact, error) {
	if l.Edges.namedContacts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := l.Edges.namedContacts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (l *Location) appendNamedContacts(name string, edges ...*LocationContact) {
	if l.Edges.namedContacts == nil {
		l.Edges.namedContacts = make(map[string][]*LocationContact)
	}
	if len(edges) == 0 {
		l.Edges.namedContacts[name] = []*LocationContact{}
	} else {
		l.Edges.namedContacts[name] = append(l.Edges.namedContacts[name], edges...)
	}
}

// Locations is a parsable slice of Location.
type Locations []*Location
