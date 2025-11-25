'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { OrderService } from '@/lib/order';
import { Order } from '@/types/order';
import OrderList from '@/components/OrderList';
import OrderFilters from '@/components/OrderFilters';

interface Filters {
  search: string;
  category: string;
  section: string;
  minBudget: string;
  maxBudget: string;
  status: string;
  sortBy: string;
}

export default function AllOrdersPage() {
  const [orders, setOrders] = useState<Order[]>([]);
  const [filteredOrders, setFilteredOrders] = useState<Order[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const router = useRouter();

  const [filters, setFilters] = useState<Filters>({
    search: '',
    category: '',
    section: '',
    minBudget: '',
    maxBudget: '',
    status: '',
    sortBy: 'newest'
  });

  useEffect(() => {
    const fetchAllOrders = async () => {
      const token = localStorage.getItem('jwt_token');
      
      if (!token) {
        router.push('/login');
        return;
      }

      try {
        const response = await OrderService.getAllOrders();
        setOrders(response.orders);
        setFilteredOrders(response.orders);
      } catch (error) {
        console.error('All orders fetch error:', error);
        setError('Ошибка загрузки заказов');
      } finally {
        setLoading(false);
      }
    };

    fetchAllOrders();
  }, [router]);

  useEffect(() => {
    let result = [...orders];

    if (filters.search) {
      result = result.filter(order => 
        order.title.toLowerCase().includes(filters.search.toLowerCase()) ||
        order.description.toLowerCase().includes(filters.search.toLowerCase())
      );
    }

    if (filters.category) {
      result = result.filter(order => order.category === filters.category);
    }

    if (filters.section) {
      result = result.filter(order => order.section === filters.section);
    }

    if (filters.minBudget) {
      result = result.filter(order => order.budget >= parseInt(filters.minBudget));
    }
    if (filters.maxBudget) {
      result = result.filter(order => order.budget <= parseInt(filters.maxBudget));
    }

    if (filters.status) {
      result = result.filter(order => order.status === filters.status);
    }

    result.sort((a, b) => {
      switch (filters.sortBy) {
        case 'newest':
          return new Date(b.created_at).getTime() - new Date(a.created_at).getTime();
        case 'oldest':
          return new Date(a.created_at).getTime() - new Date(b.created_at).getTime();
        case 'budget_high':
          return b.budget - a.budget;
        case 'budget_low':
          return a.budget - b.budget;
        default:
          return 0;
      }
    });

    setFilteredOrders(result);
  }, [filters, orders]);

  const handleFilterChange = (newFilters: Partial<Filters>) => {
    setFilters(prev => ({ ...prev, ...newFilters }));
  };

  const handleResetFilters = () => {
    setFilters({
      search: '',
      category: '',
      section: '',
      minBudget: '',
      maxBudget: '',
      status: '',
      sortBy: 'newest'
    });
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-gray-100 flex items-center justify-center">
        <div className="text-xl">Загрузка заказов...</div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-100 py-8">
      <div className="max-w-7xl mx-auto px-4">
        {error && (
          <div className="mb-6 p-4 bg-red-100 border border-red-400 text-red-700 rounded">
            {error}
          </div>
        )}
        
        <div className="flex justify-between items-center mb-8">
          <div>
            <h1 className="text-3xl font-bold text-gray-900">Все заказы</h1>
            <p className="text-gray-600 mt-2">Найдите подходящий проект для работы</p>
          </div>
        </div>

        <OrderFilters
          filters={filters}
          onFilterChange={handleFilterChange}
          onReset={handleResetFilters}
          orders={orders}
        />

        <div className="mb-6 flex justify-between items-center">
          <div className="text-gray-600">
            Найдено заказов: <span className="font-semibold">{filteredOrders.length}</span>
            {filteredOrders.length !== orders.length && (
              <span className="text-gray-400 ml-2">
                (из {orders.length})
              </span>
            )}
          </div>
        </div>

        <OrderList
          orders={filteredOrders}
          showActions={false}
        />
      </div>
    </div>
  );
}