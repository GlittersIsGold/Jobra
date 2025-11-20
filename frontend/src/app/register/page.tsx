'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import AuthForm from '@/components/AuthForm';
import { AuthService } from '@/lib/auth';
import { RegisterRequest } from '@/types/auth';

export default function RegisterPage() {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  const router = useRouter();

  const handleRegister = async (userData: RegisterRequest) => {
    setLoading(true);
    setError('');

    try {
      const response = await AuthService.register(userData);
      AuthService.saveToken(response.token!);
      router.push('/profile');
    } catch (err) {
      setError('Ошибка регистрации. Попробуйте другой username или email.');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-gray-100 py-8">
      {error && (
        <div className="max-w-md mx-auto mb-4 p-4 bg-red-100 border border-red-400 text-red-700 rounded">
          {error}
        </div>
      )}
      <AuthForm type="register" onSubmit={handleRegister} loading={loading} />
    </div>
  );
}