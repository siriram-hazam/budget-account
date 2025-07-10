-- This SQL structure is for use with POSTGRESQL
-- Create database named: auth
-- To force migration to version 1, run:
-- migrate -path db/migrations -database "postgres://postgres:123456@localhost:5432/auth?sslmode=disable" force 1
-- To apply all up migrations, run:
-- migrate -path db/migrations -database "postgres://postgres:123456@localhost:5432/auth?sslmode=disable" up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    last_login_at TIMESTAMP NULL
);