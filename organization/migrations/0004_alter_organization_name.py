# Generated by Django 4.1.5 on 2023-01-19 02:35

from django.db import migrations, models


class Migration(migrations.Migration):
    dependencies = [
        ("organization", "0003_remove_organization_authentication_bg_and_more"),
    ]

    operations = [
        migrations.AlterField(
            model_name="organization",
            name="name",
            field=models.CharField(max_length=255, verbose_name="Organization Name"),
        ),
    ]
