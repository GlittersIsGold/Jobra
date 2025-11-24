-- name: CreateOrder :one
INSERT INTO orders (customer_id, title, description, budget, category, section, status)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetOrderByID :one
SELECT 
    o.*,
    u.username as customer_username,
    u.email as customer_email
FROM orders o
JOIN users u ON o.customer_id = u.id
WHERE o.id = $1;

-- name: GetOrdersByCustomer :many
SELECT 
    o.*,
    u.username as customer_username,
    u.email as customer_email
FROM orders o
JOIN users u ON o.customer_id = u.id
WHERE o.customer_id = $1
ORDER BY o.created_at DESC;

-- name: GetAllOrders :many
SELECT 
    o.*,
    u.username as customer_username,
    u.email as customer_email
FROM orders o
JOIN users u ON o.customer_id = u.id
ORDER BY o.created_at DESC;

-- name: UpdateOrder :one
UPDATE orders 
SET title = $2, description = $3, budget = $4, category = $5, section = $6, status = $7, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM orders 
WHERE id = $1 AND customer_id = $2;

-- name: DeleteOrderByID :exec
DELETE FROM orders 
WHERE id = $1;