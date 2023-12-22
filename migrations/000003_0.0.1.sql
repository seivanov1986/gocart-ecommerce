-- +goose Up
-- +goose StatementBegin

CREATE TABLE `attribute` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `name` varchar(255) NOT NULL DEFAULT '',
    `sort` INTEGER NOT NULL DEFAULT 0,
    `signature` varchar(256) DEFAULT NULL
);

CREATE TABLE `product` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `name` varchar(256) NOT NULL DEFAULT '',
    `content` text DEFAULT NULL,
    `id_meta` INTEGER DEFAULT NULL,
    `sort` INTEGER NOT NULL DEFAULT 0,
    `price` varchar(50) DEFAULT NULL,
    `id_image` INTEGER DEFAULT NULL,
    `disabled` INTEGER NOT NULL DEFAULT 0,
    `created_at` INTEGER NOT NULL,
    `updated_at` INTEGER NOT NULL
);

CREATE TABLE `page` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `name` varchar(255) NOT NULL DEFAULT '',
    `content` text DEFAULT NULL,
    `id_meta` INTEGER DEFAULT NULL,
    `type` INTEGER NOT NULL DEFAULT 0,
    `sort` INTEGER NOT NULL DEFAULT 0,
    `short_content` text DEFAULT NULL,
    `id_image` INTEGER DEFAULT NULL,
    `created_at` INTEGER NOT NULL,
    `updated_at` INTEGER NOT NULL
);

CREATE TABLE `category` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `name` varchar(256) DEFAULT NULL,
    `id_parent` INTEGER DEFAULT NULL,
    `content` text DEFAULT NULL,
    `id_image` INTEGER DEFAULT NULL,
    `id_meta` INTEGER DEFAULT NULL,
    `sort` INTEGER DEFAULT NULL,
    `price` varchar(50) DEFAULT NULL,
    `disabled` INTEGER NOT NULL DEFAULT 0,
    `created_at` INTEGER NOT NULL,
    `updated_at` INTEGER NOT NULL
);

-- +goose StatementEnd
-- +goose Downs
