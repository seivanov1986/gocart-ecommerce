-- +goose Up
-- +goose StatementBegin

CREATE TABLE `meta` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `title` varchar(255) DEFAULT NULL,
    `keywords` varchar(256) DEFAULT NULL,
    `description` varchar(256) DEFAULT NULL
)

-- +goose StatementEnd
-- +goose Downs
