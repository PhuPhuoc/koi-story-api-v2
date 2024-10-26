-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `markets` (
    `id` VARCHAR(100) PRIMARY KEY,
    `post_id` VARCHAR(100) NOT NULL,
    `seller_id` VARCHAR(100) NOT NULL,
    `product_name` VARCHAR(255) NOT NULL,
    `price` INT NOT NULL,
    `product_type` ENUM('koi', 'decoration', 'other'),
    `color` VARCHAR(255) NOT NULL,
    `origin` VARCHAR(255) NOT NULL,
    `describe` TEXT NOT NULL,
    CONSTRAINT `fk_markets_posts`
    FOREIGN KEY (`post_id`) REFERENCES `posts`(`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `markets` DROP FOREIGN KEY `fk_markets_posts`;
DROP TABLE IF EXISTS `markets`;
-- +goose StatementEnd
