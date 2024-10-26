-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `categories` (
    `id` VARCHAR(100) PRIMARY KEY,
    `name` VARCHAR(200) NOT NULL,
    `describe` TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `categories`;
-- +goose StatementEnd
