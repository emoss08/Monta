# Generated by Django 4.1.7 on 2023-03-07 01:19

import encrypted_model_fields.fields
from django.db import migrations


class Migration(migrations.Migration):
    dependencies = [
        ("worker", "0004_alter_workerprofile_license_number"),
    ]

    operations = [
        migrations.AlterField(
            model_name="workerprofile",
            name="license_number",
            field=encrypted_model_fields.fields.EncryptedCharField(
                blank=True,
                help_text="Driver License Number",
                verbose_name="License Number",
            ),
        ),
    ]
