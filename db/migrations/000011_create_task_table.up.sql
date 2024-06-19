CREATE TABLE IF NOT EXISTS `task` (
  `id` VARCHAR(255) PRIMARY KEY,
  `project_id` varchar(255) NOT NULL,
  `assigned_to` varchar(255) NOT NULL,
  `is_approved` TINYINT(1) NOT NULL DEFAULT 0,
  `status` enum('todo','inprogress','done') NOT NULL DEFAULT 'todo',
  `name` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `start_date` datetime,
  `end_date` datetime,
  `estimated_effort` varchar(255),
  `actual_effort` varchar(255),
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime
);

ALTER TABLE `task`
ADD CONSTRAINT `fk_task_project_id`
FOREIGN KEY (`project_id`) REFERENCES `project`(`id`);

ALTER TABLE `task`
ADD CONSTRAINT `fk_task_assigned_to`
FOREIGN KEY (`assigned_to`) REFERENCES `intern`(`id`);