CREATE TABLE IF NOT EXISTS `intern` (
  `id` varchar(255) PRIMARY KEY,
  `account_id` varchar(255) NOT NULL,
  `phone_number` varchar(255),
  `ojt_id` int,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `deleted_ad` datetime,
  UNIQUE KEY `intern_account_id_unique` (`account_id`)
);

ALTER TABLE `intern`
ADD CONSTRAINT `fk_intern_ojt_id`
FOREIGN KEY (`ojt_id`) REFERENCES `ojt`(`id`);
