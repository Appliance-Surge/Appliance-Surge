-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    image VARCHAR(255),
    banner VARCHAR(255),
    score INTEGER DEFAULT 0,
    location VARCHAR(255),
    about TEXT,
    website VARCHAR(255),
    social_id VARCHAR(255),
    social_type VARCHAR(255),
    email VARCHAR(100) NOT NULL UNIQUE,
    email_verified_at TIMESTAMP,
    password VARCHAR(255),
    remember_token VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
