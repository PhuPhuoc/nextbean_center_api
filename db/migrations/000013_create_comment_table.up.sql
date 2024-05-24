CREATE TABLE IF NOT EXISTS `comment` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `from_user_name` varchar(255) NOT NULL,
  `to_report_id` int NOT NULL,
  `content` text NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `deleted_ad` datetime
);

ALTER TABLE `comment`
ADD CONSTRAINT `fk_comment_to_report_id`
FOREIGN KEY (`to_report_id`) REFERENCES `report`(`id`);