CREATE SCHEMA service_order

CREATE TABLE service_order.orders (
    id BIGSERIAL PRIMARY KEY,
    product_id BIGINT NOT NULL,
    quantity BIGINT NOT NULL,
    total_price BIGINT NOT NULL,
    status INTEGER NOT NULL,
    time_created TIMESTAMP NOT NULL DEFAULT NOW()
);