# Generated by Django 4.1.7 on 2023-02-20 20:07

from django.db import migrations, models


class Migration(migrations.Migration):
    dependencies = [
        ("dispatch", "0003_alter_commenttype_table_alter_delaycode_table_and_more"),
    ]

    operations = [
        migrations.AddField(
            model_name="dispatchcontrol",
            name="enforce_driver_ta",
            field=models.BooleanField(
                default=True,
                help_text="Disallow assignments if the driver is on Time Away",
                verbose_name="Enforce Driver Time Away",
            ),
        ),
    ]
