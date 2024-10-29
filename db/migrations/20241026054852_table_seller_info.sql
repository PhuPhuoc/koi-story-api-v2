-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `seller_info` (
    `user_id` VARCHAR(100) PRIMARY KEY,
    `phone_number` VARCHAR(20) NOT NULL,
    `location` VARCHAR(40) NOT NULL,
    `address` VARCHAR(100) NOT NULL,
    CONSTRAINT `fk_user_seller_info`
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `seller_info` DROP FOREIGN KEY `fk_user_seller_info`;
DROP TABLE IF EXISTS `seller_info`;
-- +goose StatementEnd
