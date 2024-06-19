CREATE TABLE `timetable` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `intern_id` varchar(255) NOT NULL,
  `office_time` varchar(255) NOT NULL,
  `est_start` time NOT NULL,
  `est_end` time NOT NULL,
  `act_start` time,
  `act_end` time,
  `status` ENUM('processing', 'denied', 'approved') DEFAULT 'processing',
  `created_at` datetime  DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime
);

ALTER TABLE `timetable`
ADD CONSTRAINT `fk_timetable_intern_id`
FOREIGN KEY (`intern_id`) REFERENCES `intern`(`id`);