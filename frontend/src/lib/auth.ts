import { 
  UserProfile, 
  UpdateProfileRequest, 
  ChangePasswordRequest 
} from '@/types/user';

import {
  LoginRequest, 
  RegisterRequest, 
  AuthResponse, 
  User,
} from '@/types/auth';

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

export class AuthService {
  static async login(credentials: LoginRequest): Promise<AuthResponse> {
    const response = await fetch(`${API_BASE_URL}/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(credentials),
    });

    if (!response.ok) {
      throw new Error('Login failed');
    }

    return response.json();
  }

  static async register(userData: RegisterRequest): Promise<AuthResponse> {
    const response = await fetch(`${API_BASE_URL}/register`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(userData),
    });

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(`Registration failed: ${response.status} - ${errorText}`);
    }

    return response.json();
  }

  static async getProfile(token: string): Promise<{ user: User }> {
    const response = await fetch(`${API_BASE_URL}/api/profile`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      throw new Error('Failed to fetch profile');
    }

    return response.json();
  }

  static async updateProfile(userData: UpdateProfileRequest): Promise<{ message: string; user: UserProfile }> {
    const token = this.getToken();
    if (!token) {
      throw new Error('Not authenticated');
    }

    const response = await fetch(`${API_BASE_URL}/api/profile`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
      body: JSON.stringify(userData),
    });

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(`Profile update failed: ${response.status} - ${errorText}`);
    }

    return response.json();
  }

  static async changePassword(passwordData: ChangePasswordRequest): Promise<{ message: string }> {
    const token = this.getToken();
    if (!token) {
      throw new Error('Not authenticated');
    }

    const response = await fetch(`${API_BASE_URL}/api/profile/password`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
      body: JSON.stringify(passwordData),
    });

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(`Password change failed: ${response.status} - ${errorText}`);
    }

    return response.json();
  }

  static saveToken(token: string): void {
    if (typeof window !== 'undefined') {
      localStorage.setItem('jwt_token', token);
    }
  }

  static getToken(): string | null {
    if (typeof window !== 'undefined') {
      return localStorage.getItem('jwt_token');
    }
    return null;
  }

  static removeToken(): void {
    if (typeof window !== 'undefined') {
      localStorage.removeItem('jwt_token');
    }
  }

  static isAuthenticated(): boolean {
    return this.getToken() !== null;
  }
}