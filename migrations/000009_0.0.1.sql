-- +goose Up
-- +goose StatementBegin

ALTER TABLE `image` ADD uid text;
ALTER TABLE `image` ADD origin_path text;

-- +goose StatementEnd
-- +goose Downs
