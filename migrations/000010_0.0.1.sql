-- +goose Up
-- +goose StatementBegin

CREATE TABLE `image_to_category` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `id_category` INTEGER DEFAULT NULL,
    `id_image` INTEGER DEFAULT NULL
);

CREATE TABLE `image_to_product` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `id_product` INTEGER DEFAULT NULL,
    `id_image` INTEGER DEFAULT NULL
)

-- +goose StatementEnd
-- +goose Downs
