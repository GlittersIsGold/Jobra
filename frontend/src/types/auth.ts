export interface User {
    id: string;
    username: string;
    email: string;
  }

export interface UserProfile {
  id: string;
  username: string;
  email: string;
  created_at?: string;
  updated_at?: string;
}

export interface UpdateProfileRequest {
  username: string;
  email: string;
}

export interface LoginRequest {
  username: string;
  password: string;
}

export interface RegisterRequest {
  username: string;
  email: string;
  password: string;
}

export interface AuthResponse {
  message: string;
  token?: string;
  user: User;
}