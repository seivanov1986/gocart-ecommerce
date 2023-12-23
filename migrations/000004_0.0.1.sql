-- +goose Up
-- +goose StatementBegin

CREATE TABLE `sefurl` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `url` varchar(255) NOT NULL DEFAULT '',
    `path` varchar(255) NOT NULL DEFAULT '',
    `name` varchar(255) NOT NULL DEFAULT '',
    `type` INTEGER NOT NULL,
    `id_object` INTEGER NOT NULL,
    `template` varchar(255) DEFAULT NULL,
    `created_at` INTEGER NOT NULL,
    `updated_at` INTEGER NOT NULL
)

-- +goose StatementEnd
-- +goose Downs
