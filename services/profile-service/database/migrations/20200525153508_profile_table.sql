-- +goose Up
-- +goose StatementBegin
CREATE TABLE profiles (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NULL,
    `address` varchar(255) NULL,
    `work_at` varchar(255) NULL,
    `phone_number` varchar(255) NULL,
    `gender` varchar(255) NULL,
    `date_of_birth` varchar(255) NULL,
    `is_active` boolean DEFAULT false,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE profiles;
-- +goose StatementEnd
