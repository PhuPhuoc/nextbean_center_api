CREATE TABLE IF NOT EXISTS `intern` (
  `id` varchar(255) PRIMARY KEY,
  `account_id` varchar(255) NOT NULL,
  `ojt_id` int,
  `avatar` varchar(255),
  `gender` varchar(10),
  `date_of_birth`datetime,
  `phone_number` varchar(12),
  `address` varchar(255),
  UNIQUE KEY `intern_account_id_unique` (`account_id`)
);

ALTER TABLE `intern`
ADD CONSTRAINT `fk_intern_ojt_id`
FOREIGN KEY (`ojt_id`) REFERENCES `ojt`(`id`);
