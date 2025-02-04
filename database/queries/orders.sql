-- name: CreateOrder :one
INSERT INTO orders (account_email, delivery_address, cancel_reason, total_price, paid, submitted, completed, canceled)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, account_email, delivery_address, cancel_reason, total_price, paid, submitted, completed, canceled;

-- name: GetOrder :one
SELECT id, account_email, delivery_address, cancel_reason, total_price, paid, submitted, completed, canceled
FROM orders
WHERE id = $1;

-- name: UpdateOrder :exec
UPDATE orders
SET account_email = $1, delivery_address = $2, cancel_reason = $3, total_price = $4, paid = $5, submitted = $6, completed = $7, canceled = $8
WHERE id = $9;

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE id = $1;