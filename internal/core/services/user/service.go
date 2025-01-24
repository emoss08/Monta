package user

import (
	"context"

	"github.com/emoss08/trenova/internal/core/domain/permission"
	"github.com/emoss08/trenova/internal/core/domain/user"
	"github.com/emoss08/trenova/internal/core/ports"
	"github.com/emoss08/trenova/internal/core/ports/repositories"
	"github.com/emoss08/trenova/internal/core/ports/services"
	"github.com/emoss08/trenova/internal/pkg/errors"
	"github.com/emoss08/trenova/internal/pkg/logger"
	"github.com/emoss08/trenova/pkg/types"
	"github.com/rotisserie/eris"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

type ServiceParams struct {
	fx.In

	Logger      *logger.Logger
	Repo        repositories.UserRepository
	PermService services.PermissionService
}

type Service struct {
	repo repositories.UserRepository
	l    *zerolog.Logger
	ps   services.PermissionService
}

func NewService(p ServiceParams) *Service {
	log := p.Logger.With().
		Str("service", "user").
		Logger()

	return &Service{
		repo: p.Repo,
		l:    &log,
		ps:   p.PermService,
	}
}

func (s *Service) SelectOptions(ctx context.Context, opts *ports.LimitOffsetQueryOptions) ([]*types.SelectOption, error) {
	result, err := s.repo.List(ctx, opts)
	if err != nil {
		return nil, eris.Wrap(err, "select users")
	}

	options := make([]*types.SelectOption, len(result.Items))
	for i, u := range result.Items {
		options[i] = &types.SelectOption{
			Value: u.ID.String(),
			Label: u.Name,
		}
	}

	return options, nil
}

func (s *Service) GetByID(ctx context.Context, opts repositories.GetUserByIDOptions) (*user.User, error) {
	result, err := s.ps.HasPermission(ctx, &services.PermissionCheck{
		UserID:         opts.UserID,
		BusinessUnitID: opts.BuID,
		OrganizationID: opts.OrgID,
		Resource:       permission.ResourceUser,
		Action:         permission.ActionRead,
		ResourceID:     opts.UserID,
	})
	if err != nil {
		return nil, eris.Wrap(err, "check permission")
	}

	if !result.Allowed {
		s.l.Error().Msg("user does not have permission to view user")
		return nil, errors.NewAuthorizationError("You do not have permission to view this user")
	}

	return s.repo.GetByID(ctx, opts)
}
