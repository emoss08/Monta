# Generated by Django 4.1.5 on 2023-02-02 00:53

import django.db.models.deletion
from django.db import migrations, models


class Migration(migrations.Migration):
    dependencies = [
        ("customer", "0008_alter_customerruleprofile_document_class"),
    ]

    operations = [
        migrations.AlterField(
            model_name="customerbillingprofile",
            name="customer",
            field=models.OneToOneField(
                help_text="Customer",
                on_delete=django.db.models.deletion.CASCADE,
                related_name="billing_profile",
                to="customer.customer",
                verbose_name="Customer",
            ),
        ),
        migrations.AlterField(
            model_name="customerbillingprofile",
            name="email_profile",
            field=models.ForeignKey(
                blank=True,
                help_text="Customer Email Profile",
                null=True,
                on_delete=django.db.models.deletion.CASCADE,
                related_name="billing_profile",
                to="customer.customeremailprofile",
                verbose_name="Customer Email Profile",
            ),
        ),
        migrations.AlterField(
            model_name="customerbillingprofile",
            name="rule_profile",
            field=models.ForeignKey(
                blank=True,
                help_text="Rule Profile",
                null=True,
                on_delete=django.db.models.deletion.CASCADE,
                related_name="billing_profile",
                to="customer.customerruleprofile",
                verbose_name="Rule Profile",
            ),
        ),
    ]
