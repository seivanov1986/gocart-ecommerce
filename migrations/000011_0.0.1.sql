-- +goose Up
-- +goose StatementBegin

CREATE UNIQUE INDEX sefurl_name_type
    ON sefurl (name, type);

-- +goose StatementEnd
-- +goose Downs
