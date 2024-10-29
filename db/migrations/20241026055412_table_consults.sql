-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `consults` (
    `post_id` VARCHAR(100) PRIMARY KEY,
    `user_fengshui_id` VARCHAR(100) NOT NULL,
    `title` VARCHAR(200) NOT NULL,
    `content` TEXT NOT NULL,
    CONSTRAINT `fk_consults_posts`
    FOREIGN KEY (`post_id`) REFERENCES `posts`(`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `consults` DROP FOREIGN KEY `fk_consults_posts`;
DROP TABLE IF EXISTS `consults`;
-- +goose StatementEnd
