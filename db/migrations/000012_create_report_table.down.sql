ALTER TABLE `report`
DROP FOREIGN KEY `fk_report_task_id`;

ALTER TABLE `report`
DROP FOREIGN KEY `fk_report_intern_id`;

DROP TABLE IF EXISTS `report`