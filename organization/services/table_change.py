"""
COPYRIGHT 2023 MONTA

This file is part of Monta.

Monta is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Monta is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Monta.  If not, see <https://www.gnu.org/licenses/>.
"""

from organization.models import TableChangeAlert
from organization.services.psql_triggers import (
    check_trigger_exists,
    create_insert_trigger,
    create_update_trigger,
    drop_trigger_and_function,
)

ACTION_NAMES = {
    "INSERT": {
        "function": "notify_new",
        "trigger": "after_insert",
        "listener": "new_added",
    },
    "UPDATE": {
        "function": "notify_updated",
        "trigger": "after_update",
        "listener": "updated",
    },
    "BOTH": {
        "function": "notify_new_or_updated",
        "trigger": "after_insert_or_update",
        "listener": "new_or_updated",
    },
}


def set_trigger_name_requirements(*, instance: TableChangeAlert) -> None:
    """Sets the name requirements for a table change alert instance.

    This function sets the name requirements for the provided table change alert instance
    based on the database action specified in the instance. The function, trigger, and
    listener names are all updated to reflect the appropriate names for the specified
    action. The updated instance is returned by reference.

    Args:
        instance (TableChangeAlert): The table change alert instance to update.

    Returns:
        None: This function does not return a value, but updates the provided instance
        by reference.

    Raises:
        KeyError: If the database action specified in the instance is not a valid action.
    """
    action_names = ACTION_NAMES[instance.database_action]

    instance.function_name = f"{action_names['function']}_{instance.table}"
    instance.trigger_name = f"{action_names['trigger']}_{instance.table}"
    instance.listener_name = f"{action_names['listener']}_{instance.table}"


def create_trigger_based_on_db_action(*, instance: TableChangeAlert) -> None:
    """
    Creates a trigger function and trigger for a PostgreSQL database table based on a specified database action.

    The function uses the `instance` argument to determine the type of database action that should be taken for a
    PostgreSQL table. It then calls the appropriate trigger function creation function (`create_insert_trigger`,
    `create_update_trigger`, or both) to create a trigger function that will listen for change events on the table
    and send notifications to a specified listener.

    Args:
        instance (TableChangeAlert): An instance of the `TableChangeAlert` model that specifies the properties of
            the trigger function and trigger to be created.

    Returns:
        None: This function has no return value.

    Raises:
        ValueError: If an invalid database action is specified in `instance`.

    """
    trigger_actions = {
        "INSERT": create_insert_trigger,
        "UPDATE": create_update_trigger,
        "BOTH": lambda **kwargs: (
            create_insert_trigger(**kwargs),
            create_update_trigger(**kwargs),
        ),
    }

    action = trigger_actions.get(instance.database_action)
    if action is None:
        raise ValueError(f"Invalid database action: {instance.database_action}")
    else:
        action(  # type: ignore
            trigger_name=instance.trigger_name,
            function_name=instance.function_name,
            listener_name=instance.listener_name,
            table_name=instance.table,
        )


def drop_trigger_and_create(*, instance: TableChangeAlert) -> None:
    """Drops the trigger and function associated with the given `TableChangeAlert` instance, if it exists.

    This function first checks if a trigger with the specified name exists for the given table, and if so,
    drops both the trigger and function associated with it. After dropping the trigger and function,
    a new trigger and function are created using the `create_trigger_based_on_db_action` function.

    Args:
        instance (TableChangeAlert): The `TableChangeAlert` instance associated with the trigger and function
        to be dropped and created.

    Returns:
        None: This function has no return value.

    """
    trigger_exists: bool = check_trigger_exists(
        table_name=instance.table, trigger_name=instance.trigger_name
    )

    if trigger_exists:
        drop_trigger_and_function(
            function_name=instance.function_name,
            table_name=instance.table,
            trigger_name=instance.trigger_name,
        )
    create_trigger_based_on_db_action(instance=instance)
