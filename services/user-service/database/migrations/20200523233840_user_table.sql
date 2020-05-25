-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `password` varchar(255) NULL,
    `email` varchar(255) NULL,
    `username` varchar(255) NULL,
    `role` varchar(255) NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
