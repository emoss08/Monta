# Generated by Django 4.1.6 on 2023-02-11 18:58

from django.db import migrations, models
import django.db.models.deletion
import django_extensions.db.fields
import django_lifecycle.mixins
import localflavor.us.models
import phonenumber_field.modelfields
import utils.models
import uuid


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ("billing", "0001_initial"),
        ("organization", "0001_initial"),
    ]

    operations = [
        migrations.CreateModel(
            name="Customer",
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
                    "is_active",
                    models.BooleanField(
                        default=True,
                        help_text="Designates whether this customer should be treated as active. Unselect this instead of deleting customers.",
                        verbose_name="Active",
                    ),
                ),
                (
                    "code",
                    models.CharField(
                        editable=False,
                        help_text="Customer code",
                        max_length=10,
                        unique=True,
                        verbose_name="Code",
                    ),
                ),
                (
                    "name",
                    models.CharField(
                        help_text="Customer name", max_length=150, verbose_name="Name"
                    ),
                ),
                (
                    "address_line_1",
                    models.CharField(
                        blank=True,
                        help_text="Address line 1",
                        max_length=150,
                        verbose_name="Address Line 1",
                    ),
                ),
                (
                    "address_line_2",
                    models.CharField(
                        blank=True,
                        help_text="Address line 2",
                        max_length=150,
                        verbose_name="Address Line 2",
                    ),
                ),
                (
                    "city",
                    models.CharField(
                        blank=True,
                        help_text="City",
                        max_length=150,
                        verbose_name="City",
                    ),
                ),
                (
                    "state",
                    localflavor.us.models.USStateField(
                        blank=True,
                        help_text="State",
                        max_length=2,
                        verbose_name="State",
                    ),
                ),
                (
                    "zip_code",
                    localflavor.us.models.USZipCodeField(
                        blank=True,
                        help_text="Zip code",
                        max_length=10,
                        verbose_name="Zip Code",
                    ),
                ),
            ],
            options={
                "verbose_name": "Customer",
                "verbose_name_plural": "Customers",
                "ordering": ["-code"],
            },
            bases=(django_lifecycle.mixins.LifecycleModelMixin, models.Model),
        ),
        migrations.CreateModel(
            name="CustomerBillingProfile",
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
                    "is_active",
                    models.BooleanField(
                        default=True,
                        help_text="Designates whether this customer billing profile should be treated as active. Unselect this instead of deleting customer billing profiles.",
                        verbose_name="Active",
                    ),
                ),
            ],
            options={
                "verbose_name": "Customer Billing Profile",
                "verbose_name_plural": "Customer Billing Profiles",
                "ordering": ["customer"],
            },
        ),
        migrations.CreateModel(
            name="CustomerContact",
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
                    "is_active",
                    models.BooleanField(
                        default=True,
                        help_text="Designates whether this customer contact should be treated as active. Unselect this instead of deleting customer contacts.",
                        verbose_name="Active",
                    ),
                ),
                (
                    "name",
                    models.CharField(
                        help_text="Contact name", max_length=150, verbose_name="Name"
                    ),
                ),
                (
                    "email",
                    models.EmailField(
                        blank=True,
                        help_text="Contact email",
                        max_length=150,
                        verbose_name="Email",
                    ),
                ),
                (
                    "title",
                    models.CharField(
                        blank=True,
                        help_text="Contact title",
                        max_length=100,
                        verbose_name="Title",
                    ),
                ),
                (
                    "phone",
                    phonenumber_field.modelfields.PhoneNumberField(
                        blank=True,
                        help_text="Contact phone",
                        max_length=20,
                        null=True,
                        region=None,
                        verbose_name="Phone Number",
                    ),
                ),
                (
                    "is_payable_contact",
                    models.BooleanField(
                        default=False,
                        help_text="Designates whether this contact is the payable contact",
                        verbose_name="Payable Contact",
                    ),
                ),
            ],
            options={
                "verbose_name": "Customer Contact",
                "verbose_name_plural": "Customer Contacts",
                "ordering": ["customer", "name"],
            },
        ),
        migrations.CreateModel(
            name="CustomerEmailProfile",
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
                        help_text="Name", max_length=50, verbose_name="Name"
                    ),
                ),
                (
                    "subject",
                    models.CharField(
                        blank=True,
                        help_text="Subject",
                        max_length=100,
                        verbose_name="Subject",
                    ),
                ),
                (
                    "comment",
                    models.CharField(
                        blank=True,
                        help_text="Comment",
                        max_length=100,
                        verbose_name="Comment",
                    ),
                ),
                (
                    "from_address",
                    models.CharField(
                        blank=True,
                        help_text="From Address",
                        max_length=255,
                        verbose_name="From Address",
                    ),
                ),
                (
                    "blind_copy",
                    models.CharField(
                        blank=True,
                        help_text="Blind Copy",
                        max_length=255,
                        verbose_name="Blind Copy",
                    ),
                ),
                (
                    "read_receipt",
                    models.BooleanField(
                        default=False,
                        help_text="Read Receipt",
                        verbose_name="Read Receipt",
                    ),
                ),
                (
                    "read_receipt_to",
                    models.CharField(
                        blank=True,
                        help_text="Read Receipt To",
                        max_length=255,
                        verbose_name="Read Receipt To",
                    ),
                ),
                (
                    "attachment_name",
                    models.CharField(
                        blank=True,
                        help_text="Attachment Name",
                        max_length=255,
                        verbose_name="Attachment Name",
                    ),
                ),
            ],
            options={
                "verbose_name": "Customer Email Profile",
                "verbose_name_plural": "Customer Email Profiles",
                "ordering": ["-name"],
            },
        ),
        migrations.CreateModel(
            name="CustomerFuelTable",
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
                        help_text="Customer Fuel Profile Name",
                        max_length=10,
                        unique=True,
                        verbose_name="Name",
                    ),
                ),
                (
                    "description",
                    models.CharField(
                        blank=True,
                        help_text="Customer Fuel Profile Description",
                        max_length=150,
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
                "verbose_name": "Customer Fuel Table",
                "verbose_name_plural": "Customer Fuel Table",
                "ordering": ["id"],
            },
        ),
        migrations.CreateModel(
            name="CustomerRuleProfile",
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
                        help_text="Name", max_length=50, verbose_name="Name"
                    ),
                ),
                (
                    "document_class",
                    models.ManyToManyField(
                        help_text="Document class",
                        related_name="customer_rule_profile",
                        to="billing.documentclassification",
                        verbose_name="Document Class",
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
                "verbose_name": "Customer Rule Profile",
                "verbose_name_plural": "Customer Rule Profiles",
                "ordering": ["-name"],
            },
        ),
        migrations.CreateModel(
            name="CustomerFuelTableDetail",
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
                    "amount",
                    models.DecimalField(
                        blank=True,
                        decimal_places=5,
                        help_text="Amount",
                        max_digits=16,
                        null=True,
                        verbose_name="Amount",
                    ),
                ),
                (
                    "method",
                    utils.models.ChoiceField(
                        choices=[("D", "Distance"), ("F", "Flat"), ("P", "Percentage")],
                        help_text="Method",
                        max_length=1,
                        verbose_name="Method",
                    ),
                ),
                (
                    "start_price",
                    models.DecimalField(
                        blank=True,
                        decimal_places=3,
                        help_text="Start Price",
                        max_digits=5,
                        null=True,
                        verbose_name="Start Price",
                    ),
                ),
                (
                    "percentage",
                    models.DecimalField(
                        blank=True,
                        decimal_places=2,
                        help_text="Percentage",
                        max_digits=6,
                        null=True,
                        verbose_name="Percentage",
                    ),
                ),
                (
                    "customer_fuel_table",
                    models.ForeignKey(
                        help_text="Customer Fuel Profile",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="customer_fuel_table_details",
                        related_query_name="customer_fuel_table_detail",
                        to="customer.customerfueltable",
                        verbose_name="Customer Fuel Profile",
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
                "verbose_name": "Customer Fuel Profile Detail",
                "verbose_name_plural": "Customer Fuel Profile Details",
                "ordering": ["customer_fuel_table"],
            },
        ),
        migrations.CreateModel(
            name="CustomerFuelProfile",
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
                    "start_date",
                    models.DateField(help_text="Start Date", verbose_name="Start Date"),
                ),
                (
                    "end_date",
                    models.DateField(
                        blank=True,
                        help_text="End Date",
                        null=True,
                        verbose_name="End Date",
                    ),
                ),
                (
                    "days_to_use",
                    utils.models.ChoiceField(
                        choices=[
                            ("D", "Delivery Date"),
                            ("S", "Actual Shipment Date"),
                            ("C", "Scheduled Shipment Date"),
                            ("E", "Date Entered"),
                        ],
                        help_text="Days to Use",
                        max_length=1,
                        verbose_name="Days to Use",
                    ),
                ),
                (
                    "fuel_region",
                    utils.models.ChoiceField(
                        choices=[
                            ("USA", "US National Average"),
                            ("EAST", "East Coast"),
                            ("NE", "New England"),
                            ("GA", "General Atlantic"),
                            ("LA", "Lower Atlantic"),
                            ("MW", "Midwest"),
                            ("GC", "Gulf Coast"),
                            ("RM", "Rocky Mountain"),
                            ("WC", "West Coast"),
                            ("CA", "California"),
                            ("WCL", "West Coast (No LA)"),
                        ],
                        help_text="Fuel Region",
                        max_length=4,
                        verbose_name="Fuel Region",
                    ),
                ),
                (
                    "fsc_method",
                    utils.models.ChoiceField(
                        choices=[
                            ("P", "Percentage"),
                            ("F", "Flat"),
                            ("D", "Distance"),
                            ("T", "Table"),
                        ],
                        help_text="FSC Method",
                        max_length=1,
                        verbose_name="FSC Method",
                    ),
                ),
                (
                    "base_price",
                    models.DecimalField(
                        blank=True,
                        decimal_places=3,
                        help_text="Base Price",
                        max_digits=16,
                        null=True,
                        verbose_name="Base Price",
                    ),
                ),
                (
                    "fuel_variance",
                    models.DecimalField(
                        blank=True,
                        decimal_places=3,
                        help_text="Fuel Variance ex: 0.02",
                        max_digits=16,
                        null=True,
                        verbose_name="Fuel Variance",
                    ),
                ),
                (
                    "amount",
                    models.DecimalField(
                        blank=True,
                        decimal_places=5,
                        help_text="Amount",
                        max_digits=16,
                        null=True,
                        verbose_name="Amount",
                    ),
                ),
                (
                    "percentage",
                    models.DecimalField(
                        blank=True,
                        decimal_places=2,
                        help_text="Percentage",
                        max_digits=6,
                        null=True,
                        verbose_name="Percentage",
                    ),
                ),
                (
                    "customer",
                    models.ForeignKey(
                        help_text="Customer",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="customer_fuel_profiles",
                        related_query_name="customer_fuel_profile",
                        to="customer.customer",
                        verbose_name="Customer",
                    ),
                ),
                (
                    "customer_fuel_table",
                    models.ForeignKey(
                        blank=True,
                        help_text="Customer Fuel Table",
                        null=True,
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="customer_fuel_profiles",
                        related_query_name="customer_fuel_profile",
                        to="customer.customerfueltable",
                        verbose_name="Customer Fuel Table",
                    ),
                ),
                (
                    "fsc_code",
                    models.ForeignKey(
                        help_text="FSC Code",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="customer_fuel_profiles",
                        related_query_name="customer_fuel_profile",
                        to="billing.accessorialcharge",
                        verbose_name="FSC Code",
                    ),
                ),
            ],
            options={
                "verbose_name": "Customer Fuel Profile",
                "verbose_name_plural": "Customer Fuel Profiles",
                "ordering": ["customer"],
            },
        ),
    ]
