# Generated by Django 4.1.7 on 2023-03-02 15:05

import uuid

import django.db.models.deletion
import django_extensions.db.fields
from django.conf import settings
from django.db import migrations, models

import utils.models


class Migration(migrations.Migration):
    initial = True

    dependencies = [
        ("movements", "0001_initial"),
        ("location", "0001_initial"),
        ("organization", "0001_initial"),
        ("dispatch", "0002_initial"),
        migrations.swappable_dependency(settings.AUTH_USER_MODEL),
    ]

    operations = [
        migrations.CreateModel(
            name="QualifierCode",
            fields=[
                (
                    "created",
                    django_extensions.db.fields.CreationDateTimeField(
                        auto_now_add=True, verbose_name="created"
                    ),
                ),
                (
                    "modified",
                    django_extensions.db.fields.ModificationDateTimeField(
                        auto_now=True, verbose_name="modified"
                    ),
                ),
                (
                    "id",
                    models.UUIDField(
                        default=uuid.uuid4,
                        editable=False,
                        primary_key=True,
                        serialize=False,
                        unique=True,
                    ),
                ),
                (
                    "code",
                    models.CharField(
                        help_text="Code of the Qualifier Code",
                        max_length=255,
                        unique=True,
                        verbose_name="Code",
                    ),
                ),
                (
                    "description",
                    models.CharField(
                        help_text="Description of the Qualifier Code",
                        max_length=100,
                        verbose_name="Description",
                    ),
                ),
                (
                    "organization",
                    models.ForeignKey(
                        help_text="Organization",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="%(class)ss",
                        related_query_name="%(class)s",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
            ],
            options={
                "verbose_name": "Qualifier Code",
                "verbose_name_plural": "Qualifier Codes",
                "db_table": "qualifier_code",
                "ordering": ["code"],
            },
        ),
        migrations.CreateModel(
            name="Stop",
            fields=[
                (
                    "created",
                    django_extensions.db.fields.CreationDateTimeField(
                        auto_now_add=True, verbose_name="created"
                    ),
                ),
                (
                    "modified",
                    django_extensions.db.fields.ModificationDateTimeField(
                        auto_now=True, verbose_name="modified"
                    ),
                ),
                (
                    "id",
                    models.UUIDField(
                        default=uuid.uuid4,
                        editable=False,
                        primary_key=True,
                        serialize=False,
                        unique=True,
                    ),
                ),
                (
                    "status",
                    utils.models.ChoiceField(
                        choices=[
                            ("N", "New"),
                            ("P", "In Progress"),
                            ("C", "Completed"),
                            ("H", "Hold"),
                            ("B", "Billed"),
                            ("V", "Voided"),
                        ],
                        default="N",
                        help_text="The status of the stop.",
                        max_length=1,
                    ),
                ),
                (
                    "sequence",
                    models.PositiveIntegerField(
                        default=1,
                        help_text="The sequence of the stop in the movement.",
                        verbose_name="Sequence",
                    ),
                ),
                (
                    "pieces",
                    models.PositiveIntegerField(
                        blank=True,
                        default=0,
                        help_text="Pieces",
                        null=True,
                        verbose_name="Pieces",
                    ),
                ),
                (
                    "weight",
                    models.PositiveIntegerField(
                        blank=True,
                        default=0,
                        help_text="Weight",
                        null=True,
                        verbose_name="Weight",
                    ),
                ),
                (
                    "address_line",
                    models.CharField(
                        blank=True,
                        help_text="Stop Address",
                        max_length=255,
                        verbose_name="Stop Address",
                    ),
                ),
                (
                    "appointment_time",
                    models.DateTimeField(
                        help_text="The time the equipment is expected to arrive at the stop.",
                        verbose_name="Stop Appointment Time",
                    ),
                ),
                (
                    "arrival_time",
                    models.DateTimeField(
                        blank=True,
                        help_text="The time the equipment actually arrived at the stop.",
                        null=True,
                        verbose_name="Stop Arrival Time",
                    ),
                ),
                (
                    "departure_time",
                    models.DateTimeField(
                        blank=True,
                        help_text="The time the equipment actually departed from the stop.",
                        null=True,
                        verbose_name="Stop Departure Time",
                    ),
                ),
                (
                    "stop_type",
                    utils.models.ChoiceField(
                        choices=[
                            ("P", "Pickup"),
                            ("SP", "Split Pickup"),
                            ("SD", "Split Drop Off"),
                            ("D", "Delivery"),
                            ("DO", "Drop Off"),
                        ],
                        help_text="The type of stop.",
                        max_length=2,
                    ),
                ),
                (
                    "location",
                    models.ForeignKey(
                        blank=True,
                        help_text="The location of the stop.",
                        null=True,
                        on_delete=django.db.models.deletion.PROTECT,
                        related_name="stops",
                        related_query_name="stop",
                        to="location.location",
                        verbose_name="Location",
                    ),
                ),
                (
                    "movement",
                    models.ForeignKey(
                        help_text="The movement that the stop belongs to.",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="stops",
                        related_query_name="stop",
                        to="movements.movement",
                        verbose_name="Movement",
                    ),
                ),
                (
                    "organization",
                    models.ForeignKey(
                        help_text="Organization",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="%(class)ss",
                        related_query_name="%(class)s",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
            ],
            options={
                "verbose_name": "Stop",
                "verbose_name_plural": "Stops",
                "db_table": "stop",
                "ordering": ["movement", "sequence"],
            },
        ),
        migrations.CreateModel(
            name="StopComment",
            fields=[
                (
                    "created",
                    django_extensions.db.fields.CreationDateTimeField(
                        auto_now_add=True, verbose_name="created"
                    ),
                ),
                (
                    "modified",
                    django_extensions.db.fields.ModificationDateTimeField(
                        auto_now=True, verbose_name="modified"
                    ),
                ),
                (
                    "id",
                    models.UUIDField(
                        default=uuid.uuid4,
                        editable=False,
                        primary_key=True,
                        serialize=False,
                        unique=True,
                    ),
                ),
                (
                    "comment",
                    models.TextField(help_text="Comment text.", verbose_name="Comment"),
                ),
                (
                    "comment_type",
                    models.ForeignKey(
                        help_text="The type of comment.",
                        on_delete=django.db.models.deletion.PROTECT,
                        related_name="stop_comments",
                        related_query_name="stop_comment",
                        to="dispatch.commenttype",
                        verbose_name="Comment Type",
                    ),
                ),
                (
                    "entered_by",
                    models.ForeignKey(
                        help_text="User who entered the comment.",
                        on_delete=django.db.models.deletion.PROTECT,
                        related_name="stop_comments",
                        related_query_name="stop_comment",
                        to=settings.AUTH_USER_MODEL,
                        verbose_name="Entered By",
                    ),
                ),
                (
                    "organization",
                    models.ForeignKey(
                        help_text="Organization",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="%(class)ss",
                        related_query_name="%(class)s",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
                (
                    "qualifier_code",
                    models.ForeignKey(
                        help_text="Qualifier code for the comment.",
                        on_delete=django.db.models.deletion.PROTECT,
                        related_name="stop_comments",
                        related_query_name="stop_comment",
                        to="stops.qualifiercode",
                        verbose_name="Qualifier Code",
                    ),
                ),
                (
                    "stop",
                    models.ForeignKey(
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="comments",
                        to="stops.stop",
                        verbose_name="Stop",
                    ),
                ),
            ],
            options={
                "verbose_name": "Stop Comment",
                "verbose_name_plural": "Stop Comments",
                "db_table": "stop_comment",
            },
        ),
        migrations.CreateModel(
            name="ServiceIncident",
            fields=[
                (
                    "created",
                    django_extensions.db.fields.CreationDateTimeField(
                        auto_now_add=True, verbose_name="created"
                    ),
                ),
                (
                    "modified",
                    django_extensions.db.fields.ModificationDateTimeField(
                        auto_now=True, verbose_name="modified"
                    ),
                ),
                (
                    "id",
                    models.UUIDField(
                        default=uuid.uuid4,
                        editable=False,
                        primary_key=True,
                        serialize=False,
                        unique=True,
                    ),
                ),
                (
                    "delay_reason",
                    models.CharField(
                        blank=True, max_length=100, verbose_name="Delay Reason"
                    ),
                ),
                (
                    "delay_time",
                    models.DurationField(
                        blank=True, null=True, verbose_name="Delay Time"
                    ),
                ),
                (
                    "delay_code",
                    models.ForeignKey(
                        blank=True,
                        null=True,
                        on_delete=django.db.models.deletion.PROTECT,
                        related_name="service_incidents",
                        related_query_name="service_incident",
                        to="dispatch.delaycode",
                        verbose_name="Delay Code",
                    ),
                ),
                (
                    "movement",
                    models.ForeignKey(
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="service_incidents",
                        related_query_name="service_incident",
                        to="movements.movement",
                        verbose_name="Movement",
                    ),
                ),
                (
                    "organization",
                    models.ForeignKey(
                        help_text="Organization",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="%(class)ss",
                        related_query_name="%(class)s",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
                (
                    "stop",
                    models.ForeignKey(
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="service_incidents",
                        related_query_name="service_incident",
                        to="stops.stop",
                        verbose_name="Stop",
                    ),
                ),
            ],
            options={
                "verbose_name": "Service Incident",
                "verbose_name_plural": "Service Incidents",
                "db_table": "service_incident",
            },
        ),
    ]
