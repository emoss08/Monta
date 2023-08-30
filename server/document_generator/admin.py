# --------------------------------------------------------------------------------------------------
#  COPYRIGHT(c) 2023 MONTA                                                                         -
#                                                                                                  -
#  This file is part of Monta.                                                                     -
#                                                                                                  -
#  The Monta software is licensed under the Business Source License 1.1. You are granted the right -
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

from django.contrib import admin

from document_generator import models
from utils.admin import GenericAdmin, GenericStackedInline, GenericTabularInline


class TemplateFieldInline(
    GenericTabularInline[models.DocumentTemplate, models.TemplateField]
):
    model: type[models.TemplateField] = models.TemplateField


class DocumentDataBindingInline(
    GenericTabularInline[models.DocumentTemplate, models.DocumentDataBinding]
):
    model: type[models.DocumentDataBinding] = models.DocumentDataBinding


class DocumentTableColumnBindingInline(
    GenericTabularInline[models.DocumentTemplate, models.DocumentTableColumnBinding]
):
    model: type[models.DocumentTableColumnBinding] = models.DocumentTableColumnBinding


class DocTemplateCustomizationInline(
    GenericStackedInline[models.DocTemplateCustomization, models.DocumentTemplate]
):
    model: type[models.DocTemplateCustomization] = models.DocTemplateCustomization


@admin.register(models.DocumentTemplate)
class DocumentTemplateAdmin(GenericAdmin[models.DocumentTemplate]):
    list_display = ("name",)
    search_fields = ("name",)

    inlines = [
        TemplateFieldInline,
        DocTemplateCustomizationInline,
        DocumentDataBindingInline,
    ]


@admin.register(models.DocumentTheme)
class DocumentThemeAdmin(GenericAdmin[models.DocumentTheme]):
    list_display = ("name",)
    search_fields = ("name",)


@admin.register(models.DocumentTemplateVersion)
class DocumentTemplateVersionAdmin(GenericAdmin[models.DocumentTemplateVersion]):
    list_display = ("template",)
    search_fields = ("version_number",)