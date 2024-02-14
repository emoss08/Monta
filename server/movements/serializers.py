# --------------------------------------------------------------------------------------------------
#  COPYRIGHT(c) 2024 Trenova                                                                       -
#                                                                                                  -
#  This file is part of Trenova.                                                                   -
#                                                                                                  -
#  The Trenova software is licensed under the Business Source License 1.1. You are granted the right
#  to copy, modify, and redistribute the software, but only for non-production use or with a total -
#  of less than three server instances. Starting from the Change Date (November 16, 2026), the     -
#  software will be made available under version 2 or later of the GNU General Public License.     -
#  If you use the software in violation of this license, your rights under the license will be     -
#  terminated automatically. The software is provided "as is," and the Licensor disclaims all      -
#  warranties and conditions. If you use this license's text or the "Business Source License" name -
#  and trademark, you must comply with the Licensor's covenants, which include specifying the      -
#  Change License as the GPL Version 2.0 or a compatible license, specifying an Additional Use     -
#  Grant, and not modifying the license in any other way.                                          -
# --------------------------------------------------------------------------------------------------


from movements import models
from utils.serializers import GenericSerializer
from stops.serializers import StopSerializer


class MovementSerializer(GenericSerializer):
    """A serializer for the `Movement` model.

    A serializer class for the Movement Model. This serializer is used
    to convert the Movement model instances into a Python dictionary
    format that can be rendered into a JSON response. It also defines the fields
    that should be included in the serialized representation of the model.
    """

    stops = StopSerializer(
        many=True,
        required=False,
    )

    class Meta:
        """Metaclass for Movement Serializer.

        Attributes:
            model (models.Shipment): The model that the serializer is for.
        """

        model = models.Movement
        fields = "__all__"
        read_only_fields = (
            "id",
            "ref_num",
            "organization",
            "business_unit",
        )
        extra_kwargs = {
            "organization": {"required": False},
            "business_unit": {"required": False},
        }
