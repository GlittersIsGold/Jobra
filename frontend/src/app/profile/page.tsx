'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { AuthService } from '@/lib/auth';
import { UserProfile, UpdateProfileRequest } from '@/types/auth';
import EditProfileForm from '@/components/EditProfileForm';

export default function ProfilePage() {
  const [user, setUser] = useState<UserProfile | null>(null);
  const [isEditing, setIsEditing] = useState(false);
  const [loading, setLoading] = useState(true);
  const [updateLoading, setUpdateLoading] = useState(false);
  const [error, setError] = useState('');
  const router = useRouter();

  useEffect(() => {
    const fetchProfile = async () => {
      const token = AuthService.getToken();
      
      if (!token) {
        router.push('/login');
        return;
      }

      try {
        const profile = await AuthService.getProfile(token);
        setUser(profile.user);
      } catch (error) {
        console.error('Profile fetch error:', error);
        AuthService.removeToken();
        router.push('/login');
      } finally {
        setLoading(false);
      }
    };

    fetchProfile();
  }, [router]);

  const handleUpdateProfile = async (userData: UpdateProfileRequest) => {
    setUpdateLoading(true);
    setError('');

    try {
      const response = await AuthService.updateProfile(userData);
      setUser(response.user);
      setIsEditing(false);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Ошибка обновления профиля');
    } finally {
      setUpdateLoading(false);
    }
  };

  const handleLogout = () => {
    AuthService.removeToken();
    router.push('/login');
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-gray-100 flex items-center justify-center">
        <div className="text-xl">Загрузка...</div>
      </div>
    );
  }

  if (!user) {
    return null;
  }

  if (isEditing) {
    return (
      <div className="min-h-screen bg-gray-100 py-8">
        {error && (
          <div className="max-w-md mx-auto mb-4 p-4 bg-red-100 border border-red-400 text-red-700 rounded">
            {error}
          </div>
        )}
        <EditProfileForm
          user={user}
          onUpdate={handleUpdateProfile}
          onCancel={() => setIsEditing(false)}
          loading={updateLoading}
        />
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-100 py-8">
      <div className="max-w-md mx-auto bg-white rounded-lg shadow-md p-6">
        <div className="flex justify-between items-center mb-6">
          <h1 className="text-2xl font-bold">Профиль пользователя</h1>
          <div className="space-x-2">
            <button
              onClick={() => setIsEditing(true)}
              className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded text-sm"
            >
              Редактировать
            </button>
            <button
              onClick={handleLogout}
              className="bg-gray-600 hover:bg-gray-700 text-white px-4 py-2 rounded text-sm"
            >
              Выйти
            </button>
          </div>
        </div>
        
        <div className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-gray-700">ID</label>
            <p className="mt-1 text-sm text-gray-900">{user.id}</p>
          </div>
          
          <div>
            <label className="block text-sm font-medium text-gray-700">Имя пользователя</label>
            <p className="mt-1 text-sm text-gray-900">{user.username}</p>
          </div>
          
          <div>
            <label className="block text-sm font-medium text-gray-700">Email</label>
            <p className="mt-1 text-sm text-gray-900">{user.email}</p>
          </div>

          {user.created_at && (
            <div>
              <label className="block text-sm font-medium text-gray-700">Зарегистрирован</label>
              <p className="mt-1 text-sm text-gray-900">
                {new Date(user.created_at).toLocaleDateString('ru-RU')}
              </p>
            </div>
          )}

          {user.updated_at && user.updated_at !== user.created_at && (
            <div>
              <label className="block text-sm font-medium text-gray-700">Обновлен</label>
              <p className="mt-1 text-sm text-gray-900">
                {new Date(user.updated_at).toLocaleDateString('ru-RU')}
              </p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}