import BrandBackground from "@/components/BrandBackground";
import BrandInfo from "@/components/BrandInfo";
import GlassCard from "@/components/GlassCard";
import LoginForm from "@/components/LoginForm";


export default function LoginPage() {
    return (
        <BrandBackground>
          <GlassCard className="w-full max-w-4xl mx-auto">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              {/* Kolom Kiri */}
              <div className="hidden md:flex">
                 <BrandInfo />
              </div>
    
              {/* Kolom Kanan */}
              <div className="flex items-center justify-center md:border-l md:border-white/10 md:pl-8">
                <LoginForm />
              </div>
            </div>
          </GlassCard>
        </BrandBackground>
      );
}