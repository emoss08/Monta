package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.63

import (
	"context"
	"fmt"

	"github.com/emoss08/trenova/internal/api/graph/generated"
	"github.com/emoss08/trenova/internal/core/domain/organization"
)

// CreateOrganization is the resolver for the createOrganization field.
func (r *mutationResolver) CreateOrganization(ctx context.Context, input generated.CreateOrganizationInput) (*organization.Organization, error) {
	panic(fmt.Errorf("not implemented: CreateOrganization - createOrganization"))
}

// Organizations is the resolver for the organizations field.
func (r *queryResolver) Organizations(ctx context.Context, includeBusinessUnit *bool) ([]*organization.Organization, error) {
	panic(fmt.Errorf("not implemented: Organizations - organizations"))
}
