-- +goose Up
-- +goose StatementBegin

DROP INDEX sefurl_name_type;
CREATE UNIQUE INDEX sefurl_url ON sefurl (url);

-- +goose StatementEnd
-- +goose Downs
