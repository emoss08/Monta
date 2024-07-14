package fixtures

import (
	"context"
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/emoss08/trenova/pkg/models"
	"github.com/fatih/color"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

func LoadAdminAccount(ctx context.Context, db *bun.DB, enforcer *casbin.Enforcer, org *models.Organization, bu *models.BusinessUnit) error {
	exists, err := db.NewSelect().Model((*models.User)(nil)).Where("username = ?", "admin").Exists(ctx)
	if err != nil {
		return err
	}

	if !exists {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		user := &models.User{
			OrganizationID: org.ID,
			Organization:   org,
			BusinessUnitID: bu.ID,
			BusinessUnit:   bu,
			Status:         "Active",
			Username:       "admin",
			Password:       string(hashedPassword),
			Email:          "admin@trenova.app",
			Name:           "System Administrator",
			IsAdmin:        true,
			Timezone:       "America/New_York",
		}

		_, err = db.NewInsert().Model(user).Exec(ctx)
		if err != nil {
			return err
		}

		// Assign the Admin role to the user
		_, err = enforcer.AddGroupingPolicy(user.ID.String(), "Admin", "role")
		if err != nil {
			return fmt.Errorf("failed to assign Admin role: %w", err)
		}

		log.Printf("Assigned Admin role to user: %s\n", user.ID.String())

		// Print out the admin account credentials
		color.Yellow("✅ Admin account created successfully")
		color.Yellow("-----------------------------")
		color.Yellow("Admin account credentials:")
		color.Yellow("Email: admin@trenova.app")
		color.Yellow("Password: admin")
		color.Yellow("-----------------------------")
	}

	return enforcer.SavePolicy()
}

// Normal Account is an account with no permissions assigned
func LoadNormalAccount(ctx context.Context, db *bun.DB, org *models.Organization, bu *models.BusinessUnit) error {
	exists, err := db.NewSelect().Model((*models.User)(nil)).Where("username = ?", "user").Exists(ctx)
	if err != nil {
		return err
	}

	if !exists {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("user"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		user := &models.User{
			OrganizationID: org.ID,
			Organization:   org,
			BusinessUnitID: bu.ID,
			BusinessUnit:   bu,
			Status:         "Active",
			Username:       "user",
			Password:       string(hashedPassword),
			Email:          "user@trenova.app",
			Name:           "Normal User",
			IsAdmin:        false,
			Timezone:       "America/New_York",
		}

		_, err = db.NewInsert().Model(user).Exec(ctx)
		if err != nil {
			return err
		}

		// Print out the normal account credentials
		color.Yellow("✅ Normal account created successfully")
		color.Yellow("-----------------------------")
		color.Yellow("Normal account credentials:")
		color.Yellow("Email: user@trenova.app")
		color.Yellow("Password: user")
		color.Yellow("-----------------------------")
	}

	return nil
}
