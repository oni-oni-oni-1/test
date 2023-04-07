package model

import "time"

// CREATE TABLE monster (
//     id INT AUTO_INCREMENT PRIMARY KEY,
//     user_id INT,
//     name VARCHAR(255) NOT NULL,
//     level INT DEFAULT 1,
//     FOREIGN KEY (user_id) REFERENCES user(id)
// );

type MyMonster struct {
	ID        int64  `json:"id"`
	MonsterID int64  `json:"monster_id"`
	UserID    int64  `json:"user_id"`
	Name      string `json:"name"`
	Level     int    `json:"level"`
}

// CREATE TABLE user_monster_defeated (
// 	id INT AUTO_INCREMENT PRIMARY KEY,
// 	user_id INT NOT NULL,
// 	monster_id INT NOT NULL,
// 	count INT NOT NULL DEFAULT 1,
// 	defeated_at TIMESTAMP NOT NULL,
// 	last_reset_at TIMESTAMP NOT NULL
// );

type EnemyMonster struct {
	ID           int64     `json:"id"`
	MonsterID    int64     `json:"monster_id"`
	UserID       int64     `json:"user_id"`
	Name         string    `json:"name"`
	Count        int       `json:"count"`
	DefeatedAt   time.Time `json:"defeated_at"`
	LastResetAts time.Time `json:"last_reset_at"`
}
