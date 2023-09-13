-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS groups (
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	name TEXT NOT NULL
);

-- +goose StatementEnd