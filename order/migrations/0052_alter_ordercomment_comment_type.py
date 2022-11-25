# Generated by Django 4.1.3 on 2022-11-25 00:50

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    dependencies = [
        ("dispatch", "0008_commenttype"),
        ("order", "0051_alter_additionalcharge_charge_amount"),
    ]

    operations = [
        migrations.AlterField(
            model_name="ordercomment",
            name="comment_type",
            field=models.ForeignKey(
                help_text="Comment Type",
                on_delete=django.db.models.deletion.CASCADE,
                related_name="order_comments",
                related_query_name="order_comment",
                to="dispatch.commenttype",
                verbose_name="Comment Type",
            ),
        ),
    ]
