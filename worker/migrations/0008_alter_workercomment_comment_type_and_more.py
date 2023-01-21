# Generated by Django 4.1.5 on 2023-01-21 07:24

import django.db.models.deletion
from django.conf import settings
from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        migrations.swappable_dependency(settings.AUTH_USER_MODEL),
        ("dispatch", "0001_initial"),
        ("worker", "0007_alter_workerprofile_date_of_birth"),
    ]

    operations = [
        migrations.AlterField(
            model_name="workercomment",
            name="comment_type",
            field=models.ForeignKey(
                help_text="Related comment type.",
                on_delete=django.db.models.deletion.CASCADE,
                related_name="comment",
                related_query_name="comments",
                to="dispatch.commenttype",
                verbose_name="Comment Type",
            ),
        ),
        migrations.AlterField(
            model_name="workercomment",
            name="entered_by",
            field=models.ForeignKey(
                help_text="User who entered the comment.",
                on_delete=django.db.models.deletion.CASCADE,
                related_name="comment",
                related_query_name="comments",
                to=settings.AUTH_USER_MODEL,
                verbose_name="Entered By",
            ),
        ),
        migrations.AlterField(
            model_name="workercomment",
            name="worker",
            field=models.ForeignKey(
                help_text="Related worker.",
                on_delete=django.db.models.deletion.CASCADE,
                related_name="comment",
                related_query_name="comments",
                to="worker.worker",
                verbose_name="worker",
            ),
        ),
    ]
