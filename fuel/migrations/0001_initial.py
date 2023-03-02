# Generated by Django 4.1.7 on 2023-03-02 15:05

import uuid

import django.db.models.deletion
import django_extensions.db.fields
from django.db import migrations, models

import utils.models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ("accounting", "0001_initial"),
        ("organization", "0001_initial"),
    ]

    operations = [
        migrations.CreateModel(
            name="FuelVendor",
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
                        help_text="Unique identifier for the vendor",
                        primary_key=True,
                        serialize=False,
                        verbose_name="ID",
                    ),
                ),
                (
                    "name",
                    models.CharField(
                        help_text="Name of the vendor",
                        max_length=255,
                        verbose_name="Name",
                    ),
                ),
                (
                    "account_code",
                    models.CharField(
                        help_text="Account Code for the vendor",
                        max_length=255,
                        verbose_name="Account Code",
                    ),
                ),
                (
                    "sub_account_code",
                    models.CharField(
                        help_text="Sub Account Code for the vendor",
                        max_length=255,
                        verbose_name="Sub Account Code",
                    ),
                ),
                (
                    "communication_type",
                    utils.models.ChoiceField(
                        choices=[
                            ("ftp", "File Transfer Protocol"),
                            ("sftp", "Secure File Transfer Protocol"),
                            ("local", "Local File System"),
                            ("https", "Hypertext Transfer Protocol Secure"),
                        ],
                        default="ftp",
                        help_text="Communication type used to connect to the vendor",
                        max_length=5,
                        verbose_name="Communication Type",
                    ),
                ),
                (
                    "login",
                    models.CharField(
                        blank=True,
                        help_text="Login for the vendor",
                        max_length=255,
                        verbose_name="Login",
                    ),
                ),
                (
                    "password",
                    models.CharField(
                        blank=True,
                        help_text="Password for the vendor",
                        max_length=255,
                        verbose_name="Password",
                    ),
                ),
                (
                    "port",
                    models.PositiveIntegerField(
                        blank=True,
                        help_text="Port for the vendor",
                        null=True,
                        verbose_name="Port",
                    ),
                ),
                (
                    "server_address",
                    models.CharField(
                        blank=True,
                        help_text="Server address for the vendor",
                        max_length=255,
                        verbose_name="Server Address",
                    ),
                ),
                (
                    "directory",
                    models.CharField(
                        blank=True,
                        help_text="Directory for the vendor",
                        max_length=255,
                        verbose_name="Directory",
                    ),
                ),
                (
                    "proxy_server",
                    models.CharField(
                        blank=True,
                        help_text="Proxy server for the vendor",
                        max_length=255,
                        verbose_name="Proxy Server",
                    ),
                ),
                (
                    "email_address",
                    models.EmailField(
                        blank=True,
                        help_text="Email address for the vendor",
                        max_length=254,
                        verbose_name="Email Address",
                    ),
                ),
                (
                    "communication_mode",
                    utils.models.ChoiceField(
                        blank=True,
                        choices=[
                            ("push", "Push"),
                            ("pull", "Pull"),
                            ("push_pull", "Push and Pull"),
                        ],
                        default="push",
                        help_text="Communication mode used to connect to the vendor",
                        max_length=9,
                        verbose_name="Communication Mode",
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
                "verbose_name": "Fuel Vendor",
                "verbose_name_plural": "Fuel Vendors",
                "db_table": "fuel_vendor",
            },
        ),
        migrations.CreateModel(
            name="FuelVendorFuelDetail",
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
                        help_text="Unique identifier for the vendor",
                        primary_key=True,
                        serialize=False,
                        verbose_name="ID",
                    ),
                ),
                (
                    "create_ap_voucher",
                    utils.models.ChoiceField(
                        choices=[
                            ("none", "None"),
                            ("regular", "Regular"),
                            ("manual", "Manual"),
                        ],
                        default="none",
                        help_text="Create AP Voucher",
                        max_length=7,
                        verbose_name="Create AP Voucher",
                    ),
                ),
                (
                    "ap_division_code",
                    models.ForeignKey(
                        help_text="AP Division Code",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="ap_division_code",
                        to="accounting.divisioncode",
                    ),
                ),
                (
                    "fuel_vendor",
                    models.OneToOneField(
                        help_text="Fuel Vendor",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="fuel_vendor",
                        to="fuel.fuelvendor",
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
                "verbose_name": "Fuel Vendor Fuel Detail",
                "verbose_name_plural": "Fuel Vendor Fuel Details",
                "db_table": "fuel_vendor_fuel_detail",
            },
        ),
    ]
