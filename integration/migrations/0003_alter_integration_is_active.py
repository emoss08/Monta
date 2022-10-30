# -*- coding: utf-8 -*-
# Generated by Django 4.1.2 on 2022-10-30 04:15

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("integration", "0002_alter_integration_auth_type_and_more"),
    ]

    operations = [
        migrations.AlterField(
            model_name="integration",
            name="is_active",
            field=models.BooleanField(
                default=True,
                help_text="Is the integration active?",
                verbose_name="Is Active",
            ),
        ),
    ]
