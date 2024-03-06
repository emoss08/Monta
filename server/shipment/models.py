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

from __future__ import annotations

import textwrap
import uuid
from datetime import datetime
from typing import Any, final

from auditlog.models import AuditlogHistoryField
from django.conf import settings
from django.core.exceptions import ValidationError
from django.core.validators import FileExtensionValidator
from django.db import models
from django.db.models.functions import Lower
from django.urls import reverse
from django.utils.translation import gettext_lazy as _

from commodities.models import HazardousClassChoices
from location.models import Location
from utils.models import (
    ChoiceField,
    GenericModel,
    PrimaryStatusChoices,
    RatingMethodChoices,
    StatusChoices,
)

User = settings.AUTH_USER_MODEL


def shipment_documentation_upload_to(
    instance: ShipmentDocumentation, filename: str
) -> str:
    """Uploads the Shipment Documentation to the correct location.

    Args:
        instance (ShipmentDocumentation): The Shipment Documentation instance.
        filename (str): The filename of the Shipment Documentation.

    Returns:
        str: The path to the Shipment Documentation.
    """
    return f"{instance.organization_id}/shipment/docs/{instance.shipment_id}/{filename}"


@final
class EntryMethodChoices(models.TextChoices):
    """
    Entry method choices used for shipments.
    """

    MANUAL = "MANUAL", _("Manual")
    EDI = "EDI", _("EDI")
    API = "API", _("API")


class ShipmentControl(GenericModel):
    """
    Stores the shipment control information for a related :model:`organization.Organization`.
    """

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
        related_name="shipment_control",
        related_query_name="shipment_controls",
    )
    auto_rate_shipment = models.BooleanField(
        _("Auto Rate Shipments"),
        default=True,
        help_text=_("Auto rate Shipments"),
    )
    calculate_distance = models.BooleanField(
        _("Calculate Distance"),
        default=True,
        help_text=_("Automatically Calculate distance for the shipment."),
    )
    enforce_rev_code = models.BooleanField(
        _("Enforce Rev Code"),
        default=False,
        help_text=_("Enforce rev code when entering a shipment."),
    )
    enforce_voided_comm = models.BooleanField(
        _("Enforce Voided Comm"),
        default=False,
        help_text=_("Enforce comment when voiding a shipment."),
    )
    generate_routes = models.BooleanField(
        _("Generate Routes"),
        default=False,
        help_text=_("Automatically generate routing information for the shipment."),
    )
    enforce_commodity = models.BooleanField(
        _("Enforce Commodity Code"),
        default=False,
        help_text=_("Enforce the commodity input on the entry of an shipment."),
    )
    auto_sequence_stops = models.BooleanField(
        _("Auto Sequence Stops"),
        default=True,
        help_text=_("Auto Sequence stops for the shipment and movements."),
    )
    auto_shipment_total = models.BooleanField(
        _("Auto shipment Total"),
        default=True,
        help_text=_("Automate the shipment total amount calculation."),
    )
    enforce_origin_destination = models.BooleanField(
        _("Compare Origin Destination"),
        default=False,
        help_text=_(
            "Compare and validate that origin and destination are not the same."
        ),
    )
    check_for_duplicate_bol = models.BooleanField(
        _("Check for Duplicate BOL"),
        default=False,
        help_text=_("Check for duplicate BOL numbers when entering an shipment."),
    )
    send_placard_info = models.BooleanField(
        _("Send Placard Info"),
        default=False,
        help_text=_(
            "If checked, the system will apply a comment to the shipment with the placard info."
        ),
    )
    enforce_hazmat_seg_rules = models.BooleanField(
        _("Enforce Hazardous Material Segregation Rules"),
        default=True,
        help_text=_(
            "If checked, the system will enforce hazardous material segregation rules. When entering an order."
        ),
    )

    class Meta:
        """
        Metaclass for ShipmentControl
        """

        verbose_name = _("Shipment Control")
        verbose_name_plural = _("Shipment Controls")
        db_table = "shipment_control"

    def __str__(self) -> str:
        """Shipment control string representation

        Returns:
            str: Shipment control string representation
        """
        return textwrap.shorten(f"{self.organization}", 50, placeholder="...")

    def get_absolute_url(self) -> str:
        """Shipment Control absolute url

        Returns:
            Absolute url for the Shipment Control object. For example,
            `/shipment_control/edd1e612-cdd4-43d9-b3f3-bc099872088b/`
        """
        return reverse("shipment_control:detail", kwargs={"pk": self.pk})


class ShipmentType(GenericModel):
    """
    Stores the shipment type information for a related :model:`organization.Organization`.
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
        help_text=_("Status of the Shipment Type."),
        default=PrimaryStatusChoices.ACTIVE,
    )
    code = models.CharField(
        _("Code"),
        max_length=10,
        help_text=_("Code of the Shipment Type"),
    )
    description = models.TextField(
        _("Description"),
        blank=True,
        help_text=_("Description of the Shipment Type"),
    )

    class Meta:
        """
        Metaclass for ShipmentType model
        """

        verbose_name = _("Shipment Type")
        verbose_name_plural = _("Shipment Types")
        db_table = "shipment_type"
        db_table_comment = (
            "Stores the shipment type information for a related organization."
        )
        constraints = [
            models.UniqueConstraint(
                Lower("code"),
                "organization",
                name="unique_shipment_type_code",
            )
        ]

    def __str__(self) -> str:
        """Shipment Type String Representation

        Returns:
            str: Shipment type Code
        """
        return textwrap.shorten(
            f"{self.code} - {self.description}", 50, placeholder="..."
        )

    def get_absolute_url(self) -> str:
        """Shipment Type Absolute URL

        Returns:
            str: Shipment Type Absolute URL
        """
        return reverse("shipment-types-detail", kwargs={"pk": self.pk})


class ServiceType(GenericModel):
    """
    Stores the service type information for a related :model:`organization.Organization`.
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
        help_text=_("Status of the service type."),
        default=PrimaryStatusChoices.ACTIVE,
    )
    code = models.CharField(
        _("Code"),
        max_length=4,
        help_text=_("Code of the Service Type"),
    )
    description = models.TextField(
        _("Description"),
        blank=True,
        help_text=_("Description of the Service Type"),
    )

    class Meta:
        """
        Metaclass for ServiceType model
        """

        verbose_name = _("Service Type")
        verbose_name_plural = _("Service Types")
        db_table = "service_type"
        db_table_comment = (
            "Stores the service type information for a related organization."
        )
        constraints = [
            models.UniqueConstraint(
                Lower("code"),
                "organization",
                name="unique_service_type_code",
            )
        ]

    def __str__(self) -> str:
        """Service Type String Representation

        Returns:
            str: service type Name
        """
        return textwrap.shorten(
            f"{self.code} - {self.description}", 50, placeholder="..."
        )

    def get_absolute_url(self) -> str:
        """Service Type Absolute URL

        Returns:
            str: Service Type Absolute URL
        """
        return reverse("service-types-detail", kwargs={"pk": self.pk})


class Shipment(GenericModel):
    """
    Stores shipment information related to a :model:`organization.Organization`.
    """

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    # General Information
    pro_number = models.CharField(
        _("Pro Number"),
        max_length=13,
        editable=False,
        help_text=_("Pro Number of the Shipment"),
    )
    shipment_type = models.ForeignKey(
        ShipmentType,
        on_delete=models.PROTECT,
        verbose_name=_("Shipment type"),
        related_name="shipments",
        related_query_name="shipment",
        help_text=_("Shipment type of the Shipment"),
    )
    service_type = models.ForeignKey(
        ServiceType,
        on_delete=models.PROTECT,
        verbose_name=_("Service type"),
        related_name="shipments",
        related_query_name="shipment",
        blank=True,
        null=True,
        help_text=_("Service type of the Shipment"),
    )
    status = ChoiceField(
        _("Status"),
        choices=StatusChoices.choices,
        default=StatusChoices.NEW,
    )
    revenue_code = models.ForeignKey(
        "accounting.RevenueCode",
        on_delete=models.PROTECT,
        related_name="shipments",
        related_query_name="shipment",
        verbose_name=_("Revenue Code"),
        help_text=_("Revenue Code of the Shipment"),
        blank=True,
        null=True,
    )
    origin_location = models.ForeignKey(
        "location.Location",
        on_delete=models.PROTECT,
        related_name="origin_shipment",
        related_query_name="origin_shipments",
        verbose_name=_("Origin Location"),
        help_text=_("Origin Location of the Shipment"),
        blank=True,
        null=True,
    )
    origin_address = models.CharField(
        _("Origin Address"),
        max_length=255,
        blank=True,
        help_text=_("Origin Address of the Shipment"),
    )
    origin_appointment_window_start = models.DateTimeField(
        _("Origin Appointment"),
        help_text=_(
            "The time the equipment is expected to arrive at the origin/pickup."
        ),
    )
    origin_appointment_window_end = models.DateTimeField(
        _("Origin Appointment Window End"),
        help_text=_(
            "The time the equipment is expected to arrive at the origin/pickup."
        ),
    )
    destination_location = models.ForeignKey(
        "location.Location",
        on_delete=models.PROTECT,
        related_name="destination_shipment",
        related_query_name="destination_shipments",
        verbose_name=_("Destination Location"),
        blank=True,
        null=True,
    )
    destination_address = models.CharField(
        _("Destination Address"),
        max_length=255,
        blank=True,
    )
    destination_appointment_window_start = models.DateTimeField(
        _("Destination Appointment Time"),
        help_text=_(
            "The time the equipment is expected to arrive at the destination/drop-off."
        ),
    )
    destination_appointment_window_end = models.DateTimeField(
        _("Destination Appointment Window End"),
        help_text=_(
            "The time the equipment is expected to arrive at the destination/drop-off."
        ),
    )

    # Billing Information for the Shipment
    rating_units = models.PositiveIntegerField(
        _("Rating Units"),
        default=1,
        help_text=_("Number of units to be rated."),
    )
    rate = models.ForeignKey(
        "dispatch.Rate",
        on_delete=models.RESTRICT,
        related_name="shipments",
        related_query_name="shipment",
        verbose_name=_("Rate"),
        help_text=_("Associated Rate to the shipment."),
        blank=True,
        null=True,
    )
    mileage = models.FloatField(
        _("Total Mileage"),
        default=0,
        help_text=_("Total Mileage"),
        blank=True,
        null=True,
    )
    other_charge_amount = models.DecimalField(
        _("Additional Charge Amount"),
        max_digits=19,
        decimal_places=4,
        default=0,
        help_text=_("Additional Charge Amount"),
        null=True,
        blank=True,
    )
    freight_charge_amount = models.DecimalField(
        _("Freight Charge Amount"),
        max_digits=19,
        decimal_places=4,
        default=0,
        help_text=_("Freight Charge Amount"),
        blank=True,
        null=True,
    )
    rate_method = ChoiceField(
        _("Rating Method"),
        blank=True,
        choices=RatingMethodChoices.choices,
        default=RatingMethodChoices.FLAT,
        help_text=_("Rating Method"),
    )
    customer = models.ForeignKey(
        "customer.Customer",
        on_delete=models.PROTECT,
        related_name="shipments",
        related_query_name="shipment",
        verbose_name=_("Customer"),
        help_text=_("Customer of the Shipment"),
    )
    pieces = models.PositiveIntegerField(
        _("Pieces"),
        help_text=_("Total Piece Count of the Shipment"),
        null=True,
        blank=True,
    )
    weight = models.DecimalField(
        _("Weight"),
        max_digits=10,
        decimal_places=2,
        help_text=_("Total Weight of the Shipment"),
        blank=True,
        null=True,
    )
    ready_to_bill = models.BooleanField(
        _("Ready to Bill"),
        default=False,
        help_text=_("Ready to Bill"),
    )
    bill_date = models.DateField(
        _("Billed Date"),
        null=True,
        blank=True,
        help_text=_("Date shipment was billed to the Customer."),
    )
    ship_date = models.DateField(
        _("Ship Date"),
        null=True,
        blank=True,
        help_text=_("Date shipment was shipped."),
    )
    billed = models.BooleanField(
        _("Billed"),
        default=False,
        help_text=_("Billed"),
    )
    transferred_to_billing = models.BooleanField(
        _("Transferred to Billing"),
        default=False,
        help_text=_("Transferred to Billing"),
    )
    billing_transfer_date = models.DateTimeField(
        _("Billing Transfer Date"),
        null=True,
        blank=True,
        help_text=_("Billing Transfer Date"),
    )
    sub_total = models.DecimalField(
        _("Sub Total Amount"),
        max_digits=19,
        decimal_places=4,
        default=0,
        help_text=_("Sub Total Amount"),
    )

    # Dispatch Information
    trailer_type = models.ForeignKey(
        "equipment.EquipmentType",
        on_delete=models.PROTECT,
        related_name="shipments",
        related_query_name="shipment",
        verbose_name=_("Trailer Type"),
        help_text=_("Type of trailer for the shipment."),
    )
    tractor_type = models.ForeignKey(
        "equipment.EquipmentType",
        on_delete=models.PROTECT,
        null=True,
        blank=True,
        related_name="tractor_type_shipments",
        related_query_name="tractor_type_shipment",
        verbose_name=_("Tractor Type"),
        help_text=_("Type of tractor for the shipment."),
    )
    temperature_min = models.PositiveIntegerField(
        _("Minimum Temperature"),
        null=True,
        blank=True,
        help_text=_("Minimum Temperature"),
    )
    temperature_max = models.PositiveIntegerField(
        _("Maximum Temperature"),
        null=True,
        blank=True,
        help_text=_("Maximum Temperature"),
    )
    bol_number = models.CharField(
        _("BOL Number"),
        max_length=255,
        help_text=_("BOL Number"),
    )
    consignee_ref_number = models.CharField(
        _("Consignee Reference Number"),
        max_length=255,
        blank=True,
        help_text=_("Consignee Reference Number"),
    )
    comment = models.CharField(
        _("Comment"),
        max_length=100,
        blank=True,
        help_text=_("Planning Comment"),
    )
    voided_comm = models.CharField(
        _("Voided Comment"),
        max_length=100,
        blank=True,
        help_text=_("Voided Comment"),
    )
    auto_rate = models.BooleanField(
        _("Auto Rate"),
        default=True,
        help_text=_("Determines whether shipment will be auto-rated by entered rate."),
    )
    current_suffix = models.CharField(
        _("Current Suffix"),
        max_length=2,
        default="",
        help_text=_("Current suffix for shipment in the BillingQueue."),
        blank=True,
    )
    formula_template = models.ForeignKey(
        "FormulaTemplate",
        on_delete=models.RESTRICT,
        related_name="shipments",
        null=True,
        blank=True,
        help_text=_("Selected formula template for this shipment."),
    )
    entry_method = ChoiceField(
        _("Entry Method"),
        default=EntryMethodChoices.MANUAL,
        choices=EntryMethodChoices.choices,
        help_text=_("Method of entry for this shipment."),
        editable=False,
    )
    entered_by = models.ForeignKey(
        User,
        on_delete=models.PROTECT,
        related_name="shipments",
        related_query_name="shipment",
        verbose_name=_("User"),
        help_text=_("Shipment entered by User"),
    )
    is_hazmat = models.BooleanField(
        _("Is Hazmat Shipment"),
        default=False,
        help_text=_("Shipment contains hazardous materials."),
    )

    class Meta:
        """
        Metaclass for the Shipment model
        """

        verbose_name = _("Shipment")
        verbose_name_plural = _("Shipments")
        db_table = "shipment"
        db_table_comment = (
            "Stores shipment information related to a related organization."
        )
        constraints = [
            models.UniqueConstraint(
                Lower("pro_number"),
                "organization",
                name="unique_shipment_number_per_organization",
            )
        ]
        indexes = [
            models.Index(fields=["status"], name="shipment_status_idx"),
            models.Index(
                fields=["bill_date", "organization"], name="bill_date_org_idx"
            ),
            models.Index(
                fields=["ship_date", "organization"], name="ship_date_org_idx"
            ),
            models.Index(
                fields=["bol_number", "organization"], name="bol_number_org_idx"
            ),
        ]

    def __str__(self) -> str:
        """String representation of the Shipment

        Returns:
            str: String representation of the Shipment
        """
        return textwrap.shorten(
            f"{self.pro_number} - {self.customer}", 50, placeholder="..."
        )

    def save(self, *args: Any, **kwargs: Any) -> None:
        """Overrides the default Django save method to provide custom save behavior for the Shipment model.

        Args:
            *args (Any): Variable length argument list.
            **kwargs (Any): Arbitrary keyword arguments.

        Returns:
            None: This function does not return anything.
        """
        from dispatch.services import transfer_rate_details
        from shipment.services import calculate_total, handle_voided_shipment
        from stops.selectors import (
            get_total_piece_count_by_shipment,
            get_total_weight_by_shipment,
        )

        if not self.pro_number:
            self.pro_number = self.generate_pro_number()

        if self.auto_rate:
            transfer_rate_details(shipment=self)

        if (
            self.status == StatusChoices.COMPLETED
            and not self.pieces
            and not self.weight
        ):
            self.weight = get_total_weight_by_shipment(shipment=self)
            self.pieces = get_total_piece_count_by_shipment(shipment=self)

        if self.organization.shipment_control.auto_shipment_total:
            self.sub_total = calculate_total(shipment=self)

        if self.origin_location and not self.origin_address:
            self.origin_address = self.origin_location.get_address_combination

        if self.destination_location and not self.destination_address:
            self.destination_address = self.destination_location.get_address_combination

        if self.status == StatusChoices.VOIDED:
            handle_voided_shipment(shipment=self)

        self.calculate_other_charge_amount()

        self.set_shipment_mileage_and_create_route()

        super().save(*args, **kwargs)

    def get_absolute_url(self) -> str:
        """Get the absolute url for the Shipment

        Returns:
            str: Absolute url for the Shipment
        """
        return reverse("shipment-detail", kwargs={"pk": self.pk})

    def clean(self) -> None:
        """Shipment clean method

        Returns:
            None: This function does not return anything.

        Raises:
            ValidationError: If the Shipment is not valid
        """
        from shipment.validation import ShipmentValidator

        ShipmentValidator(shipment=self)

        # The validate_delivery_slot method will now raise an error directly if the slots don't match.
        if self.origin_location:
            self.validate_delivery_slot(
                self.origin_appointment_window_start,
                self.origin_appointment_window_end,
                self.origin_location,
            )

        if self.destination_location:
            self.validate_delivery_slot(
                self.destination_appointment_window_start,
                self.destination_appointment_window_end,
                self.destination_location,
            )

        super().clean()

    def set_shipment_mileage_and_create_route(self) -> None:
        """Set the mileage for a shipment and create a route.

        This function is called as a signal when a shipment model instance is saved.
        If the shipment has an origin and destination location, it sets the mileage
        for the shipment and creates a route using the generate_route().

        Returns:
            None: This function does not return anything.
        """
        from route.services import get_shipment_mileage

        if (
            self.origin_location
            and self.destination_location
            and self.origin_location.is_geocoded
            and self.destination_location.is_geocoded
        ):
            self.mileage = get_shipment_mileage(shipment=self)

    def generate_pro_number(self) -> str:
        """Generates a unique pro number for the shipment.

        Returns:
            str: The generated pro number.
        """
        today = datetime.now().strftime("%y%m%d")
        count_for_today = (
            self.__class__.objects.filter(
                pro_number__startswith=today, organization=self.organization
            ).count()
            + 1
        )

        return f"{today}-{count_for_today:04d}"

    def validate_delivery_slot(
        self, start_time: datetime, end_time: datetime, location: Location
    ) -> None:
        """Validates if a delivery slot is available for a given time interval and location.

        This method checks the existence of a delivery slot during a given timespan and location for a specific
        customer.

        The customer is attached to the object instance (self). In the absence of matching slots, it raises a
        ValidationError, indicating that the selected appointment window doesn't match with the customer’s allowed
        schedule.

        Args:
            start_time (datetime): The starting point of the time interval for the delivery slot.
            end_time (datetime): The ending point of the time interval for the delivery slot.
            location (Location): The location where the delivery should occur.

        Raises:
            ValidationError: If no delivery slots are found that align with the desired time interval or location.
        """
        from customer.models import DeliverySlot

        # Convert start_time to a weekday number (0=Monday, 6=Sunday)
        day_of_week = start_time.weekday()
        allowed_slots = DeliverySlot.objects.filter(
            customer=self.customer, day_of_week=day_of_week, location=location
        )

        if not allowed_slots.exists():
            return  # If no slots exist, bypass validation

        # Ensure both start and end times are within a single slot's duration
        if not any(
            slot.start_time <= start_time.time() <= slot.end_time
            and slot.start_time <= end_time.time() <= slot.end_time
            for slot in allowed_slots
        ):
            raise ValidationError(
                {
                    "origin_appointment_window_start": _(
                        "The chosen appointment window for the location is not allowed by the customer. Please try "
                        "again."
                    ),
                }
            )

    def calculate_other_charge_amount(self):
        if other_charges := self.additional_charges.all():
            self.other_charge_amount = sum(
                other_charge.charge_amount for other_charge in other_charges
            )
        else:
            self.other_charge_amount = 0

    @property
    def temperature_differential(self) -> int:
        if self.temperature_min and self.temperature_max:
            return int(self.temperature_max - self.temperature_min)
        return 0


class ShipmentDocumentation(GenericModel):
    """
    Stores documentation related to a :model:`shipment.Shipmentt`.
    """

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    shipment = models.ForeignKey(
        Shipment,
        on_delete=models.CASCADE,
        related_name="shipment_documentation",
        verbose_name=_("Shipment"),
    )
    document = models.FileField(
        _("Document"),
        upload_to=shipment_documentation_upload_to,
        validators=[FileExtensionValidator(allowed_extensions=["pdf"])],
    )
    document_class = models.ForeignKey(
        "billing.DocumentClassification",
        on_delete=models.CASCADE,
        related_name="shipment_documentation",
        verbose_name=_("Document Class"),
        help_text=_("Document Class"),
    )

    class Meta:
        """
        ShipmentDocumentation Metaclass
        """

        verbose_name = _("Shipment Documentation")
        verbose_name_plural = _("Shipment Documentation")
        db_table = "shipment_documentation"
        db_table_comment = "Stores documentation related to a related shipment."

    def __str__(self) -> str:
        """String representation of the ShipmentDocumentation

        Returns:
            str: String representation of the ShipmentDocumentation
        """
        return textwrap.shorten(
            f"{self.shipment} - {self.document_class}", 50, placeholder="..."
        )

    def get_absolute_url(self) -> str:
        """Get the absolute url for the ShipmentDocumentation

        Returns:
            str: Absolute url for the ShipmentDocumentation
        """
        return reverse("shipment-documentation-detail", kwargs={"pk": self.pk})


class ShipmentComment(GenericModel):
    """
    Stores comments related to a :model:`shipment.Shipment`.
    """

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    shipment = models.ForeignKey(
        Shipment,
        on_delete=models.CASCADE,
        related_name="shipment_comments",
        related_query_name="shipment_comment",
        verbose_name=_("shipment"),
    )
    comment_type = models.ForeignKey(
        "dispatch.CommentType",
        on_delete=models.CASCADE,
        related_name="shipment_comments",
        related_query_name="shipment_comment",
        verbose_name=_("Comment Type"),
        help_text=_("Comment Type"),
    )
    comment = models.TextField(
        _("Comment"),
        help_text=_("Comment"),
    )
    entered_by = models.ForeignKey(
        User,
        on_delete=models.CASCADE,
        related_name="shipment_comments",
        related_query_name="shipment_comment",
        verbose_name=_("Entered By"),
        help_text=_("Entered By"),
    )

    class Meta:
        """
        ShipmentComment Metaclass
        """

        verbose_name = _("Shipment Comment")
        verbose_name_plural = _("Shipment Comments")
        db_table = "shipment_comment"
        db_table_comment = "Stores comments related to a related shipment."

    def __str__(self) -> str:
        """String representation of the ShipmentComment

        Returns:
            str: String representation of the ShipmentComment
        """
        return textwrap.shorten(
            f"{self.shipment} - {self.comment_type}", 50, placeholder="..."
        )

    def get_absolute_url(self) -> str:
        """Get the absolute url for the ShipmentComment

        Returns:
            str: Absolute url for the ShipmentComment
        """
        return reverse("shipment-comment-detail", kwargs={"pk": self.pk})


class AdditionalCharge(GenericModel):
    """
    Stores Additional Charge related to a :model:`shipment.Shipment`.
    """

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    shipment = models.ForeignKey(
        Shipment,
        on_delete=models.CASCADE,
        related_name="additional_charges",
        related_query_name="additional_charge",
        verbose_name=_("shipment"),
    )
    accessorial_charge = models.ForeignKey(
        "billing.AccessorialCharge",
        on_delete=models.CASCADE,
        related_name="additional_charges",
        related_query_name="additional_charge",
        verbose_name=_("Charge"),
        help_text=_("Charge"),
    )
    description = models.CharField(
        _("Description"),
        max_length=100,
        help_text=_("Description of the charge"),
        blank=True,
    )
    charge_amount = models.DecimalField(
        _("Charge Amount"),
        max_digits=19,
        decimal_places=4,
        help_text=_("Charge Amount"),
        null=True,
        blank=True,
    )
    unit = models.PositiveIntegerField(
        _("Unit"),
        default=1,
        help_text=_("Number of units to be charged"),
    )
    sub_total = models.DecimalField(
        _("Sub Total"),
        max_digits=19,
        decimal_places=4,
        help_text=_("Sub Total"),
        null=True,
        blank=True,
    )
    entered_by = models.ForeignKey(
        User,
        on_delete=models.CASCADE,
        related_name="additional_charges",
        related_query_name="additional_charge",
        verbose_name=_("Entered By"),
        help_text=_("Entered By"),
    )

    class Meta:
        """
        AdditionalCharges Metaclass
        """

        verbose_name = _("Additional Charge")
        verbose_name_plural = _("Additional Charges")
        db_table = "additional_charge"
        db_table_comment = "Stores Additional Charge related to a related shipment."

    def __str__(self) -> str:
        """String representation of the AdditionalCharges

        Returns:
            str: String representation of the AdditionalCharges
        """
        return textwrap.shorten(
            f"{self.shipment} - {self.accessorial_charge}", 50, placeholder="..."
        )

    def save(self, *args: Any, **kwargs: Any) -> None:
        """
        Save the AdditionalCharge
        """
        self.charge_amount = self.accessorial_charge.charge_amount
        self.sub_total = self.charge_amount * self.unit

        super().save(*args, **kwargs)

    def get_absolute_url(self) -> str:
        """Get the absolute url for the AdditionalCharges

        Returns:
            str: Absolute url for the AdditionalCharges
        """
        return reverse("additional-charges-detail", kwargs={"pk": self.pk})


class ReasonCode(GenericModel):
    """
    Stores Reason code information for when a load is voided or cancelled.
    """

    @final
    class CodeTypeChoices(models.TextChoices):
        """
        Code Type choices for Reason Code model
        """

        VOIDED = "VOIDED", _("Voided")
        CANCELLED = "CANCELLED", _("Cancelled")

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    status = ChoiceField(
        _("Status"),
        choices=PrimaryStatusChoices.choices,
        help_text=_("Status of the Reason Code"),
        default=PrimaryStatusChoices.ACTIVE,
    )
    code = models.CharField(
        _("Code"),
        max_length=10,
        help_text=_("Code of the Reason Code"),
    )
    code_type = ChoiceField(
        _("Code Type"),
        choices=CodeTypeChoices.choices,
        help_text=_("Code Type of the Reason Code"),
    )
    description = models.CharField(
        _("Description"),
        max_length=100,
        help_text=_("Description of the Reason Code"),
    )

    class Meta:
        """
        Reason Code Metaclass
        """

        verbose_name = _("Reason Code")
        verbose_name_plural = _("Reason Codes")
        db_table = "reason_code"
        db_table_comment = (
            "Stores Reason code information for when a load is voided or cancelled."
        )
        constraints = [
            models.UniqueConstraint(
                Lower("code"),
                "organization",
                name="unique_reason_code_organization",
            )
        ]

    def __str__(self) -> str:
        """Reason Code String Representation

        Returns:
            str: Code of the Reason
        """
        return textwrap.shorten(
            f"{self.code} - {self.description}", 50, placeholder="..."
        )

    def get_absolute_url(self) -> str:
        """Reason Code Absolute URL

        Returns:
            str: Reason Code Absolute URL
        """
        return reverse("reason-codes-detail", kwargs={"pk": self.pk})


@final
class TemplateTypeChoices(models.TextChoices):
    """
    Template Type choices for Formula Template model
    """

    REFRIGERATED = "REFRIGERATED", _("Refrigerated Shipments")
    HAZMAT = "HAZMAT", _("Hazardous Material Shipments")


class FormulaTemplate(GenericModel):
    """
    Stores formula template information for a related :model:`organization.Organization`.
    """

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    name = models.CharField(
        verbose_name=_("Name"),
        max_length=255,
        help_text=_("Name of the formula template"),
    )
    formula_text = models.TextField(
        verbose_name=_("Formula Text"), help_text=_("Formula text")
    )
    description = models.TextField(
        verbose_name=_("Description"),
        blank=True,
        help_text=_("Description or notes about the formula"),
    )
    template_type = models.CharField(
        verbose_name=_("Template Type"),
        max_length=50,
        choices=TemplateTypeChoices.choices,
        help_text=_("Type of the formula template"),
    )
    customer = models.ForeignKey(
        "customer.Customer",
        on_delete=models.RESTRICT,
        related_name="formula_templates",
        verbose_name=_("Customer"),
        help_text=_("Customer of the formula template"),
        blank=True,
        null=True,
    )
    shipment_type = models.ForeignKey(
        ShipmentType,
        on_delete=models.RESTRICT,
        related_name="formula_templates",
        verbose_name=_("shipment type"),
        help_text=_("shipment type of the formula template"),
        blank=True,
        null=True,
    )
    auto_apply = models.BooleanField(
        verbose_name=_("Auto Apply"),
        default=False,
        help_text=_(
            "Auto apply formula template to shipments, based on customer, shipment_type, and template_type."
        ),
    )
    history = AuditlogHistoryField()

    class Meta:
        """
        Metaclass for FormulaTemplate model
        """

        verbose_name = _("Formula Template")
        verbose_name_plural = _("Formula Templates")
        db_table = "formula_template"
        db_table_comment = (
            "Stores formula template information for a related organization."
        )

    def __str__(self) -> str:
        """String representation of the FormulaTemplate

        Returns:
            str: Formula Template Name
        """
        return textwrap.shorten(self.name, 50, placeholder="...")

    def get_absolute_url(self) -> str:
        """Formula Template Absolute URL

        Returns:
            str: Formula Template Absolute URL
        """
        return reverse("formula-template-detail", kwargs={"pk": self.pk})

    def clean(self) -> None:
        """Clean method for FormulaTemplate model.

        Returns:
            None: This function does not return anything.
        """
        from shipment import helpers

        if self.formula_text:
            # Get the list of variables in the formula.
            formula_variables = helpers.extract_variable_from_formula(
                formula=self.formula_text
            )

            if invalid_variables := [
                var
                for var in formula_variables
                if var not in helpers.FORMULA_ALLOWED_VARIABLES
            ]:
                raise ValidationError(
                    {
                        "formula_text": _(
                            f"Formula template contains invalid variables: {', '.join(invalid_variables)}. Please try "
                            "again."
                        )
                    },
                    code="invalid",
                )

        super().clean()


class ShipmentCommodity(GenericModel):
    shipment = models.ForeignKey(
        Shipment,
        on_delete=models.CASCADE,
        related_name="shipment_commodities",
        related_query_name="shipment_commodity",
        verbose_name=_("Shipment"),
    )
    commodity = models.ForeignKey(
        "commodities.Commodity",
        on_delete=models.PROTECT,
        related_name="shipment_commodities",
        related_query_name="shipment_commodity",
        verbose_name=_("Commodity"),
        help_text=_("Commodity"),
    )
    hazardous_material = models.ForeignKey(
        "commodities.HazardousMaterial",
        on_delete=models.PROTECT,
        related_name="shipment_commodities",
        related_query_name="shipment_commodity",
        verbose_name=_("Hazardous Material"),
        help_text=_("Hazardous Material"),
        blank=True,
        null=True,
    )
    quantity = models.DecimalField(
        _("Quantity"),
        max_digits=10,
        decimal_places=2,
        help_text=_("Quantity of the commodity in the shipment"),
    )
    placard_needed = models.BooleanField(
        _("Placard Needed"),
        default=False,
        help_text=_("Placard Needed"),
    )

    class Meta:
        """
        ShipmentCommodity Metaclass
        """

        verbose_name = _("Shipment Commodity")
        verbose_name_plural = _("Shipment Commodities")
        db_table = "shipment_commodity"
        db_table_comment = "Stores the commodities for a related shipment."

    def __str__(self) -> str:
        """String representation of the ShipmentCommodity

        Returns:
            str: String representation of the ShipmentCommodity
        """
        return textwrap.shorten(
            f"{self.shipment} - {self.commodity}", 50, placeholder="..."
        )

    def clean(self) -> None:
        from shipment.validation import check_hazardous_material_compatibility

        check_hazardous_material_compatibility(commodity=self)
        super().clean()

    def save(self, *args: Any, **kwargs: Any) -> None:
        """Overrides the default Django save method to provide custom save behavior for the Shipment Commodity model.

        Args:
            *args (Any): Variable length argument list.
            **kwargs (Any): Arbitrary keyword arguments.

        Returns:
            None: This function does not return anything.
        """
        super().full_clean()

        # If the commodity is hazardous, set the hazardous_material field to the commodity's hazardous_material field.
        # Also, set the shipment's is_hazmat field to True.
        if self.commodity and self.commodity.hazardous_material:
            self.hazardous_material = self.commodity.hazardous_material
            self.shipment.is_hazmat = True

        if self.hazardous_material:
            self.placard_needed = True

        super().save(*args, **kwargs)

    def get_absolute_url(self) -> str:
        """Get the absolute url for the ShipmentCommodity

        Returns:
            str: Absolute url for the ShipmentCommodity
        """
        return reverse("shipment-commodity-detail", kwargs={"pk": self.pk})


class HazardousMaterialSegregation(GenericModel):
    """
    Stores hazardous material segregation information for a related :model:`organization.Organization`.
    """

    @final
    class SegregationTypeChoices(models.TextChoices):
        """
        Segregation Type choices for Hazardous Material Segregation model
        """

        NOT_ALLOWED = "X", _("Not Allowed")
        ALLOWED_WITH_CONDITIONS = "O", _("Allowed with Conditions")

    class_a = ChoiceField(
        verbose_name=_("Class/Division A"),
        help_text=_("First hazardous material class or division."),
        choices=HazardousClassChoices.choices,
    )
    class_b = ChoiceField(
        verbose_name=_("Class/Division B"),
        help_text=_("Second hazardous material class or division."),
        choices=HazardousClassChoices.choices,
    )
    segregation_type = models.CharField(
        max_length=1,
        choices=SegregationTypeChoices.choices,
        verbose_name=_("Segregation Type"),
        help_text=_(
            "Indicates if the materials are allowed to be transported together."
        ),
    )

    class Meta:
        """
        Metaclass for HazardousMaterialSegregation
        """

        verbose_name = _("Hazardous Material Segregation")
        verbose_name_plural = _("Hazardous Material Segregations")
        constraints = [
            models.UniqueConstraint(
                fields=["organization", "class_a", "class_b"],
                name="unique_hazmat_segregation",
            )
        ]
        db_table = "hazardous_material_segregation"

    def __str__(self) -> str:
        """String representation of the HazardousMaterialSegregation

        Returns:
            str: String representation of the HazardousMaterialSegregation
        """
        return textwrap.shorten(
            f"Class A: {self.class_a}, Class B: {self.class_b} - {self.segregation_type}",
            50,
            placeholder="...",
        )
