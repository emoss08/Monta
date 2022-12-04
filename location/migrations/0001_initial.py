# Generated by Django 4.1.3 on 2022-12-04 00:03

from django.conf import settings
from django.db import migrations, models
import django.db.models.deletion
import django_extensions.db.fields
import localflavor.us.models
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
            name="Location",
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
                    "description",
                    models.TextField(
                        blank=True,
                        help_text="Location description",
                        verbose_name="Description",
                    ),
                ),
                (
                    "address_line_1",
                    models.CharField(
                        help_text="Address Line 1",
                        max_length=255,
                        verbose_name="Address Line 1",
                    ),
                ),
                (
                    "address_line_2",
                    models.CharField(
                        blank=True,
                        help_text="Address Line 2",
                        max_length=255,
                        verbose_name="Address Line 2",
                    ),
                ),
                (
                    "city",
                    models.CharField(
                        help_text="City", max_length=255, verbose_name="City"
                    ),
                ),
                (
                    "state",
                    localflavor.us.models.USStateField(
                        help_text="State", max_length=2, verbose_name="State"
                    ),
                ),
                (
                    "zip_code",
                    localflavor.us.models.USZipCodeField(
                        help_text="Zip Code", max_length=10, verbose_name="Zip Code"
                    ),
                ),
                (
                    "longitude",
                    models.FloatField(
                        blank=True,
                        help_text="Longitude",
                        null=True,
                        verbose_name="Longitude",
                    ),
                ),
                (
                    "latitude",
                    models.FloatField(
                        blank=True,
                        help_text="Latitude",
                        null=True,
                        verbose_name="Latitude",
                    ),
                ),
                (
                    "place_id",
                    models.CharField(
                        blank=True,
                        help_text="Place ID",
                        max_length=255,
                        verbose_name="Place ID",
                    ),
                ),
                (
                    "is_geocoded",
                    models.BooleanField(
                        default=False,
                        help_text="Is the location geocoded?",
                        verbose_name="Is Geocoded",
                    ),
                ),
                (
                    "depot",
                    models.ForeignKey(
                        blank=True,
                        help_text="Location depot",
                        null=True,
                        on_delete=django.db.models.deletion.PROTECT,
                        related_name="location",
                        related_query_name="locations",
                        to="organization.depot",
                        verbose_name="Depot",
                    ),
                ),
            ],
            options={
                "verbose_name": "Location",
                "verbose_name_plural": "Locations",
                "ordering": ("id",),
            },
        ),
        migrations.CreateModel(
            name="LocationContact",
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
                        verbose_name="Name",
                    ),
                ),
                (
                    "email",
                    models.EmailField(
                        blank=True,
                        help_text="Email of the contact.",
                        max_length=255,
                        verbose_name="Email",
                    ),
                ),
                (
                    "phone",
                    models.PositiveIntegerField(
                        blank=True,
                        help_text="Phone of the contact.",
                        null=True,
                        verbose_name="Phone",
                    ),
                ),
                (
                    "fax",
                    models.PositiveIntegerField(
                        blank=True,
                        help_text="Fax of the contact.",
                        null=True,
                        verbose_name="Fax",
                    ),
                ),
                (
                    "location",
                    models.ForeignKey(
                        help_text="Location",
                        on_delete=django.db.models.deletion.PROTECT,
                        related_name="location_contact",
                        related_query_name="location_contacts",
                        to="location.location",
                        verbose_name="Location",
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
                "verbose_name": "Location Contact",
                "verbose_name_plural": "Location Contacts",
                "ordering": ("name",),
            },
        ),
        migrations.CreateModel(
            name="LocationComment",
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
                        on_delete=django.db.models.deletion.PROTECT,
                        related_name="location_comments",
                        related_query_name="location_comment",
                        to="dispatch.commenttype",
                        verbose_name="Comment Type",
                    ),
                ),
                (
                    "entered_by",
                    models.ForeignKey(
                        on_delete=django.db.models.deletion.PROTECT,
                        related_name="location_comments",
                        related_query_name="location_comment",
                        to=settings.AUTH_USER_MODEL,
                        verbose_name="Entered By",
                    ),
                ),
                (
                    "location",
                    models.ForeignKey(
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="location_comments",
                        related_query_name="location_comment",
                        to="location.location",
                        verbose_name="Location",
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
                "verbose_name": "Location Comment",
                "verbose_name_plural": "Location Comments",
                "ordering": ("location",),
            },
        ),
        migrations.CreateModel(
            name="LocationCategory",
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
                    models.CharField(max_length=100, unique=True, verbose_name="Name"),
                ),
                (
                    "description",
                    models.TextField(blank=True, verbose_name="Description"),
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
                "verbose_name": "Location Category",
                "verbose_name_plural": "Location Categories",
                "ordering": ("name",),
            },
        ),
        migrations.AddField(
            model_name="location",
            name="location_category",
            field=models.ForeignKey(
                blank=True,
                help_text="Location category",
                null=True,
                on_delete=django.db.models.deletion.PROTECT,
                related_name="location",
                related_query_name="locations",
                to="location.locationcategory",
                verbose_name="Category",
            ),
        ),
        migrations.AddField(
            model_name="location",
            name="organization",
            field=models.ForeignKey(
                help_text="Organization",
                on_delete=django.db.models.deletion.CASCADE,
                related_name="%(class)ss",
                related_query_name="%(class)s",
                to="organization.organization",
                verbose_name="Organization",
            ),
        ),
        migrations.AddIndex(
            model_name="locationcontact",
            index=models.Index(fields=["name"], name="location_lo_name_81e740_idx"),
        ),
    ]
