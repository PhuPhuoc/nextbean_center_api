CREATE TABLE IF NOT EXISTS `project_manager` (
  `project_id` int NOT NULL,
  `account_id` varchar(255) NOT NULL,
  PRIMARY KEY (`project_id`, `account_id`)
);

ALTER TABLE `project_manager`
ADD CONSTRAINT `fk_project_manager_project_id`
FOREIGN KEY (`project_id`) REFERENCES `project`(`id`);

ALTER TABLE `project_manager`
ADD CONSTRAINT `fk_project_manager_account_id`
FOREIGN KEY (`account_id`) REFERENCES `account`(`id`);