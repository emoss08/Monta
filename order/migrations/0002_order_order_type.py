# Generated by Django 4.1.5 on 2023-01-10 06:57

import django.db.models.deletion
from django.db import migrations, models


class Migration(migrations.Migration):
    dependencies = [
        ("order", "0001_initial"),
    ]

    operations = [
        migrations.AddField(
            model_name="order",
            name="order_type",
            field=models.ForeignKey(
                default=1,
                help_text="Order Type of the Order",
                on_delete=django.db.models.deletion.PROTECT,
                related_name="orders",
                related_query_name="order",
                to="order.ordertype",
                verbose_name="Order Type",
            ),
            preserve_default=False,
        ),
    ]
