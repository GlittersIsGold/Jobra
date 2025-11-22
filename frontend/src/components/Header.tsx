'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';

export default function Header() {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const router = useRouter();

  useEffect(() => {
    // Проверяем аутентификацию только на клиенте
    const token = localStorage.getItem('jwt_token');
    setIsAuthenticated(!!token);
  }, []);

  const handleLogout = () => {
    localStorage.removeItem('jwt_token');
    setIsAuthenticated(false);
    router.push('/login');
  };

  return (
    <header className="bg-blue-600 text-white shadow-lg">
      <div className="container mx-auto px-4 py-4 flex justify-between items-center">
        <h1 className="text-xl font-bold">My App</h1>
        
        <nav>
          <ul className="flex space-x-4 items-center">
            {isAuthenticated ? (
              <>
                <li>
                  <a href="/profile" className="hover:text-blue-200">
                    Профиль
                  </a>
                </li>
                <li>
                  <button
                    onClick={handleLogout}
                    className="bg-blue-700 hover:bg-blue-800 px-3 py-1 rounded"
                  >
                    Выйти
                  </button>
                </li>
              </>
            ) : (
              <>
                <li>
                  <a href="/login" className="hover:text-blue-200">
                    Вход
                  </a>
                </li>
                <li>
                  <a href="/register" className="hover:text-blue-200">
                    Регистрация
                  </a>
                </li>
              </>
            )}
          </ul>
        </nav>
      </div>
    </header>
  );
}