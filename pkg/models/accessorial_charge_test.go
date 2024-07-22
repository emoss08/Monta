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

package models_test

import (
	"testing"

	"github.com/emoss08/trenova/pkg/models"
	"github.com/emoss08/trenova/pkg/models/property"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAccessorialCharge_Validate(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		ac := &models.AccessorialCharge{
			Status:         property.StatusActive,
			OrganizationID: uuid.New(),
			BusinessUnitID: uuid.New(),
			Code:           "CODE",
			Method:         "Distance",
			Description:    "Test Accessorial Charge",
		}

		err := ac.Validate()
		require.NoError(t, err)
	})

	t.Run("invalid", func(t *testing.T) {
		ac := &models.AccessorialCharge{
			Status:         property.StatusActive,
			OrganizationID: uuid.New(),
			BusinessUnitID: uuid.New(),
		}

		err := ac.Validate()
		require.Error(t, err)
	})
}