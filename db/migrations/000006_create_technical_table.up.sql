CREATE TABLE IF NOT EXISTS `technical` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `technical_skill` varchar(255) NOT NULL,
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` DATETIME
);
