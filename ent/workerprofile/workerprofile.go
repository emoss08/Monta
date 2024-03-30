// Code generated by ent, DO NOT EDIT.

package workerprofile

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the workerprofile type in the database.
	Label = "worker_profile"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBusinessUnitID holds the string denoting the business_unit_id field in the database.
	FieldBusinessUnitID = "business_unit_id"
	// FieldOrganizationID holds the string denoting the organization_id field in the database.
	FieldOrganizationID = "organization_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldWorkerID holds the string denoting the worker_id field in the database.
	FieldWorkerID = "worker_id"
	// FieldRace holds the string denoting the race field in the database.
	FieldRace = "race"
	// FieldSex holds the string denoting the sex field in the database.
	FieldSex = "sex"
	// FieldDateOfBirth holds the string denoting the date_of_birth field in the database.
	FieldDateOfBirth = "date_of_birth"
	// FieldLicenseNumber holds the string denoting the license_number field in the database.
	FieldLicenseNumber = "license_number"
	// FieldLicenseStateID holds the string denoting the license_state_id field in the database.
	FieldLicenseStateID = "license_state_id"
	// FieldLicenseExpirationDate holds the string denoting the license_expiration_date field in the database.
	FieldLicenseExpirationDate = "license_expiration_date"
	// FieldEndorsements holds the string denoting the endorsements field in the database.
	FieldEndorsements = "endorsements"
	// FieldHazmatExpirationDate holds the string denoting the hazmat_expiration_date field in the database.
	FieldHazmatExpirationDate = "hazmat_expiration_date"
	// FieldHireDate holds the string denoting the hire_date field in the database.
	FieldHireDate = "hire_date"
	// FieldTerminationDate holds the string denoting the termination_date field in the database.
	FieldTerminationDate = "termination_date"
	// FieldPhysicalDueDate holds the string denoting the physical_due_date field in the database.
	FieldPhysicalDueDate = "physical_due_date"
	// FieldMedicalCertDate holds the string denoting the medical_cert_date field in the database.
	FieldMedicalCertDate = "medical_cert_date"
	// FieldMvrDueDate holds the string denoting the mvr_due_date field in the database.
	FieldMvrDueDate = "mvr_due_date"
	// EdgeBusinessUnit holds the string denoting the business_unit edge name in mutations.
	EdgeBusinessUnit = "business_unit"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// EdgeWorker holds the string denoting the worker edge name in mutations.
	EdgeWorker = "worker"
	// Table holds the table name of the workerprofile in the database.
	Table = "worker_profiles"
	// BusinessUnitTable is the table that holds the business_unit relation/edge.
	BusinessUnitTable = "worker_profiles"
	// BusinessUnitInverseTable is the table name for the BusinessUnit entity.
	// It exists in this package in order to avoid circular dependency with the "businessunit" package.
	BusinessUnitInverseTable = "business_units"
	// BusinessUnitColumn is the table column denoting the business_unit relation/edge.
	BusinessUnitColumn = "business_unit_id"
	// OrganizationTable is the table that holds the organization relation/edge.
	OrganizationTable = "worker_profiles"
	// OrganizationInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationInverseTable = "organizations"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "organization_id"
	// WorkerTable is the table that holds the worker relation/edge.
	WorkerTable = "worker_profiles"
	// WorkerInverseTable is the table name for the Worker entity.
	// It exists in this package in order to avoid circular dependency with the "worker" package.
	WorkerInverseTable = "workers"
	// WorkerColumn is the table column denoting the worker relation/edge.
	WorkerColumn = "worker_id"
)

// Columns holds all SQL columns for workerprofile fields.
var Columns = []string{
	FieldID,
	FieldBusinessUnitID,
	FieldOrganizationID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldWorkerID,
	FieldRace,
	FieldSex,
	FieldDateOfBirth,
	FieldLicenseNumber,
	FieldLicenseStateID,
	FieldLicenseExpirationDate,
	FieldEndorsements,
	FieldHazmatExpirationDate,
	FieldHireDate,
	FieldTerminationDate,
	FieldPhysicalDueDate,
	FieldMedicalCertDate,
	FieldMvrDueDate,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// LicenseNumberValidator is a validator for the "license_number" field. It is called by the builders before save.
	LicenseNumberValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Endorsements defines the type for the "endorsements" enum field.
type Endorsements string

// EndorsementsNone is the default value of the Endorsements enum.
const DefaultEndorsements = EndorsementsNone

// Endorsements values.
const (
	EndorsementsNone         Endorsements = "None"
	EndorsementsTanker       Endorsements = "Tanker"
	EndorsementsHazmat       Endorsements = "Hazmat"
	EndorsementsTankerHazmat Endorsements = "TankerHazmat"
)

func (e Endorsements) String() string {
	return string(e)
}

// EndorsementsValidator is a validator for the "endorsements" field enum values. It is called by the builders before save.
func EndorsementsValidator(e Endorsements) error {
	switch e {
	case EndorsementsNone, EndorsementsTanker, EndorsementsHazmat, EndorsementsTankerHazmat:
		return nil
	default:
		return fmt.Errorf("workerprofile: invalid enum value for endorsements field: %q", e)
	}
}

// OrderOption defines the ordering options for the WorkerProfile queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByBusinessUnitID orders the results by the business_unit_id field.
func ByBusinessUnitID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBusinessUnitID, opts...).ToFunc()
}

// ByOrganizationID orders the results by the organization_id field.
func ByOrganizationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrganizationID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByWorkerID orders the results by the worker_id field.
func ByWorkerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWorkerID, opts...).ToFunc()
}

// ByRace orders the results by the race field.
func ByRace(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRace, opts...).ToFunc()
}

// BySex orders the results by the sex field.
func BySex(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSex, opts...).ToFunc()
}

// ByDateOfBirth orders the results by the date_of_birth field.
func ByDateOfBirth(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDateOfBirth, opts...).ToFunc()
}

// ByLicenseNumber orders the results by the license_number field.
func ByLicenseNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLicenseNumber, opts...).ToFunc()
}

// ByLicenseStateID orders the results by the license_state_id field.
func ByLicenseStateID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLicenseStateID, opts...).ToFunc()
}

// ByLicenseExpirationDate orders the results by the license_expiration_date field.
func ByLicenseExpirationDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLicenseExpirationDate, opts...).ToFunc()
}

// ByEndorsements orders the results by the endorsements field.
func ByEndorsements(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndorsements, opts...).ToFunc()
}

// ByHazmatExpirationDate orders the results by the hazmat_expiration_date field.
func ByHazmatExpirationDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHazmatExpirationDate, opts...).ToFunc()
}

// ByHireDate orders the results by the hire_date field.
func ByHireDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHireDate, opts...).ToFunc()
}

// ByTerminationDate orders the results by the termination_date field.
func ByTerminationDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTerminationDate, opts...).ToFunc()
}

// ByPhysicalDueDate orders the results by the physical_due_date field.
func ByPhysicalDueDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhysicalDueDate, opts...).ToFunc()
}

// ByMedicalCertDate orders the results by the medical_cert_date field.
func ByMedicalCertDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMedicalCertDate, opts...).ToFunc()
}

// ByMvrDueDate orders the results by the mvr_due_date field.
func ByMvrDueDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMvrDueDate, opts...).ToFunc()
}

// ByBusinessUnitField orders the results by business_unit field.
func ByBusinessUnitField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBusinessUnitStep(), sql.OrderByField(field, opts...))
	}
}

// ByOrganizationField orders the results by organization field.
func ByOrganizationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrganizationStep(), sql.OrderByField(field, opts...))
	}
}

// ByWorkerField orders the results by worker field.
func ByWorkerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkerStep(), sql.OrderByField(field, opts...))
	}
}
func newBusinessUnitStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BusinessUnitInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
	)
}
func newOrganizationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrganizationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
	)
}
func newWorkerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, WorkerTable, WorkerColumn),
	)
}
