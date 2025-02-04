-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE "orders" (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "account_email" VARCHAR(320) NOT NULL,
    "delivery_address" TEXT NOT NULL,
    "cancel_reason" TEXT NOT NULL,
    "total_price" DOUBLE PRECISION NOT NULL DEFAULT 0,
    "paid" BOOLEAN DEFAULT FALSE,
    "submitted" BOOLEAN DEFAULT FALSE,
    "completed" BOOLEAN DEFAULT FALSE,
    "canceled" BOOLEAN DEFAULT FALSE
);

CREATE TABLE "items" (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "title" TEXT NOT NULL,
    "description" TEXT,
    "quantity" BIGINT NOT NULL DEFAULT 1,
    "price" DOUBLE PRECISION NOT NULL DEFAULT 0,
    "order_id" UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE
);

CREATE TABLE "payments" (
    "payment_id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "paymented_time" TIMESTAMPTZ NOT NULL,
    "order_id" UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE "items";
DROP TABLE "orders";
DROP TABLE "payments";