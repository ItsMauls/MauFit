"use client";
import { useState } from "react";
import useAuthGuard from "@/hooks/useAuthGuard";
import GlassCard from "@/components/GlassCard";
import Modal from "@/components/Modal";
import UserForm from "@/components/UserForm";

const sidebarItems = [
  { label: "Dashboard", icon: "🏠" },
  { label: "User", icon: "👤" },
  { label: "Absensi", icon: "📝" },
  { label: "Locker", icon: "🔒" },
];

const modalTabs = [
  { key: "user", label: "Kelola User" },
  { key: "absensi", label: "Kelola Absensi" },
  { key: "locker", label: "Kelola Locker" },
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
          <GlassCard className="flex flex-col items-center justify-center min-h-[180px] cursor-pointer hover:scale-105" onClick={() => { setModalOpen(true); setActiveTab("user"); }}>
            <div className="text-4xl mb-2">👤</div>
            <div className="text-xl font-semibold text-white mb-1">Kelola User</div>
            <div className="text-gray-300 text-sm">Tambah, edit, dan hapus user</div>
          </GlassCard>
          <GlassCard className="flex flex-col items-center justify-center min-h-[180px] cursor-pointer hover:scale-105">
            <div className="text-4xl mb-2">📝</div>
            <div className="text-xl font-semibold text-white mb-1">Kelola Absensi</div>
            <div className="text-gray-300 text-sm">Lihat dan kelola data absensi</div>
          </GlassCard>
          <GlassCard className="flex flex-col items-center justify-center min-h-[180px] cursor-pointer hover:scale-105">
            <div className="text-4xl mb-2">🔒</div>
            <div className="text-xl font-semibold text-white mb-1">Kelola Locker</div>
            <div className="text-gray-300 text-sm">Atur dan monitor locker</div>
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
            {activeTab === "user" && (
              <UserForm />
            )}
            {activeTab === "absensi" && (
              <div> <span className="text-lg font-bold">Kelola Absensi (Coming Soon)</span> </div>
            )}
            {activeTab === "locker" && (
              <div> <span className="text-lg font-bold">Kelola Locker (Coming Soon)</span> </div>
            )}
          </div>
        </div>
      </Modal>
    </main>
  );
}
