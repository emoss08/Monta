// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/usstate"
	"github.com/emoss08/trenova/ent/worker"
	"github.com/emoss08/trenova/ent/workerprofile"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// WorkerProfile is the model entity for the WorkerProfile schema.
type WorkerProfile struct {
	config `json:"-" validate:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// BusinessUnitID holds the value of the "business_unit_id" field.
	BusinessUnitID uuid.UUID `json:"businessUnitId"`
	// OrganizationID holds the value of the "organization_id" field.
	OrganizationID uuid.UUID `json:"organizationId"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt"`
	// WorkerID holds the value of the "worker_id" field.
	WorkerID uuid.UUID `json:"workerId" validate:"required,uuid"`
	// Race holds the value of the "race" field.
	Race string `json:"race" validate:"omitempty"`
	// Sex holds the value of the "sex" field.
	Sex string `json:"sex,omitempty" validate:"omitempty"`
	// DateOfBirth holds the value of the "date_of_birth" field.
	DateOfBirth *pgtype.Date `json:"dateOfBirth" validate:"omitempty"`
	// LicenseNumber holds the value of the "license_number" field.
	LicenseNumber string `json:"licenseNumber" validate:"required"`
	// LicenseStateID holds the value of the "license_state_id" field.
	LicenseStateID uuid.UUID `json:"licenseStateId" validate:"omitempty,uuid"`
	// LicenseExpirationDate holds the value of the "license_expiration_date" field.
	LicenseExpirationDate *pgtype.Date `json:"licenseExpirationDate" validate:"omitempty"`
	// Endorsements holds the value of the "endorsements" field.
	Endorsements workerprofile.Endorsements `json:"endorsements" validate:"omitempty"`
	// HazmatExpirationDate holds the value of the "hazmat_expiration_date" field.
	HazmatExpirationDate *pgtype.Date `json:"hazmatExpirationDate" validate:"omitempty"`
	// HireDate holds the value of the "hire_date" field.
	HireDate *pgtype.Date `json:"hireDate" validate:"omitempty"`
	// TerminationDate holds the value of the "termination_date" field.
	TerminationDate *pgtype.Date `json:"terminationDate" validate:"omitempty"`
	// PhysicalDueDate holds the value of the "physical_due_date" field.
	PhysicalDueDate *pgtype.Date `json:"physicalDueDate" validate:"omitempty"`
	// MedicalCertDate holds the value of the "medical_cert_date" field.
	MedicalCertDate *pgtype.Date `json:"medicalCertDate" validate:"omitempty"`
	// MvrDueDate holds the value of the "mvr_due_date" field.
	MvrDueDate *pgtype.Date `json:"mvrDueDate" validate:"omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkerProfileQuery when eager-loading is set.
	Edges        WorkerProfileEdges `json:"edges"`
	selectValues sql.SelectValues
}

// WorkerProfileEdges holds the relations/edges for other nodes in the graph.
type WorkerProfileEdges struct {
	// BusinessUnit holds the value of the business_unit edge.
	BusinessUnit *BusinessUnit `json:"business_unit,omitempty"`
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// Worker holds the value of the worker edge.
	Worker *Worker `json:"worker,omitempty"`
	// State holds the value of the state edge.
	State *UsState `json:"state"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// BusinessUnitOrErr returns the BusinessUnit value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkerProfileEdges) BusinessUnitOrErr() (*BusinessUnit, error) {
	if e.BusinessUnit != nil {
		return e.BusinessUnit, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: businessunit.Label}
	}
	return nil, &NotLoadedError{edge: "business_unit"}
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkerProfileEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// WorkerOrErr returns the Worker value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkerProfileEdges) WorkerOrErr() (*Worker, error) {
	if e.Worker != nil {
		return e.Worker, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: worker.Label}
	}
	return nil, &NotLoadedError{edge: "worker"}
}

// StateOrErr returns the State value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkerProfileEdges) StateOrErr() (*UsState, error) {
	if e.State != nil {
		return e.State, nil
	} else if e.loadedTypes[3] {
		return nil, &NotFoundError{label: usstate.Label}
	}
	return nil, &NotLoadedError{edge: "state"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WorkerProfile) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case workerprofile.FieldDateOfBirth, workerprofile.FieldLicenseExpirationDate, workerprofile.FieldHazmatExpirationDate, workerprofile.FieldHireDate, workerprofile.FieldTerminationDate, workerprofile.FieldPhysicalDueDate, workerprofile.FieldMedicalCertDate, workerprofile.FieldMvrDueDate:
			values[i] = new(pgtype.Date)
		case workerprofile.FieldRace, workerprofile.FieldSex, workerprofile.FieldLicenseNumber, workerprofile.FieldEndorsements:
			values[i] = new(sql.NullString)
		case workerprofile.FieldCreatedAt, workerprofile.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case workerprofile.FieldID, workerprofile.FieldBusinessUnitID, workerprofile.FieldOrganizationID, workerprofile.FieldWorkerID, workerprofile.FieldLicenseStateID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WorkerProfile fields.
func (wp *WorkerProfile) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workerprofile.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				wp.ID = *value
			}
		case workerprofile.FieldBusinessUnitID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field business_unit_id", values[i])
			} else if value != nil {
				wp.BusinessUnitID = *value
			}
		case workerprofile.FieldOrganizationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value != nil {
				wp.OrganizationID = *value
			}
		case workerprofile.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				wp.CreatedAt = value.Time
			}
		case workerprofile.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				wp.UpdatedAt = value.Time
			}
		case workerprofile.FieldWorkerID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field worker_id", values[i])
			} else if value != nil {
				wp.WorkerID = *value
			}
		case workerprofile.FieldRace:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field race", values[i])
			} else if value.Valid {
				wp.Race = value.String
			}
		case workerprofile.FieldSex:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sex", values[i])
			} else if value.Valid {
				wp.Sex = value.String
			}
		case workerprofile.FieldDateOfBirth:
			if value, ok := values[i].(*pgtype.Date); !ok {
				return fmt.Errorf("unexpected type %T for field date_of_birth", values[i])
			} else if value != nil {
				wp.DateOfBirth = value
			}
		case workerprofile.FieldLicenseNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field license_number", values[i])
			} else if value.Valid {
				wp.LicenseNumber = value.String
			}
		case workerprofile.FieldLicenseStateID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field license_state_id", values[i])
			} else if value != nil {
				wp.LicenseStateID = *value
			}
		case workerprofile.FieldLicenseExpirationDate:
			if value, ok := values[i].(*pgtype.Date); !ok {
				return fmt.Errorf("unexpected type %T for field license_expiration_date", values[i])
			} else if value != nil {
				wp.LicenseExpirationDate = value
			}
		case workerprofile.FieldEndorsements:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field endorsements", values[i])
			} else if value.Valid {
				wp.Endorsements = workerprofile.Endorsements(value.String)
			}
		case workerprofile.FieldHazmatExpirationDate:
			if value, ok := values[i].(*pgtype.Date); !ok {
				return fmt.Errorf("unexpected type %T for field hazmat_expiration_date", values[i])
			} else if value != nil {
				wp.HazmatExpirationDate = value
			}
		case workerprofile.FieldHireDate:
			if value, ok := values[i].(*pgtype.Date); !ok {
				return fmt.Errorf("unexpected type %T for field hire_date", values[i])
			} else if value != nil {
				wp.HireDate = value
			}
		case workerprofile.FieldTerminationDate:
			if value, ok := values[i].(*pgtype.Date); !ok {
				return fmt.Errorf("unexpected type %T for field termination_date", values[i])
			} else if value != nil {
				wp.TerminationDate = value
			}
		case workerprofile.FieldPhysicalDueDate:
			if value, ok := values[i].(*pgtype.Date); !ok {
				return fmt.Errorf("unexpected type %T for field physical_due_date", values[i])
			} else if value != nil {
				wp.PhysicalDueDate = value
			}
		case workerprofile.FieldMedicalCertDate:
			if value, ok := values[i].(*pgtype.Date); !ok {
				return fmt.Errorf("unexpected type %T for field medical_cert_date", values[i])
			} else if value != nil {
				wp.MedicalCertDate = value
			}
		case workerprofile.FieldMvrDueDate:
			if value, ok := values[i].(*pgtype.Date); !ok {
				return fmt.Errorf("unexpected type %T for field mvr_due_date", values[i])
			} else if value != nil {
				wp.MvrDueDate = value
			}
		default:
			wp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WorkerProfile.
// This includes values selected through modifiers, order, etc.
func (wp *WorkerProfile) Value(name string) (ent.Value, error) {
	return wp.selectValues.Get(name)
}

// QueryBusinessUnit queries the "business_unit" edge of the WorkerProfile entity.
func (wp *WorkerProfile) QueryBusinessUnit() *BusinessUnitQuery {
	return NewWorkerProfileClient(wp.config).QueryBusinessUnit(wp)
}

// QueryOrganization queries the "organization" edge of the WorkerProfile entity.
func (wp *WorkerProfile) QueryOrganization() *OrganizationQuery {
	return NewWorkerProfileClient(wp.config).QueryOrganization(wp)
}

// QueryWorker queries the "worker" edge of the WorkerProfile entity.
func (wp *WorkerProfile) QueryWorker() *WorkerQuery {
	return NewWorkerProfileClient(wp.config).QueryWorker(wp)
}

// QueryState queries the "state" edge of the WorkerProfile entity.
func (wp *WorkerProfile) QueryState() *UsStateQuery {
	return NewWorkerProfileClient(wp.config).QueryState(wp)
}

// Update returns a builder for updating this WorkerProfile.
// Note that you need to call WorkerProfile.Unwrap() before calling this method if this WorkerProfile
// was returned from a transaction, and the transaction was committed or rolled back.
func (wp *WorkerProfile) Update() *WorkerProfileUpdateOne {
	return NewWorkerProfileClient(wp.config).UpdateOne(wp)
}

// Unwrap unwraps the WorkerProfile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wp *WorkerProfile) Unwrap() *WorkerProfile {
	_tx, ok := wp.config.driver.(*txDriver)
	if !ok {
		panic("ent: WorkerProfile is not a transactional entity")
	}
	wp.config.driver = _tx.drv
	return wp
}

// String implements the fmt.Stringer.
func (wp *WorkerProfile) String() string {
	var builder strings.Builder
	builder.WriteString("WorkerProfile(")
	builder.WriteString(fmt.Sprintf("id=%v, ", wp.ID))
	builder.WriteString("business_unit_id=")
	builder.WriteString(fmt.Sprintf("%v", wp.BusinessUnitID))
	builder.WriteString(", ")
	builder.WriteString("organization_id=")
	builder.WriteString(fmt.Sprintf("%v", wp.OrganizationID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(wp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(wp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("worker_id=")
	builder.WriteString(fmt.Sprintf("%v", wp.WorkerID))
	builder.WriteString(", ")
	builder.WriteString("race=")
	builder.WriteString(wp.Race)
	builder.WriteString(", ")
	builder.WriteString("sex=")
	builder.WriteString(wp.Sex)
	builder.WriteString(", ")
	builder.WriteString("date_of_birth=")
	builder.WriteString(fmt.Sprintf("%v", wp.DateOfBirth))
	builder.WriteString(", ")
	builder.WriteString("license_number=")
	builder.WriteString(wp.LicenseNumber)
	builder.WriteString(", ")
	builder.WriteString("license_state_id=")
	builder.WriteString(fmt.Sprintf("%v", wp.LicenseStateID))
	builder.WriteString(", ")
	builder.WriteString("license_expiration_date=")
	builder.WriteString(fmt.Sprintf("%v", wp.LicenseExpirationDate))
	builder.WriteString(", ")
	builder.WriteString("endorsements=")
	builder.WriteString(fmt.Sprintf("%v", wp.Endorsements))
	builder.WriteString(", ")
	builder.WriteString("hazmat_expiration_date=")
	builder.WriteString(fmt.Sprintf("%v", wp.HazmatExpirationDate))
	builder.WriteString(", ")
	builder.WriteString("hire_date=")
	builder.WriteString(fmt.Sprintf("%v", wp.HireDate))
	builder.WriteString(", ")
	builder.WriteString("termination_date=")
	builder.WriteString(fmt.Sprintf("%v", wp.TerminationDate))
	builder.WriteString(", ")
	builder.WriteString("physical_due_date=")
	builder.WriteString(fmt.Sprintf("%v", wp.PhysicalDueDate))
	builder.WriteString(", ")
	builder.WriteString("medical_cert_date=")
	builder.WriteString(fmt.Sprintf("%v", wp.MedicalCertDate))
	builder.WriteString(", ")
	builder.WriteString("mvr_due_date=")
	builder.WriteString(fmt.Sprintf("%v", wp.MvrDueDate))
	builder.WriteByte(')')
	return builder.String()
}

// WorkerProfiles is a parsable slice of WorkerProfile.
type WorkerProfiles []*WorkerProfile
