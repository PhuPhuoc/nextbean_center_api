CREATE TABLE IF NOT EXISTS `notification` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `account_id` varchar(255) NOT NULL,
  `content` text,
  `created_at` datetime  DEFAULT CURRENT_TIMESTAMP,
  `viewed` BOOLEAN DEFAULT false
);

ALTER TABLE `notification`
ADD CONSTRAINT `fk_notification_account_id`
FOREIGN KEY (`account_id`) REFERENCES `account`(`id`);