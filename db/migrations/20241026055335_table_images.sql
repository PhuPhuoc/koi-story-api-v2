-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `images` (
    `id` VARCHAR(100) PRIMARY KEY,
    `post_id` VARCHAR(100) NOT NULL,
    `image_url` VARCHAR(100),
    `deleted_at` DATETIME,
    CONSTRAINT `fk_images_posts`
    FOREIGN KEY (`post_id`) REFERENCES `posts`(`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `images` DROP FOREIGN KEY `fk_images_posts`;
DROP TABLE IF EXISTS `images`;
-- +goose StatementEnd
