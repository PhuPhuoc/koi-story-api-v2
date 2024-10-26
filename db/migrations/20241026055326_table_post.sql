-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `posts` (
    `id` VARCHAR(100) PRIMARY KEY,
    `user_id` VARCHAR(100),
    `post_type` ENUM('blog','consult','market') NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` DATETIME,
    CONSTRAINT `fk_posts_users`
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `posts` DROP FOREIGN KEY `fk_posts_users`;
DROP TABLE IF EXISTS `posts`;
-- +goose StatementEnd
