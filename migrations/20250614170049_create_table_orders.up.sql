CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    order_date TIMESTAMP DEFAULT now(),
    status VARCHAR(50) DEFAULT 'pending',
    total_amount NUMERIC(12,2) NOT NULL
);