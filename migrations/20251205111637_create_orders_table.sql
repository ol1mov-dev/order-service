-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
    id BIGSERIAL PRIMARY KEY,
    public_order_number VARCHAR(20) UNIQUE NOT NULL,
    user_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    quantity INTEGER NOT NULL CHECK (quantity > 0),

    status VARCHAR(50) NOT NULL DEFAULT 'created' CHECK (
        status IN (
                   'created',           -- заказ создан
                   'pending_payment',   -- ожидает оплаты
                   'paid',              -- оплачен
                   'processing',        -- в обработке
                   'ready_for_shipment',-- готов к отгрузке
                   'shipped',           -- отправлен
                   'in_transit',        -- в пути
                   'out_for_delivery',  -- у курьера
                   'delivered',         -- доставлен
                   'cancelled',         -- отменен
                   'refunded',          -- возвращен
                   'failed'             -- не удался
            )
        ),
    total_amount DECIMAL(15,2) NOT NULL CHECK (total_amount >= 0),

    -- Связь со складом
    warehouse_id BIGINT,

    -- Timestamps
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS warehouses;
-- +goose StatementEnd
