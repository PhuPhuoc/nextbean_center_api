INSERT INTO account (id, user_name, email, password, role, created_at) 
VALUES ('1', 'phuoc', 'phuoc@gmail.com', 'e10adc3949ba59abbe56e057f20f883e', 'admin', '2024-06-01 00:00:00'),
('2', 'thu', 'thu@gmail.com', 'e10adc3949ba59abbe56e057f20f883e', 'admin', '2024-05-01 00:00:00'),
('3', 'tam', 'tam@gmail.com', 'e10adc3949ba59abbe56e057f20f883e', 'pm', '2024-04-01 00:00:00'),
('4', 'danh', 'danh@gmail.com', 'e10adc3949ba59abbe56e057f20f883e', 'pm', '2024-05-01 00:00:00'),
('5', 'dat', 'dat@gmail.com', 'e10adc3949ba59abbe56e057f20f883e', 'pm', '2024-04-01 00:00:00'),
('6', 'nghi', 'nghi@gmail.com', 'e10adc3949ba59abbe56e057f20f883e', 'pm', '2024-03-01 00:00:00'),
('7', 'tien', 'tien@gmail.com', 'e10adc3949ba59abbe56e057f20f883e', 'user', '2024-03-01 00:00:00'),
('8', 'duy', 'duy@gmail.com', 'e10adc3949ba59abbe56e057f20f883e', 'user', '2024-04-02 00:00:00'),
('9', 'quan', 'quan@gmail.com', 'e10adc3949ba59abbe56e057f20f883e', 'user', '2024-05-03 00:00:00'),
('10', 'nhan', 'nhan@gmail.com', 'e10adc3949ba59abbe56e057f20f883e', 'user', '2024-09-03 00:00:00'),
('11', 'minh', 'minh@gmail.com', 'e10adc3949ba59abbe56e057f20f883e', 'user', '2024-02-03 00:00:00'),
('12', 'tuan', 'tuan@gmail.com', 'e10adc3949ba59abbe56e057f20f883e', 'user', '2024-07-03 00:00:00'),
('13', 'hoang', 'hoangn@gmail.com', 'e10adc3949ba59abbe56e057f20f883e', 'user', '2024-03-03 00:00:00'),
('14', 'bao', 'bao@gmail.com', 'e10adc3949ba59abbe56e057f20f883e', 'user', '2024-02-03 00:00:00'),
('15', 'ninh', 'ninh@gmail.com', 'e10adc3949ba59abbe56e057f20f883e', 'user', '2024-09-03 00:00:00'),
('16', 'van', 'van@gmail.com', 'e10adc3949ba59abbe56e057f20f883e', 'user', '2024-03-03 00:00:00'),
('17', 'mai', 'mai@gmail.com', 'e10adc3949ba59abbe56e057f20f883e', 'user', '2024-07-03 00:00:00'),
('18', 'thanh', 'thanh@gmail.com', 'e10adc3949ba59abbe56e057f20f883e', 'manager', '2024-04-01 00:00:00'),
('19', 'yen', 'yen@gmail.com', 'e10adc3949ba59abbe56e057f20f883e', 'manager', '2024-06-01 00:00:00'),
('20', 'nhu', 'nhu@gmail.com', 'e10adc3949ba59abbe56e057f20f883e', 'manager', '2024-07-01 00:00:00');


INSERT INTO ojt (semester, university, start_at, end_at)
VALUES ('spring24', 'FPT', '2024-01-01 00:00:00', '2024-04-15 00:00:00'),
('summer24', 'FPT', '2024-05-01 00:00:00', '2024-09-15 00:00:00'),
('fall24', 'FPT', '2024-10-01 00:00:00', '2024-12-15 00:00:00');


INSERT INTO intern (id, student_code, account_id, ojt_id,  avatar, gender, date_of_birth, phone_number, address)
VALUES ('intern1', 'SE170001', '7', 1, 'avatar_url', 'male', '2003-01-15', '090000001', 'Khu A Duong B Quan C Phuong D, HCM'),
('intern2','SE170002', '8', 1, 'avatar_url', 'male', '2003-01-15', '090000002', 'Khu A Duong B Quan C Phuong D, HCM'),
('intern3','SE170003', '9', 1, 'avatar_url', 'male', '2003-02-15', '090000003', 'Khu A Duong B Quan C Phuong D, HCM'),
('intern4','SE170004', '10', 1, 'avatar_url', 'male', '2003-02-15', '090000004', 'Khu A Duong B Quan C Phuong D, HCM'),
('intern5','SE170005', '11', 1, 'avatar_url', 'female', '2003-03-15', '090000005', 'Khu A Duong B Quan C Phuong D, HCM'),
('intern6','SE170006', '12', 1, 'avatar_url', 'male', '2003-03-15', '090000006', 'Khu A Duong B Quan C Phuong D, HCM'),
('intern7','SE170007', '13', 1, 'avatar_url', 'male', '2003-04-15', '090000007', 'Khu A Duong B Quan C Phuong D, HCM'),
('intern8','SE170008', '14', 1, 'avatar_url', 'male', '2003-04-15', '090000008', 'Khu A Duong B Quan C Phuong D, HCM'),
('intern9','SE170009', '15', 1, 'avatar_url', 'male', '2003-05-15', '090000009', 'Khu A Duong B Quan C Phuong D, HCM'),
('intern10','SE170010', '16', 1, 'avatar_url', 'female', '2003-06-15', '090000010', 'Khu A Duong B Quan C Phuong D, HCM'),
('intern11','SE170011', '17', 1, 'avatar_url', 'female', '2003-07-15', '090000011', 'Khu A Duong B Quan C Phuong D, HCM');


INSERT INTO technical (technical_skill)
VALUES ('ReactJS'),
('Angular'),
('VueJS'),
('Flutter'),
('React Native'),
('Asp dotNet'),
('Winform dotNet'),
('Golang'),
('Python'),
('NestJS'),
('NextJS');

INSERT INTO `project` (`id`, `name`, `status`, `description`, `start_date`, `duration`)
VALUES ('proj1', 'Project Alpha', 'not_start', 'This is the first project', '2024-06-15 08:00:00', '3 months'),
('proj2', 'Project Beta', 'doing', 'This project is currently in progress', '2024-05-01 09:00:00', '6 months'),
('proj3', 'Project Gamma', 'done', 'This project has been completed', '2023-12-01 10:00:00', '1 year');