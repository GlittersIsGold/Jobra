import { User } from '@/types/user';

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

export class UserService {
  private static getToken(): string {
    if (typeof window !== 'undefined') {
      const token = localStorage.getItem('jwt_token');
      if (!token) {
        throw new Error('No JWT token found');
      }
      return token;
    }
    throw new Error('Window is not defined');
  }

  static async getUsers(): Promise<{ users: User[] }> {
    const token = this.getToken();

    const response = await fetch(`${API_BASE_URL}/api/users`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(`Failed to fetch users: ${response.status} - ${errorText}`);
    }

    return response.json();
  }

  static async getUser(userId: string): Promise<{ user: User }> {
    const token = this.getToken();

    const response = await fetch(`${API_BASE_URL}/api/users/${userId}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(`Failed to fetch user: ${response.status} - ${errorText}`);
    }

    return response.json();
  }

  static async updateUser(userId: string, userData: Partial<User>): Promise<{ user: User }> {
    const token = this.getToken();

    const response = await fetch(`${API_BASE_URL}/api/users/${userId}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
      body: JSON.stringify(userData),
    });

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(`Failed to update user: ${response.status} - ${errorText}`);
    }

    return response.json();
  }

  static async deleteUser(userId: string): Promise<{ message: string }> {
    const token = this.getToken();

    const response = await fetch(`${API_BASE_URL}/api/users/${userId}`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(`Failed to delete user: ${response.status} - ${errorText}`);
    }

    return response.json();
  }

  static async getProfile(): Promise<{ user: User }> {
    const token = this.getToken();

    const response = await fetch(`${API_BASE_URL}/api/profile`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(`Failed to fetch profile: ${response.status} - ${errorText}`);
    }

    return response.json();
  }

  static async updateProfile(userData: Partial<User>): Promise<{ user: User }> {
    const token = this.getToken();

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
      throw new Error(`Failed to update profile: ${response.status} - ${errorText}`);
    }

    return response.json();
  }
}