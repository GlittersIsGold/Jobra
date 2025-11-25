-- sql/queries/order_users.sql
-- name: CreateOrderUser :one
INSERT INTO order_users (order_id, user_id, message)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetOrderUsers :many
SELECT 
    ou.*,
    u.username as user_username,
    u.email as user_email
FROM order_users ou
JOIN users u ON ou.user_id = u.id
WHERE ou.order_id = $1
ORDER BY ou.created_at DESC;

-- name: GetUserOrderUser :one
SELECT * FROM order_users 
WHERE order_id = $1 AND user_id = $2;

-- name: UpdateOrderUserStatus :one
UPDATE order_users 
SET status = $1, updated_at = NOW()
WHERE id = $2
RETURNING *;

-- name: DeleteOrderUser :exec
DELETE FROM order_users 
WHERE id = $1 AND user_id = $2;

-- name: AssignOrderExecutor :one
UPDATE orders 
SET executor_id = $1, 
    executor_assigned_at = NOW(),
    status = 'in_progress', 
    updated_at = NOW()
WHERE id = $2
RETURNING *;

-- name: RemoveOrderExecutor :one
UPDATE orders 
SET executor_id = NULL,
    executor_assigned_at = NULL,
    status = 'active',
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: CheckOrderUserExists :one
SELECT EXISTS(
    SELECT 1 FROM order_users 
    WHERE order_id = $1 AND user_id = $2
) as has_responded;

-- name: GetOrderWithDetails :one
SELECT 
    o.*,
    u_customer.username as customer_username,
    u_customer.email as customer_email,
    u_executor.username as executor_username,
    u_executor.email as executor_email
FROM orders o
JOIN users u_customer ON o.customer_id = u_customer.id
LEFT JOIN users u_executor ON o.executor_id = u_executor.id
WHERE o.id = $1;