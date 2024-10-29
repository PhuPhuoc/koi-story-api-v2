-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `users` (
    `id` VARCHAR(100) PRIMARY KEY,
    `fb_id` VARCHAR(100),
    `email` VARCHAR(30) NOT NULL,
    `password` VARCHAR(30) NOT NULL,
    `name` VARCHAR(100) NOT NULL,
    `avatar` VARCHAR(255),
    `role` ENUM('admin', 'user', 'member'),
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` DATETIME,
    UNIQUE(`email`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `users`;
-- +goose StatementEnd
