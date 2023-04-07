package repository

import (
	"gaudy_code/domain/model"
)

type MissionRepository interface {
	FindAllMissions() ([]model.Mission, error)
	FindAllMyMissions(userID int64) ([]model.Mission, error)
	FindCompletedMissions(userID int64) ([]model.Mission, error)
}
