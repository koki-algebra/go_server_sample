-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS groups (
	id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
	name TEXT NOT NULL
);

-- +goose StatementEnd