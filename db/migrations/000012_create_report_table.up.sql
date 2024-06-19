CREATE TABLE IF NOT EXISTS `report` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `task_id` varchar(255) NOT NULL,
  `intern_id` varchar(255) NOT NULL,
  `content` text NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime
);
  
ALTER TABLE `report`
ADD CONSTRAINT `fk_report_task_id`
FOREIGN KEY (`task_id`) REFERENCES `task`(`id`);

ALTER TABLE `report`
ADD CONSTRAINT `fk_report_intern_id`
FOREIGN KEY (`intern_id`) REFERENCES `intern`(`id`);