'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';

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
        <h1 className="text-xl font-bold">Джобра</h1>
        
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