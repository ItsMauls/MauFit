'use client'
import { useEffect, useState } from "react";

const BrandInfo = () => {
    const illustrations = [
        {
            src: 'https://placehold.co/600x400/166534/FFFFFF?text=Hidrasi+Penting',
            alt: 'Ilustrasi orang minum air',
            trivia: 'Trivia: Tetap terhidrasi tidak hanya meningkatkan energi, tetapi juga melumasi sendi untuk mencegah cedera.'
        },
        {
            src: 'https://placehold.co/600x400/15803d/FFFFFF?text=Pemanasan',
            alt: 'Ilustrasi orang melakukan pemanasan',
            trivia: 'Trivia: Pemanasan sebelum berolahraga meningkatkan aliran darah ke otot dan mengurangi risiko cedera.'
        },
        {
            src: 'https://placehold.co/600x400/16a34a/FFFFFF?text=Istirahat+Cukup',
            alt: 'Ilustrasi orang tidur',
            trivia: 'Trivia: Otot Anda tumbuh dan memperbaiki diri saat Anda tidur, bukan saat Anda berolahraga.'
        }
    ];

    const [currentIndex, setCurrentIndex] = useState(0);

    useEffect(() => {
        const interval = setInterval(() => {
            setCurrentIndex(prevIndex => (prevIndex + 1) % illustrations.length);
        }, 5000); // Ganti slide setiap 5 detik
        return () => clearInterval(interval);
    }, [illustrations.length]);

    return (
        <div className="flex flex-col justify-center text-white h-full pr-8">
            {/* UPDATED: Font dan gradasi warna diterapkan di sini */}
            <h1 className="text-6xl font-bold drop-shadow-lg font-poppins bg-gradient-to-r from-white to-green-400 bg-clip-text text-transparent pb-2">
                MauFit
            </h1>
            
            <div className="mt-8 relative h-48 w-full overflow-hidden rounded-lg">
                {illustrations.map((item, index) => (
                    <img
                        key={index}
                        src={item.src}
                        alt={item.alt}
                        className={`absolute top-0 left-0 w-full h-full object-cover transition-opacity duration-1000 ${index === currentIndex ? 'opacity-100' : 'opacity-0'}`}
                    />
                ))}
            </div>
            
            <div className="mt-4 text-sm text-gray-300 italic h-16">
                 <p className="transition-opacity duration-1000">{illustrations[currentIndex].trivia}</p>
            </div>
        </div>
    );
};

export default BrandInfo;