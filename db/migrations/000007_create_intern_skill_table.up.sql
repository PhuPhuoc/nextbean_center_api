CREATE TABLE IF NOT EXISTS `intern_skill` (
  `intern_id` varchar(255) NOT NULL,
  `technical_id` int NOT NULL,
  PRIMARY KEY (`intern_id`, `technical_id`)
);

ALTER TABLE `intern_skill`
ADD CONSTRAINT `fk_intern_skill_intern_id`
FOREIGN KEY (`intern_id`) REFERENCES `intern`(`id`);

ALTER TABLE `intern_skill`
ADD CONSTRAINT `fk_intern_skill_technical_id`
FOREIGN KEY (`technical_id`) REFERENCES `technical`(`id`);