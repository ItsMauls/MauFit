// components/GlassCard.tsx
import React from "react";

interface GlassCardProps extends React.HTMLAttributes<HTMLDivElement> {
  children: React.ReactNode;
  className?: string;
}

const GlassCard: React.FC<GlassCardProps> = ({ children, className, ...rest }) => {
  return (
    <div
      className={`
        bg-black/5 backdrop-blur-lg 
        border border-white/10 
        rounded-2xl 
        shadow-lg 
        p-8
        transition-all duration-300
        ${className} 
      `}
      {...rest}
    >
      {children}
    </div>
  );
};

export default GlassCard;