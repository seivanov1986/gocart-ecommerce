-- +goose Up
-- +goose StatementBegin

INSERT INTO user(login, password) VALUES('root', '$2a$10$MV8/hMpBwKzMrsn5pYHsV.emsnS61c9W6IozuoFYudo5gNHf8rzNa')

-- +goose StatementEnd
-- +goose Downs
