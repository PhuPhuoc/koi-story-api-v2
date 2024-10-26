-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `blogs` (
    `id` VARCHAR(100) PRIMARY KEY,
    `post_id` VARCHAR(100) NOT NULL,
    `category_id` VARCHAR(100) NOT NULL,
    `author_name` VARCHAR(100) NOT NULL,
    `content` TEXT NOT NULL,
    CONSTRAINT `fk_blogs_posts`
    FOREIGN KEY (`post_id`) REFERENCES `posts`(`id`)
    ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_blogs_categories`
    FOREIGN KEY (`category_id`) REFERENCES `categories`(`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `blogs` DROP FOREIGN KEY `fk_blogs_posts`;
ALTER TABLE `blogs` DROP FOREIGN KEY `fk_blogs_categories`;
DROP TABLE IF EXISTS `blogs`;
-- +goose StatementEnd
