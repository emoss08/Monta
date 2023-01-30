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

from typing import Any
from collections.abc import Generator

import pytest

from accounting.tests.factories import RevenueCodeFactory
from billing.tests.factories import ChargeTypeFactory, DocumentClassificationFactory
from commodities.factories import CommodityFactory
from customer.factories import (
    CustomerBillingProfileFactory,
    CustomerContactFactory,
    CustomerFactory,
)
from order.tests.factories import OrderFactory, OrderTypeFactory
from organization.factories import EmailProfileFactory
from worker.factories import WorkerFactory

pytestmark = pytest.mark.django_db


@pytest.fixture
def document_classification():
    """
    Document classification fixture
    """
    yield DocumentClassificationFactory()


@pytest.fixture
def charge_type():
    """
    Charge type fixture
    """
    yield ChargeTypeFactory()


@pytest.fixture
def order_type():
    """
    Order Type Fixture
    """
    yield OrderTypeFactory()


@pytest.fixture
def order():
    """
    Order Fixture
    """
    yield OrderFactory()


@pytest.fixture
def revenue_code():
    """
    Revenue Code Fixture
    """
    yield RevenueCodeFactory()


@pytest.fixture
def customer():
    """
    Customer Fixture
    """
    yield CustomerFactory()


@pytest.fixture
def worker():
    """
    Worker Fixture
    """
    yield WorkerFactory()


@pytest.fixture
def commodity():
    """
    Commodity Fixture
    """
    yield CommodityFactory()


@pytest.fixture
def email_profile() -> Generator[Any, Any, None]:
    """
    Email Profile fixture
    """
    yield EmailProfileFactory()


@pytest.fixture
def order() -> Generator[Any, Any, None]:
    """
    Order fixture
    """
    yield OrderFactory()


@pytest.fixture
def customer() -> Generator[Any, Any, None]:
    """
    Customer fixture
    """
    yield CustomerFactory()


@pytest.fixture
def customer_contact() -> Generator[Any, Any, None]:
    """
    Customer Contact fixture
    """
    yield CustomerContactFactory()


@pytest.fixture
def customer_billing_profile() -> Generator[Any, Any, None]:
    """
    Customer Billing Profile fixture
    """
    yield CustomerBillingProfileFactory()
