ALTER TABLE `comment`
DROP FOREIGN KEY `fk_comment_from_account_id`;

ALTER TABLE `comment`
DROP FOREIGN KEY `fk_comment_to_report_id`;

DROP TABLE IF EXISTS `comment`