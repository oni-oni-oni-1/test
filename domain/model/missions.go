package model

import "time"

// CREATE TABLE master_mission (
//     id INT AUTO_INCREMENT PRIMARY KEY,
//     name VARCHAR(255) NOT NULL,
//     reward_type ENUM('coins', 'item') NOT NULL,
//     reward_value INT NOT NULL,
//     reset_type ENUM('none', 'daily', 'weekly') DEFAULT 'none',
//     dependency_ids JSON
// );

// CREATE TABLE user_mission (
//     id INT AUTO_INCREMENT PRIMARY KEY,
//     user_id INT,
//     mission_id INT,
//     status ENUM('locked', 'in_progress', 'completed') DEFAULT 'locked',
//     progress INT DEFAULT 0,
//     last_updated TIMESTAMP,
//     FOREIGN KEY (user_id) REFERENCES user(id),
//     FOREIGN KEY (mission_id) REFERENCES master_mission(id)
// );

type Mission struct {
	ID            int64     `json:"id"`
	UserID        int64     `json:"user_id"`
	Name          string    `json:"name"`
	RewardCoin    int       `json:"reward_coin"`
	RewardItems   []int64   `json:"reward_item"`
	DependencyIDs []int64   `json:"dependencies"`
	IsRewarded    bool      `json:"is_rewarded"`
	LastUpdated   time.Time `json:"last_updated"`
}
