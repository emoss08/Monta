# Generated by Django 4.1.5 on 2023-02-01 06:01

import django.db.models.deletion
from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("organization", "0006_emaillog"),
    ]

    operations = [
        migrations.AlterField(
            model_name="emaillog",
            name="organization",
            field=models.ForeignKey(
                blank=True,
                help_text="The organization that the email log belongs to.",
                null=True,
                on_delete=django.db.models.deletion.CASCADE,
                related_name="email_logs",
                to="organization.organization",
                verbose_name="Organization",
            ),
        ),
    ]
