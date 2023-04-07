package main

import (
	"database/sql"
	"fmt"
	"gaudy_code/config"
	"gaudy_code/domain/model"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	user     = "root"
	password = "my-secret-pw"
	host     = "localhost"
	port     = "3306"
	dbname   = "game_db"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", user, password, host, port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE DATABASE IF NOT EXISTS game_db;`)
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS user (
			id INT AUTO_INCREMENT PRIMARY KEY,
			login_date TIMESTAMP,
			coins INT DEFAULT 0,
			is_login_rewarded BOOLEAN DEFAULT FALSE
		);`)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS monster (
			id INT AUTO_INCREMENT PRIMARY KEY,
			monster_id INT,
			user_id INT,
			name VARCHAR(255) NOT NULL,
			level INT DEFAULT 1,
			FOREIGN KEY (user_id) REFERENCES user(id)
		);`)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS user_monster_defeated (
			id INT AUTO_INCREMENT PRIMARY KEY,
			user_id INT NOT NULL,
			monster_id INT NOT NULL,
			count INT NOT NULL DEFAULT 1,
			defeated_at TIMESTAMP NOT NULL,
			last_reset_at TIMESTAMP NOT NULL
		);`)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}

	_, err = db.Exec(`
		INSERT INTO user (login_date, coins) VALUES ('2023-04-07 10:00:00', 1800);`)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}

	_, err = db.Exec(`
		INSERT INTO monster (user_id, monster_id, name, level) VALUES (1, 1, 'モンスターA', 4), (1, 2, 'モンスターB', 7);`)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}

	_, err = db.Exec(`
		INSERT INTO user_monster_defeated (user_id, monster_id, count, defeated_at, last_reset_at) VALUES (1, 1, 8, '2023-04-07 09:45:00', '2023-04-03 00:00:00');`)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}

	missions := []model.Mission{
		{ID: 1, Name: "特定のモンスター(ID:-1)を倒す", RewardCoin: 100, RewardItems: []int64{}, DependencyIDs: []int64{}},
		{ID: 2, Name: "2000コイン貯まる", RewardCoin: 0, RewardItems: []int64{config.ITEM_A_ID}, DependencyIDs: []int64{}},
		{ID: 3, Name: "モンスターAのレベルが5になる", RewardCoin: 100, RewardItems: []int64{}, DependencyIDs: []int64{}},
		{ID: 4, Name: "レベル５以上のモンスターが２体", RewardCoin: 100, RewardItems: []int64{}, DependencyIDs: []int64{}},
		{ID: 5, Name: "任意のモンスターを10回倒す", RewardCoin: 100, RewardItems: []int64{}, DependencyIDs: []int64{}},
		{ID: 6, Name: "ログイン（毎日午前4時にリセットされる）", RewardCoin: 100, RewardItems: []int64{}, DependencyIDs: []int64{}},
		{ID: 7, Name: "アイテムAを所有する", RewardCoin: 100, RewardItems: []int64{}, DependencyIDs: []int64{4}},
	}

	for _, mission := range missions {
		_, err = db.Exec(`
			INSERT INTO user_mission (user_id, mission_id, status, progress, last_updated)
			VALUES (1, ?, 'locked', 0, ?)`,
			mission.ID, time.Now(),
		)
		if err != nil {
			log.Fatalf("Failed to execute query: %v", err)
		}
	}

	fmt.Println("Database and tables created, and records inserted successfully.")
}
