CREATE TABLE IF NOT EXISTS `account` (
    `id` VARCHAR(255) PRIMARY KEY,
    `user_name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `role` ENUM('admin', 'manager', 'pm', 'user'),
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` DATETIME,
    UNIQUE KEY `account_email_unique` (`email`)
);
