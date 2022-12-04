# Generated by Django 4.1.3 on 2022-12-04 00:03

from django.conf import settings
import django.core.validators
from django.db import migrations, models
import django.db.models.deletion
import django.utils.timezone
import django_extensions.db.fields
import localflavor.us.models
import utils.models
import uuid


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ("dispatch", "0001_initial"),
        ("organization", "0001_initial"),
        migrations.swappable_dependency(settings.AUTH_USER_MODEL),
    ]

    operations = [
        migrations.CreateModel(
            name="Worker",
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
                        blank=True,
                        editable=False,
                        help_text="The code of the worker. This field is required and must be unique.",
                        max_length=10,
                        unique=True,
                        verbose_name="Code",
                    ),
                ),
                (
                    "is_active",
                    models.BooleanField(
                        default=True,
                        help_text="Designates whether this worker should be treated as active. Unselect this instead of deleting workers.",
                        verbose_name="Active",
                    ),
                ),
                (
                    "worker_type",
                    utils.models.ChoiceField(
                        choices=[
                            ("EMPLOYEE", "Employee"),
                            ("CONTRACTOR", "Contractor"),
                        ],
                        default="EMPLOYEE",
                        help_text="The type of worker.",
                        max_length=10,
                        verbose_name="Worker type",
                    ),
                ),
                (
                    "first_name",
                    models.CharField(
                        help_text="The first name of the worker.",
                        max_length=255,
                        verbose_name="First name",
                    ),
                ),
                (
                    "last_name",
                    models.CharField(
                        help_text="The last name of the worker.",
                        max_length=255,
                        verbose_name="Last name",
                    ),
                ),
                (
                    "address_line_1",
                    models.CharField(
                        help_text="The address line 1 of the worker.",
                        max_length=255,
                        verbose_name="Address line 1",
                    ),
                ),
                (
                    "address_line_2",
                    models.CharField(
                        blank=True,
                        help_text="The address line 2 of the worker.",
                        max_length=255,
                        verbose_name="Address line 2",
                    ),
                ),
                (
                    "city",
                    models.CharField(
                        help_text="The city of the worker.",
                        max_length=255,
                        verbose_name="City",
                    ),
                ),
                (
                    "state",
                    localflavor.us.models.USStateField(
                        help_text="The state of the worker.",
                        max_length=2,
                        verbose_name="State",
                    ),
                ),
                (
                    "zip_code",
                    localflavor.us.models.USZipCodeField(
                        help_text="The zip code of the worker.",
                        max_length=10,
                        verbose_name="zip code",
                    ),
                ),
                (
                    "depot",
                    models.ForeignKey(
                        blank=True,
                        help_text="The depot of the worker.",
                        null=True,
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="worker",
                        related_query_name="workers",
                        to="organization.depot",
                        verbose_name="Depot",
                    ),
                ),
                (
                    "manager",
                    models.ForeignKey(
                        blank=True,
                        help_text="The manager of the worker.",
                        null=True,
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="worker",
                        related_query_name="workers",
                        to=settings.AUTH_USER_MODEL,
                        verbose_name="Manager",
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
                "verbose_name": "worker",
                "verbose_name_plural": "workers",
                "ordering": ["code"],
            },
        ),
        migrations.CreateModel(
            name="WorkerContact",
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
                        help_text="Name of the contact.",
                        max_length=255,
                        verbose_name="name",
                    ),
                ),
                (
                    "phone",
                    models.PositiveIntegerField(
                        blank=True,
                        help_text="Phone number in the format 1234567890",
                        null=True,
                        validators=[
                            django.core.validators.MinValueValidator(1000000000),
                            django.core.validators.MaxValueValidator(9999999999),
                        ],
                        verbose_name="phone",
                    ),
                ),
                (
                    "email",
                    models.EmailField(
                        blank=True,
                        help_text="Email address of the contact.",
                        max_length=255,
                        verbose_name="Email Address",
                    ),
                ),
                (
                    "relationship",
                    models.CharField(
                        blank=True,
                        help_text="Relationship to the worker.",
                        max_length=255,
                        verbose_name="Relationship",
                    ),
                ),
                (
                    "is_primary",
                    models.BooleanField(
                        default=False,
                        help_text="Is this the primary contact?",
                        verbose_name="Primary",
                    ),
                ),
                (
                    "mobile_phone",
                    models.PositiveIntegerField(
                        blank=True,
                        help_text="Mobile phone number.",
                        null=True,
                        verbose_name="mobile phone",
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
                    "worker",
                    models.ForeignKey(
                        help_text="Related Worker.",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="contacts",
                        related_query_name="contacts",
                        to="worker.worker",
                        verbose_name="Worker",
                    ),
                ),
            ],
            options={
                "verbose_name": "worker contact",
                "verbose_name_plural": "worker contacts",
                "ordering": ["worker"],
            },
        ),
        migrations.CreateModel(
            name="WorkerComment",
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
                    models.TextField(help_text="Comment", verbose_name="Comment"),
                ),
                (
                    "comment_type",
                    models.ForeignKey(
                        help_text="Related comment type.",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="comments",
                        related_query_name="comments",
                        to="dispatch.commenttype",
                        verbose_name="Comment Type",
                    ),
                ),
                (
                    "entered_by",
                    models.ForeignKey(
                        help_text="User who entered the comment.",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="worker_comments",
                        related_query_name="worker_comments",
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
                    "worker",
                    models.ForeignKey(
                        help_text="Related worker.",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="comments",
                        related_query_name="comments",
                        to="worker.worker",
                        verbose_name="worker",
                    ),
                ),
            ],
            options={
                "verbose_name": "worker comment",
                "verbose_name_plural": "worker comments",
                "ordering": ["worker"],
            },
        ),
        migrations.CreateModel(
            name="WorkerProfile",
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
                    "worker",
                    models.OneToOneField(
                        help_text="The worker of the profile.",
                        on_delete=django.db.models.deletion.CASCADE,
                        primary_key=True,
                        related_name="profile",
                        related_query_name="profiles",
                        serialize=False,
                        to="worker.worker",
                        verbose_name="Worker",
                    ),
                ),
                (
                    "race",
                    models.CharField(
                        blank=True,
                        help_text="Race/Ethnicity",
                        max_length=100,
                        verbose_name="Race/Ethnicity",
                    ),
                ),
                (
                    "sex",
                    utils.models.ChoiceField(
                        blank=True,
                        choices=[
                            ("male", "Male"),
                            ("female", "Female"),
                            ("non-binary", "Non-binary"),
                            ("other", "Other"),
                        ],
                        help_text="Sex/Gender of the worker.",
                        max_length=10,
                        verbose_name="Sex/Gender",
                    ),
                ),
                (
                    "date_of_birth",
                    models.DateField(
                        blank=True,
                        help_text="Date of Birth of the worker.",
                        null=True,
                        verbose_name="Date of Birth",
                    ),
                ),
                (
                    "license_number",
                    models.CharField(
                        blank=True,
                        help_text="License Number.",
                        max_length=20,
                        verbose_name="License Number",
                    ),
                ),
                (
                    "license_state",
                    localflavor.us.models.USStateField(
                        blank=True,
                        help_text="License State.",
                        max_length=2,
                        null=True,
                        verbose_name="License State",
                    ),
                ),
                (
                    "license_expiration_date",
                    models.DateField(
                        blank=True,
                        help_text="License Expiration Date.",
                        null=True,
                        verbose_name="License Expiration Date",
                    ),
                ),
                (
                    "endorsements",
                    utils.models.ChoiceField(
                        blank=True,
                        choices=[
                            ("N", "None"),
                            ("H", "Hazmat"),
                            ("T", "Tanker"),
                            ("X", "Tanker and Hazmat"),
                        ],
                        default="N",
                        help_text="Endorsements.",
                        max_length=1,
                        verbose_name="Endorsements",
                    ),
                ),
                (
                    "hazmat_expiration_date",
                    models.DateField(
                        blank=True,
                        help_text="Hazmat Endorsement Expiration Date.",
                        null=True,
                        verbose_name="Hazmat Expiration Date",
                    ),
                ),
                (
                    "hm_126_expiration_date",
                    models.DateField(
                        blank=True,
                        help_text="HM126GF Training Expiration Date.",
                        null=True,
                        verbose_name="HM-126 Expiration Date",
                    ),
                ),
                (
                    "hire_date",
                    models.DateField(
                        blank=True,
                        default=django.utils.timezone.now,
                        help_text="Date of hire.",
                        null=True,
                        verbose_name="Hire Date",
                    ),
                ),
                (
                    "termination_date",
                    models.DateField(
                        blank=True,
                        help_text="Date of termination.",
                        null=True,
                        verbose_name="Termination Date",
                    ),
                ),
                (
                    "review_date",
                    models.DateField(
                        blank=True,
                        help_text="Next Review Date.",
                        null=True,
                        verbose_name="Review Date",
                    ),
                ),
                (
                    "physical_due_date",
                    models.DateField(
                        blank=True,
                        help_text="Next Physical Due Date.",
                        null=True,
                        verbose_name="Physical Due Date",
                    ),
                ),
                (
                    "mvr_due_date",
                    models.DateField(
                        blank=True,
                        help_text="Next MVR Due Date.",
                        null=True,
                        verbose_name="MVR Due Date",
                    ),
                ),
                (
                    "medical_cert_date",
                    models.DateField(
                        blank=True,
                        help_text="Medical Certification Expiration Date.",
                        null=True,
                        verbose_name="Medical Cert Date",
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
                "verbose_name": "Worker profile",
                "verbose_name_plural": "Worker profiles",
                "ordering": ["worker"],
            },
        ),
    ]
