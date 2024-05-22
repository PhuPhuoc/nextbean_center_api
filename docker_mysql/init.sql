CREATE TABLE IF NOT EXISTS account (
    `id` VARCHAR(255) PRIMARY KEY,
    `user_name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `role` ENUM('admin', 'pm', 'user'),
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` DATETIME
);
INSERT INTO account (id, user_name, email, password, role)
VALUES (
        '1',
        'phu_phuoc',
        'blessforwork@gmail.com',
        'pass1',
        'admin'
    );
INSERT INTO account (id, user_name, email, password, role)
VALUES (
        '2',
        'phuoc_pm',
        'blessing.cover@gmail.com',
        'pass2',
        'pm'
    );