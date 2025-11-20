'use client';

import { AuthService } from '@/lib/auth';
import Link from 'next/link';
import { useRouter } from 'next/navigation';

export default function Header() {
  const router = useRouter();
  const isAuthenticated = AuthService.isAuthenticated();

  const handleLogout = () => {
    AuthService.removeToken();
    router.push('/login');
  };

  return (
    <header className="bg-blue-600 text-white shadow-lg">
      <div className="container mx-auto px-4 py-4 flex justify-between items-center">
        <Link href="/" className="text-xl font-bold">Джобра</Link>
        
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