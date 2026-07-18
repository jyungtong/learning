"use client";

import {
  focusManager,
  QueryClient,
  QueryClientProvider,
} from "@tanstack/react-query";
import { useEffect, useState } from "react";

export function Providers({ children }: { children: React.ReactNode }) {
  const [queryClient] = useState(
    () =>
      new QueryClient({
        defaultOptions: {
          queries: { staleTime: 5_000, refetchOnWindowFocus: "always" },
        },
      }),
  );

  useEffect(
    () =>
      focusManager.setEventListener((onFocus) => {
        const handleFocus = () => onFocus();

        window.addEventListener("focus", handleFocus);
        document.addEventListener("visibilitychange", handleFocus);

        return () => {
          window.removeEventListener("focus", handleFocus);
          document.removeEventListener("visibilitychange", handleFocus);
        };
      }),
    [],
  );

  return (
    <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
  );
}
