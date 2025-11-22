'use client';

import { useState } from 'react';
import { LoginRequest, RegisterRequest } from '@/types/auth';

type LoginFormProps = {
  type: 'login';
  onSubmit: (data: LoginRequest) => void;
  loading?: boolean;
}

type RegisterFormProps = {
  type: 'register';
  onSubmit: (data: RegisterRequest) => void;
  loading?: boolean;
}

type AuthFormProps = LoginFormProps | RegisterFormProps;

export default function AuthForm(props: AuthFormProps) {
  const { type, onSubmit, loading = false } = props;
  
  const [formData, setFormData] = useState({
    username: '',
    email: '',
    password: '',
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    
    if (type === 'login') {
      const { username, password } = formData;
      (onSubmit as (data: LoginRequest) => void)({ username, password });
    } else {
      (onSubmit as (data: RegisterRequest) => void)(formData);
    }
  };

  return (
    <div className="max-w-md mx-auto mt-8 p-6 bg-white rounded-lg shadow-md">
      <h2 className="text-2xl font-bold mb-6 text-center">
        {type === 'login' ? 'Вход' : 'Регистрация'}
      </h2>
      
      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label htmlFor="username" className="block text-sm font-medium text-gray-700">
            Имя пользователя
          </label>
          <input
            type="text"
            id="username"
            name="username"
            value={formData.username}
            onChange={handleChange}
            required
            className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        {type === 'register' && (
          <div>
            <label htmlFor="email" className="block text-sm font-medium text-gray-700">
              Email
            </label>
            <input
              type="email"
              id="email"
              name="email"
              value={formData.email}
              onChange={handleChange}
              required
              className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
        )}

        <div>
          <label htmlFor="password" className="block text-sm font-medium text-gray-700">
            Пароль
          </label>
          <input
            type="password"
            id="password"
            name="password"
            value={formData.password}
            onChange={handleChange}
            required
            className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <button
          type="submit"
          disabled={loading}
          className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50"
        >
          {loading ? 'Загрузка...' : (type === 'login' ? 'Войти' : 'Зарегистрироваться')}
        </button>
      </form>

      <div className="mt-4 text-center">
        {type === 'login' ? (
          <p className="text-sm text-gray-600">
            Нет аккаунта?{' '}
            <a href="/register" className="text-blue-600 hover:text-blue-500">
              Зарегистрируйтесь
            </a>
          </p>
        ) : (
          <p className="text-sm text-gray-600">
            Уже есть аккаунт?{' '}
            <a href="/login" className="text-blue-600 hover:text-blue-500">
              Войдите
            </a>
          </p>
        )}
      </div>
    </div>
  );
}