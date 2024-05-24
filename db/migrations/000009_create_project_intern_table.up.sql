CREATE TABLE IF NOT EXISTS `project_intern` (
  `project_id` int NOT NULL,
  `intern_id` varchar(255) NOT NULL,
  PRIMARY KEY (`project_id`, `intern_id`)
);

ALTER TABLE `project_intern`
ADD CONSTRAINT `fk_project_intern_project_id`
FOREIGN KEY (`project_id`) REFERENCES `project`(`id`);

ALTER TABLE `project_intern`
ADD CONSTRAINT `fk_project_intern_intern_id`
FOREIGN KEY (`intern_id`) REFERENCES `intern`(`id`);