// COPYRIGHT(c) 2024 Trenova
//
// This file is part of Trenova.
//
// The Trenova software is licensed under the Business Source License 1.1. You are granted the right
// to copy, modify, and redistribute the software, but only for non-production use or with a total
// of less than three server instances. Starting from the Change Date (November 16, 2026), the
// software will be made available under version 2 or later of the GNU General Public License.
// If you use the software in violation of this license, your rights under the license will be
// terminated automatically. The software is provided "as is," and the Licensor disclaims all
// warranties and conditions. If you use this license's text or the "Business Source License" name
// and trademark, you must comply with the Licensor's covenants, which include specifying the
// Change License as the GPL Version 2.0 or a compatible license, specifying an Additional Use
// Grant, and not modifying the license in any other way.

package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/emoss08/trenova/internal/server"
	"github.com/emoss08/trenova/pkg/models"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
)

// EquipmentTypeService handles business logic for EquipmentType
type EquipmentTypeService struct {
	db     *bun.DB
	logger *zerolog.Logger
}

// NewEquipmentTypeService creates a new instance of EquipmentTypeService
func NewEquipmentTypeService(s *server.Server) *EquipmentTypeService {
	return &EquipmentTypeService{
		db:     s.DB,
		logger: s.Logger,
	}
}

// QueryFilter defines the filter parameters for querying EquipmentType
type EquipmentTypeQueryFilter struct {
	Query          string
	OrganizationID uuid.UUID
	BusinessUnitID uuid.UUID
	Limit          int
	Offset         int
}

// filterQuery applies filters to the query
func (s EquipmentTypeService) filterQuery(q *bun.SelectQuery, f *EquipmentTypeQueryFilter) *bun.SelectQuery {
	q = q.Where("et.organization_id = ?", f.OrganizationID).
		Where("et.business_unit_id = ?", f.BusinessUnitID)

	if f.Query != "" {
		q = q.Where("et.code = ? OR et.description ILIKE ?", f.Query, "%"+strings.ToLower(f.Query)+"%")
	}

	q = q.OrderExpr("CASE WHEN et.code = ? THEN 0 ELSE 1 END", f.Query).
		Order("et.created_at DESC")

	return q.Limit(f.Limit).Offset(f.Offset)
}

// GetAll retrieves all EquipmentType based on the provided filter
func (s EquipmentTypeService) GetAll(ctx context.Context, filter *EquipmentTypeQueryFilter) ([]*models.EquipmentType, int, error) {
	var entities []*models.EquipmentType

	q := s.db.NewSelect().
		Model(&entities)

	q = s.filterQuery(q, filter)

	count, err := q.ScanAndCount(ctx)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to fetch EquipmentType")
		return nil, 0, fmt.Errorf("failed to fetch EquipmentType: %w", err)
	}

	return entities, count, nil
}

// Get retrieves a single EquipmentType by ID
func (s EquipmentTypeService) Get(ctx context.Context, id, orgID, buID uuid.UUID) (*models.EquipmentType, error) {
	entity := new(models.EquipmentType)
	err := s.db.NewSelect().
		Model(entity).
		Where("et.organization_id = ?", orgID).
		Where("et.business_unit_id = ?", buID).
		Where("et.id = ?", id).
		Scan(ctx)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to fetch EquipmentType")
		return nil, fmt.Errorf("failed to fetch EquipmentType: %w", err)
	}

	return entity, nil
}

// Create creates a new EquipmentType
func (s EquipmentTypeService) Create(ctx context.Context, entity *models.EquipmentType) (*models.EquipmentType, error) {
	err := s.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewInsert().
			Model(entity).
			Returning("*").
			Exec(ctx)
		return err
	})
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to create EquipmentType")
		return nil, fmt.Errorf("failed to create EquipmentType: %w", err)
	}

	return entity, nil
}

// UpdateOne updates an existing EquipmentType
func (s EquipmentTypeService) UpdateOne(ctx context.Context, entity *models.EquipmentType) (*models.EquipmentType, error) {
	err := s.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if err := entity.OptimisticUpdate(ctx, tx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to update EquipmentType")
		return nil, fmt.Errorf("failed to update EquipmentType: %w", err)
	}

	return entity, nil
}