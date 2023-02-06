# Generated by Django 4.1.5 on 2023-01-27 05:35

from django.db import migrations, models


class Migration(migrations.Migration):
    dependencies = [
        ("billing", "0007_billingqueue_bill_date_billingqueue_bol_number_and_more"),
    ]

    operations = [
        migrations.AlterField(
            model_name="billingexception",
            name="exception_message",
            field=models.TextField(
                blank=True,
                default=1,
                help_text="Message for the billing exception",
                verbose_name="Exception Message",
            ),
            preserve_default=False,
        ),
        migrations.AlterField(
            model_name="billingqueue",
            name="commodity_descr",
            field=models.CharField(
                blank=True,
                default=1,
                max_length=255,
                verbose_name="Commodity Description",
            ),
            preserve_default=False,
        ),
        migrations.AlterField(
            model_name="billingqueue",
            name="invoice_number",
            field=models.CharField(
                blank=True,
                default=1,
                help_text="Invoice number for the billing queue",
                max_length=50,
                verbose_name="Invoice Number",
            ),
            preserve_default=False,
        ),
    ]
