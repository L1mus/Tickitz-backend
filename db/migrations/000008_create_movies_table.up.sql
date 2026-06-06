CREATE TABLE movies
(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    duration INTERVAL,
    poster VARCHAR(255),
    realase_data DATE,
    synopsis TEXT,
    category VARCHAR(155),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);