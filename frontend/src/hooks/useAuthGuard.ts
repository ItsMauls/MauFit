"use client";
import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { getAccessToken } from "@/utils/auth";

/**
 * useAuthGuard
 * Redirects to /login if access_token is not present in cookies.
 * Use in protected pages/components.
 */
export default function useAuthGuard() {
  const router = useRouter();
  useEffect(() => {
    if (!getAccessToken()) {
      router.replace("/login");
    }
  }, [router]);
} 