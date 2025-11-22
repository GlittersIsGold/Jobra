'use client';

import { useEffect, useState } from 'react';
import Link from 'next/link';

export default function Home() {
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  useEffect(() => {
    const token = localStorage.getItem('jwt_token');
    setIsAuthenticated(!!token);
  }, []);

  return (
    <div className="min-h-screen bg-gray-100">
      <main className="container mx-auto px-4 py-8">
        <div className="text-center">
          <h1 className="text-4xl font-bold text-gray-900 mb-4">
            Добро пожаловать
          </h1>
          <p className="text-xl text-gray-600 mb-8">
            {isAuthenticated 
              ? 'Вы успешно вошли в систему' 
              : 'Пожалуйста, войдите или зарегистрируйтесь'
            }
          </p>
          
          {isAuthenticated ? (
            <Link
              href="/profile"
              className="bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700"
            >
              Перейти в профиль
            </Link>
          ) : (
            <div className="space-x-4">
              <Link
                href="/login"
                className="bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700"
              >
                Войти
              </Link>
              <Link
                href="/register"
                className="bg-green-600 text-white px-6 py-3 rounded-lg hover:bg-green-700"
              >
                Зарегистрироваться
              </Link>
            </div>
          )}
        </div>
      </main>
    </div>
  );
}