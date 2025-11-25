'use client';

import { useEffect, useState } from 'react';
import { useParams, useRouter } from 'next/navigation';
import { OrderUserService } from '@/lib/orderUser';
import { OrderService } from '@/lib/order';
import { OrderUser } from '@/types/orderUser';
import { Order } from '@/types/order';

export default function OrderResponsesPage() {
  const [orderUsers, setOrderUsers] = useState<OrderUser[]>([]);
  const [order, setOrder] = useState<Order | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [assigning, setAssigning] = useState<string | null>(null);
  const [removing, setRemoving] = useState(false);
  const params = useParams();
  const router = useRouter();
  const orderId = params.id as string;

  useEffect(() => {
    const fetchData = async () => {
      const token = localStorage.getItem('jwt_token');
      
      if (!token) {
        router.push('/login');
        return;
      }

      try {
        const [orderResponse, orderUsersResponse] = await Promise.all([
          OrderService.getOrder(orderId),
          OrderUserService.getOrderUsers(orderId)
        ]);
        
        setOrder(orderResponse.order);
        setOrderUsers(orderUsersResponse.order_users);
      } catch (err) {
        setError(err instanceof Error ? err.message : 'Ошибка загрузки данных');
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, [orderId, router]);

  const handleAssignExecutor = async (orderUserId: string) => {
    if (!confirm('Назначить этого пользователя исполнителем заказа?')) {
      return;
    }

    setAssigning(orderUserId);
    try {
      await OrderUserService.assignExecutor(orderId, { order_user_id: orderUserId });
      const [orderResponse, orderUsersResponse] = await Promise.all([
        OrderService.getOrder(orderId),
        OrderUserService.getOrderUsers(orderId)
      ]);
      setOrder(orderResponse.order);
      setOrderUsers(orderUsersResponse.order_users);
      alert('Исполнитель успешно назначен!');
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Ошибка назначения исполнителя');
    } finally {
      setAssigning(null);
    }
  };

  const handleRemoveExecutor = async () => {
    if (!confirm('Отменить назначение исполнителя?')) {
      return;
    }

    setRemoving(true);
    try {
      await OrderUserService.removeExecutor(orderId);
      const [orderResponse, orderUsersResponse] = await Promise.all([
        OrderService.getOrder(orderId),
        OrderUserService.getOrderUsers(orderId)
      ]);
      setOrder(orderResponse.order);
      setOrderUsers(orderUsersResponse.order_users);
      alert('Исполнитель отменен!');
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Ошибка отмены исполнителя');
    } finally {
      setRemoving(false);
    }
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'accepted':
        return 'bg-green-100 text-green-800';
      case 'rejected':
        return 'bg-red-100 text-red-800';
      default:
        return 'bg-yellow-100 text-yellow-800';
    }
  };

  const getStatusText = (status: string) => {
    switch (status) {
      case 'accepted':
        return 'Принят';
      case 'rejected':
        return 'Отклонен';
      default:
        return 'Ожидает';
    }
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-gray-100 flex items-center justify-center">
        <div className="text-xl">Загрузка...</div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-100 py-8">
      <div className="max-w-6xl mx-auto px-4">
        {error && (
          <div className="mb-6 p-4 bg-red-100 border border-red-400 text-red-700 rounded">
            {error}
          </div>
        )}

        <div className="mb-8">
          <button
            onClick={() => router.push('/my-orders')}
            className="text-blue-600 hover:text-blue-800 mb-4 flex items-center"
          >
            ← Назад к моим заказам
          </button>
          
          <div className="bg-white rounded-lg shadow-md p-6 mb-6">
            <h1 className="text-3xl font-bold text-gray-900 mb-2">
              Отклики на заказ: {order?.title}
            </h1>
            <p className="text-gray-600 mb-4">{order?.description}</p>
            
            <div className="flex items-center justify-between">
              <div className="text-gray-600">
                <span className="font-semibold">{orderUsers.length}</span> отклик(ов)
                {order?.executor_id && (
                  <span className="ml-4 text-green-600 font-semibold">
                    ✓ Исполнитель назначен
                  </span>
                )}
              </div>
              
              {order?.executor_id && (
                <button
                  onClick={handleRemoveExecutor}
                  disabled={removing}
                  className="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700 disabled:opacity-50 text-sm"
                >
                  {removing ? 'Отмена...' : 'Отменить исполнителя'}
                </button>
              )}
            </div>
          </div>
        </div>

        {orderUsers.length === 0 ? (
          <div className="bg-white rounded-lg shadow-md p-8 text-center">
            <p className="text-gray-500 text-lg">Пока нет откликов на этот заказ</p>
            <p className="text-gray-400 mt-2">Отклики появятся здесь, когда пользователи откликнутся на ваш заказ</p>
          </div>
        ) : (
          <div className="space-y-4">
            {orderUsers.map((orderUser) => (
              <div key={orderUser.id} className="bg-white rounded-lg shadow-md p-6 border border-gray-200">
                <div className="flex justify-between items-start mb-4">
                  <div className="flex-1">
                    <div className="flex items-center gap-4 mb-2">
                      <h3 className="text-xl font-semibold text-gray-900">
                        {orderUser.user_username}
                      </h3>
                      <span className={`px-3 py-1 rounded-full text-xs font-medium ${getStatusColor(orderUser.status)}`}>
                        {getStatusText(orderUser.status)}
                      </span>
                    </div>
                    
                    <p className="text-gray-600 text-sm mb-2">{orderUser.user_email}</p>
                    
                    {orderUser.message && (
                      <div className="mt-3">
                        <p className="text-gray-700 bg-gray-50 rounded-lg p-3">
                          {orderUser.message}
                        </p>
                      </div>
                    )}
                  </div>
                </div>

                <div className="flex justify-between items-center text-sm text-gray-500">
                  <span>
                    Откликнулся: {new Date(orderUser.created_at).toLocaleDateString('ru-RU', {
                      day: 'numeric',
                      month: 'long',
                      year: 'numeric',
                      hour: '2-digit',
                      minute: '2-digit'
                    })}
                  </span>
                  
                  <div className="flex gap-2">
                    {orderUser.status === 'pending' && !order?.executor_id && (
                      <button
                        onClick={() => handleAssignExecutor(orderUser.id)}
                        disabled={assigning === orderUser.id}
                        className="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 disabled:opacity-50 text-sm"
                      >
                        {assigning === orderUser.id ? 'Назначаем...' : 'Назначить исполнителем'}
                      </button>
                    )}
                    
                    {orderUser.status === 'accepted' && (
                      <span className="px-3 py-1 bg-green-100 text-green-800 rounded-full text-xs font-medium">
                        Исполнитель
                      </span>
                    )}
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}