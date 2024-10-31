-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `fates` (
    `id` VARCHAR(100) PRIMARY KEY,
    `element` VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS `user_fengshui` (
    `id` VARCHAR(100) PRIMARY KEY,
    `user_id` VARCHAR(100),
    `gender` BOOLEAN NOT NULL,
    `year_of_birth` SMALLINT NOT NULL,
    `fate_id` VARCHAR(100) NOT NULL,
    `direction` VARCHAR(50),
    `cung_phi` VARCHAR(50),
    CONSTRAINT `fk_user_fengshui_users`
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
    ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_user_fengshui_fate`
    FOREIGN KEY (`fate_id`) REFERENCES `fates`(`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `user_fengshui` DROP FOREIGN KEY `fk_user_fengshui_users`;
ALTER TABLE `user_fengshui` DROP FOREIGN KEY `fk_user_fengshui_fate`;
DROP TABLE IF EXISTS `fates`;
DROP TABLE IF EXISTS `user_fengshui`;
-- +goose StatementEnd
