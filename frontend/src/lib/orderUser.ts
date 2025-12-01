import { OrderUser, CreateOrderUserRequest, AssignExecutorRequest } from '@/types/orderUser';

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

export class OrderUserService {
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

  static async getOrderUsers(orderId: string): Promise<{ order_users: OrderUser[]; count: number }> {
    const token = this.getToken();
    if (!token) {
      throw new Error('Not authenticated');
    }

    const response = await fetch(`${API_BASE_URL}/api/orders/users?order_id=${orderId}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(`Failed to fetch order users: ${response.status} - ${errorText}`);
    }

    return response.json();
  }

  static async createOrderUser(orderId: string, request: CreateOrderUserRequest): Promise<{ message: string; order_user: OrderUser }> {
    const token = this.getToken();
  
    const response = await fetch(`${API_BASE_URL}/api/orders/users?order_id=${orderId}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
      body: JSON.stringify(request),
    });
  
    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(`Order user creation failed: ${response.status} - ${errorText}`);
    }
  
    return response.json();
  }

  static async deleteOrderUser(orderUserId: string): Promise<{ message: string }> {
    const token = this.getToken();
    if (!token) {
      throw new Error('Not authenticated');
    }

    const response = await fetch(`${API_BASE_URL}/api/orders/users/${orderUserId}`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(`Delete order user failed: ${response.status} - ${errorText}`);
    }

    return response.json();
  }

  static async checkOrderUserExists(orderId: string): Promise<{ has_responded: boolean }> {
    const token = this.getToken();
    if (!token) {
      throw new Error('Not authenticated');
    }
  
    const url = `${API_BASE_URL}/api/orders/users/check?order_id=${orderId}`;
  
    try {
      const response = await fetch(url, {
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      });

      if (!response.ok) {
        const errorText = await response.text();
        throw new Error(`Failed to check order user: ${response.status} - ${errorText}`);
      }
  
      const result = await response.json();
      return result;
    } catch (error) {
      throw error;
    }
  }

static async assignExecutor(orderId: string, request: AssignExecutorRequest): Promise<{ message: string; order: any }> {
    const token = this.getToken();
    if (!token) {
      throw new Error('Not authenticated');
    }

    const requestBody = {
      order_id: orderId,
      user_id: request.user_id
    };
    
    const url = `${API_BASE_URL}/api/assign`;
    console.log('Making request to:', url);
    console.log('Request body:', requestBody);
    
    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`,
            },
            body: JSON.stringify(requestBody),
        });

        console.log('Response status:', response.status);
        
        const responseText = await response.text();
        console.log('Response text:', responseText);
        
        if (!response.ok) {
            throw new Error(`Assignment failed: ${response.status} - ${responseText}`);
        }

        if (!responseText) {
            console.warn('Empty response body');
            return {
                message: 'Executor assigned successfully',
                order: null
            };
        }

        const result = JSON.parse(responseText);
        console.log('Success response:', result);
        return result;
    } catch (error) {
        console.error('Fetch error:', error);
        throw error;
    }
}

static async removeExecutor(orderId: string): Promise<{ message: string; order: any }> {
    const token = this.getToken();
    if (!token) {
        throw new Error('Not authenticated');
    }

    const requestBody = {
        order_id: orderId
    };
    
    // ИЗМЕНЕННЫЙ URL:
    const url = `${API_BASE_URL}/api/assign`;
    console.log('Making DELETE request to:', url);
    
    try {
        const response = await fetch(url, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`,
            },
            body: JSON.stringify(requestBody),
        });

        console.log('Response status:', response.status);
        
        if (!response.ok) {
            const errorText = await response.text();
            console.error('Error response:', errorText);
            throw new Error(`Removal failed: ${response.status} - ${errorText}`);
        }

        const result = await response.json();
        console.log('Success response:', result);
        return result;
    } catch (error) {
        console.error('Fetch error:', error);
        throw error;
    }
  }

}