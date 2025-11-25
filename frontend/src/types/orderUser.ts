import { Order } from "@/types/order";

export interface OrderUser {
  id: string;
  order_id: string;
  user_id: string;
  user_username: string;
  user_email: string;
  message?: string;
  status: 'pending' | 'accepted' | 'rejected';
  created_at: string;
  updated_at: string;
}

export interface CreateOrderUserRequest {
  message?: string;
}

export interface AssignExecutorRequest {
  order_user_id: string;
}

export interface OrderWithDetails extends Order {
  executor_id?: string;
  executor_assigned_at?: string;
  customer_username?: string;
  customer_email?: string;
  executor_username?: string;
  executor_email?: string;
}