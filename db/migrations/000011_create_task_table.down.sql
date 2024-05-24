ALTER TABLE `task`
DROP FOREIGN KEY `fk_task_project_id`;

ALTER TABLE `task`
DROP FOREIGN KEY `fk_task_assigned_to`;

DROP TABLE IF EXISTS `task` 

