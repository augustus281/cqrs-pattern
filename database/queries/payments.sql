-- name: CreatePayment :one
INSERT INTO payments (paymented_time)
VALUES ($1)
RETURNING payment_id, paymented_time;

-- name: GetPayment :one
SELECT payment_id, paymented_time
FROM payments
WHERE payment_id = $1;

-- name: UpdatePayment :exec
UPDATE payments
SET paymented_time = $1
WHERE payment_id = $2;

-- Delete a payment
-- name: DeletePayment :exec
DELETE FROM payments
WHERE payment_id = $1;

-- Get all payments for a given order
-- name: GetPaymentsByOrderID :many
SELECT payment_id, paymented_time
FROM payments
WHERE paymented_time >= $1;