import type { JobFunctionChoiceProps, TimezoneChoices } from "@/lib/choices";
import { type StatusChoiceProps } from "@/types/index";
import { type BaseModel } from "./organization";

interface UserRole extends BaseModel {
  name: string;
  description: string;
  edges: {
    permissions: UserPermission[];
  };
}

export interface UserPermission extends BaseModel {
  name: string;
  description: string;
  action: string;
  nameHumanized: string;
  resourceId: string;
}

export type UserFavorite = {
  id: string;
  userID: string;
  created: string;
  pageLink: string;
};

export type User = {
  id: string;
  businessUnitId: string;
  organizationId: string;
  username: string;
  name: string;
  email: string;
  isSuperAdmin: boolean;
  isAdmin: boolean;
  status: StatusChoiceProps;
  timezone: TimezoneChoices;
  phoneNumber?: string;
  profilePicUrl?: string;
  thumbnailUrl?: string;
  lastLogin?: string | null;
  edges: {
    roles?: UserRole[];
  };
};

export type UserFormValues = {
  organization: string;
  username: string;
  department?: string;
  email: string;
  isSuperAdmin: boolean;
};

export type JobTitle = {
  id: string;
  organization: string;
  name: string;
  description?: string | null;
  status: StatusChoiceProps;
  jobFunction: JobFunctionChoiceProps | "";
  created: string;
  modified: string;
};

export type JobTitleFormValues = Omit<
  JobTitle,
  "id" | "organization" | "created" | "modified"
>;

export type UserReport = {
  id: string;
  user: string;
  report: string;
  created: string;
  fileName: string;
  modified: string;
};

export type UserReportResponse = {
  count: number;
  next?: string | null;
  previous?: string | null;
  results: UserReport[];
};

export type Notification = {
  id: number;
  userID: string;
  isRead: boolean;
  title: string;
  description: string;
  actionUrl: string;
  createdAt: string;
};

export type UserNotification = {
  unreadCount: number;
  unreadList: Notification[];
};

export type GroupType = {
  id: string;
  name: string;
  codename: string;
};
