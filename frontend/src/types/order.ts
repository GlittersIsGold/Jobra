export interface Order {
    id: string;
    customer_id: string;
    title: string;
    description: string;
    budget: number;
    category: string;
    section: string;
    status: string;
    created_at: string;
    updated_at: string;
    executor_id?: string;
    executor_assigned_at?: string;
}
  
export interface CreateOrderRequest {
    title: string;
    description: string;
    budget: number;
    category: string;
    section: string;
}
  
export interface UpdateOrderRequest {
    title: string;
    description: string;
    budget: number;
    category: string;
    section: string;
    status: string;
}