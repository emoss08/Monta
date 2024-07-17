/**
 * Copyright (c) 2024 Trenova Technologies, LLC
 *
 * Licensed under the Business Source License 1.1 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://trenova.app/pricing/
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *
 * Key Terms:
 * - Non-production use only
 * - Change Date: 2026-11-16
 * - Change License: GNU General Public License v2 or later
 *
 * For full license text, see the LICENSE file in the root directory.
 */

import { upperFirst } from "@/lib/utils";
import { routes } from "@/routing/AppRoutes";
import { useBreadcrumbStore } from "@/stores/BreadcrumbStore";
import { pathToRegexp } from "path-to-regexp";
import { useEffect, useMemo } from "react";
import { useLocation } from "react-router-dom";
import { Skeleton } from "../ui/skeleton";
import { FavoriteIcon } from "./user-favorite";

const useRouteMatching = (
  setLoading: (loading: boolean) => void,
  setCurrentRoute: (route: any) => void,
) => {
  const location = useLocation();

  useEffect(() => {
    setLoading(true);
    const matchedRoute = routes.find((route) => {
      return (
        route.path !== "*" && pathToRegexp(route.path).test(location.pathname)
      );
    });

    setCurrentRoute(matchedRoute || null);
    setLoading(false);
  }, [location, setLoading, setCurrentRoute]);
};

const useDocumentTitle = (currentRoute: any) => {
  useEffect(() => {
    if (currentRoute) {
      document.title = currentRoute.title;
    }
  }, [currentRoute]);
};

export function Breadcrumb({ children }: { children?: React.ReactNode }) {
  const [currentRoute, setCurrentRoute] =
    useBreadcrumbStore.use("currentRoute");
  const [loading, setLoading] = useBreadcrumbStore.use("loading");
  useRouteMatching(setLoading, setCurrentRoute);
  useDocumentTitle(currentRoute);

  // Construct breadcrumb text
  const breadcrumbText = useMemo(() => {
    if (!currentRoute) return "";
    const parts = [currentRoute.group, currentRoute.subMenu, currentRoute.title]
      .filter((str): str is string => str !== undefined)
      .map(upperFirst);
    return parts.join(" - ");
  }, [currentRoute]);

  if (loading) {
    return (
      <>
        <Skeleton className="h-[30px] w-[200px]" />
        <Skeleton className="mt-5 h-[30px] w-[200px]" />
      </>
    );
  }

  // If the current route is not found or is an excluded path, return null
  if (!currentRoute) {
    return null;
  }

  return (
    <div className="pb-4 pt-5 md:py-4">
      <div>
        <h2 className="mt-10 flex scroll-m-20 items-center pb-2 text-xl font-semibold tracking-tight transition-colors first:mt-0">
          {currentRoute?.title}
          <FavoriteIcon />
        </h2>
        <div className="flex items-center">
          <a className="text-muted-foreground hover:text-muted-foreground/80 text-sm font-medium">
            {breadcrumbText}
          </a>
        </div>
      </div>
      <div className="mt-3 flex">{children}</div>
    </div>
  );
}
