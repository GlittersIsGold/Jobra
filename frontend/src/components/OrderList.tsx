'use client';

import { Order } from '@/types/order';

interface OrderListProps {
  orders: Order[];
  onEdit?: (order: Order) => void;
  onDelete?: (orderId: string) => void;
  showActions?: boolean;
}

export default function OrderList({ orders, onEdit, onDelete, showActions = true }: OrderListProps) {
  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString('ru-RU', {
      day: 'numeric',
      month: 'long',
      year: 'numeric'
    });
  };

  const formatBudget = (budget: number) => {
    return new Intl.NumberFormat('ru-RU', {
      style: 'currency',
      currency: 'RUB'
    }).format(budget);
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'active':
        return 'bg-green-100 text-green-800';
      case 'completed':
        return 'bg-blue-100 text-blue-800';
      case 'cancelled':
        return 'bg-red-100 text-red-800';
      default:
        return 'bg-gray-100 text-gray-800';
    }
  };

  const getStatusText = (status: string) => {
    switch (status) {
      case 'active':
        return 'Активный';
      case 'completed':
        return 'Завершен';
      case 'cancelled':
        return 'Отменен';
      default:
        return status;
    }
  };

  if (orders.length === 0) {
    return (
      <div className="text-center py-12">
        <div className="text-gray-500 text-lg">Заказов пока нет</div>
        <div className="text-gray-400 mt-2">Создайте первый заказ</div>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      {orders.map((order) => (
        <div key={order.id} className="bg-white rounded-lg shadow-md p-6 border border-gray-200">
          <div className="flex justify-between items-start mb-4">
            <div className="flex-1">
              <h3 className="text-xl font-semibold text-gray-900 mb-2">{order.title}</h3>
              <p className="text-gray-600 mb-3">{order.description}</p>
            </div>
            <span className={`px-3 py-1 rounded-full text-xs font-medium ${getStatusColor(order.status)}`}>
              {getStatusText(order.status)}
            </span>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-4 gap-4 text-sm text-gray-600 mb-4">
            <div>
              <span className="font-medium">Бюджет:</span>
              <div className="text-green-600 font-semibold">{formatBudget(order.budget)}</div>
            </div>
            <div>
              <span className="font-medium">Категория:</span>
              <div>{order.category}</div>
            </div>
            <div>
              <span className="font-medium">Раздел:</span>
              <div>{order.section}</div>
            </div>
            <div>
              <span className="font-medium">Создан:</span>
              <div>{formatDate(order.created_at)}</div>
            </div>
          </div>

          {showActions && onEdit && onDelete && (
            <div className="flex space-x-3 pt-4 border-t border-gray-200">
              <button
                onClick={() => onEdit(order)}
                className="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 text-sm font-medium"
              >
                Редактировать
              </button>
              <button
                onClick={() => onDelete(order.id)}
                className="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700 text-sm font-medium"
              >
                Удалить
              </button>
            </div>
          )}
        </div>
      ))}
    </div>
  );
}