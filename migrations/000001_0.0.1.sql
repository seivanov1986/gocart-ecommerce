-- +goose Up
-- +goose StatementBegin

CREATE TABLE `user` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `login` varchar(50) DEFAULT NULL,
    `email` varchar(50) DEFAULT NULL,
    `password` char(60) DEFAULT NULL,
    `created_at` INTEGER DEFAULT NULL,
    `active` tinyint(1) DEFAULT '0',
    `activation_hash` char(64) DEFAULT NULL,
    `hash_sended_at` INTEGER DEFAULT NULL,
    `forgot_hash` char(64) DEFAULT NULL,
    `recreate_at` INTEGER DEFAULT NULL
)

-- +goose StatementEnd
-- +goose Downs
