package usermodel

/*
CREATE TABLE IF NOT EXISTS `users` (
    `id` VARCHAR(100) PRIMARY KEY,
    `fb_id` VARCHAR(100),
    `email` VARCHAR(30) NOT NULL,
    `password` VARCHAR(30) NOT NULL,
    `name` VARCHAR(100) NOT NULL,
    `avatar` VARCHAR(255),
    `role` ENUM('admin', 'user', 'member'),
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` DATETIME,
    UNIQUE(`email`)
);
*/

type User struct {
	ID        string  `db:"id" json:"id"`
	FBID      *string `db:"fb_id" json:"fb_id"`
	Email     string  `db:"email" json:"email"`
	Password  string  `db:"password" json:"-"`
	Name      string  `db:"name" json:"user_name"`
	Avatar    string  `db:"avatar" json:"avatar"`
	Role      string  `db:"role" json:"role"`
	CreatedAt string  `db:"created_at" json:"-"`
}
