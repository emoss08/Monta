# Generated by Django 4.1.5 on 2023-01-06 21:08

from django.db import migrations


class Migration(migrations.Migration):
    dependencies = [
        ("accounts", "0007_remove_token_description_alter_userprofile_user"),
    ]

    operations = [
        migrations.RemoveField(
            model_name="userprofile",
            name="bio",
        ),
    ]
