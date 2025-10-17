-- 003_create_payments.up.sql

CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    order_uid TEXT REFERENCES orders(order_uid),
    transaction TEXT,
    request_id TEXT,
    currency TEXT,
    provider TEXT,
    amount INT,
    payment_dt BIGINT,
    bank TEXT,
    delivery_cost INT,
    goods_total INT,
    custom_fee INT
);