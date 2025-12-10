-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
                        id BIGSERIAL PRIMARY KEY,
                        public_order_number VARCHAR(20) UNIQUE NOT NULL,
                        user_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    quantity INTEGER NOT NULL CHECK (quantity > 0),

-- Изменено на INTEGER с маппингом на enum
    status INTEGER NOT NULL DEFAULT 1 CHECK (
        status BETWEEN 0 AND 12  -- от ORDER_STATUS_UNSPECIFIED до ORDER_STATUS_FAILED
        ),
    total_amount DECIMAL(15,2) NOT NULL CHECK (total_amount >= 0),

-- Связь со складом
    warehouse_id BIGINT NOT NULL,

-- Timestamps
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Создаем индекс для статуса если будет много фильтраций по нему
CREATE INDEX idx_orders_status ON orders(status);
CREATE INDEX idx_orders_user_id ON orders(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
