package repositories

import (
	"context"

	"github.com/trenova-app/transport/internal/core/domain/shipmenttype"
	"github.com/trenova-app/transport/internal/core/ports"
	"github.com/trenova-app/transport/pkg/types/pulid"
)

type GetShipmentTypeByIDOptions struct {
	ID     pulid.ID
	OrgID  pulid.ID
	BuID   pulid.ID
	UserID pulid.ID
}

type ShipmentTypeRepository interface {
	List(ctx context.Context, opts *ports.LimitOffsetQueryOptions) (*ports.ListResult[*shipmenttype.ShipmentType], error)
	GetByID(ctx context.Context, opts GetShipmentTypeByIDOptions) (*shipmenttype.ShipmentType, error)
	Create(ctx context.Context, st *shipmenttype.ShipmentType) (*shipmenttype.ShipmentType, error)
	Update(ctx context.Context, st *shipmenttype.ShipmentType) (*shipmenttype.ShipmentType, error)
}