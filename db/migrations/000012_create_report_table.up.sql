CREATE TABLE IF NOT EXISTS `report` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `project_id` int NOT NULL,
  `intern_id` varchar(255) NOT NULL,
  `intern_name` varchar(255) NOT NULL,
  `content` text NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `deleted_ad` datetime
);

ALTER TABLE `report`
ADD CONSTRAINT `fk_report_project_id`
FOREIGN KEY (`project_id`) REFERENCES `project`(`id`);

ALTER TABLE `report`
ADD CONSTRAINT `fk_report_intern_id`
FOREIGN KEY (`intern_id`) REFERENCES `intern`(`id`);