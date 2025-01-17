package organization

import (
	"context"
	"mime/multipart"
	"net/http"

	"github.com/rotisserie/eris"
	"github.com/rs/zerolog"
	"github.com/trenova-app/transport/internal/core/domain/organization"
	"github.com/trenova-app/transport/internal/core/domain/permission"
	"github.com/trenova-app/transport/internal/core/ports"
	"github.com/trenova-app/transport/internal/core/ports/repositories"
	"github.com/trenova-app/transport/internal/core/ports/services"
	"github.com/trenova-app/transport/internal/pkg/errors"
	"github.com/trenova-app/transport/internal/pkg/logger"
	"github.com/trenova-app/transport/internal/pkg/utils/fileutils"
	"github.com/trenova-app/transport/pkg/types"
	"github.com/trenova-app/transport/pkg/types/pulid"
	"go.uber.org/fx"
)

type ServiceParams struct {
	fx.In

	Logger      *logger.Logger
	Repo        repositories.OrganizationRepository
	PermService services.PermissionService
	FileService services.FileService
}

type Service struct {
	repo repositories.OrganizationRepository
	l    *zerolog.Logger
	ps   services.PermissionService
	fs   services.FileService
}

func NewService(p ServiceParams) *Service {
	log := p.Logger.With().
		Str("service", "organization").
		Logger()

	return &Service{
		repo: p.Repo,
		ps:   p.PermService,
		fs:   p.FileService,
		l:    &log,
	}
}

// SelectOptions returns a list of select options for organizations.
func (s *Service) SelectOptions(ctx context.Context, opts *ports.LimitOffsetQueryOptions) ([]types.SelectOption, error) {
	result, err := s.repo.List(ctx, opts)
	if err != nil {
		return nil, eris.Wrap(err, "failed to list organizations")
	}

	// Convert the organizations to select options
	options := make([]types.SelectOption, len(result.Organizations))
	for i, org := range result.Organizations {
		options[i] = types.SelectOption{
			Value: org.ID.String(),
			Label: org.Name,
		}
	}

	return options, nil
}

// List returns a list of organizations.
func (s *Service) List(ctx context.Context, opts *ports.LimitOffsetQueryOptions) (*ports.ListResult[*organization.Organization], error) {
	log := s.l.With().
		Str("operation", "List").
		Logger()

	result, err := s.ps.HasAnyPermissions(ctx,
		[]*services.PermissionCheck{
			{
				UserID:         opts.TenantOpts.UserID,
				Resource:       permission.ResourceOrganization,
				Action:         permission.ActionRead,
				BusinessUnitID: opts.TenantOpts.BuID,
			},
		},
	)
	if err != nil {
		s.l.Error().Err(err).Msg("failed to check permissions")
		return nil, eris.Wrap(err, "failed to check permissions")
	}

	if !result.Allowed {
		return nil, errors.NewAuthorizationError("You do not have permission to read organizations")
	}

	entities, err := s.repo.List(ctx, opts)
	if err != nil {
		log.Error().Err(err).Msg("failed to list organizations")
		return nil, eris.Wrap(err, "failed to list organizations")
	}

	return &ports.ListResult[*organization.Organization]{
		Items: entities.Organizations,
		Total: entities.Total,
	}, nil
}

// Get returns an organization by its ID.
func (s *Service) Get(ctx context.Context, opts repositories.GetOrgByIDOptions) (*organization.Organization, error) {
	log := s.l.With().
		Str("operation", "Get").
		Str("orgID", opts.OrgID.String()).
		Logger()

	org, err := s.repo.GetByID(ctx, opts)
	if err != nil {
		log.Error().Err(err).Msg("failed to get organization by id")
		return nil, eris.Wrap(err, "failed to get organization by id")
	}

	return org, nil
}

// Create creates an organization.
func (s *Service) Create(ctx context.Context, org *organization.Organization, userID pulid.ID) (*organization.Organization, error) {
	log := s.l.With().Str("operation", "Create").
		Str("orgID", org.ID.String()).
		Str("userID", userID.String()).
		Logger()

	result, err := s.ps.HasAnyPermissions(ctx, []*services.PermissionCheck{
		{
			UserID:         userID,
			Resource:       permission.ResourceOrganization,
			Action:         permission.ActionCreate,
			BusinessUnitID: org.BusinessUnitID,
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to check permissions")
		return nil, eris.Wrap(err, "failed to check permissions")
	}

	if !result.Allowed {
		return nil, errors.NewAuthorizationError("You do not have permission to create an organization")
	}

	createdOrg, err := s.repo.Create(ctx, org, userID)
	if err != nil {
		s.l.Error().Err(err).Interface("org", org).Msg("failed to create organization")
		return nil, eris.Wrap(err, "failed to create organization")
	}

	return createdOrg, nil
}

// Update updates an organization.
func (s *Service) Update(ctx context.Context, org *organization.Organization, requesterID pulid.ID) (*organization.Organization, error) {
	log := s.l.With().
		Str("operation", "Update").
		Interface("org", org).
		Logger()

	result, err := s.ps.HasAnyPermissions(ctx, []*services.PermissionCheck{
		{
			UserID:         requesterID,
			Resource:       permission.ResourceOrganization,
			Action:         permission.ActionUpdate,
			BusinessUnitID: org.BusinessUnitID,
			ResourceID:     org.ID,
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to check permissions")
		return nil, eris.Wrap(err, "failed to check permissions")
	}

	if !result.Allowed {
		return nil, errors.NewAuthorizationError("You do not have permission to update this organization")
	}

	updatedOrg, err := s.repo.Update(ctx, org, requesterID)
	if err != nil {
		log.Error().Err(err).Interface("org", org).Msg("failed to update organization")
		return nil, eris.Wrap(err, "failed to update organization")
	}

	return updatedOrg, nil
}

func (s *Service) SetLogo(ctx context.Context, orgID, buID, userID pulid.ID, logo *multipart.FileHeader) (*organization.Organization, error) {
	result, err := s.ps.HasFieldPermission(ctx, &services.PermissionCheck{
		UserID:     userID,
		Resource:   permission.ResourceOrganization,
		Action:     permission.ActionModifyField,
		ResourceID: orgID,
		Field:      "logo_url",
	})
	if err != nil {
		return nil, eris.Wrap(err, "check field permission")
	}

	if !result.Allowed {
		return nil, errors.NewAuthorizationError("You do not have permission to set the logo for this organization")
	}

	fileData, err := fileutils.ReadFileData(logo)
	if err != nil {
		return nil, eris.Wrap(err, "read file data")
	}

	fileName, err := fileutils.RenameFile(logo, orgID.String())
	if err != nil {
		return nil, eris.Wrap(err, "rename file")
	}

	s.l.Info().Str("fileName", fileName).Msg("file renamed")

	// Get the organization from the DB
	org, err := s.repo.GetByID(ctx, repositories.GetOrgByIDOptions{
		OrgID: orgID,
		BuID:  buID,
	})
	if err != nil {
		return nil, eris.Wrap(err, "get organization")
	}

	if org.BucketName == "" {
		return nil, ErrOrgBucketNameNotSet
	}

	org.Metadata = &organization.Metadata{
		ObjectID: fileName,
	}

	return s.uploadLogo(ctx, org, userID, &services.SaveFileRequest{
		File:           fileData,
		FileName:       fileName,
		FileType:       services.ImageFile,
		Classification: services.ClassificationPublic,
		Category:       services.CategoryBranding,
		Tags:           map[string]string{"organization_id": org.ID.String()},
		OrgID:          org.ID.String(),
		BucketName:     org.BucketName,
		UserID:         userID.String(),
		Metadata: http.Header{
			"organization_id": []string{org.ID.String()},
			"user_id":         []string{userID.String()},
			"file_type":       []string{string(services.ImageFile)},
		},
	})
}

func (s *Service) uploadLogo(ctx context.Context, org *organization.Organization, userID pulid.ID, params *services.SaveFileRequest) (*organization.Organization, error) {
	ui, err := s.fs.SaveFile(ctx, params)
	if err != nil {
		return nil, eris.Wrap(err, "save file")
	}

	// Set the logo URL in the organization
	org.LogoURL = ui.Location

	updatedOrg, err := s.repo.SetLogo(ctx, org, userID)
	if err != nil {
		return nil, eris.Wrap(err, "set logo")
	}

	return updatedOrg, nil
}

func (s *Service) ClearLogo(ctx context.Context, orgID, buID, userID pulid.ID) (*organization.Organization, error) {
	result, err := s.ps.HasFieldPermission(ctx, &services.PermissionCheck{
		UserID:     userID,
		Resource:   permission.ResourceOrganization,
		Action:     permission.ActionModifyField,
		ResourceID: orgID,
		Field:      "logo_url",
	})
	if err != nil {
		return nil, eris.Wrap(err, "check field permission")
	}

	if !result.Allowed {
		return nil, errors.NewAuthorizationError("You do not have permission to clear the logo for this organization")
	}

	org, err := s.repo.GetByID(ctx, repositories.GetOrgByIDOptions{
		OrgID: orgID,
		BuID:  buID,
	})
	if err != nil {
		return nil, eris.Wrap(err, "get organization")
	}

	if org.Metadata != nil && org.Metadata.ObjectID != "" {
		err = s.fs.DeleteFile(ctx, org.BucketName, org.Metadata.ObjectID)
		if err != nil {
			s.l.Warn().
				Err(err).
				Str("orgID", orgID.String()).
				Str("bucketName", org.BucketName).
				Str("objectID", org.Metadata.ObjectID).
				Msg("failed to delete file")
			// ! Non-crticial error, continue with clearing the logo
		}
	}

	updatedOrg, err := s.repo.ClearLogo(ctx, org, userID)
	if err != nil {
		s.l.Error().Err(err).Str("orgID", orgID.String()).Msg("failed to clear logo")
		return nil, eris.Wrap(err, "failed to clear logo")
	}

	return updatedOrg, nil
}

func (s *Service) GetUserOrganizations(
	ctx context.Context, opts *ports.LimitOffsetQueryOptions,
) (*ports.ListResult[*organization.Organization], error) {
	result, err := s.repo.GetUserOrganizations(ctx, opts)
	if err != nil {
		return nil, eris.Wrap(err, "get user organizations")
	}

	return &ports.ListResult[*organization.Organization]{
		Items: result.Organizations,
		Total: result.Total,
	}, nil
}