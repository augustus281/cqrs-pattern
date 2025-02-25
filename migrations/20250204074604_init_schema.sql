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

CREATE TABLE "events" (
    event_id UUID DEFAULT gen_random_uuid(),
    aggregate_id   VARCHAR(250) NOT NULL CHECK ( aggregate_id <> '' ),
    aggregate_type VARCHAR(250) NOT NULL CHECK ( aggregate_type <> '' ),
    event_type     VARCHAR(250) NOT NULL CHECK ( event_type <> '' ),
    data           BYTEA,
    metadata       BYTEA,
    version        SERIAL       NOT NULL,
    timestamp      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (aggregate_id, version)
) PARTITION BY HASH (aggregate_id);

CREATE INDEX IF NOT EXISTS aggregate_id_aggregate_version_idx ON events(aggregate_id, version);

CREATE TABLE IF NOT EXISTS events_partition_hash_1 PARTITION OF "events"
    FOR VALUES WITH (MODULUS 3, REMAINDER 0);

CREATE TABLE IF NOT EXISTS events_partition_hash_2 PARTITION OF "events"
    FOR VALUES WITH (MODULUS 3, REMAINDER 1);

CREATE TABLE IF NOT EXISTS events_partition_hash_3 PARTITION OF "events"
    FOR VALUES WITH (MODULUS 3, REMAINDER 2);

CREATE TABLE "snapshots" (
    snapshot_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    aggregate_id   VARCHAR(250) UNIQUE NOT NULL CHECK ( aggregate_id <> '' ),
    aggregate_type VARCHAR(250)        NOT NULL CHECK ( aggregate_type <> '' ),
    data           BYTEA,
    metadata       BYTEA,
    version        SERIAL              NOT NULL,
    timestamp      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (aggregate_id)
);

CREATE INDEX IF NOT EXISTS aggregate_id_aggregate_version_idx ON snapshots(aggregate_id, version);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE "items";
DROP TABLE "orders";
DROP TABLE "payments";
DROP TABLE "events";
DROP TABLE "snapshots";