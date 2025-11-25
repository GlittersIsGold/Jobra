'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import Link from 'next/link';

export default function Header() {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [username, setUsername] = useState('');
  const router = useRouter();

  useEffect(() => {
    const token = localStorage.getItem('jwt_token');
    setIsAuthenticated(!!token);
    
    const userData = localStorage.getItem('user_data');
    if (userData) {
      try {
        const user = JSON.parse(userData);
        setUsername(user.username);
      } catch (e) {
        console.error('Failed to parse user data');
      }
    }
  }, []);

  const handleLogout = () => {
    localStorage.removeItem('jwt_token');
    localStorage.removeItem('user_data');
    setIsAuthenticated(false);
    setUsername('');
    router.push('/login');
  };

  return (
    <header className="bg-blue-600 text-white shadow-lg">
      <div className="container mx-auto px-4 py-4 flex justify-between items-center">
        <Link href="/" className="text-xl font-bold">
          Джобра
        </Link>
        
        <nav>
          <ul className="flex space-x-6 items-center">
            {isAuthenticated ? (
              <>
                <li>
                  <Link href="/orders" className="hover:text-blue-200 transition-colors">
                    Все заказы
                  </Link>
                </li>
                <li>
                  <Link href="/my-orders" className="hover:text-blue-200 transition-colors">
                    Мои заказы
                  </Link>
                </li>
                <li>
                  <Link href="/profile" className="hover:text-blue-200 transition-colors">
                    Профиль
                  </Link>
                </li>
                <li className="flex items-center space-x-2">
                  <button
                    onClick={handleLogout}
                    className="bg-blue-700 hover:bg-blue-800 px-3 py-1 rounded transition-colors"
                  >
                    Выйти
                  </button>
                </li>
              </>
            ) : (
              <>
                <li>
                  <Link href="/login" className="hover:text-blue-200 transition-colors">
                    Вход
                  </Link>
                </li>
                <li>
                  <Link href="/register" className="hover:text-blue-200 transition-colors">
                    Регистрация
                  </Link>
                </li>
              </>
            )}
          </ul>
        </nav>
      </div>
    </header>
  );
}