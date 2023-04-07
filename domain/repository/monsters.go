package repository

import (
	"gaudy_code/domain/model"
)

type MonsterRepository interface {
	FindMyMonsters(userID int64) ([]model.MyMonster, error)
	FindDefeatedMonsters(userID int64) ([]model.EnemyMonster, error)
	UpdateMyMonster(monster model.MyMonster) error
	UpdateEnemyMonster(monster model.EnemyMonster) error
}
