ALTER TABLE `comment`
DROP FOREIGN KEY `fk_comment_to_report_id`;

DROP TABLE IF EXISTS `comment`