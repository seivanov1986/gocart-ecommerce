-- +goose Up
-- +goose StatementBegin

CREATE TABLE `attribute_to_product` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `id_product` INTEGER NOT NULL,
    `id_attribute` INTEGER NOT NULL,
    `value` varchar(256) NOT NULL DEFAULT ''
);

CREATE TABLE `product_to_category` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `id_product` INTEGER NOT NULL,
    `id_category` INTEGER NOT NULL,
    `main_category` INTEGER DEFAULT NULL
)

-- +goose StatementEnd
-- +goose Downs
