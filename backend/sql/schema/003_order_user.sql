CREATE TABLE order_users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    message TEXT,
    status VARCHAR(20) DEFAULT 'pending' CHECK (status IN ('pending', 'accepted', 'rejected')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(order_id, user_id)
);

ALTER TABLE orders ADD COLUMN executor_id UUID REFERENCES users(id);
ALTER TABLE orders ADD COLUMN executor_assigned_at TIMESTAMP WITH TIME ZONE;

CREATE INDEX idx_order_users_order_id ON order_users(order_id);
CREATE INDEX idx_order_users_user_id ON order_users(user_id);
CREATE INDEX idx_order_users_status ON order_users(status);
CREATE INDEX idx_orders_executor_id ON orders(executor_id);