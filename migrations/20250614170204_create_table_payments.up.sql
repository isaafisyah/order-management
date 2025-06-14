CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(id),
    payment_date TIMESTAMP DEFAULT now(),
    amount NUMERIC(12,2) NOT NULL,
    payment_method VARCHAR(50)
);