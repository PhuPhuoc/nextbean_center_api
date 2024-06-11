CREATE TABLE IF NOT EXISTS `project` (
  `id` VARCHAR(255) PRIMARY KEY,
  `name` varchar(255) NOT NULL,
  `status` enum('not_start', 'doing', 'done', 'cancel'),
  `description` text,
  `start_date` datetime NOT NULL,
  `duration` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime
);