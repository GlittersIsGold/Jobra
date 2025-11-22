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
      console.error('Registration error:', err);
      
      if (err instanceof Error) {
        if (err.message.includes('min') || err.message.includes('Password')) {
          setError('Пароль должен содержать минимум 6 символов');
        } else if (err.message.includes('Username') || err.message.includes('username')) {
          setError('Логин должен содержать минимум 3 символа');
        } else if (err.message.includes('email') || err.message.includes('Email')) {
          setError('Введите корректный email адрес');
        } else if (err.message.includes('409') || err.message.includes('Conflict')) {
          setError('Пользователь с таким именем или email уже существует');
        } else if (err.message.includes('400') || err.message.includes('validation')) {
          setError('Проверьте правильность введенных данных');
        } else {
          setError('Ошибка регистрации. Попробуйте позже.');
        }
      } else {
        setError('Неизвестная ошибка');
      }
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