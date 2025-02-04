import { type ChoiceProps, Gender, Status } from "@/types/common";
import { EquipmentClass } from "@/types/equipment-type";
import {
  HazardousClassChoiceProps,
  PackingGroupChoiceProps,
} from "@/types/hazardous-material";
import { FacilityType, LocationCategoryType } from "@/types/location-category";
import { EquipmentStatus } from "@/types/tractor";
import { Endorsement, PTOStatus, PTOType, WorkerType } from "@/types/worker";

/**
 * Returns status choices for a select input.
 * @returns An array of status choices.
 */
export const statusChoices = [
  { value: Status.Active, label: "Active", color: "#15803d" },
  { value: Status.Inactive, label: "Inactive", color: "#b91c1c" },
] satisfies ReadonlyArray<ChoiceProps<Status>>;

export const equipmentStatusChoices = [
  { value: EquipmentStatus.Available, label: "Available", color: "#15803d" },
  {
    value: EquipmentStatus.OOS,
    label: "Out of Service",
    color: "#b91c1c",
  },
  {
    value: EquipmentStatus.AtMaintenance,
    label: "At Maintenance",
    color: "#7e22ce",
  },
  { value: EquipmentStatus.Sold, label: "Sold", color: "#f59e0b" },
] satisfies ReadonlyArray<ChoiceProps<EquipmentStatus>>;

/**
 * Returns status choices for a select input.
 * @returns An array of status choices.
 */
export const workerTypeChoices = [
  { value: WorkerType.Employee, label: "Employee", color: "#15803d" },
  { value: WorkerType.Contractor, label: "Contractor", color: "#7e22ce" },
] satisfies ReadonlyArray<ChoiceProps<WorkerType>>;

export const endorsementChoices = [
  { value: Endorsement.None, label: "None", color: "#15803d" },
  { value: Endorsement.Tanker, label: "Tanker", color: "#7e22ce" },
  { value: Endorsement.Hazmat, label: "Hazmat", color: "#dc2626" },
  { value: Endorsement.TankerHazmat, label: "Tanker/Hazmat", color: "#f59e0b" },
  { value: Endorsement.Passenger, label: "Passenger", color: "#1d4ed8" },
  {
    value: Endorsement.DoublesTriples,
    label: "Doubles/Triples",
    color: "#0369a1",
  },
] satisfies ReadonlyArray<ChoiceProps<Endorsement>>;

export const equipmentClassChoices = [
  { value: EquipmentClass.Tractor, label: "Tractor", color: "#15803d" },
  { value: EquipmentClass.Trailer, label: "Trailer", color: "#7e22ce" },
  { value: EquipmentClass.Container, label: "Container", color: "#dc2626" },
  { value: EquipmentClass.Other, label: "Other", color: "#f59e0b" },
] satisfies ReadonlyArray<ChoiceProps<EquipmentClass>>;

export const genderChoices = [
  { value: Gender.Male, label: "Male", color: "#1d4ed8" },
  { value: Gender.Female, label: "Female", color: "#ec4899" },
] satisfies ReadonlyArray<ChoiceProps<Gender>>;

export const ptoStatusChoices = [
  { value: PTOStatus.Requested, label: "Requested", color: "#15803d" },
  { value: PTOStatus.Approved, label: "Approved", color: "#7e22ce" },
  { value: PTOStatus.Rejected, label: "Rejected", color: "#b91c1c" },
  { value: PTOStatus.Cancelled, label: "Cancelled", color: "#f59e0b" },
] satisfies ReadonlyArray<ChoiceProps<PTOStatus>>;

export const ptoTypeChoices = [
  { value: PTOType.Vacation, label: "Vacation", color: "#15803d" },
  { value: PTOType.Sick, label: "Sick", color: "#7e22ce" },
  { value: PTOType.Holiday, label: "Holiday", color: "#b91c1c" },
  { value: PTOType.Bereavement, label: "Bereavement", color: "#f59e0b" },
  { value: PTOType.Maternity, label: "Maternity", color: "#0369a1" },
  { value: PTOType.Paternity, label: "Paternity", color: "#0369a1" },
] satisfies ReadonlyArray<ChoiceProps<PTOType>>;

export const hazardousClassChoices = [
  {
    value: HazardousClassChoiceProps.HazardClass1And1,
    label: "Division 1.1: Mass Explosive Hazard",
  },
  {
    value: HazardousClassChoiceProps.HazardClass1And2,
    label: "Division 1.2: Projection Hazard",
  },
  {
    value: HazardousClassChoiceProps.HazardClass1And3,
    label: "Division 1.3: Fire and/or Minor Blast/Minor Projection Hazard",
  },
  {
    value: HazardousClassChoiceProps.HazardClass1And4,
    label: "Division 1.4: Minor Explosion Hazard",
  },
  {
    value: HazardousClassChoiceProps.HazardClass1And5,
    label: "Division 1.5: Very Insensitive With Mass Explosion Hazard",
  },
  {
    value: HazardousClassChoiceProps.HazardClass1And6,
    label: "Division 1.6: Extremely Insensitive; No Mass Explosion Hazard",
  },
  {
    value: HazardousClassChoiceProps.HazardClass2And1,
    label: "Division 2.1: Flammable Gases",
  },
  {
    value: HazardousClassChoiceProps.HazardClass2And2,
    label: "Division 2.2: Non-Flammable Gases",
  },
  {
    value: HazardousClassChoiceProps.HazardClass2And3,
    label: "Division 2.3: Poisonous Gases",
  },
  {
    value: HazardousClassChoiceProps.HazardClass3,
    label: "Division 3: Flammable Liquids",
  },
  {
    value: HazardousClassChoiceProps.HazardClass4And1,
    label: "Division 4.1: Flammable Solids",
  },
  {
    value: HazardousClassChoiceProps.HazardClass4And2,
    label: "Division 4.2: Spontaneously Combustible Solids",
  },
  {
    value: HazardousClassChoiceProps.HazardClass4And3,
    label: "Division 4.3: Dangerous When Wet",
  },
  {
    value: HazardousClassChoiceProps.HazardClass5And1,
    label: "Division 5.1: Oxidizing Substances",
  },
  {
    value: HazardousClassChoiceProps.HazardClass5And2,
    label: "Division 5.2: Organic Peroxides",
  },
  {
    value: HazardousClassChoiceProps.HazardClass6And1,
    label: "Division 6.1: Toxic Substances",
  },
  {
    value: HazardousClassChoiceProps.HazardClass6And2,
    label: "Division 6.2: Infectious Substances",
  },
  {
    value: HazardousClassChoiceProps.HazardClass7,
    label: "Division 7: Radioactive Material",
  },
  {
    value: HazardousClassChoiceProps.HazardClass8,
    label: "Division 8: Corrosive Substances",
  },
  {
    value: HazardousClassChoiceProps.HazardClass9,
    label: "Division 9: Miscellaneous Hazardous Substances and Articles",
  },
] satisfies ReadonlyArray<ChoiceProps<HazardousClassChoiceProps>>;

export const packingGroupChoices = [
  {
    value: PackingGroupChoiceProps.PackingGroupI,
    label: "I (High Danger)",
    color: "#b91c1c",
  },
  {
    value: PackingGroupChoiceProps.PackingGroupII,
    label: "II (Medium Danger)",
    color: "#ca8a04",
  },
  {
    value: PackingGroupChoiceProps.PackingGroupIII,
    label: "III (Low Danger)",
    color: "#16a34a",
  },
] satisfies ReadonlyArray<ChoiceProps<PackingGroupChoiceProps>>;

export const locationCategoryTypeChoices = [
  { value: LocationCategoryType.Terminal, label: "Terminal", color: "#15803d" },
  {
    value: LocationCategoryType.Warehouse,
    label: "Warehouse",
    color: "#7e22ce",
  },
  {
    value: LocationCategoryType.DistributionCenter,
    label: "Distribution Center",
    color: "#dc2626",
  },
  {
    value: LocationCategoryType.TruckStop,
    label: "Truck Stop",
    color: "#f59e0b",
  },
  {
    value: LocationCategoryType.RestArea,
    label: "Rest Area",
    color: "#0369a1",
  },
  {
    value: LocationCategoryType.CustomerLocation,
    label: "Customer Location",
    color: "#10b981",
  },
  { value: LocationCategoryType.Port, label: "Port", color: "#6366f1" },
  {
    value: LocationCategoryType.RailYard,
    label: "Rail Yard",
    color: "#ec4899",
  },
  {
    value: LocationCategoryType.MaintenanceFacility,
    label: "Maintenance Facility",
    color: "#14b8a6",
  },
] satisfies ReadonlyArray<ChoiceProps<LocationCategoryType>>;

export const facilityTypeChoices = [
  { value: FacilityType.CrossDock, label: "Cross Dock", color: "#7e22ce" },
  {
    value: FacilityType.StorageWarehouse,
    label: "Storage Warehouse",
    color: "#dc2626",
  },
  { value: FacilityType.ColdStorage, label: "Cold Storage", color: "#f59e0b" },
  {
    value: FacilityType.HazmatFacility,
    label: "Hazmat Facility",
    color: "#0369a1",
  },
  {
    value: FacilityType.IntermodalFacility,
    label: "Intermodal Facility",
    color: "#10b981",
  },
] satisfies ReadonlyArray<ChoiceProps<FacilityType>>;
