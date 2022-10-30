# -*- coding: utf-8 -*-
"""
COPYRIGHT 2022 MONTA

This file is part of Monta.

Monta is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Monta is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Monta.  If not, see <https://www.gnu.org/licenses/>.
"""

from typing import final

from django.core.exceptions import ValidationError
from django.db import models
from django.urls import reverse
from django.utils.translation import gettext_lazy as _
from django_extensions.db.models import TimeStampedModel

from organization.models import Organization


@final
class IntegrationChoices(models.TextChoices):
    """
    Integration Choices
    """

    GOOGLE_MAPS = "google_maps", _("Google Maps")
    GOOGLE_PLACES = "google_places", _("Google Places")


@final
class IntegrationAuthTypes(models.TextChoices):
    """
    Integration Auth Types
    """

    NO_AUTH = "no_auth", _("No Auth")
    API_KEY = "api_key", _("API Key")
    BEARER_TOKEN = "bearer_token", _("Bearer Token")
    BASIC_AUTH = "basic_auth", _("Basic Auth")


class Integration(TimeStampedModel):
    """
    Integration Model Fields
    """

    organization = models.ForeignKey(
        Organization,
        on_delete=models.CASCADE,
        related_name="integrations",
        related_query_name="integration",
        verbose_name=_("Organization"),
    )
    is_active = models.BooleanField(
        _("Is Active"), default=True, help_text=_("Is the integration active?")
    )
    name = models.CharField(
        _("Name"),
        max_length=100,
        choices=IntegrationChoices.choices,
        unique=True,
        help_text=_("Name of the integration"),
    )
    auth_type = models.CharField(
        _("Auth Type"),
        max_length=100,
        choices=IntegrationAuthTypes.choices,
        help_text=_("Authentication type for the integration"),
        default=IntegrationAuthTypes.NO_AUTH,
    )
    login_url = models.URLField(
        _("Login URL"),
        max_length=255,
        blank=True,
        null=True,
        help_text=_("Login URL for the integration"),
    )
    auth_token = models.CharField(
        _("Auth Token"),
        max_length=255,
        help_text=_("Api or Bearer Token  for the specified integration"),
        null=True,
        blank=True,
    )
    username = models.CharField(
        _("Username"),
        max_length=255,
        help_text=_("Username for the specified integration"),
        null=True,
        blank=True,
    )
    password = models.CharField(
        _("Password"),
        max_length=255,
        help_text=_("Password for the specified integration"),
        null=True,
        blank=True,
    )
    client_id = models.CharField(
        _("Client ID"),
        max_length=255,
        null=True,
        blank=True,
        help_text=_("Client ID for the specified integration"),
    )
    client_secret = models.CharField(
        _("Client Secret"),
        max_length=255,
        null=True,
        blank=True,
        help_text=_("Client Secret for the specified integration"),
    )

    class Meta:
        """
        Metaclass for Integration
        """

        verbose_name = _("Integration")
        verbose_name_plural = _("Integrations")
        ordering: list[str] = ["name"]
        indexes: list[models.Index] = [
            models.Index(fields=["name"]),
        ]

    def __str__(self) -> str:
        """
        Returns:
            (str): String representation of the Integration
        """
        return self.name

    def clean(self) -> None:
        """Clean method to validate the Integration Model

        Returns:
            None

        Raises:
            ValidationError: If the Integration is not active and the auth_type is not NO_AUTH.
        """

        if self.name in [
            IntegrationChoices.GOOGLE_MAPS,
            IntegrationChoices.GOOGLE_PLACES,
        ]:
            if not self.auth_type == IntegrationAuthTypes.API_KEY:
                raise ValidationError(
                    {
                        "auth_type": ValidationError(
                            _("API Key is required for Google Integrations"),
                            code="invalid",
                        )
                    }
                )

        if self.auth_type in [
            IntegrationAuthTypes.BEARER_TOKEN,
            IntegrationAuthTypes.API_KEY,
        ]:
            if not self.auth_token:
                raise ValidationError(
                    {
                        "auth_token": ValidationError(
                            _("Auth Token required for Authentication Type."),
                            code="required",
                        )
                    }
                )

        if self.auth_type == IntegrationAuthTypes.BASIC_AUTH:
            if not self.username or not self.password:
                raise ValidationError(
                    {
                        "username": ValidationError(
                            _(
                                "Username and Password required for Authentication Type."
                            ),
                            code="required",
                        ),
                        "password": ValidationError(
                            _(
                                "Username and Password required for Authentication Type."
                            ),
                            code="required",
                        ),
                    }
                )

        if self.auth_type == IntegrationAuthTypes.NO_AUTH:
            if self.auth_token or self.username or self.password:
                raise ValidationError(
                    {
                        "auth_token": ValidationError(
                            _("Auth Token not required for Authentication Type."),
                            code="invalid",
                        ),
                        "username": ValidationError(
                            _("Username not required for Authentication Type."),
                            code="invalid",
                        ),
                        "password": ValidationError(
                            _("Password not required for Authentication Type."),
                            code="invalid",
                        ),
                    }
                )

    def get_absolute_url(self) -> str:
        """
        Returns:
            (str): Absolute URL for the Integration
        """
        return reverse("integration:detail", kwargs={"pk": self.pk})
