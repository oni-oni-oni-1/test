package model

import "time"

// CREATE TABLE user (
//     id INT AUTO_INCREMENT PRIMARY KEY,
//     login_date DATE,
//     coins INT DEFAULT 0
//     is_login_rewarded ddd
// );

type User struct {
	ID                 int64          `json:"id"`
	LastLoginDate      time.Time      `json:"last_login_date"`
	Coin               int64          `json:"coin"`
	IsLoginRewaredNeed bool           `json:"is_login_rewarded"`
	RewardedMissionIDs map[int64]bool `json:"rewarded_mission_ids"`
}

func NewUser() *User {
	return &User{}
}
