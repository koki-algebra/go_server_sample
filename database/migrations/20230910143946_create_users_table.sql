-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	name TEXT
);

-- +goose StatementEnd