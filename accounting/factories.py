"""
COPYRIGHT 2022 MONTA

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

import factory


class GeneralLedgerAccountFactory(factory.django.DjangoModelFactory):
    """
    General Ledger Account factory
    """

    organization = factory.SubFactory("organization.factories.OrganizationFactory")
    account_number = factory.Faker(
        "numerify",
        text="####-####-####-####",
    )
    account_type = factory.Faker(
        "random_element",
        elements=("ASSET", "LIABILITY", "EQUITY", "REVENUE", "EXPENSE"),
    )

    class Meta:
        """
        Metaclass for General
        Ledger Account Factory
        """

        model = "accounting.GeneralLedgerAccount"


class RevenueCodeFactory(factory.django.DjangoModelFactory):
    """
    Revenue Code factory
    """

    organization = factory.SubFactory("organization.factories.OrganizationFactory")
    code = factory.Faker("numerify", text="####")
    description = factory.Faker("text", locale="en_US", max_nb_chars=100)
    expense_account = factory.SubFactory(GeneralLedgerAccountFactory)
    revenue_account = factory.SubFactory(GeneralLedgerAccountFactory)

    class Meta:
        """
        Metaclass for Revenue
        Code Factory
        """

        model = "accounting.RevenueCode"
