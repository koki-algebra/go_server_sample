-- +goose Up
CREATE TABLE IF NOT EXISTS users (
	id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
	name TEXT NOT NULL
);