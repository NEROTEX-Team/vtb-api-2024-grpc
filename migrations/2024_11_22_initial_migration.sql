CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT TIMEZONE('utc', now()),
    updated_at TIMESTAMP NOT NULL DEFAULT TIMEZONE('utc', now()),
    deleted_at TIMESTAMP
);