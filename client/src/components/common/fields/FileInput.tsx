/*
 * COPYRIGHT(c) 2023 MONTA
 *
 * This file is part of Monta.
 *
 * The Monta software is licensed under the Business Source License 1.1. You are granted the right
 * to copy, modify, and redistribute the software, but only for non-production use or with a total
 * of less than three server instances. Starting from the Change Date (November 16, 2026), the
 * software will be made available under version 2 or later of the GNU General Public License.
 * If you use the software in violation of this license, your rights under the license will be
 * terminated automatically. The software is provided "as is," and the Licensor disclaims all
 * warranties and conditions. If you use this license's text or the "Business Source License" name
 * and trademark, you must comply with the Licensor's covenants, which include specifying the
 * Change License as the GPL Version 2.0 or a compatible license, specifying an Additional Use
 * Grant, and not modifying the license in any other way.
 */

import React from "react";
import { Center, FileInput, FileInputProps, Group, rem } from "@mantine/core";
import { IconAlertTriangle, IconPhoto } from "@tabler/icons-react";
import { UseFormReturnType } from "@mantine/form";
import { useFormStyles } from "@/assets/styles/FormStyles";
import { InputFieldNameProp } from "@/types";

function Value({ file }: { file: File | string }) {
  let displayValue: string;

  if (file instanceof File) {
    // This is an actual file object
    displayValue = file.name;
  } else {
    // This is a URL string
    const urlParts = file.split("/");
    displayValue = urlParts[urlParts.length - 1] || "Unknown File"; // Taking the last part of the URL as the file name
  }

  return (
    <Center
      inline
      sx={(theme) => ({
        backgroundColor:
          theme.colorScheme === "dark"
            ? theme.colors.dark[7]
            : theme.colors.gray[2],
        fontSize: theme.fontSizes.xs,
        padding: `${rem(3)} ${rem(7)}`,
        borderRadius: theme.radius.sm,
      })}
    >
      <IconPhoto size={rem(14)} style={{ marginRight: rem(5) }} />
      <span
        style={{
          whiteSpace: "nowrap",
          textOverflow: "ellipsis",
          overflow: "hidden",
          maxWidth: rem(200),
          display: "inline-block",
        }}
      >
        {displayValue}
      </span>
    </Center>
  );
}

const ValueComponent: FileInputProps["valueComponent"] = ({ value }) => {
  if (Array.isArray(value)) {
    return (
      <Group spacing="sm" py="xs">
        {value.map(
          (file, _) => file && <Value file={file} key={`input-file-${file}`} />,
        )}
      </Group>
    );
  }

  return value && <Value file={value} />;
};

interface ValidatedFileInputProps<TFormValues>
  extends Omit<FileInputProps, "form" | "name"> {
  form: UseFormReturnType<TFormValues, (values: TFormValues) => TFormValues>;
  name: InputFieldNameProp<TFormValues>;
}

export function ValidatedFileInput<TFormValues extends object>({
  form,
  name,
  ...rest
}: ValidatedFileInputProps<TFormValues>) {
  const { classes } = useFormStyles();
  const error = form.errors[name as string];

  return (
    <FileInput
      {...rest}
      {...form.getInputProps(name as string)}
      error={error}
      variant="filled"
      className={classes.fields}
      placeholder={rest.placeholder || "Select file"}
      rightSection={
        error && (
          <IconAlertTriangle
            stroke={1.5}
            size="1.1rem"
            className={classes.invalidIcon}
          />
        )
      }
      valueComponent={ValueComponent}
    />
  );
}
