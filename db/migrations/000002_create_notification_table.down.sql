ALTER TABLE `notification`
DROP FOREIGN KEY `fk_notification_account_id`;

DROP TABLE IF EXISTS `notification`;