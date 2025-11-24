'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { AuthService } from '@/lib/auth';
import { UserProfile, UpdateProfileRequest, ChangePasswordRequest } from '@/types/auth';
import EditProfileForm from '@/components/EditProfileForm';
import ChangePasswordForm from '@/components/ChangePasswordForm';

export default function ProfilePage() {
  const [user, setUser] = useState<UserProfile | null>(null);
  const [isEditing, setIsEditing] = useState(false);
  const [isChangingPassword, setIsChangingPassword] = useState(false);
  const [loading, setLoading] = useState(true);
  const [updateLoading, setUpdateLoading] = useState(false);
  const [passwordLoading, setPasswordLoading] = useState(false);
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');
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
    setSuccess('');

    try {
      const response = await AuthService.updateProfile(userData);
      setUser(response.user);
      setIsEditing(false);
      setSuccess('Профиль успешно обновлен');
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Ошибка обновления профиля');
    } finally {
      setUpdateLoading(false);
    }
  };

  const handleChangePassword = async (passwordData: ChangePasswordRequest) => {
    setPasswordLoading(true);
    setError('');
    setSuccess('');
  
    try {
      await AuthService.changePassword(passwordData);
      setIsChangingPassword(false);
      setSuccess('Пароль успешно изменен');
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Ошибка смены пароля');
    } finally {
      setPasswordLoading(false);
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

  if (isChangingPassword) {
    return (
      <div className="min-h-screen bg-gray-100 py-8">
        {error && (
          <div className="max-w-md mx-auto mb-4 p-4 bg-red-100 border border-red-400 text-red-700 rounded">
            {error}
          </div>
        )}
        <ChangePasswordForm
          onChangePassword={handleChangePassword}
          onCancel={() => setIsChangingPassword(false)}
          loading={passwordLoading}
        />
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-100 py-8">
      <div className="max-w-2xl mx-auto bg-white rounded-lg shadow-md p-6">
        {success && (
          <div className="mb-6 p-4 bg-green-100 border border-green-400 text-green-700 rounded">
            {success}
          </div>
        )}
        {error && (
          <div className="mb-6 p-4 bg-red-100 border border-red-400 text-red-700 rounded">
            {error}
          </div>
        )}
        
        {/* Заголовок в одну строку */}
        <div className="mb-6">
          <h1 className="text-2xl font-bold text-gray-900">Профиль</h1>
        </div>
        
        {/* Информация о пользователе */}
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
          <div className="space-y-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">ID пользователя</label>
              <p className="text-sm text-gray-900 bg-gray-50 p-2 rounded">{user.id}</p>
            </div>
            
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Имя пользователя</label>
              <p className="text-sm text-gray-900 bg-gray-50 p-2 rounded">{user.username}</p>
            </div>
          </div>
          
          <div className="space-y-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Email</label>
              <p className="text-sm text-gray-900 bg-gray-50 p-2 rounded">{user.email}</p>
            </div>
            
            {user.created_at && (
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Зарегистрирован</label>
                <p className="text-sm text-gray-900 bg-gray-50 p-2 rounded">
                  {new Date(user.created_at).toLocaleDateString('ru-RU')}
                </p>
              </div>
            )}
          </div>
        </div>

        <div className="pt-6">
          <div className="flex flex-wrap gap-3 justify-center">
            <button
              onClick={() => setIsEditing(true)}
              className="bg-blue-600 hover:bg-blue-700 text-white px-6 py-2 rounded-lg font-medium transition-colors"
            >
              Редактировать профиль
            </button>
            <button
              onClick={() => setIsChangingPassword(true)}
              className="bg-green-600 hover:bg-green-700 text-white px-6 py-2 rounded-lg font-medium transition-colors"
            >
              Сменить пароль
            </button>
            <button
              onClick={handleLogout}
              className="bg-gray-600 hover:bg-gray-700 text-white px-6 py-2 rounded-lg font-medium transition-colors"
            >
              Выйти
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}