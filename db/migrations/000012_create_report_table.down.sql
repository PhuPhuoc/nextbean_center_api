ALTER TABLE `report`
DROP FOREIGN KEY `fk_report_task_id`;

ALTER TABLE `report`
DROP FOREIGN KEY `fk_report_account_id`;

DROP TABLE IF EXISTS `report`