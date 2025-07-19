"use client";
import { useState, useRef } from "react";
import { createUserApi, createUserProfileApi } from "@/utils/api";

export default function UserForm() {
  // Step 1: User
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  // Step 2: Profile
  const [address, setAddress] = useState("");
  const [phone, setPhone] = useState("");
  const [bio, setBio] = useState("");
  const [photoProfileUrl, setPhotoProfileUrl] = useState("");
  const [uploading, setUploading] = useState(false);
  const [photoPreview, setPhotoPreview] = useState<string | null>(null);

  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [step, setStep] = useState<"user" | "profile">("user");
  const [createdUserId, setCreatedUserId] = useState<number | null>(null);

  const fileInputRef = useRef<HTMLInputElement>(null);

  // Step 1: Register user
  const handleUserSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setMessage(null);
    setError(null);
    try {
      const res = await createUserApi({ name, email });
      setMessage("User berhasil didaftarkan! Lanjutkan melengkapi profil.");
      setCreatedUserId(res.data.id || res.data.user?.id); // backend response
      setStep("profile");
    } catch (err: any) {
      setError(err.message || "Gagal mendaftarkan user.");
    } finally {
      setLoading(false);
    }
  };

  // Step 2: Complete profile
  const handleProfileSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!createdUserId) return;
    setLoading(true);
    setMessage(null);
    setError(null);
    try {
      await createUserProfileApi(createdUserId, {
        address,
        phone,
        bio,
        photo_profile_url: photoProfileUrl,
      });
      console.log(photoProfileUrl, 'tesss');
      
      setMessage("Profil user berhasil dilengkapi!");
      setStep("user");
      setName(""); setEmail(""); setAddress(""); setPhone(""); setBio(""); setPhotoProfileUrl("");
      setPhotoPreview(null);
      setCreatedUserId(null);
      if (fileInputRef.current) fileInputRef.current.value = "";
    } catch (err: any) {
      setError(err.message || "Gagal melengkapi profil user.");
    } finally {
      setLoading(false);
    }
  };

  const handlePhotoUpload = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;

    setUploading(true);
    setPhotoPreview(URL.createObjectURL(file));

    const formData = new FormData();
    formData.append("file", file);
    formData.append("upload_preset", process.env.NEXT_PUBLIC_CLOUDINARY_UPLOAD_PRESET!);

    try {
      const res = await fetch(`https://api.cloudinary.com/v1_1/${process.env.NEXT_PUBLIC_CLOUDINARY_CLOUD_NAME}/image/upload`, {
        method: "POST",
        body: formData,
      });
      const data = await res.json();
      console.log(data, 'data cloudinary');
      
      setPhotoProfileUrl(data.secure_url);
    } catch (err) {
      setError("Gagal mengunggah foto.");
    } finally {
      setUploading(false);
    }
  };

  return (
    <div>
      {step === "user" && (
        <form onSubmit={handleUserSubmit} className="space-y-6">
          <div>
            <label className="block text-sm font-medium text-green-200 mb-1.5">Nama Lengkap</label>
            <input
              type="text"
              className="w-full bg-black/30 rounded-lg px-4 py-2.5 text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-green-500 transition-all duration-300"
              placeholder="Nama lengkap member"
              value={name}
              onChange={e => setName(e.target.value)}
              required
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-green-200 mb-1.5">Email</label>
            <input
              type="email"
              className="w-full bg-black/30 rounded-lg px-4 py-2.5 text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-green-500 transition-all duration-300"
              placeholder="Email member"
              value={email}
              onChange={e => setEmail(e.target.value)}
              required
            />
          </div>
          {message && <div className="text-green-400 text-sm text-center p-2 rounded-md bg-black/20">{message}</div>}
          {error && <div className="text-red-400 text-sm text-center p-2 rounded-md bg-black/20">{error}</div>}
          <button
            type="submit"
            className="w-full bg-green-500 text-black font-bold py-2.5 px-4 rounded-lg hover:bg-green-400 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-black focus:ring-green-400 transition-colors duration-300 shadow-[0_0_15px_rgba(0,255,102,0.5)] hover:shadow-[0_0_25px_rgba(0,255,102,0.7)]"
            disabled={loading}
          >
            {loading ? "Mendaftarkan..." : "Daftarkan User"}
          </button>
        </form>
      )}
      {step === "profile" && (
        <form onSubmit={handleProfileSubmit} className="space-y-6">
          <div>
            <label className="block text-sm font-medium text-green-200 mb-1.5">Alamat</label>
            <input
              type="text"
              className="w-full bg-black/30 rounded-lg px-4 py-2.5 text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-green-500 transition-all duration-300"
              placeholder="Alamat member"
              value={address}
              onChange={e => setAddress(e.target.value)}
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-green-200 mb-1.5">No. Telepon</label>
            <input
              type="text"
              className="w-full bg-black/30 rounded-lg px-4 py-2.5 text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-green-500 transition-all duration-300"
              placeholder="Nomor telepon member"
              value={phone}
              onChange={e => setPhone(e.target.value)}
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-green-200 mb-1.5">Bio</label>
            <textarea
              className="w-full bg-black/30 rounded-lg px-4 py-2.5 text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-green-500 transition-all duration-300"
              placeholder="Bio singkat member"
              value={bio}
              onChange={e => setBio(e.target.value)}
              rows={2}
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-green-200 mb-1.5">Foto Profil</label>
            <input
              type="file"
              accept="image/*"
              onChange={handlePhotoUpload}
              className="w-full text-sm text-gray-400 file:mr-4 file:py-2 file:px-4 file:rounded-lg file:border-0 file:text-sm file:font-semibold file:bg-green-500 file:text-black hover:file:bg-green-400"
              ref={fileInputRef}
            />
            {uploading && <p className="text-sm text-green-300 mt-2">Mengunggah...</p>}
            {photoPreview && !uploading && (
              <div className="mt-4">
                <img src={photoPreview} alt="Preview" className="w-32 h-32 rounded-full object-cover mx-auto" />
              </div>
            )}
          </div>
          {message && <div className="text-green-400 text-sm text-center p-2 rounded-md bg-black/20">{message}</div>}
          {error && <div className="text-red-400 text-sm text-center p-2 rounded-md bg-black/20">{error}</div>}
          <button
            type="submit"
            className="w-full bg-green-500 text-black font-bold py-2.5 px-4 rounded-lg hover:bg-green-400 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-black focus:ring-green-400 transition-colors duration-300 shadow-[0_0_15px_rgba(0,255,102,0.5)] hover:shadow-[0_0_25px_rgba(0,255,102,0.7)]"
            disabled={loading || uploading}
          >
            {loading ? "Menyimpan..." : "Simpan Profil"}
          </button>
        </form>
      )}
    </div>
  );
}