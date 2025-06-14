CREATE TABLE logs (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    action VARCHAR(255),
    created_at TIMESTAMP DEFAULT now()
);