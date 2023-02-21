# Generated by Django 4.1.7 on 2023-02-16 02:42

from django.db import migrations, models


class Migration(migrations.Migration):
    dependencies = [
        ("organization", "0010_tablechangealert_listener_name"),
    ]

    operations = [
        migrations.AlterField(
            model_name="tablechangealert",
            name="function_name",
            field=models.CharField(
                blank=True,
                help_text="The function name that the table change alert will use.",
                max_length=50,
                verbose_name="Function Name",
            ),
        ),
        migrations.AlterField(
            model_name="tablechangealert",
            name="listener_name",
            field=models.CharField(
                blank=True,
                help_text="The listener name that the table change alert will use.",
                max_length=50,
                verbose_name="Listener Name",
            ),
        ),
        migrations.AlterField(
            model_name="tablechangealert",
            name="trigger_name",
            field=models.CharField(
                blank=True,
                help_text="The trigger name that the table change alert will use.",
                max_length=50,
                verbose_name="Trigger Name",
            ),
        ),
    ]
