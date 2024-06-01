CREATE TABLE `timetable` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `intern_id` varchar(255) NOT NULL,
  `office_time` varchar(255) NOT NULL,
  `est_start_at` datetime NOT NULL,
  `est_end_at` datetime NOT NULL,
  `act_start_at` datetime,
  `act_end_at` datetime,
  `status` ENUM('processing', 'notapproved', 'approved') DEFAULT 'processing',
  `created_at` datetime  DEFAULT CURRENT_TIMESTAMP,
  `deleted_ad` datetime
);

ALTER TABLE `timetable`
ADD CONSTRAINT `fk_timetable_intern_id`
FOREIGN KEY (`intern_id`) REFERENCES `intern`(`id`);