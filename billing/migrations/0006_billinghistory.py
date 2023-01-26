# Generated by Django 4.1.5 on 2023-01-25 04:44

import uuid

import django.db.models.deletion
import django_extensions.db.fields
from django.db import migrations, models

import utils.models


class Migration(migrations.Migration):

    dependencies = [
        ("organization", "0004_alter_organization_name"),
        ("order", "0012_alter_order_mileage"),
        ("billing", "0005_billingexception"),
    ]

    operations = [
        migrations.CreateModel(
            name="BillingHistory",
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
                        help_text="Unique identifier for the billing history",
                        primary_key=True,
                        serialize=False,
                        unique=True,
                    ),
                ),
                (
                    "batch_name",
                    models.CharField(
                        blank=True,
                        help_text="Name of the batch",
                        max_length=100,
                        null=True,
                        unique=True,
                        verbose_name="Batch Name",
                    ),
                ),
                (
                    "bill_type",
                    utils.models.ChoiceField(
                        choices=[
                            ("INVOICE", "Invoice"),
                            ("CREDIT", "Credit"),
                            ("DEBIT", "Debit"),
                            ("PREPAID", "Prepaid"),
                            ("OTHER", "Other"),
                        ],
                        default="INVOICE",
                        help_text="Type of bill",
                        max_length=7,
                        verbose_name="Bill Type",
                    ),
                ),
                (
                    "sub_total",
                    models.DecimalField(
                        blank=True,
                        decimal_places=2,
                        help_text="Sub total for Order",
                        max_digits=10,
                        null=True,
                        verbose_name="Sub Total",
                    ),
                ),
                (
                    "order",
                    models.ForeignKey(
                        help_text="Assigned order to the billing history",
                        on_delete=django.db.models.deletion.RESTRICT,
                        related_name="billing_history",
                        to="order.order",
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
                "verbose_name": "Billing History",
                "verbose_name_plural": "Billing Histories",
                "ordering": ["order"],
            },
        ),
    ]
