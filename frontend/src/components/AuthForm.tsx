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

interface FormErrors {
  username?: string;
  password?: string;
  email?: string;
}

export default function AuthForm(props: AuthFormProps) {
  const { type, onSubmit, loading = false } = props;
  
  const [formData, setFormData] = useState({
    username: '',
    email: '',
    password: '',
  });
  const [errors, setErrors] = useState<FormErrors>({});

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    
    setFormData({
      ...formData,
      [name]: value,
    });

    if (type === 'register') {
      if (name === 'username') {
        if (value.length > 0 && value.length < 3) {
          setErrors(prev => ({ ...prev, username: 'Логин должен содержать минимум 3 символа' }));
        } else {
          setErrors(prev => ({ ...prev, username: undefined }));
        }
      }

      if (name === 'password') {
        if (value.length > 0 && value.length < 6) {
          setErrors(prev => ({ ...prev, password: 'Пароль должен содержать минимум 6 символов' }));
        } else {
          setErrors(prev => ({ ...prev, password: undefined }));
        }
      }

      if (name === 'email') {
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        if (value.length > 0 && !emailRegex.test(value)) {
          setErrors(prev => ({ ...prev, email: 'Введите корректный email адрес' }));
        } else {
          setErrors(prev => ({ ...prev, email: undefined }));
        }
      }
    }
  };

  const validateForm = (): boolean => {
    if (type === 'login') {
      return formData.username.length > 0 && formData.password.length > 0;
    }

    const newErrors: FormErrors = {};

    if (formData.username.length < 3) {
      newErrors.username = 'Логин должен содержать минимум 3 символа';
    }

    if (formData.password.length < 6) {
      newErrors.password = 'Пароль должен содержать минимум 6 символов';
    }

    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(formData.email)) {
      newErrors.email = 'Введите корректный email адрес';
    }

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const isFormValid = (): boolean => {
    if (type === 'login') {
      return formData.username.length > 0 && formData.password.length > 0;
    }

    if (formData.username.length < 3) return false;
    if (formData.password.length < 6) return false;
    
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(formData.email)) return false;
    
    return true;
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    
    if (!validateForm()) {
      return;
    }
    
    if (type === 'login') {
      const { username, password } = formData;
      (onSubmit as (data: LoginRequest) => void)({ username, password });
    } else {
      (onSubmit as (data: RegisterRequest) => void)(formData);
    }
  };

  const isSubmitDisabled = loading || !isFormValid();

  return (
    <div className="max-w-md mx-auto mt-8 p-6 bg-white rounded-lg shadow-md">
      <h2 className="text-2xl font-bold mb-6 text-center">
        {type === 'login' ? 'Вход' : 'Регистрация'}
      </h2>
      
      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label htmlFor="username" className="block text-sm font-medium text-gray-700">
            Имя пользователя {type === 'register' && <span className="text-gray-500 text-xs">(минимум 3 символа)</span>}
          </label>
          <input
            type="text"
            id="username"
            name="username"
            value={formData.username}
            onChange={handleChange}
            required
            minLength={type === 'register' ? 3 : undefined}
            className={`mt-1 block w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 ${
              errors.username ? 'border-red-500' : 'border-gray-300'
            }`}
          />
          {errors.username && (
            <p className="mt-1 text-sm text-red-600">{errors.username}</p>
          )}
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
              className={`mt-1 block w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 ${
                errors.email ? 'border-red-500' : 'border-gray-300'
              }`}
            />
            {errors.email && (
              <p className="mt-1 text-sm text-red-600">{errors.email}</p>
            )}
          </div>
        )}

        <div>
          <label htmlFor="password" className="block text-sm font-medium text-gray-700">
            Пароль {type === 'register' && <span className="text-gray-500 text-xs">(минимум 6 символов)</span>}
          </label>
          <input
            type="password"
            id="password"
            name="password"
            value={formData.password}
            onChange={handleChange}
            required
            minLength={type === 'register' ? 6 : undefined}
            className={`mt-1 block w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 ${
              errors.password ? 'border-red-500' : 'border-gray-300'
            }`}
          />
          {errors.password && (
            <p className="mt-1 text-sm text-red-600">{errors.password}</p>
          )}
        </div>

        <button
          type="submit"
          disabled={isSubmitDisabled}
          className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed"
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