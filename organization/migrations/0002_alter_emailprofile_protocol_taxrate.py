# Generated by Django 4.1.6 on 2023-02-11 22:35

from django.db import migrations, models
import django.db.models.deletion
import django_extensions.db.fields
import uuid


class Migration(migrations.Migration):

    dependencies = [
        ("organization", "0001_initial"),
    ]

    operations = [
        migrations.AlterField(
            model_name="emailprofile",
            name="protocol",
            field=models.CharField(
                blank=True,
                choices=[
                    ("TLS", "TLS"),
                    ("SSL", "SSL"),
                    ("UNENCRYPTED", "Unencrypted SMTP"),
                ],
                help_text="The protocol that will be used for outgoing email.",
                max_length=12,
                verbose_name="Protocol",
            ),
        ),
        migrations.CreateModel(
            name="TaxRate",
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
                    "name",
                    models.CharField(
                        help_text="The name of the tax rate.",
                        max_length=255,
                        verbose_name="Name",
                    ),
                ),
                (
                    "rate",
                    models.DecimalField(
                        decimal_places=2,
                        help_text="The rate of the tax rate.",
                        max_digits=5,
                        verbose_name="Rate",
                    ),
                ),
                (
                    "organization",
                    models.ForeignKey(
                        help_text="The organization that the tax rate belongs to.",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="tax_rates",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
            ],
            options={
                "verbose_name": "Tax Rate",
                "verbose_name_plural": "Tax Rates",
                "ordering": ["name"],
            },
        ),
    ]
