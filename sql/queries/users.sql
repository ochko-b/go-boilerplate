-- name: CreateUser :one
INSERT INTO users (email, password_hash, first_name, last_name)
VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 AND is_active = true;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 AND is_active = true;

-- name: UpdateUser :one
UPDATE users
SET first_name = $2, last_name = $3, updated_at = NOW()
WHERE id = $1 AND is_active = true
RETURNING *;

-- name: DeactiviateUser :exec
UPDATE users
SET is_active = false, updated_at = NOW()
WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users
WHERE is_active = true
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;
