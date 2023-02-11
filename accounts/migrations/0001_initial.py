# Generated by Django 4.1.6 on 2023-02-11 18:58

from django.conf import settings
import django.core.validators
from django.db import migrations, models
import django.db.models.deletion
import django.utils.timezone
import django_extensions.db.fields
import localflavor.us.models
import utils.validators
import uuid


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ("auth", "0012_alter_user_first_name_max_length"),
        ("organization", "0001_initial"),
    ]

    operations = [
        migrations.CreateModel(
            name="User",
            fields=[
                ("password", models.CharField(max_length=128, verbose_name="password")),
                (
                    "last_login",
                    models.DateTimeField(
                        blank=True, null=True, verbose_name="last login"
                    ),
                ),
                (
                    "is_superuser",
                    models.BooleanField(
                        default=False,
                        help_text="Designates that this user has all permissions without explicitly assigning them.",
                        verbose_name="superuser status",
                    ),
                ),
                (
                    "id",
                    models.UUIDField(
                        default=uuid.uuid4,
                        editable=False,
                        help_text="Unique ID for the user.",
                        primary_key=True,
                        serialize=False,
                        unique=True,
                    ),
                ),
                (
                    "username",
                    models.CharField(
                        help_text="Required. 30 characters or fewer. Letters, digits and @/./+/-/_ only.",
                        max_length=30,
                        unique=True,
                        verbose_name="Username",
                    ),
                ),
                (
                    "email",
                    models.EmailField(
                        help_text="Required. A valid email address.",
                        max_length=254,
                        unique=True,
                        verbose_name="Email Address",
                    ),
                ),
                (
                    "is_staff",
                    models.BooleanField(
                        default=False,
                        help_text="Designates whether the user can log into this admin site.",
                        verbose_name="Staff Status",
                    ),
                ),
                (
                    "date_joined",
                    models.DateTimeField(
                        default=django.utils.timezone.now, verbose_name="Date Joined"
                    ),
                ),
                (
                    "department",
                    models.ForeignKey(
                        blank=True,
                        null=True,
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="users",
                        related_query_name="user",
                        to="organization.department",
                        verbose_name="Department",
                    ),
                ),
                (
                    "groups",
                    models.ManyToManyField(
                        blank=True,
                        help_text="The groups this user belongs to. A user will get all permissions granted to each of their groups.",
                        related_name="user_set",
                        related_query_name="user",
                        to="auth.group",
                        verbose_name="groups",
                    ),
                ),
                (
                    "organization",
                    models.ForeignKey(
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="users",
                        related_query_name="user",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
                (
                    "user_permissions",
                    models.ManyToManyField(
                        blank=True,
                        help_text="Specific permissions for this user.",
                        related_name="user_set",
                        related_query_name="user",
                        to="auth.permission",
                        verbose_name="user permissions",
                    ),
                ),
            ],
            options={
                "verbose_name": "User",
                "verbose_name_plural": "Users",
                "ordering": ["-date_joined"],
            },
        ),
        migrations.CreateModel(
            name="JobTitle",
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
                        help_text="If the job title is active",
                        verbose_name="Is Active",
                    ),
                ),
                (
                    "name",
                    models.CharField(
                        help_text="Name of the job title",
                        max_length=100,
                        unique=True,
                        verbose_name="Name",
                    ),
                ),
                (
                    "description",
                    models.TextField(
                        blank=True,
                        help_text="Description of the job title",
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
                "verbose_name": "Job Title",
                "verbose_name_plural": "Job Titles",
                "ordering": ["name"],
            },
        ),
        migrations.CreateModel(
            name="UserProfile",
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
                    "first_name",
                    models.CharField(
                        help_text="The first name of the user",
                        max_length=255,
                        verbose_name="First Name",
                    ),
                ),
                (
                    "last_name",
                    models.CharField(
                        help_text="The last name of the user",
                        max_length=255,
                        verbose_name="Last Name",
                    ),
                ),
                (
                    "profile_picture",
                    models.ImageField(
                        blank=True,
                        help_text="The profile picture of the user",
                        null=True,
                        upload_to="profiles/",
                        validators=[
                            utils.validators.ImageSizeValidator(600, 600, False, True)
                        ],
                        verbose_name="Profile Picture",
                    ),
                ),
                (
                    "address_line_1",
                    models.CharField(
                        help_text="The address line 1 of the user",
                        max_length=100,
                        verbose_name="Address",
                    ),
                ),
                (
                    "address_line_2",
                    models.CharField(
                        blank=True,
                        help_text="The address line 2 of the user",
                        max_length=100,
                        verbose_name="Address Line 2",
                    ),
                ),
                (
                    "city",
                    models.CharField(
                        help_text="The city of the user",
                        max_length=100,
                        verbose_name="City",
                    ),
                ),
                (
                    "state",
                    localflavor.us.models.USStateField(
                        help_text="The state of the user",
                        max_length=2,
                        verbose_name="State",
                    ),
                ),
                (
                    "zip_code",
                    localflavor.us.models.USZipCodeField(
                        help_text="The zip code of the user",
                        max_length=10,
                        verbose_name="Zip Code",
                    ),
                ),
                (
                    "phone",
                    models.CharField(
                        blank=True,
                        help_text="The phone number of the user",
                        max_length=15,
                        verbose_name="Phone Number",
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
                    "title",
                    models.ForeignKey(
                        blank=True,
                        null=True,
                        on_delete=django.db.models.deletion.PROTECT,
                        related_name="profile",
                        related_query_name="profiles",
                        to="accounts.jobtitle",
                        verbose_name="Job Title",
                    ),
                ),
                (
                    "user",
                    models.OneToOneField(
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="profile",
                        to=settings.AUTH_USER_MODEL,
                        verbose_name="User",
                    ),
                ),
            ],
            options={
                "verbose_name": "Profile",
                "verbose_name_plural": "Profiles",
                "ordering": ["-created"],
            },
        ),
        migrations.CreateModel(
            name="Token",
            fields=[
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
                    "created",
                    models.DateTimeField(
                        auto_now_add=True,
                        help_text="The date and time the token was created",
                        verbose_name="Created",
                    ),
                ),
                (
                    "expires",
                    models.DateTimeField(
                        blank=True,
                        help_text="The date and time the token expires",
                        null=True,
                        verbose_name="Expires",
                    ),
                ),
                (
                    "last_used",
                    models.DateTimeField(
                        blank=True,
                        help_text="The date and time the token was last used",
                        null=True,
                        verbose_name="Last Used",
                    ),
                ),
                (
                    "key",
                    models.CharField(
                        max_length=40,
                        unique=True,
                        validators=[django.core.validators.MinLengthValidator(40)],
                    ),
                ),
                (
                    "user",
                    models.ForeignKey(
                        help_text="The user that the token belongs to",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="tokens",
                        related_query_name="token",
                        to=settings.AUTH_USER_MODEL,
                    ),
                ),
            ],
            options={
                "verbose_name": "Token",
                "verbose_name_plural": "Tokens",
            },
        ),
        migrations.AddIndex(
            model_name="userprofile",
            index=models.Index(
                fields=["-created"], name="accounts_us_created_2ea57d_idx"
            ),
        ),
    ]
