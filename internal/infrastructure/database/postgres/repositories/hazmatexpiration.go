package repositories

import (
	"context"

	"github.com/rotisserie/eris"
	"github.com/rs/zerolog"
	"github.com/trenova-app/transport/internal/core/domain/compliance"
	"github.com/trenova-app/transport/internal/core/ports/db"
	"github.com/trenova-app/transport/internal/core/ports/repositories"
	"github.com/trenova-app/transport/internal/pkg/logger"
	"github.com/trenova-app/transport/pkg/types/pulid"
	"go.uber.org/fx"
)

// TODO(Wolfred): We should add caching to this. Since the hazmat expiration is expected to never change.
type HazmatExpirationRepositoryParams struct {
	fx.In

	DB     db.Connection
	Logger *logger.Logger
}

type hazmatExpirationRepository struct {
	db     db.Connection
	logger *zerolog.Logger
}

func NewHazmatExpirationRepository(p HazmatExpirationRepositoryParams) repositories.HazmatExpirationRepository {
	log := p.Logger.With().
		Str("repository", "hazmat_expiration").
		Logger()

	return &hazmatExpirationRepository{
		db:     p.DB,
		logger: &log,
	}
}

func (r *hazmatExpirationRepository) GetHazmatExpirationByStateID(ctx context.Context, stateID pulid.ID) (*compliance.HazmatExpiration, error) {
	dba, err := r.db.DB(ctx)
	if err != nil {
		return nil, eris.Wrap(err, "get database connection")
	}

	r.logger.Debug().Str("state_id", stateID.String()).Msg("getting hazmat expiration")

	expiration := new(compliance.HazmatExpiration)
	err = dba.NewSelect().Model(expiration).
		Where("state_id = ?", stateID).
		Scan(ctx)
	if err != nil {
		return nil, eris.Wrap(err, "get hazmat expiration by state id")
	}

	r.logger.Debug().Interface("expiration", expiration).Msg("hazmat expiration")

	return expiration, nil
}