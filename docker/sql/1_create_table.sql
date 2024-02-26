CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
	id UUID DEFAULT uuid_generate_v4(),
	name TEXT NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS groups (
	id UUID DEFAULT uuid_generate_v4(),
	name TEXT NOT NULL,
	PRIMARY KEY (id)
);