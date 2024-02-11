# --------------------------------------------------------------------------------------------------
#  COPYRIGHT(c) 2024 Trenova                                                                       -
#                                                                                                  -
#  This file is part of Trenova.                                                                   -
#                                                                                                  -
#  The Trenova software is licensed under the Business Source License 1.1. You are granted the right
#  to copy, modify, and redistribute the software, but only for non-production use or with a total -
#  of less than three server instances. Starting from the Change Date (November 16, 2026), the     -
#  software will be made available under version 2 or later of the GNU General Public License.     -
#  If you use the software in violation of this license, your rights under the license will be     -
#  terminated automatically. The software is provided "as is," and the Licensor disclaims all      -
#  warranties and conditions. If you use this license's text or the "Business Source License" name -
#  and trademark, you must comply with the Licensor's covenants, which include specifying the      -
#  Change License as the GPL Version 2.0 or a compatible license, specifying an Additional Use     -
#  Grant, and not modifying the license in any other way.                                          -
# --------------------------------------------------------------------------------------------------

import textwrap
import uuid
from typing import Any, final

from django.conf import settings
from django.core.exceptions import ValidationError
from django.db import models
from django.db.models.functions import Lower
from django.urls import reverse
from django.utils import timezone
from django.utils.translation import gettext_lazy as _
from organization.models import Organization
from utils.models import (
    ChoiceField,
    GenericModel,
    PrimaryStatusChoices,
    RatingMethodChoices,
)

User = settings.AUTH_USER_MODEL


class DispatchControl(GenericModel):
    """Stores dispatch control information for a related :model:organization.Organization.

    The DispatchControl model stores dispatch control information for a related organization. It is used to store
    information such as the record service incident control, grace period, deadhead target,
    driver assign, trailer continuity, distance method, duplicate trailer check, regulatory check,
    prevention of shipments on hold, and the generation of routes.
    """

    @final
    class ServiceIncidentControlChoices(models.TextChoices):
        """
        Service Incident Control Choices
        """

        NEVER = "Never", _("Never")
        PICKUP = "Pickup", _("Pickup")
        DELIVERY = "Delivery", _("Delivery")
        PICKUP_DELIVERY = "Pickup and Delivery", _("Pickup and Delivery")
        ALL_EX_SHIPPER = "All except shipper", _("All except shipper")

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    organization = models.OneToOneField(
        Organization,
        on_delete=models.CASCADE,
        verbose_name=_("Organization"),
        related_name="dispatch_control",
        related_query_name="dispatch_controls",
    )
    record_service_incident = ChoiceField(
        _("Record Service Incident"),
        choices=ServiceIncidentControlChoices.choices,
        default=ServiceIncidentControlChoices.NEVER,
    )
    grace_period = models.PositiveIntegerField(
        _("Grace Period"),
        default=0,
        help_text=_("Grace period for the service incident in minutes."),
    )
    deadhead_target = models.DecimalField(
        _("Deadhead Target"),
        max_digits=5,
        decimal_places=2,
        default=0.00,
        help_text=_("Deadhead Mileage target for the company."),
    )
    enforce_worker_assign = models.BooleanField(
        _("Enforce Worker Assign"),
        default=True,
        help_text=_("Enforce worker assign to shipments for the company."),
    )
    trailer_continuity = models.BooleanField(
        _("Enforce Trailer Continuity"),
        default=False,
        help_text=_("Enforce trailer continuity for the company."),
    )
    dupe_trailer_check = models.BooleanField(
        _("Enforce Duplicate Trailer Check"),
        default=False,
        help_text=_("Enforce duplicate trailer check for the company."),
    )
    maintenance_compliance = models.BooleanField(
        _("Vehicle Maintenance Compliance"),
        default=True,
        help_text=_(
            "Ensures that all vehicles are compliant with maintenance standards before dispatch. Dispatch is "
            "disallowed if the assigned vehicle is not compliant."
        ),
    )
    max_shipment_weight_limit = models.IntegerField(
        _("Max Shipment Weight Limit"),
        default=80000,
        help_text=_(
            "Sets the maximum allowable weight (in pounds) for any shipment. Dispatch is prevented if a shipment "
            "exceeds this limit."
        ),
    )
    regulatory_check = models.BooleanField(
        _("Enforce Regulatory Check"),
        default=False,
        help_text=_("Enforce regulatory check for the company."),
    )
    prev_shipments_on_hold = models.BooleanField(
        _("Prevent Shipments On Hold"),
        default=False,
        help_text=_("Prevent dispatch of shipments on hold for the company."),
    )
    worker_time_away_restriction = models.BooleanField(
        _("Enforce Worker Time Away"),
        default=True,
        help_text=_("Disallow assignments if the worker is on Time Away"),
    )
    tractor_worker_fleet_constraint = models.BooleanField(
        _("Enforce Tractor and Worker Fleet Continuity "),
        default=False,
        help_text=_(
            "Enforce Worker and Tractor must be in the same fleet to be assigned to a dispatch."
        ),
    )

    class Meta:
        """
        Metaclass for DispatchControl
        """

        verbose_name = _("Dispatch Control")
        verbose_name_plural = _("Dispatch Controls")
        ordering = ["organization"]
        db_table = "dispatch_control"
        db_table_comment = (
            "Stores dispatch control information for a related organization."
        )

    def __str__(self) -> str:
        """Dispatch control string representation

        Returns:
            str: Dispatch control string representation
        """
        return textwrap.shorten(
            f"Dispatch Control for {self.organization.name}",
            width=50,
            placeholder="...",
        )

    def get_absolute_url(self) -> str:
        """Dispatch control absolute URL

        Returns:
            str: Dispatch control absolute URL
        """
        return reverse("dispatch-control-detail", kwargs={"pk": self.pk})


class DelayCode(GenericModel):
    """A model to store delay codes for a service incident.

    The DelayCode model stores codes and descriptions for a delay that occurs during a service incident. The fault of
        the delay can be recorded as either the fault of the carrier or driver.

    Attributes:
        code (CharField): The primary key, unique, and four character code for the delay. Help text is "Delay code for
            the service incident."
        description (CharField): A 100-character description for the delay code. Help text is "Description for the
            delay code."
        f_carrier_or_driver (BooleanField): A boolean value indicating if the fault of the delay is the carrier or
            driver. Default value is False.
        Help text is "Fault is carrier or driver."

    Class Attributes:
        Meta (class): A metaclass for the DelayCode model with verbose name "Delay Code" and verbose name plural
            "Delay Codes".
        The ordering is based on the code attribute.

    Methods:
        str(self) -> str:
            Returns the string representation of the DelayCode instance, which is the first 50 characters of the
            code attribute.
        get_absolute_url(self) -> str:
            Returns the URL for the DelayCode instance's detail view.
    """

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    status = ChoiceField(
        _("Status"),
        choices=PrimaryStatusChoices.choices,
        help_text=_("Status of the delay code."),
        default=PrimaryStatusChoices.ACTIVE,
    )
    code = models.CharField(
        _("Delay Code"),
        max_length=4,
        help_text=_("Delay code for the service incident."),
    )
    description = models.CharField(
        _("Description"),
        max_length=100,
        help_text=_("Description for the delay code."),
    )
    f_carrier_or_driver = models.BooleanField(
        _("Fault of Carrier or Driver"),
        default=False,
        help_text=_("Fault is carrier or driver."),
    )

    class Meta:
        """
        Metaclass for DelayCode
        """

        verbose_name = _("Delay Code")
        verbose_name_plural = _("Delay Codes")
        ordering = ["code"]
        db_table = "delay_code"
        db_table_comment = "Stores delay codes for a service incident."
        constraints = [
            models.UniqueConstraint(
                Lower("code"),
                "organization",
                name="unique_delay_code_organization",
            )
        ]

    def __str__(self) -> str:
        """Delay code string representation

        Returns:
            str: Delay code string representation
        """
        return textwrap.shorten(
            f"{self.code} - {self.description}", width=50, placeholder="..."
        )

    def get_absolute_url(self) -> str:
        """Delay code absolute URL

        Returns:
            str: Delay code absolute URL
        """
        return reverse("delay-codes-detail", kwargs={"pk": self.pk})


class FleetCode(GenericModel):
    """Model for storing fleet codes for service incidents.

    A FleetCode instance represents a code used to identify a fleet of vehicles for service incidents.
    This allows for tracking and reporting on specific fleets of vehicles, including their revenue goals,
    deadhead goals, and mileage goals.
    """

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    status = ChoiceField(
        _("Status"),
        choices=PrimaryStatusChoices.choices,
        help_text=_("Status of the general ledger account."),
        default=PrimaryStatusChoices.ACTIVE,
    )
    code = models.CharField(
        _("Fleet Code"),
        max_length=10,
        help_text=_("Fleet code for the service incident."),
    )
    description = models.CharField(
        _("Description"),
        max_length=100,
        help_text=_("Description for the fleet code."),
    )
    revenue_goal = models.DecimalField(
        _("Revenue Goal"),
        max_digits=10,
        decimal_places=2,
        blank=True,
        null=True,
        help_text=_("Revenue goal for the fleet code."),
    )
    deadhead_goal = models.DecimalField(
        _("Deadhead Goal"),
        max_digits=10,
        decimal_places=2,
        blank=True,
        null=True,
        help_text=_("Deadhead goal for the fleet code."),
    )
    mileage_goal = models.DecimalField(
        _("Mileage Goal"),
        max_digits=10,
        decimal_places=2,
        blank=True,
        null=True,
        help_text=_("Mileage goal for the fleet code."),
    )
    manager = models.ForeignKey(
        User,
        on_delete=models.RESTRICT,
        related_name="fleet_code_manager",
        help_text=_("Manager for the fleet code."),
        limit_choices_to={"is_active": True},
        null=True,
        blank=True,
    )

    class Meta:
        """
        Metaclass for FleetCode model.
        """

        verbose_name = _("Fleet Code")
        verbose_name_plural = _("Fleet Codes")
        db_table = "fleet_code"
        constraints = [
            models.UniqueConstraint(
                Lower("code"),
                "organization",
                name="unique_fleet_code_organization",
            )
        ]

    def __str__(self) -> str:
        """
        Return a string representation of the FleetCode instance.

        Returns:
            str: A string representation of the FleetCode instance, wrapped to a maximum of 4 characters.
        """
        return textwrap.shorten(
            f"{self.code} - {self.description}", width=50, placeholder="..."
        )

    def get_absolute_url(self) -> str:
        """
        Return the absolute URL for the FleetCode instance's detail view.

        Returns:
            str: The absolute URL for the FleetCode instance's detail view.
        """
        return reverse("fleet-codes-detail", kwargs={"pk": self.pk})


class CommentType(GenericModel):
    """Model for storing different types of comments.

    A CommentType instance represents a type of comment that can be associated with a comment.
    This allows for categorization and grouping of comments based on their type.

    Attributes:
        id (UUIDField): Primary key and default value is a randomly generated UUID.
            Editable and unique.
        name (CharField): Name of the comment type.
            Has a max length of 255 characters and help text of "Comment type name".
        description (TextField): Description of the comment type.
            Has a max length of 255 characters and help text of "Comment type description".

    Methods:
        __str__(self) -> str:
            Returns a string representation of the comment type, wrapped to a maximum of 50 characters.

        get_absolute_url(self) -> str:
            Returns the URL for this object's detail view.
    """

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    status = ChoiceField(
        _("Status"),
        choices=PrimaryStatusChoices.choices,
        help_text=_("Status of the comment type."),
        default=PrimaryStatusChoices.ACTIVE,
    )
    severity = models.CharField(
        _("Severity"),
        max_length=50,
        choices=[("H", "High"), ("M", "Medium"), ("L", "Low")],
        default="L",
        help_text=_("Priority or severity level of the comment type."),
    )
    name = models.CharField(
        _("Name"),
        max_length=10,
        help_text=_("Comment type name"),
    )
    description = models.CharField(
        _("Description"),
        max_length=100,
        help_text=_("Comment type description"),
    )

    class Meta:
        """
        Metaclass for the Comment Type model.

        The Meta class defines some options for the Comment Type model.
        """

        verbose_name = _("Comment Type")
        verbose_name_plural = _("Comment Types")
        ordering = ["organization"]
        db_table = "comment_type"
        db_table_comment = "Stores different types of comments."
        constraints = [
            models.UniqueConstraint(
                Lower("name"),
                "organization",
                name="unique_comment_type_name_organization",
            )
        ]

    def __str__(self) -> str:
        """
        Return a string representation of the CommentType instance.

        Returns:
            str: A string representation of the CommentType instance, wrapped to a maximum of 50 characters.
        """
        return textwrap.shorten(
            f"{self.name} - {self.description}", 50, placeholder="..."
        )

    def get_absolute_url(self) -> str:
        """
        Return the absolute URL for the CommentType instance's detail view.

        Returns:
            str: The absolute URL for the CommentType instance's detail view.
        """
        return reverse("comment-types-detail", kwargs={"pk": self.pk})


class Rate(GenericModel):
    """Django model representing a Rate. This model stores information about the rates for a related customer,
    commodity, shipment type, and equipment type.

    Attributes:
        id (UUIDField): Primary key and default value is a randomly generated UUID. Not editable and unique.
        rate_number (CharField): A unique identifier for the rate, with max length of 10 characters.
        customer (ForeignKey): A foreign key to the customer model, with a related name of "rates".
        effective_date (DateField): The date when the rate becomes effective.
        expiration_date (DateField): The date when the rate expires.
        commodity (ForeignKey): A foreign key to the commodity model, with a related name of "rates".
        shipment_type (ForeignKey): A foreign key to the shipment type model, with a related name of "rates".
        equipment_type (ForeignKey): A foreign key to the equipment type model, with a related name of "rates".
        comments (TextField): Comments about the rate.

    Methods:
        str(self) -> str:
            Returns the string representation of the Rate instance, which is the first 10 characters of the rate_number
            field.

        get_absolute_url(self) -> str:
            Returns the absolute URL for the detail view of this Rate instance.

        set_rate_number_before_create(self) -> None:
            Sets the rate_number field with the result of the generate_rate_number method before the instance is
            created.

        generate_rate_number() -> str:
            Returns a new rate number that has not been used before, generated by incrementing the count of all previous
            Rate instances.

        Class Meta:
            verbose_name (str): "Rate".
            verbose_name_plural (str): "Rates".
            ordering (list): Orders the Rate instances by the rate_number field.
    """

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    status = ChoiceField(
        _("Status"),
        choices=PrimaryStatusChoices.choices,
        help_text=_("Status of the rate."),
        default=PrimaryStatusChoices.ACTIVE,
    )
    rate_number = models.CharField(
        _("Rate Number"),
        max_length=6,
        editable=False,
        help_text=_("Rate Number for Rate"),
    )
    customer = models.ForeignKey(
        "customer.Customer",
        on_delete=models.SET_NULL,
        verbose_name=_("Customer"),
        related_name="rates",
        null=True,
        blank=True,
        help_text=_("Customer for Rate"),
    )
    effective_date = models.DateField(
        _("Effective Date"),
        help_text=_("Effective Date for Rate"),
        default=timezone.now,
    )
    expiration_date = models.DateField(
        _("Expiration Date"),
        help_text=_("Expiration Date for Rate"),
        default=timezone.now,
    )
    commodity = models.ForeignKey(
        "commodities.Commodity",
        on_delete=models.SET_NULL,
        verbose_name=_("Commodity"),
        related_name="rates",
        null=True,
        blank=True,
        help_text=_("Commodity for Rate"),
    )
    shipment_type = models.ForeignKey(
        "shipment.ShipmentType",
        on_delete=models.SET_NULL,
        verbose_name=_("shipment type"),
        related_name="rates",
        null=True,
        blank=True,
    )
    origin_location = models.ForeignKey(
        "location.Location",
        on_delete=models.PROTECT,
        verbose_name=_("Origin Location"),
        related_name="origin_rates",
        help_text=_("Origin Location for Rate"),
        blank=True,
        null=True,
    )
    destination_location = models.ForeignKey(
        "location.Location",
        on_delete=models.PROTECT,
        verbose_name=_("Destination Location"),
        related_name="destination_rates",
        help_text=_("Destination Location for Rate"),
        blank=True,
        null=True,
    )
    rate_method = ChoiceField(
        _("Rate Method"),
        choices=RatingMethodChoices.choices,
        default=RatingMethodChoices.FLAT,
        help_text=_("Rate Method for Rate"),
    )
    rate_amount = models.DecimalField(
        _("Rate Amount"),
        max_digits=19,
        decimal_places=4,
        default=0,
        help_text=_("Rate Amount for Rate"),
    )
    distance_override = models.PositiveIntegerField(
        _("Distance Override"),
        help_text=_("Distance Override for Rate"),
        blank=True,
        null=True,
    )
    comments = models.TextField(
        _("Comments"),
        max_length=255,
        blank=True,
        help_text=_("Comments for Rate"),
    )

    class Meta:
        """
        Metaclass for the Rate model.

        The Meta class defines some options for the Rate model.
        """

        verbose_name = _("Rate")
        verbose_name_plural = _("Rates")
        ordering = ["rate_number"]
        db_table = "rate"
        constraints = [
            models.UniqueConstraint(
                Lower("rate_number"),
                "organization",
                name="unique_rate_number_organization",
            )
        ]

    def __str__(self) -> str:
        """
        Return the string representation of a Rate instance.

        Returns:
            str: The first 10 characters of the rate_number field.
        """
        return textwrap.shorten(f"{self.rate_number}", width=10, placeholder="...")

    def get_absolute_url(self) -> str:
        """
        Return the absolute URL for the detail view of a Rate instance.

        Returns:
            str: The absolute URL for the detail view of the Rate instance.
        """
        return reverse("rates-detail", kwargs={"pk": self.pk})

    def save(self, *args: Any, **kwargs: Any) -> None:
        """Save method for Rate model.

        Args:
            *args(Any): Variable length argument list.
            **kwargs(Any): Arbitrary keyword arguments.

        Returns:
            None: This function does not return anything.
        """

        if not self.rate_number:
            self.rate_number = self.generate_rate_number()

        super().save(*args, **kwargs)

    def clean(self) -> None:
        """
        Clean the Rate instance.

        Returns:
            None: None
        """
        if self.expiration_date < self.effective_date:
            raise ValidationError(
                {
                    "expiration_date": _(
                        "Expiration Date must be after Effective Date. Please correct and try again."
                    )
                }
            )

    @classmethod
    def generate_rate_number(cls) -> str:
        """
        Generate a unique rate number for a Rate instance.

        This method generates a unique rate number by finding the highest rate number and
        incrementing it by 1.

        Returns:
            str: A unique rate number for a Rate instance, formatted as "R{count:05d}".
        """
        if code := cls.objects.order_by("-rate_number").first():
            return f"R{int(code.rate_number[1:]) + 1:05d}"
        return "R00001"


class RateBillingTable(GenericModel):
    """Django model representing a RateBillingTable. This model stores Billing Table information for a
    related :model:`rates.Rate`.

    Attributes:
        id (UUIDField): The primary key for the rate billing table instance.
        rate (ForeignKey): The rate associated with the rate billing table instance.
        accessorial_charge (ForeignKey): The charge code associated with the rate billing table instance.
        description (CharField): The description for the rate billing table instance.
        unit (PositiveIntegerField): The number of units for the rate billing table instance.
        charge_amount (MoneyField): The charge amount for the rate billing table instance.
        sub_total (MoneyField): The sub_total for the rate billing table instance.

    Methods:
        meta: Return the meta options for the RateBillingTable model.
        get_absolute_url: Return the absolute URL for the detail view of a RateBillingTable instance.
        __str__: Return the string representation of a RateBillingTable instance.
    """

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    rate = models.ForeignKey(
        Rate,
        on_delete=models.CASCADE,  # if rate is deleted, delete the rate billing table
        related_name="rate_billing_tables",
        verbose_name=_("Rate"),
        help_text=_("Rate for Rate Billing Table"),
    )
    accessorial_charge = models.ForeignKey(
        "billing.AccessorialCharge",
        on_delete=models.RESTRICT,  # if accessorial charge is deleted, do not delete the rate billing table
        related_name="rate_billing_tables",
        verbose_name=_("Charge Code"),
        help_text=_("Charge Code for Rate Billing Table"),
    )
    description = models.CharField(
        _("Description"),
        max_length=100,
        help_text=_("Description for Rate Billing Table"),
        blank=True,
    )
    unit = models.PositiveIntegerField(
        _("Unit"),
        help_text=_("Unit for Rate Billing Table"),
    )
    charge_amount = models.DecimalField(
        _("Charge Amount"),
        max_digits=19,
        decimal_places=4,
        default=0,
        help_text=_("Charge Amount for Rate Billing Table"),
    )
    sub_total = models.DecimalField(
        _("Total"),
        max_digits=19,
        decimal_places=4,
        default=0,
        help_text=_("Total for Rate Billing Table"),
    )

    class Meta:
        """
        Metaclass for the RateBillingTable model.
        """

        verbose_name = _("Rate Billing Table")
        verbose_name_plural = _("Rate Billing Tables")
        ordering = ["rate", "accessorial_charge"]
        db_table = "rate_billing_table"

    def __str__(self) -> str:
        """String representation of a RateBillingTable instance is the value of its description field.

        Returns:
            str: The description field of the RateBillingTable instance.
        """
        return textwrap.shorten(
            f"{self.rate}, {self.accessorial_charge}", width=50, placeholder="..."
        )

    def clean(self) -> None:
        """Clean the RateBillingTable instance.

        Returns:
            None: This function does not return anything.
        """
        if self.accessorial_charge.status == PrimaryStatusChoices.INACTIVE:
            raise ValidationError(
                {
                    "accessorial_charge": _(
                        "The accessorial charge is inactive. Please select an active accessorial charge and try again."
                    ),
                },
                code="invalid",
            )

        super().clean()

    def save(self, *args: Any, **kwargs: Any) -> None:
        """
        Save the RateBillingTable instance.

        This method overrides the default save method for the RateBillingTable model. It calculates
        the sub_total for the RateBillingTable instance.

        Returns:
            None: None
        """
        if not self.sub_total:
            self.sub_total = self.unit * self.charge_amount

        if not self.charge_amount:
            self.charge_amount = self.accessorial_charge.charge_amount

        self.sub_total = self.charge_amount * self.unit
        super().save(*args, **kwargs)

    def get_absolute_url(self) -> str:
        """
        Return the absolute URL for the detail view of a RateBillingTable instance.

        Returns:
            str: The absolute URL for the detail view of the RateBillingTable instance.
        """
        return reverse("rate-billing-tables-detail", kwargs={"pk": self.pk})


class FeasibilityToolControl(GenericModel):
    @final
    class OperatorChoices(models.TextChoices):
        EQUALS = "eq", _("Equals")
        NOT_EQUALS = "ne", _("Not Equals")
        GREATER_THAN = "gt", _("Greater Than")
        GREATER_THAN_OR_EQUAL_TO = "gte", _("Greater Than or Equal To")
        LESS_THAN = "lt", _("Less Than")
        LESS_THAN_OR_EQUAL_TO = "lte", _("Less Than or Equal To")

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    organization = models.OneToOneField(
        "organization.Organization",
        on_delete=models.CASCADE,
        verbose_name=_("Organization"),
        related_name="feasibility_tool_control",
    )
    mpw_operator = ChoiceField(
        _("MPW Operator"),
        choices=OperatorChoices.choices,
        default=OperatorChoices.EQUALS,
        help_text=_("Miles Per Week Operator (e.g. Equals, Not Equals, etc.)"),
    )
    mpw_criteria = models.FloatField(
        _("MPW Criteria"),
        default=0,
        help_text=_("Miles Per Week criteria of the feasibility control. (e.g. 0.5)"),
    )
    mpd_operator = ChoiceField(
        _("MPD Operator"),
        choices=OperatorChoices.choices,
        default=OperatorChoices.EQUALS,
        help_text=_("Miles Per Day Operator (e.g. Equals, Not Equals, etc.)"),
    )
    mpd_criteria = models.FloatField(
        _("MPD Criteria"),
        default=0,
        help_text=_("Miles Per Day criteria of the feasibility control. (e.g. 0.5)"),
    )
    mpg_operator = ChoiceField(
        _("MPG Operator"),
        choices=OperatorChoices.choices,
        default=OperatorChoices.EQUALS,
        help_text=_("Miles Per Gallon Operator (e.g. Equals, Not Equals, etc.)"),
    )
    mpg_criteria = models.FloatField(
        _("MPG Criteria"),
        default=0,
        help_text=_("Miles Per Gallon criteria of the feasibility control. (e.g. 0.5)"),
    )
    otp_operator = ChoiceField(
        _("OTP Operator"),
        choices=OperatorChoices.choices,
        default=OperatorChoices.EQUALS,
        help_text=_("On Time Performance Operator (e.g. Equals, Not Equals, etc.)"),
    )
    otp_criteria = models.FloatField(
        _("OTP Criteria"),
        default=0,
        help_text=_(
            "On Time Performance criteria of the feasibility control. (e.g. 0.5)"
        ),
    )

    class Meta:
        """
        Metaclass for the FeasibilityToolControl model.
        """

        verbose_name = _("Feasibility Tool Control")
        verbose_name_plural = _("Feasibility Tool Controls")
        db_table = "feasibility_tool_control"

    def __str__(self) -> str:
        """String representation of a FeasibilityToolControl instance is the value of its organization field.

        Returns:
            str: The organization field of the FeasibilityToolControl instance.
        """
        return textwrap.shorten(
            f"Feasibility Control for {self.organization.name}",
            width=30,
            placeholder="...",
        )

    def get_absolute_url(self) -> str:
        """Absolute URL for the detail view of a FeasibilityToolControl instance.

        Returns:
            str: The absolute URL for the detail view of the FeasibilityToolControl instance.
        """
        return reverse("feasibility-tool-control-detail", kwargs={"pk": self.pk})
