-- +goose Up
-- +goose StatementBegin
INSERT INTO
	users ("name")
VALUES
	('John Smith'),
	('Emily Johnson'),
	('Michael Williams'),
	('Sarah Brown'),
	('David Miller');

-- +goose StatementEnd