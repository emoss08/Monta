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

import { useEffect } from "react";
import { clearAllCookies, WEB_SOCKET_URL } from "@/lib/utils";
import { createWebsocketManager } from "@/utils/websockets";
import { useAuthStore } from "@/stores/AuthStore";
import { useNavigate } from "react-router-dom";

interface LogoutListenerProps {
  userId: string;
}

export function LogoutListener({ userId }: LogoutListenerProps) {
  const webSocketManager = createWebsocketManager();
  const [, setIsAuthenticated] = useAuthStore((state) => [
    state.isAuthenticated,
    state.setIsAuthenticated,
  ]);
  const navigate = useNavigate();

  useEffect(() => {
    webSocketManager.connect(
      "session",
      `${WEB_SOCKET_URL}/session/${userId}/`,
      {
        onOpen: () => console.info("Connected to notifications websocket"),
        onMessage: (message) => {
          const data = JSON.parse(message.data);
          console.log(data);
          if (data.user_id === userId) {
            sessionStorage.clear();
            setIsAuthenticated(false);
            navigate("/login");
            clearAllCookies(); // Clear all cookies to prevent auto-login
            console.info("Logged out");
          }
        },
      }
    );
    // Close the connection when the component unmounts.
    return () => {
      webSocketManager.disconnect("session");
    };
  }, [userId]);

  return null;
}
