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

export interface CreateUserInput {
  name: string;
  email: string;
}

export interface CreateUserProfileInput {
  address?: string;
  phone?: string;
  bio?: string;
  photo_profile_url?: string;
}

export async function createUserApi(input: CreateUserInput): Promise<any> {
  // Generate random fingerprint_id (6 digit number as string)
  const fingerprint_id = Math.floor(100000 + Math.random() * 900000).toString();
  const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/users`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ name: input.name, email: input.email, fingerprint_id }),
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