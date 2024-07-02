CREATE TABLE IF NOT EXISTS `project` (
  `id` VARCHAR(255) PRIMARY KEY,
  `name` varchar(255) NOT NULL,
  `status` enum('not_started', 'in_progress', 'completed', 'cancel'),
  `description` text,
  `Est_start_time` date NOT NULL,
  `Est_completion_time` date NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime
);