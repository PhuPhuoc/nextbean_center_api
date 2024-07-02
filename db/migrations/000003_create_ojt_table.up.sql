CREATE TABLE IF NOT EXISTS `ojt` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `semester` VARCHAR(255) NOT NULL,
    `university` VARCHAR(255) NOT NULL,
    `start_at` DATETIME,
    `end_at` DATETIME,
    `status` enum('not_started', 'in_progress', 'completed'),
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` DATETIME
);