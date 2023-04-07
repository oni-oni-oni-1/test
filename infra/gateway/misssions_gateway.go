package gateway

import (
	"database/sql"
	"encoding/json"
	"gaudy_code/config"
	"gaudy_code/domain/model"
)

type MissionRepository struct{}

func (repo *MissionRepository) FindAllMissions() ([]model.Mission, error) {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/game_db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, reward_type, reward_value, reset_type, dependency_ids FROM master_mission")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var missions []model.Mission
	for rows.Next() {
		var mission model.Mission
		var rewardType string
		var rewardValue int
		var resetType string
		var dependencyIDsJSON string

		err := rows.Scan(&mission.ID, &mission.Name, &rewardType, &rewardValue, &resetType, &dependencyIDsJSON)
		if err != nil {
			return nil, err
		}

		if rewardType == config.REWARD_TYPE_COIN {
			mission.RewardCoin = rewardValue
		} else {
			mission.RewardItems = append(mission.RewardItems, int64(rewardValue))
		}
		if err := json.Unmarshal([]byte(dependencyIDsJSON), &mission.DependencyIDs); err != nil {
			return nil, err
		}
		missions = append(missions, mission)
	}
	return missions, nil
}

func (repo *MissionRepository) FindAllMyMissions(userID int64) ([]model.Mission, error) {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/game_db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM user_mission WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var missions []model.Mission
	for rows.Next() {
		var mission model.Mission
		var rewardItems, dependencyIDs []int64
		var rewardItemsJSON, dependencyIDsJSON []byte

		err := rows.Scan(&mission.ID, &mission.UserID, &mission.Name, &mission.RewardCoin, &rewardItemsJSON, &dependencyIDsJSON, &mission.IsRewarded, &mission.LastUpdated)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(rewardItemsJSON, &rewardItems)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(dependencyIDsJSON, &dependencyIDs)
		if err != nil {
			return nil, err
		}
		mission.RewardItems = rewardItems
		mission.DependencyIDs = dependencyIDs
		missions = append(missions, mission)
	}
	return missions, nil
}

func (repo *MissionRepository) FindCompletedMissions(userID int64) ([]model.Mission, error) {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/game_db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM user_mission WHERE user_id = ? AND status = 'completed'", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var missions []model.Mission
	for rows.Next() {
		var mission model.Mission
		var rewardItems, dependencyIDs []int64
		var rewardItemsJSON, dependencyIDsJSON []byte
		err := rows.Scan(&mission.ID, &mission.UserID, &mission.Name, &mission.RewardCoin, &rewardItemsJSON, &dependencyIDsJSON, &mission.IsRewarded, &mission.LastUpdated)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(rewardItemsJSON, &rewardItems)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(dependencyIDsJSON, &dependencyIDs)
		if err != nil {
			return nil, err
		}
		mission.RewardItems = rewardItems
		mission.DependencyIDs = dependencyIDs
		missions = append(missions, mission)
	}
	return missions, nil
}
