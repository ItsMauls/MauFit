"use client";
import { useState } from "react";
import useAuthGuard from "@/hooks/useAuthGuard";
import GlassCard from "@/components/GlassCard";
import Modal from "@/components/Modal";
import UserForm from "@/components/UserForm";
import LockerGrid from "@/components/LockerGrid";

const sidebarItems = [
  { label: "Tambah User", icon: "➕" },
  { label: "Absen", icon: "🕒" },
  { label: "Kelola User", icon: "👤" },
  { label: "Kelola Kelas", icon: "🏋️" },
  { label: "Broadcast", icon: "📢" },
  { label: "Kelola Locker", icon: "🔒" },
];

const modalTabs = [
  { key: "tambah-user", label: "Tambah User" },
  { key: "absensi", label: "Absen" },
  { key: "kelola-user", label: "Kelola User" },
  { key: "kelas", label: "Kelola Kelas" },
  { key: "broadcast", label: "Broadcast" },
  { key: "kelola-locker", label: "Kelola Locker" },
];

export default function Home() {
  useAuthGuard();
  const [modalOpen, setModalOpen] = useState(false);
  const [activeTab, setActiveTab] = useState("user");

  return (
    <main className="min-h-screen flex bg-gradient-to-br from-green-900 via-green-900/70 to-black">
      {/* Sidebar */}
      <aside className="w-64 min-h-screen bg-black/30 backdrop-blur-lg border-r border-white/10 flex flex-col p-6 gap-4">
        <div className="text-3xl font-bold text-white mb-8 tracking-wide">MauFit Admin</div>
        <nav className="flex flex-col gap-2">
          {sidebarItems.map((item) => (
            <button
              key={item.label}
              className="flex items-center gap-3 px-4 py-2 rounded-lg text-white hover:bg-white/10 transition font-medium text-lg"
            >
              <span className="text-xl">{item.icon}</span> {item.label}
            </button>
          ))}
        </nav>
      </aside>
      {/* Main Content */}
      <section className="flex-1 flex flex-col items-center justify-center p-8">
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8 w-full max-w-5xl">
          <GlassCard className="flex flex-col items-center justify-center min-h-[180px] cursor-pointer hover:scale-105" onClick={() => { setModalOpen(true); setActiveTab("tambah-user"); }}>
            <div className="text-4xl mb-2">➕</div>
            <div className="text-xl font-semibold text-white mb-1">Tambah User</div>
            <div className="text-gray-300 text-sm">Tambah member baru</div>
          </GlassCard>
          <GlassCard className="flex flex-col items-center justify-center min-h-[180px] cursor-pointer hover:scale-105" onClick={() => { setModalOpen(true); setActiveTab("absensi"); }}>
            <div className="text-4xl mb-2">🕒</div>
            <div className="text-xl font-semibold text-white mb-1">Absen</div>
            <div className="text-gray-300 text-sm">Kelola absensi member</div>
          </GlassCard>
          <GlassCard className="flex flex-col items-center justify-center min-h-[180px] cursor-pointer hover:scale-105" onClick={() => { setModalOpen(true); setActiveTab("kelola-user"); }}>
            <div className="text-4xl mb-2">👤</div>
            <div className="text-xl font-semibold text-white mb-1">Kelola User</div>
            <div className="text-gray-300 text-sm">Edit & hapus user</div>
          </GlassCard>
          <GlassCard className="flex flex-col items-center justify-center min-h-[180px] cursor-pointer hover:scale-105" onClick={() => { setModalOpen(true); setActiveTab("kelas"); }}>
            <div className="text-4xl mb-2">🏋️</div>
            <div className="text-xl font-semibold text-white mb-1">Kelola Kelas</div>
            <div className="text-gray-300 text-sm">Atur jadwal kelas</div>
          </GlassCard>
          <GlassCard className="flex flex-col items-center justify-center min-h-[180px] cursor-pointer hover:scale-105" onClick={() => { setModalOpen(true); setActiveTab("broadcast"); }}>
            <div className="text-4xl mb-2">📢</div>
            <div className="text-xl font-semibold text-white mb-1">Broadcast</div>
            <div className="text-gray-300 text-sm">Kirim pengumuman</div>
          </GlassCard>
          <GlassCard className="flex flex-col items-center justify-center min-h-[180px] cursor-pointer hover:scale-105" onClick={() => { setModalOpen(true); setActiveTab("kelola-locker"); }}>
            <div className="text-4xl mb-2">🔒</div>
            <div className="text-xl font-semibold text-white mb-1">Kelola Locker</div>
            <div className="text-gray-300 text-sm">Atur & monitor locker</div>
          </GlassCard>
        </div>
      </section>
      <Modal isOpen={modalOpen} onClose={() => setModalOpen(false)}>
        <div className="flex flex-col w-[90vw] max-w-2xl min-h-[400px]">
          {/* Tab Navigation */}
          <div className="flex gap-2 mb-6">
            {modalTabs.map(tab => (
              <button
                key={tab.key}
                className={`px-4 py-2 rounded-lg font-semibold transition-all ${activeTab === tab.key ? 'bg-green-500 text-black shadow' : 'bg-black/20 text-green-200 hover:bg-green-700/30'}`}
                onClick={() => setActiveTab(tab.key)}
              >
                {tab.label}
              </button>
            ))}
          </div>
          {/* Tab Content */}
          <div className="flex-1">
            {(activeTab === "tambah-user" || activeTab === "kelola-user") && (
              <UserForm />
            )}
            {activeTab === "absensi" && (
              <div> <span className="text-lg font-bold">Absen (Coming Soon)</span> </div>
            )}
            {activeTab === "kelas" && (
              <div> <span className="text-lg font-bold">Kelola Kelas (Coming Soon)</span> </div>
            )}
            {activeTab === "broadcast" && (
              <div> <span className="text-lg font-bold">Broadcast (Coming Soon)</span> </div>
            )}
            {activeTab === "kelola-locker" && (
              <LockerGrid />
            )}
          </div>
        </div>
      </Modal>
    </main>
  );
}
