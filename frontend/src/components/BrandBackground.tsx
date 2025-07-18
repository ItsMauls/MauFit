import React from "react";

interface BrandBackgroundProps {
  children?: React.ReactNode;
}

const BrandBackground: React.FC<BrandBackgroundProps> = ({ children }) => {
  // URL untuk GIF gym. Ganti dengan URL GIF Anda sendiri jika perlu.
  const gymGifUrl = '';
  
  return (
    <div className="relative w-full h-screen overflow-hidden font-sans bg-gradient-to-br from-green-900 via-green-900/70 to-black">
      {/* Konten utama (Form Login) diposisikan di tengah */}
      <div className="relative z-10 w-full h-full flex items-center justify-center p-4">
        {children}
      </div>
    </div>
  );
};

export default BrandBackground; 