-- name: CreateUser :one
INSERT INTO users (username, email, password_hash)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users 
WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users 
WHERE username = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users 
ORDER BY username;

-- name: UpdateUser :one
UPDATE users 
SET username = $2, email = $3, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: UpdateUserPassword :one
UPDATE users 
SET password_hash = $2, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users 
WHERE id = $1;