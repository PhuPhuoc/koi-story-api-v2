-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `comments` (
    `id` VARCHAR(100) PRIMARY KEY,
    `post_id` VARCHAR(100) NOT NULL,
    `user_id` VARCHAR(100) NOT NULL,
    `content` TEXT NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` DATETIME,
    CONSTRAINT `fk_comments_posts`
    FOREIGN KEY (`post_id`) REFERENCES `posts`(`id`)
    ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_comments_users`
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `comments` DROP FOREIGN KEY `fk_comments_posts`;
ALTER TABLE `comments` DROP FOREIGN KEY `fk_comments_users`;
DROP TABLE IF EXISTS `comments`;
-- +goose StatementEnd
