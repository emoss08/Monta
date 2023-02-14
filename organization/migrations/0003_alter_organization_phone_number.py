# Generated by Django 4.1.6 on 2023-02-14 08:18

from django.db import migrations
import phonenumber_field.modelfields


class Migration(migrations.Migration):

    dependencies = [
        ("organization", "0002_alter_emailprofile_protocol_taxrate"),
    ]

    operations = [
        migrations.AlterField(
            model_name="organization",
            name="phone_number",
            field=phonenumber_field.modelfields.PhoneNumberField(
                blank=True,
                help_text="The phone number of the organization.",
                max_length=128,
                region="US",
                verbose_name="Phone Number",
            ),
        ),
    ]
