ALTER TABLE `project_manager`
DROP FOREIGN KEY `fk_project_manager_project_id`;

ALTER TABLE `project_manager`
DROP FOREIGN KEY `fk_project_manager_account_id`;

DROP TABLE IF EXISTS `project_manager`;