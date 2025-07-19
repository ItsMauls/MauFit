'use client';

import { useEffect, useState } from 'react';
import GlassCard from '@/components/GlassCard';
import { createAttendanceByFingerprintApi } from '@/utils/api';

const AttendanceModePage = () => {
  const [status, setStatus] = useState('Connecting...');
  const [message, setMessage] = useState('');
  const [statusColor, setStatusColor] = useState('text-yellow-400');

  useEffect(() => {
    const ws = new WebSocket('ws://localhost:8088');

    ws.onopen = () => {
      setStatus('Connected');
      setStatusColor('text-green-400');
      setMessage('Waiting for fingerprint scan...');
    };

    ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      if (data.type === 'fingerprint_scanned') {
        setMessage(`Fingerprint received! Verifying...`);
        
        createAttendanceByFingerprintApi(data.template)
          .then(response => {
            // Assuming the response contains user data in `response.data.user`
            const userName = response?.data?.user?.name || 'Member';
            setMessage(`Welcome, ${userName}! Attendance recorded.`);
            setStatusColor('text-green-400');
          })
          .catch(error => {
            setMessage(`Verification Failed: ${error.message}`);
            setStatusColor('text-orange-500');
          })
          .finally(() => {
            // Reset message after a few seconds
            setTimeout(() => {
              if (ws.readyState === ws.OPEN) {
                setMessage('Waiting for fingerprint scan...');
                setStatusColor('text-yellow-400');
              }
            }, 3000);
          });
      }
    };

    ws.onclose = () => {
      setStatus('Disconnected');
      setStatusColor('text-red-400');
      setMessage('Connection to agent lost. Please ensure the agent is running.');
    };

    ws.onerror = (error) => {
      setStatus('Error');
      setStatusColor('text-red-600');
      setMessage('Failed to connect to the fingerprint agent.');
      console.error('WebSocket Error:', error);
    };

    // Cleanup on component unmount
    return () => {
      ws.close();
    };
  }, []);

  return (
    <div className="flex items-center justify-center min-h-screen">
      <GlassCard className="w-full max-w-md text-center">
        <h1 className="text-3xl font-bold mb-4">Attendance Mode</h1>
        <p className="text-lg mb-2">
          Status: <strong className={statusColor}>{status}</strong>
        </p>
        <p className="text-md whitespace-normal break-words">{message}</p>
        <div className="mt-6">
          <p className="text-sm text-gray-400">Please scan your fingerprint on the device.</p>
        </div>
      </GlassCard>
    </div>
  );
};

export default AttendanceModePage;