// components/LoginForm.tsx
'use client'

import React, { useState } from "react";
import { loginApi } from "@/utils/api";
import { setAccessToken } from "@/utils/auth";

// Komponen Input yang reusable
const Input = (props: React.InputHTMLAttributes<HTMLInputElement>) => (
    <input
        {...props}
        className="w-full bg-black/30 rounded-lg px-4 py-2.5 text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-green-500 transition-all duration-300"
    />
);

// Komponen Button yang reusable
const Button = (props: React.ButtonHTMLAttributes<HTMLButtonElement>) => (
    <button
        {...props}
        className="w-full bg-green-500 text-black font-bold py-2.5 px-4 rounded-lg hover:bg-green-400 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-black focus:ring-green-400 transition-colors duration-300 shadow-[0_0_15px_rgba(0,255,102,0.5)] hover:shadow-[0_0_25px_rgba(0,255,102,0.7)]"
    >
        {props.children}
    </button>
);


export default function LoginForm() {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [isLoading, setIsLoading] = useState(false);
    const [message, setMessage] = useState('');


    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        if (!email || !password) {
            setMessage('email dan password harus diisi.');
            return;
        }
        setIsLoading(true);
        setMessage('Mencoba untuk login...');

        try {
            const data = await loginApi(email, password);
            
            setAccessToken(data.data.token);
            setMessage("Login berhasil!");
            // TODO: redirect ke halaman dashboard, dsb
        } catch (err: any) {
            setMessage(err.message || "Login gagal.");
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <div className="w-full">
            <h2 className="text-3xl font-bold text-white text-center">Login</h2>
            <p className="text-gray-300 mt-2 text-sm text-center mb-6">Selamat datang kembali! Silakan masuk.</p>
            <form onSubmit={handleSubmit} className="space-y-6">
                <div className="space-y-4">
                    <div>
                        <label htmlFor="email" className="block text-sm font-medium text-gray-300 mb-1.5">Email</label>
                        <Input id="email" name="email" type="text" placeholder="Masukkan email Anda" value={email} onChange={(e) => setEmail(e.target.value)} required />
                    </div>
                    <div>
                        <label htmlFor="password" className="block text-sm font-medium text-gray-300 mb-1.5">Password</label>
                        <Input id="password" name="password" type="password" placeholder="••••••••" value={password} onChange={(e) => setPassword(e.target.value)} required />
                    </div>
                </div>
                {message && (<div className={`text-center text-sm p-2 rounded-md ${isLoading ? 'text-green-200' : 'text-green-300'}`}>{message}</div>)}
                <div className="pt-2">
                    <Button type="submit">{isLoading ? 'Loading...' : 'Login'}</Button>
                </div>
            </form>
        </div>
    );
}