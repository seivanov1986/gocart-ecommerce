-- +goose Up
-- +goose StatementBegin

CREATE TABLE `category_closure` (
    `parent` INTEGER NOT NULL,
    `child` INTEGER NOT NULL,
    `depth` INTEGER NOT NULL,
    PRIMARY KEY (`parent`,`child`)
)

-- +goose StatementEnd
-- +goose Downs
