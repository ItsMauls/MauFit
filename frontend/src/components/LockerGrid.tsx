import React, { useEffect, useState } from "react";
import { getAllLockersApi, Locker } from "@/utils/api";

export default function LockerGrid() {
  const [lockers, setLockers] = useState<Locker[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    setLoading(true);
    getAllLockersApi()
      .then((data) => {
        // Sort by locker_number ascending, limit 100
        setLockers(data.sort((a, b) => a.locker_number - b.locker_number).slice(0, 100));
        setError(null);
      })
      .catch((err) => setError(err.message))
      .finally(() => setLoading(false));
  }, []);

  if (loading) return <div className="text-green-200">Memuat data locker...</div>;
  if (error) return <div className="text-red-400">{error}</div>;

  return (
    <div>
      <h3 className="text-lg font-bold mb-4 text-white">Daftar Locker</h3>
      <div className="grid grid-cols-5 sm:grid-cols-10 gap-3">
        {lockers.map((locker) => (
          <div
            key={locker.locker_number}
            className={`flex items-center justify-center w-10 h-10 rounded-lg font-bold text-lg border transition-all
              ${locker.is_used ? 'bg-red-500 text-white border-red-700' : 'bg-gray-300 text-gray-700 border-gray-400 hover:bg-green-200 cursor-pointer'}`}
          >
            {locker.locker_number}
          </div>
        ))}
      </div>
      <div className="mt-4 text-sm text-gray-400">Merah: Sudah diisi, Abu: Tersedia</div>
    </div>
  );
} 