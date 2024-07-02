CREATE TABLE `timetable` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `intern_id` varchar(255) NOT NULL,
  `office_time` DATE NOT NULL,
  `verified` ENUM('processing', 'denied', 'approved') DEFAULT 'processing',
  `est_start_time` time NOT NULL,
  `est_end_time` time NOT NULL,
  `act_clockin` time,
  `clockin_validated` ENUM('not-yet','authen-by-ip','admin-check','admin-approve') DEFAULT 'not-yet',
  `act_clockout` time,
  `clockout_validated` ENUM('not-yet','authen-by-ip','admin-check','admin-approve') DEFAULT 'not-yet',
  `status_attendance` ENUM('not-yet','absent','present') DEFAULT 'not-yet',
  `created_at` datetime  DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime
);

ALTER TABLE `timetable`
ADD CONSTRAINT `fk_timetable_intern_id`
FOREIGN KEY (`intern_id`) REFERENCES `intern`(`id`);
