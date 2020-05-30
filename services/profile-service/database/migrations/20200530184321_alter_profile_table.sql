-- +goose Up
-- +goose StatementBegin
ALTER TABLE `profiles`ADD COLUMN `user_id` int(10) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT TABLE `profiles` DROP COLUMN `user_id`;
-- +goose StatementEnd
