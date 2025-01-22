import { AsyncSelectField } from "@/components/fields/async-select";
import { CheckboxField } from "@/components/fields/checkbox-field";
import { InputField } from "@/components/fields/input-field";
import { SelectField } from "@/components/fields/select-field";
import { TextareaField } from "@/components/fields/textarea-field";
import { FormControl, FormGroup } from "@/components/ui/form";
import { statusChoices } from "@/lib/choices";
import { type CommoditySchema } from "@/lib/schemas/commodity-schema";
import { useEffect } from "react";
import { useFormContext } from "react-hook-form";

export function CommodityForm() {
  const { control, setValue, watch } = useFormContext<CommoditySchema>();

  // Watch hazardousMaterialId specifically for changes
  const hazardousMaterialId = watch("hazardousMaterialId");

  useEffect(() => {
    // Update isHazardous when hazardousMaterialId changes
    setValue("isHazardous", Boolean(hazardousMaterialId));
  }, [hazardousMaterialId, setValue]);

  return (
    <FormGroup className="gap-y-3" cols={2}>
      <FormControl>
        <SelectField
          control={control}
          rules={{ required: true }}
          name="status"
          label="Status"
          placeholder="Status"
          description="Indicates whether the commodity is Active, Inactive, or Archived."
          options={statusChoices}
        />
      </FormControl>
      <FormControl>
        <InputField
          control={control}
          rules={{ required: true }}
          name="name"
          label="Name"
          placeholder="Name"
          description="The official name used to identify the commodity."
        />
      </FormControl>
      <FormControl cols="full">
        <TextareaField
          control={control}
          rules={{ required: true }}
          name="description"
          label="Description"
          placeholder="Description"
          description="Detailed information about the commodity's characteristics and handling requirements."
        />
      </FormControl>
      <FormControl>
        <AsyncSelectField
          isClearable
          name="hazardousMaterialId"
          control={control}
          link="/hazardous-materials"
          label="Hazardous Material"
          placeholder="Select Hazardous Material"
          description="Select the hazardous material classification if this commodity contains regulated substances."
          hasPopoutWindow
          popoutLink="/shipments/configurations/hazardous-materials/"
          popoutLinkLabel="Hazardous Material"
        />
      </FormControl>
      <FormControl>
        <CheckboxField
          name="isHazardous"
          control={control}
          label="Is Hazardous"
          outlined
          description="Indicates if the commodity is classified as hazardous."
        />
      </FormControl>
      <FormControl>
        <InputField
          name="minTemperature"
          control={control}
          label="Min Temperature"
          placeholder="Min Temperature"
          type="number"
          description="The lowest temperature (°F) at which the commodity can be safely transported."
        />
      </FormControl>
      <FormControl>
        <InputField
          name="maxTemperature"
          control={control}
          label="Max Temperature"
          placeholder="Max Temperature"
          type="number"
          description="The highest temperature (°F) at which the commodity can be safely transported."
        />
      </FormControl>
      <FormControl>
        <InputField
          name="weightPerUnit"
          control={control}
          label="Weight Per Unit"
          placeholder="Weight Per Unit"
          type="number"
          description="The weight (in pounds) of a single unit of this commodity. Used for calculating total load weight."
        />
      </FormControl>
      <FormControl>
        <InputField
          name="freightClass"
          control={control}
          label="Freight Class"
          placeholder="Freight Class"
          description="The NMFC code used for pricing and handling in LTL shipping."
        />
      </FormControl>
      <FormControl cols="full">
        <InputField
          name="dotClassification"
          control={control}
          label="DOT Classification"
          placeholder="DOT Classification"
          description="The U.S. Department of Transportation classification used for regulatory compliance."
        />
      </FormControl>
      <FormControl>
        <CheckboxField
          name="stackable"
          control={control}
          label="Stackable"
          outlined
          description="Indicates if the commodity can be safely stacked during transport or storage."
        />
      </FormControl>
      <FormControl>
        <CheckboxField
          name="fragile"
          control={control}
          label="Fragile"
          outlined
          description="Specifies whether the commodity is fragile and requires special handling."
        />
      </FormControl>
    </FormGroup>
  );
}