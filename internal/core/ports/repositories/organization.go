package repositories

import (
	"context"

	"github.com/emoss08/trenova/internal/core/domain/organization"
	"github.com/emoss08/trenova/internal/core/ports"
	"github.com/emoss08/trenova/pkg/types/pulid"
)

type ListOrganizationResult struct {
	Organizations []*organization.Organization
	Total         int
}

type GetOrgByIDOptions struct {
	// ID of the organization
	OrgID pulid.ID

	// ID of the business unit
	BuID pulid.ID

	// IncludeState includes the state in the response
	IncludeState bool

	// IncludeBu includes the business unit in the response
	IncludeBu bool
}

type OrganizationRepository interface {
	List(ctx context.Context, opts *ports.LimitOffsetQueryOptions) (*ListOrganizationResult, error)
	GetByID(ctx context.Context, opts GetOrgByIDOptions) (*organization.Organization, error)
	Create(ctx context.Context, org *organization.Organization, requesterID pulid.ID) (*organization.Organization, error)
	Update(ctx context.Context, org *organization.Organization, requesterID pulid.ID) (*organization.Organization, error)
	SetLogo(ctx context.Context, org *organization.Organization, userID pulid.ID) (*organization.Organization, error)
	ClearLogo(ctx context.Context, org *organization.Organization, userID pulid.ID) (*organization.Organization, error)
	GetUserOrganizations(ctx context.Context, opts *ports.LimitOffsetQueryOptions) (*ListOrganizationResult, error)
}
