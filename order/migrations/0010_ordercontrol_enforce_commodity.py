# Generated by Django 4.1.5 on 2023-01-20 04:48

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("order", "0009_remove_ordercontrol_enforce_customer_and_more"),
    ]

    operations = [
        migrations.AddField(
            model_name="ordercontrol",
            name="enforce_commodity",
            field=models.BooleanField(
                default=False,
                help_text="Enforce the commodity input on the entry of an order.",
                verbose_name="Enforce Commodity Code",
            ),
        ),
    ]
