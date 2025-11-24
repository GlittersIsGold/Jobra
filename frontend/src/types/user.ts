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

export interface ChangePasswordRequest {
  currentPassword: string;
  newPassword: string;
}