-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
	id BIGSERIAL NOT NULL PRIMARY KEY,
	username text UNIQUE
);

INSERT INTO users (username)
	VALUES 
	('Kenny'),
	('Benny'),
	('Sunny'),
	('Gunny');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
