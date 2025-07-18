// src/utils/auth.ts
import Cookies from "js-cookie";

export function setAccessToken(token: string) {   
  Cookies.set("access_token", token, { expires: 7 });
}

export function getAccessToken(): string | undefined {
  return Cookies.get("access_token");
}

export function removeAccessToken() {
  Cookies.remove("access_token");
}