-- name: CreateItem :one
INSERT INTO items (title, description, quantity, price, order_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, title, description, quantity, price, order_id;

-- name: GetItem :one
SELECT id, title, description, quantity, price, order_id
FROM items
WHERE id = $1;

-- name: GetItemsByOrderID :many
SELECT id, title, description, quantity, price, order_id
FROM items
WHERE order_id = $1;

-- name: UpdateItem :exec
UPDATE items
SET title = $1, description = $2, quantity = $3, price = $4
WHERE id = $5;

-- name: DeleteItem :exec
DELETE FROM items
WHERE id = $1;