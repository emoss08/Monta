package repositories

import (
	"context"

	"github.com/trenova-app/transport/internal/core/domain/location"
	"github.com/trenova-app/transport/internal/core/ports"
	"github.com/trenova-app/transport/pkg/types/pulid"
)

type GetLocationCategoryByIDOptions struct {
	ID     pulid.ID
	OrgID  pulid.ID
	BuID   pulid.ID
	UserID pulid.ID
}

type LocationCategoryRepository interface {
	List(ctx context.Context, opts *ports.LimitOffsetQueryOptions) (*ports.ListResult[*location.LocationCategory], error)
	GetByID(ctx context.Context, opts GetLocationCategoryByIDOptions) (*location.LocationCategory, error)
	Create(ctx context.Context, lc *location.LocationCategory) (*location.LocationCategory, error)
	Update(ctx context.Context, lc *location.LocationCategory) (*location.LocationCategory, error)
}