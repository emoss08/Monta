package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rotisserie/eris"
	"github.com/rs/zerolog"
	"github.com/trenova-app/transport/internal/core/domain/fleetcode"
	"github.com/trenova-app/transport/internal/core/ports"
	"github.com/trenova-app/transport/internal/core/ports/db"
	"github.com/trenova-app/transport/internal/core/ports/repositories"
	"github.com/trenova-app/transport/internal/pkg/errors"
	"github.com/trenova-app/transport/internal/pkg/logger"
	"github.com/trenova-app/transport/internal/pkg/utils/queryutils/queryfilters"
	"github.com/uptrace/bun"
	"go.uber.org/fx"
)

type FleetCodeRepositoryParams struct {
	fx.In

	DB     db.Connection
	Logger *logger.Logger
}

type fleetCodeRepository struct {
	db db.Connection
	l  *zerolog.Logger
}

func NewFleetCodeRepository(p FleetCodeRepositoryParams) repositories.FleetCodeRepository {
	log := p.Logger.With().
		Str("repository", "fleetcode").
		Logger()

	return &fleetCodeRepository{
		db: p.DB,
		l:  &log,
	}
}

func (fcr *fleetCodeRepository) filterQuery(q *bun.SelectQuery, opts *ports.LimitOffsetQueryOptions) *bun.SelectQuery {
	q = queryfilters.TenantFilterQuery(&queryfilters.TenantFilterQueryOptions{
		Query:      q,
		TableAlias: "fc",
		Filter:     opts,
	})

	if opts.Query != "" {
		q = q.Where("fc.name ILIKE ? OR fc.description ILIKE ?", "%"+opts.Query+"%", "%"+opts.Query+"%")
	}

	return q.Limit(opts.Limit).Offset(opts.Offset)
}

func (fcr *fleetCodeRepository) List(ctx context.Context, opts *ports.LimitOffsetQueryOptions) (*ports.ListResult[*fleetcode.FleetCode], error) {
	dba, err := fcr.db.DB(ctx)
	if err != nil {
		return nil, eris.Wrap(err, "get database connection")
	}

	log := fcr.l.With().
		Str("operation", "List").
		Str("buID", opts.TenantOpts.BuID.String()).
		Str("userID", opts.TenantOpts.UserID.String()).
		Logger()

	fcs := make([]*fleetcode.FleetCode, 0)

	q := dba.NewSelect().Model(&fcs)
	q = fcr.filterQuery(q, opts)

	total, err := q.ScanAndCount(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to scan fleet codes")
		return nil, eris.Wrap(err, "scan fleet codes")
	}

	return &ports.ListResult[*fleetcode.FleetCode]{
		Items: fcs,
		Total: total,
	}, nil
}

func (fcr *fleetCodeRepository) GetByID(ctx context.Context, opts repositories.GetFleetCodeByIDOptions) (*fleetcode.FleetCode, error) {
	dba, err := fcr.db.DB(ctx)
	if err != nil {
		return nil, eris.Wrap(err, "get database connection")
	}

	log := fcr.l.With().
		Str("operation", "GetByID").
		Str("fleetCodeID", opts.ID.String()).
		Logger()

	fc := new(fleetcode.FleetCode)

	query := dba.NewSelect().Model(fc).
		Where("fc.id = ? AND fc.organization_id = ? AND fc.business_unit_id = ?", opts.ID, opts.OrgID, opts.BuID)

	if err = query.Scan(ctx); err != nil {
		if eris.Is(err, sql.ErrNoRows) {
			return nil, errors.NewNotFoundError("Fleet code not found within your organization")
		}

		log.Error().Err(err).Msg("failed to get fleet code")
		return nil, eris.Wrap(err, "get fleet code")
	}

	return fc, nil
}

func (fcr *fleetCodeRepository) Create(ctx context.Context, fc *fleetcode.FleetCode) (*fleetcode.FleetCode, error) {
	dba, err := fcr.db.DB(ctx)
	if err != nil {
		return nil, eris.Wrap(err, "get database connection")
	}

	log := fcr.l.With().
		Str("operation", "Create").
		Str("orgID", fc.OrganizationID.String()).
		Str("buID", fc.BusinessUnitID.String()).
		Logger()

	err = dba.RunInTx(ctx, nil, func(c context.Context, tx bun.Tx) error {
		if _, iErr := tx.NewInsert().Model(fc).Exec(c); iErr != nil {
			log.Error().
				Err(iErr).
				Interface("fleetCode", fc).
				Msg("failed to insert fleet code")
			return eris.Wrap(iErr, "insert fleet code")
		}

		return nil
	})
	if err != nil {
		return nil, eris.Wrap(err, "create fleet code")
	}

	return fc, nil
}

func (fcr *fleetCodeRepository) Update(
	ctx context.Context,
	fc *fleetcode.FleetCode,
) (*fleetcode.FleetCode, error) {
	dba, err := fcr.db.DB(ctx)
	if err != nil {
		return nil, eris.Wrap(err, "get database connection")
	}

	log := fcr.l.With().
		Str("operation", "Update").
		Str("id", fc.GetID()).
		Int64("version", fc.Version).
		Logger()

	err = dba.RunInTx(ctx, nil, func(c context.Context, tx bun.Tx) error {
		ov := fc.Version

		fc.Version++

		results, rErr := tx.NewUpdate().
			Model(fc).
			WherePK().
			Where("fc.version = ?", ov).
			Where("fc.organization_id = ?", fc.OrganizationID).
			Where("fc.business_unit_id = ?", fc.BusinessUnitID).
			Returning("*").
			Exec(c)
		if rErr != nil {
			log.Error().
				Err(rErr).
				Interface("fleetCode", fc).
				Msg("failed to update fleet code")
			return eris.Wrap(rErr, "update fleet code")
		}

		rows, roErr := results.RowsAffected()
		if roErr != nil {
			log.Error().
				Err(roErr).
				Interface("fleetCode", fc).
				Msg("failed to get rows affected")
			return eris.Wrap(roErr, "get rows affected")
		}

		if rows == 0 {
			return errors.NewValidationError(
				"version",
				errors.ErrVersionMismatch,
				fmt.Sprintf("Version mismatch. The fleet code (%s) has either been updated or deleted since the last request.", fc.ID.String()),
			)
		}

		return nil
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to update fleet code")
		return nil, eris.Wrap(err, "update fleet code")
	}

	return fc, nil
}