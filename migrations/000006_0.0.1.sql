-- +goose Up
-- +goose StatementBegin

CREATE TABLE `image` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `name` varchar(256) NOT NULL DEFAULT '',
    `id_parent` INTEGER NOT NULL DEFAULT 0,
    `path` varchar(512) NOT NULL DEFAULT '',
    `ftype` INTEGER NOT NULL DEFAULT 0,
    `created_at` INTEGER NOT NULL
)

-- +goose StatementEnd
-- +goose Downs
