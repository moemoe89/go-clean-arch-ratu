
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `users` (
    `id` char(20) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `name` varchar(50) NOT NULL DEFAULT '',
    `email` varchar(50) NOT NULL DEFAULT '',
    `phone` varchar(20) NOT NULL DEFAULT '',
    `address` text NOT NULL DEFAULT '',
    PRIMARY KEY (`id`),
    UNIQUE KEY `id` (`id`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF exists users;
