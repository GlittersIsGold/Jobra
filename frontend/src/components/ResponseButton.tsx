'use client';

import { useState, useEffect } from 'react';
import { OrderUserService } from '@/lib/orderUser';

interface ResponseButtonProps {
  orderId: string;
  customerId: string;
  currentUserId: string;
}

export default function ResponseButton({ orderId, customerId, currentUserId }: ResponseButtonProps) {
  const [hasResponded, setHasResponded] = useState(false);
  const [showModal, setShowModal] = useState(false);
  const [message, setMessage] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  useEffect(() => {
    const checkResponse = async () => {
      try {
        const response = await OrderUserService.checkOrderUserExists(orderId);
        setHasResponded(response.has_responded);
      } catch (err) {
        console.error('Failed to check response status:', err);
      }
    };

    if (currentUserId && customerId !== currentUserId) {
      checkResponse();
    }
  }, [orderId, customerId, currentUserId]);

  const handleResponse = async () => {
    setLoading(true);
    setError('');

    try {
      await OrderUserService.createOrderUser(orderId, {
        message: message || undefined
      });
      setHasResponded(true);
      setShowModal(false);
      setMessage('');
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Ошибка отклика');
    } finally {
      setLoading(false);
    }
  };

  if (customerId === currentUserId) {
    return null;
  }

  if (hasResponded) {
    return (
      <button
        disabled
        className="px-4 py-2 bg-gray-500 text-white rounded-md text-sm font-medium cursor-not-allowed"
      >
        Вы откликнулись
      </button>
    );
  }

  return (
    <>
      <button
        onClick={() => setShowModal(true)}
        className="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 text-sm font-medium"
      >
        Откликнуться
      </button>

      {showModal && (
  <div className="fixed inset-0 bg-opacity-50 flex items-center justify-center z-50 p-4">
    <div className="bg-white rounded-lg border border-gray-300 shadow-lg w-full max-w-md">
      <div className="p-6">
        <h3 className="text-lg font-semibold mb-4">Отклик на заказ</h3>
        
        {error && (
          <div className="mb-4 p-3 bg-red-100 border border-red-400 text-red-700 rounded">
            {error}
          </div>
        )}

        <div className="mb-4">
          <label className="block text-sm font-medium text-gray-700 mb-2">
            Сообщение (необязательно)
          </label>
          <textarea
            value={message}
            onChange={(e) => setMessage(e.target.value)}
            placeholder="Расскажите, почему вы подходите для этого заказа..."
            className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 h-24"
            maxLength={500}
          />
          <div className="text-right text-sm text-gray-500 mt-1">
            {message.length}/500
          </div>
        </div>

        <div className="flex justify-end space-x-3">
          <button
            onClick={() => setShowModal(false)}
            disabled={loading}
            className="px-4 py-2 text-gray-600 hover:text-gray-800 disabled:opacity-50"
          >
            Отмена
          </button>
          <button
            onClick={handleResponse}
            disabled={loading}
            className="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 disabled:opacity-50"
          >
            {loading ? 'Отправка...' : 'Отправить отклик'}
          </button>
        </div>
      </div>
    </div>
  </div>
)}
    </>
  );
}