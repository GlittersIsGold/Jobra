'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { OrderService } from '@/lib/order';
import { Order, CreateOrderRequest } from '@/types/order';
import CreateOrderForm from '@/components/CreateOrderForm';
import OrderList from '@/components/OrderList';

export default function OrdersPage() {
  const [orders, setOrders] = useState<Order[]>([]);
  const [isCreating, setIsCreating] = useState(false);
  const [loading, setLoading] = useState(true);
  const [createLoading, setCreateLoading] = useState(false);
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');
  const router = useRouter();

  useEffect(() => {
    const fetchOrders = async () => {
      const token = localStorage.getItem('jwt_token');
      
      if (!token) {
        router.push('/login');
        return;
      }

      try {
        const response = await OrderService.getUserOrders();
        setOrders(response.orders);
      } catch (error) {
        console.error('Orders fetch error:', error);
        setError('Ошибка загрузки заказов');
      } finally {
        setLoading(false);
      }
    };

    fetchOrders();
  }, [router]);

  const handleCreateOrder = async (orderData: CreateOrderRequest) => {
    setCreateLoading(true);
    setError('');
    setSuccess('');

    try {
      const response = await OrderService.createOrder(orderData);
      setOrders(prev => [response.order, ...prev]);
      setIsCreating(false);
      setSuccess('Заказ успешно создан');
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Ошибка создания заказа');
    } finally {
      setCreateLoading(false);
    }
  };

  const handleEditOrder = (order: Order) => {
    // TODO: Реализовать редактирование
    console.log('Edit order:', order);
  };

  const handleDeleteOrder = async (orderId: string) => {
    if (!confirm('Вы уверены, что хотите удалить этот заказ?')) {
      return;
    }

    try {
      await OrderService.deleteOrder(orderId);
      setOrders(prev => prev.filter(order => order.id !== orderId));
      setSuccess('Заказ успешно удален');
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Ошибка удаления заказа');
    }
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-gray-100 flex items-center justify-center">
        <div className="text-xl">Загрузка заказов...</div>
      </div>
    );
  }

  if (isCreating) {
    return (
      <div className="min-h-screen bg-gray-100 py-8">
        {error && (
          <div className="max-w-2xl mx-auto mb-4 p-4 bg-red-100 border border-red-400 text-red-700 rounded">
            {error}
          </div>
        )}
        <CreateOrderForm
          onCreate={handleCreateOrder}
          onCancel={() => setIsCreating(false)}
          loading={createLoading}
        />
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-100 py-8">
      <div className="max-w-6xl mx-auto px-4">
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
        
        <div className="flex justify-between items-center mb-8">
          <div>
            <h1 className="text-3xl font-bold text-gray-900">Мои заказы</h1>
            <p className="text-gray-600 mt-2">Управление вашими заказами</p>
          </div>
          <button
            onClick={() => setIsCreating(true)}
            className="bg-blue-600 hover:bg-blue-700 text-white px-6 py-3 rounded-lg font-medium"
          >
            Создать заказ
          </button>
        </div>

        <div className="mb-6 flex justify-between items-center">
          <div className="text-gray-600">
            Всего заказов: <span className="font-semibold">{orders.length}</span>
          </div>
        </div>

        <OrderList
          orders={orders}
          onEdit={handleEditOrder}
          onDelete={handleDeleteOrder}
          showActions={true}
        />
      </div>
    </div>
  );
}