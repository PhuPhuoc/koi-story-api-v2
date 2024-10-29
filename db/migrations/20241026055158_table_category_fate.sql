-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `category_fate` (
    `category_id` VARCHAR(100),
    `fate_id` VARCHAR(100),
    PRIMARY KEY (`category_id`, `fate_id`),
    CONSTRAINT `fk_category_face_category`
        FOREIGN KEY (`category_id`) REFERENCES `categories`(`id`)
        ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_category_face_fate`
        FOREIGN KEY (`fate_id`) REFERENCES `fates`(`id`)
        ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `category_fate` DROP FOREIGN KEY `fk_category_face_category`;
ALTER TABLE `category_fate` DROP FOREIGN KEY `fk_category_face_fate`;
DROP TABLE IF EXISTS `category_fate`;
-- +goose StatementEnd
