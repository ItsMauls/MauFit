"use client";
import { mainApiUrl } from "@/constants";
import { fetchApi } from "@/utils/api";
import { useState, useEffect, useRef } from "react";

export default function AbsensiForm() {
  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState<any>(null);
  const [error, setError] = useState<string | null>(null);
  const [scannerStatus, setScannerStatus] = useState<string>("Disconnected");
  const [isScanning, setIsScanning] = useState(false);
  const wsRef = useRef<WebSocket | null>(null);

  // Connect to fingerprint agent WebSocket
  useEffect(() => {
    const connectWebSocket = () => {
      try {
        const ws = new WebSocket('ws://localhost:8088');
        wsRef.current = ws;

        ws.onopen = () => {
          console.log('Connected to fingerprint agent');
          setScannerStatus('Connected');
          setError(null);
        };

        ws.onmessage = async (event) => {
          const message = JSON.parse(event.data);
          console.log('Received from fingerprint agent:', message);
          
          if (message.type === 'fingerprint_scanned') {
            setIsScanning(false);
            await handleFingerprintScanned(message);
          }
        };

        ws.onclose = () => {
          console.log('Disconnected from fingerprint agent');
          setScannerStatus('Disconnected');
          setError('Koneksi ke fingerprint agent terputus. Pastikan aplikasi fingerprint agent berjalan.');
        };

        ws.onerror = (error) => {
          console.error('WebSocket error:', error);
          setScannerStatus('Error');
          setError('Error koneksi ke fingerprint agent');
        };
      } catch (err) {
        console.error('Failed to connect to fingerprint agent:', err);
        setScannerStatus('Error');
        setError('Gagal terhubung ke fingerprint agent');
      }
    };

    connectWebSocket();

    return () => {
      if (wsRef.current) {
        wsRef.current.close();
      }
    };
  }, []);

  const handleFingerprintScanned = async (scanData: any) => {
    setLoading(true);
    setError(null);
    setResult(null);
    
    try {
      // Check scan quality first
      if (scanData.quality < 70) {
        throw new Error(`Kualitas sidik jari terlalu rendah (${scanData.quality}%). Silakan coba lagi.`);
      }

      // Use the fingerprint template to create attendance
      const res = await fetch(`${process.env.NEXT_PUBLIC_MAIN_API_URL}/attendances/fingerprint`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ fingerprint_template: scanData.template }),
      });
      
      const data = await res.json();
      if (!res.ok) throw new Error(data.message || "User tidak ditemukan untuk sidik jari ini");
      
      setResult({
        message: `Absensi berhasil untuk ${scanData.user_name}!`,
        attendance: data.data,
        user: { name: scanData.user_name },
        scanInfo: {
          template: scanData.template,
          quality: scanData.quality,
          timestamp: scanData.timestamp
        }
      });
    } catch (err: any) {
      setError(err.message || "Gagal memproses absensi");
    } finally {
      setLoading(false);
    }
  };

  const startFingerprint = () => {
    if (wsRef.current && wsRef.current.readyState === WebSocket.OPEN) {
      setIsScanning(true);
      setError(null);
      setResult(null);
      wsRef.current.send(JSON.stringify({ command: 'start_scan' }));
    } else {
      setError('Fingerprint agent tidak terhubung');
    }
  };

  return (
    <div className="space-y-6">
      {/* Scanner Status */}
      <div className="flex items-center justify-between p-3 rounded-lg bg-black/20">
        <span className="text-sm font-medium text-green-200">Status Scanner:</span>
        <span className={`text-sm font-bold ${
          scannerStatus === 'Connected' ? 'text-green-400' : 
          scannerStatus === 'Disconnected' ? 'text-red-400' : 'text-yellow-400'
        }`}>
          {scannerStatus}
        </span>
      </div>

      {/* Scan Instructions */}
      <div className="text-center p-4 rounded-lg bg-black/20">
        <div className="text-6xl mb-4">👆</div>
        <div className="text-lg font-semibold text-white mb-2">
          {isScanning ? "Tempelkan sidik jari Anda..." : "Siap untuk memindai sidik jari"}
        </div>
        <div className="text-sm text-gray-300">
          {isScanning ? "Sedang memindai, harap tunggu..." : "Klik tombol di bawah untuk memulai pemindaian"}
        </div>
      </div>

      {/* Scan Button */}
      <button
        onClick={startFingerprint}
        className="w-full bg-green-500 text-black font-bold py-3 px-4 rounded-lg hover:bg-green-400 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-black focus:ring-green-400 transition-colors duration-300 shadow-[0_0_15px_rgba(0,255,102,0.5)] hover:shadow-[0_0_25px_rgba(0,255,102,0.7)] disabled:opacity-50 disabled:cursor-not-allowed"
        disabled={loading || isScanning || scannerStatus !== 'Connected'}
      >
        {loading ? "Memproses Absensi..." : 
         isScanning ? "Sedang Memindai..." : 
         scannerStatus !== 'Connected' ? "Scanner Tidak Terhubung" :
         "🔍 Mulai Pindai Sidik Jari"}
      </button>

      {/* Results and Errors */}
      {error && (
        <div className="text-red-400 text-sm text-center p-3 rounded-md bg-red-900/20 border border-red-500/30">
          <div className="font-semibold mb-1">❌ Error</div>
          <div>{error}</div>
        </div>
      )}
      
      {result && (
        <div className="text-green-300 text-sm text-center p-3 rounded-md bg-green-900/20 border border-green-500/30">
          <div className="font-semibold mb-3">✅ {result.message}</div>
          
          {/* User Information */}
          {result.user && (
            <div className="mb-3 p-2 bg-black/20 rounded">
              <div className="font-semibold text-green-400">👤 {result.user.name}</div>
            </div>
          )}
          
          {/* Attendance Information */}
          {result.attendance && (
            <div className="mb-3 space-y-1">
              <div className="font-semibold text-green-400">📋 Detail Absensi:</div>
              <div>ID: <span className="font-mono text-green-300">#{result.attendance.id}</span></div>
              <div>User ID: <span className="font-mono text-green-300">{result.attendance.user_id}</span></div>
              <div>Waktu Masuk: <span className="font-mono text-green-300">
                {new Date(result.attendance.time_in).toLocaleString('id-ID')}
              </span></div>
            </div>
          )}
          
          {/* Scan Quality Information */}
          {result.scanInfo && (
            <div className="space-y-1 p-2 bg-black/20 rounded">
              <div className="font-semibold text-blue-400">🔍 Info Pemindaian:</div>
              <div>Kualitas: <span className={`font-bold ${
                result.scanInfo.quality >= 85 ? 'text-green-400' :
                result.scanInfo.quality >= 75 ? 'text-yellow-400' : 'text-orange-400'
              }`}>{result.scanInfo.quality}%</span></div>
              <div>Template: <span className="font-mono text-xs text-gray-400">
                {result.scanInfo.template.length > 30 ? 
                  result.scanInfo.template.substring(0, 30) + '...' : 
                  result.scanInfo.template
                }
              </span></div>
              <div>Waktu Scan: <span className="font-mono text-xs text-gray-400">
                {new Date(result.scanInfo.timestamp).toLocaleTimeString('id-ID')}
              </span></div>
            </div>
          )}
        </div>
      )}

      {/* Instructions */}
      <div className="text-xs text-gray-400 text-center p-2 rounded bg-black/10">
        💡 Pastikan aplikasi Fingerprint Agent sudah berjalan di background
      </div>
    </div>
  );
}