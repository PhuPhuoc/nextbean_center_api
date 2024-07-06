INSERT INTO account (id, user_name, email, password, role, created_at) 
VALUES 
('1', 'phuoc', 'blessforwork@gmail.com', 'e10adc3949ba59abbe56e057f20f883e', 'admin', '2024-06-01 00:00:00'),
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


INSERT INTO ojt (semester, university, start_at, end_at, status)
VALUES 
('spring24', 'FPT', '2024-01-01 00:00:00', '2024-04-15 00:00:00', 'completed'),
('summer24', 'FPT', '2024-05-01 00:00:00', '2024-09-15 00:00:00', 'in_progress'),
('fall24', 'FPT', '2024-10-01 00:00:00', '2024-12-15 00:00:00', 'not_started');


INSERT INTO intern (id, student_code, account_id, ojt_id,  avatar, gender, date_of_birth, phone_number, address)
VALUES 
('intern1', 'SE170001', '7', 2, 'avatar_url', 'male', '2003-01-15', '090000001', 'Khu A Duong B Quan C Phuong D, HCM'),
('intern2','SE170002', '8', 2, 'avatar_url', 'male', '2003-01-15', '090000002', 'Khu A Duong B Quan C Phuong D, HCM'),
('intern3','SE170003', '9', 2, 'avatar_url', 'male', '2003-02-15', '090000003', 'Khu A Duong B Quan C Phuong D, HCM'),
('intern4','SE170004', '10', 2, 'avatar_url', 'male', '2003-02-15', '090000004', 'Khu A Duong B Quan C Phuong D, HCM'),
('intern5','SE170005', '11', 3, 'avatar_url', 'female', '2003-03-15', '090000005', 'Khu A Duong B Quan C Phuong D, HCM'),
('intern6','SE170006', '12', 2, 'avatar_url', 'male', '2003-03-15', '090000006', 'Khu A Duong B Quan C Phuong D, HCM'),
('intern7','SE170007', '13', 2, 'avatar_url', 'male', '2003-04-15', '090000007', 'Khu A Duong B Quan C Phuong D, HCM'),
('intern8','SE170008', '14', 2, 'avatar_url', 'male', '2003-04-15', '090000008', 'Khu A Duong B Quan C Phuong D, HCM'),
('intern9','SE170009', '15', 2, 'avatar_url', 'male', '2003-05-15', '090000009', 'Khu A Duong B Quan C Phuong D, HCM'),
('intern10','SE170010', '16', 3, 'avatar_url', 'female', '2003-06-15', '090000010', 'Khu A Duong B Quan C Phuong D, HCM'),
('intern11','SE170011', '17', 1, 'avatar_url', 'female', '2003-07-15', '090000011', 'Khu A Duong B Quan C Phuong D, HCM');


INSERT INTO technical (technical_skill)
VALUES 
('ReactJS'),
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

INSERT INTO intern_skill(intern_id, technical_id, skill_level)
VALUES 
('intern1', 1, 'basic'), 
('intern1', 2, 'intermediate'), 
('intern1', 3, 'advanced'),
('intern2', 4, 'basic'), 
('intern2', 5, 'intermediate'), 
('intern2', 6, 'advanced'),
('intern3', 7, 'basic'), 
('intern3', 8, 'intermediate'), 
('intern3', 9, 'advanced'),
('intern4', 10, 'basic'), 
('intern4', 11, 'intermediate'), 
('intern4', 1, 'advanced'),
('intern5', 2, 'basic'), 
('intern5', 3, 'intermediate'), 
('intern5', 4, 'advanced'), 
('intern5', 5, 'basic'),
('intern6', 6, 'basic'), 
('intern6', 7, 'intermediate'), 
('intern6', 8, 'advanced'),
('intern7', 9, 'basic'), 
('intern7', 10, 'intermediate'), 
('intern7', 11, 'advanced'),
('intern8', 1, 'basic'), 
('intern8', 2, 'intermediate'), 
('intern8', 3, 'advanced'), 
('intern8', 4, 'basic'),
('intern9', 5, 'basic'), 
('intern9', 6, 'intermediate'), 
('intern9', 7, 'advanced'),
('intern10', 8, 'basic'), 
('intern10', 9, 'intermediate'), 
('intern10', 10, 'advanced'),
('intern11', 11, 'basic'), 
('intern11', 1, 'intermediate'), 
('intern11', 2, 'advanced'), 
('intern11', 3, 'basic');


insert into timetable(intern_id, office_time, est_start_time, est_end_time)
values
("intern1", "2024-07-10", "8:00", "17:00"),
("intern2", "2024-07-10", "8:00", "17:00"),
("intern3", "2024-07-10", "8:00", "17:00"),
("intern4", "2024-07-10", "8:00", "17:00"),
("intern5", "2024-07-10", "8:00", "17:00"),
("intern6", "2024-07-10", "8:00", "17:00"),
("intern1", "2024-07-12", "8:00", "17:00"),
("intern2", "2024-07-12", "8:00", "17:00"),
("intern3", "2024-07-12", "8:00", "17:00"),
("intern4", "2024-07-12", "8:00", "17:00"),
("intern5", "2024-07-12", "8:00", "17:00"),
("intern6", "2024-07-12", "8:00", "17:00"),
("intern1", "2024-07-13", "8:00", "17:00"),
("intern2", "2024-07-13", "8:00", "17:00"),
("intern3", "2024-07-13", "8:00", "17:00"),
("intern4", "2024-07-13", "8:00", "17:00"),
("intern5", "2024-07-13", "8:00", "17:00");


INSERT INTO `project` (`id`, `name`, `status`, `description`, `est_start_time`, `est_completion_time`)
VALUES 
('proj1', 'Project Alpha', 'not_started', 'This is the first project', '2024-07-15', '2024-012-15'),
('proj2', 'Project Beta', 'in_progress', 'This project is currently in progress', '2024-05-01', '2024-06-15'),
('proj3', 'Project Gamma', 'completed', 'This project has been completed', '2023-12-01', '2024-04-15');

insert into project_manager (project_id, account_id)
values
('proj1','3'),
('proj2','3'),
('proj2','4'),
('proj3','5'),
('proj3','6');

insert into project_intern (project_id, intern_id, join_at, status) 
values 
('proj1', 'intern1', '2024-06-15 08:00:00', 'in_progress'),
('proj1', 'intern2', '2024-06-15 08:00:00', 'in_progress'),
('proj1', 'intern3', '2024-06-15 08:00:00', 'in_progress'),
('proj2', 'intern3', '2024-06-15 08:00:00', 'in_progress'),
('proj2', 'intern4', '2024-06-15 08:00:00', 'in_progress'),
('proj2', 'intern5', '2024-06-15 08:00:00', 'in_progress'),
('proj3', 'intern6', '2024-06-15 08:00:00', 'in_progress'),
('proj3', 'intern7', '2024-06-15 08:00:00', 'in_progress'),
('proj3', 'intern3', '2024-06-15 08:00:00', 'in_progress'),
('proj3', 'intern2', '2024-06-15 08:00:00', 'in_progress');


insert into task(id, project_id, assigned_to, is_approved, name, description, estimated_effort, created_at)
values
('task1', 'proj1' , 'intern1', 1, 'task 1', "description task 1", "1 days", '2024-06-16 00:00:00'),
('task2', 'proj1' , 'intern2', 1, 'task 2', "description task 2", "2 days", '2024-06-16 00:00:00'),
('task3', 'proj1' , 'intern3', 1, 'task 3', "description task 3", "1 days", '2024-06-16 00:00:00'),
('task4', 'proj2' , 'intern3', 1, 'task 4', "description task 4", "2 days", '2024-06-16 00:00:00'),
('task5', 'proj2' , 'intern4', 1, 'task 5', "description task 5", "1 days", '2024-06-16 00:00:00'),
('task6', 'proj2' , 'intern5', 1, 'task 6', "description task 6", "4 days", '2024-06-16 00:00:00'),
('task7', 'proj3' , 'intern3', 1, 'task 7', "description task 7", "3 days", '2024-06-16 00:00:00'),
('task8', 'proj3' , 'intern6', 1, 'task 8', "description task 8", "1 days", '2024-06-16 00:00:00'),
('task9', 'proj3' , 'intern7', 1, 'task 9', "description task 9", "2 days", '2024-06-16 00:00:00');


