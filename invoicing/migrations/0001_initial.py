# Generated by Django 4.1.6 on 2023-02-12 02:32

import uuid

import django.db.models.deletion
import django_extensions.db.fields
from django.db import migrations, models


class Migration(migrations.Migration):
    initial = True

    dependencies = [
        ("organization", "0002_alter_emailprofile_protocol_taxrate"),
    ]

    operations = [
        migrations.CreateModel(
            name="InvoiceControl",
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
                    "invoice_number_prefix",
                    models.CharField(
                        default="INV-",
                        help_text="Define a prefix for invoice numbers.",
                        max_length=10,
                        verbose_name="Invoice Number Prefix",
                    ),
                ),
                (
                    "invoice_due_after_days",
                    models.PositiveIntegerField(
                        default=30,
                        help_text="Define the number of days after invoice date that an invoice is due.",
                        verbose_name="Invoice Due After Days",
                    ),
                ),
                (
                    "invoice_terms",
                    models.TextField(
                        blank=True,
                        default="",
                        help_text="Define invoice terms.",
                        verbose_name="Invoice Terms",
                    ),
                ),
                (
                    "invoice_footer",
                    models.TextField(
                        blank=True,
                        default="",
                        help_text="Define invoice footer.",
                        verbose_name="Invoice Footer",
                    ),
                ),
                (
                    "invoice_logo",
                    models.ImageField(
                        blank=True,
                        help_text="Define invoice logo.",
                        upload_to="invoice_logo",
                        verbose_name="Invoice Logo",
                    ),
                ),
                (
                    "organization",
                    models.OneToOneField(
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="invoice_control",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
            ],
            options={
                "verbose_name": "Invoice Control",
                "verbose_name_plural": "Invoice Controls",
            },
        ),
    ]
