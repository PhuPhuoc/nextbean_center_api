ALTER TABLE `project_intern`
DROP FOREIGN KEY `fk_project_intern_project_id`;

ALTER TABLE `project_intern`
DROP FOREIGN KEY `fk_project_intern_intern_id`;


DROP TABLE IF EXISTS `project_intern`;