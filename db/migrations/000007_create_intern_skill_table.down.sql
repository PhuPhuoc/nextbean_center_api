ALTER TABLE `intern_skill`
DROP FOREIGN KEY `fk_intern_skill_intern_id`;

ALTER TABLE `intern_skill`
DROP FOREIGN KEY `fk_intern_skill_technical_id`;


DROP TABLE IF EXISTS `intern_skill`;