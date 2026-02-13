CREATE TABLE IF NOT EXISTS health_checks (
    id SERIAL PRIMARY KEY,
    service_id INT REFERENCES services(id) ON DELETE CASCADE,
    status TEXT NOT NULL,
    latency_ms INT,
    checked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
