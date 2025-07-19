// src/utils/api.ts
import { mainApiUrl } from "@/constants";
import { getAccessToken } from "./auth";

export interface LoginResponse {
    access_token: string;
    [key: string]: any;
  }
  
  export async function loginApi(email: string, password: string): Promise<LoginResponse> {
    const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/users/login`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, password }),
    });
    
    if (!res.ok) {
      const err = await res.json();
      throw new Error(err.message || "Login gagal.");
    }
    return res.json();
  }

export interface CreateUserInput {
  name: string;
  email: string;
  fingerprint_id: string;
}

export interface CreateUserProfileInput {
  address?: string;
  phone?: string;
  bio?: string;
  photo_profile_url?: string;
}

export async function createUserApi(input: CreateUserInput): Promise<any> {
  // Generate random fingerprint_id (6 digit number as string)
  const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/users`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(input),
  });
  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.message || "Gagal menambah user.");
  }
  return res.json();
}

export async function createUserProfileApi(id: number, input: CreateUserProfileInput): Promise<any> {
  const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/users/${id}/profile`, {
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(input),
  });
  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.message || "Gagal melengkapi profil user.");
  }
  return res.json();
}

export interface Locker {
  id: number;
  locker_number: number;
  is_used: boolean;
}

export async function getAllLockersApi(): Promise<Locker[]> {
  const data = await fetchApi(mainApiUrl, "/lockers", { withAuth: true });
  // Sesuaikan jika response ada data wrapper
  return data.data || data;
}

export async function createAttendanceByFingerprintApi(fingerprintTemplate: string): Promise<any> {
 const data = await fetchApi(mainApiUrl, "/attendances/fingerprint", {
   method: "POST",
   body: { fingerprint_template: fingerprintTemplate },
   withAuth: true,
 });
 return data;
}

export interface FetchApiOptions {
  method?: string;
  headers?: Record<string, string>;
  body?: any;
  withAuth?: boolean;
}

export async function fetchApi<T = any>(
  baseUrl: string,
  endpoint: string,
  options: FetchApiOptions = {}
): Promise<T> {
  // const baseUrl = process.env.NEXT_PUBLIC_API_URL || "";
  console.log(options, 'opsi');
  
  const url = endpoint.startsWith("http") ? endpoint : `${baseUrl}${endpoint}`;
  const headers: Record<string, string> = {
    "Content-Type": "application/json",
    ...(options.headers || {}),
  };
  if (options.withAuth) {
    const token = getAccessToken();
    if (token) headers["Authorization"] = `Bearer ${token}`;
  }
  const fetchOptions: RequestInit = {
    method: options.method || "GET",
    headers,
  };
  if (options.body) {
    fetchOptions.body = typeof options.body === "string" ? options.body : JSON.stringify(options.body);
  }
  const res = await fetch(url, fetchOptions);
  if (!res.ok) {
    let errMsg = "Gagal fetch API";
    try { errMsg = (await res.json()).message || errMsg; } catch {}
    throw new Error(errMsg);
  }
  return res.json();
}