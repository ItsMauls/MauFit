// src/utils/api.ts
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