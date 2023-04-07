package controller

import (
	"gaudy_code/domain/model"
	"gaudy_code/domain/repository"
	"gaudy_code/usecase"
	"net/http"
	"strconv"
	"time"
)

type MonsterController struct {
	Interactor usecase.MonsterInteractor
}

func NewMonsterController(repo repository.MonsterRepository) *MonsterController {
	return &MonsterController{
		Interactor: usecase.MonsterInteractor{
			Repo: repo,
		},
	}
}

func (controller *MonsterController) UpdateMonsterLevel(c Context, userIDStr string,
	monsterIDStr string, level int) string {
	parseID := func(str string) (int64, error) {
		id, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return 0, err
		}
		return id, nil
	}
	userID, _ := parseID(userIDStr)
	monsterID, _ := parseID(monsterIDStr)

	if err := controller.Interactor.UpdateMyMonsterLevel(userID, monsterID, level); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
	return "ok"
}

func (controller *MonsterController) AttackToEnemy(c Context, userIDStr string, targetEnemyMonsterIDStr string) string {
	parseID := func(str string) int64 {
		id, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return 0
		}
		return id
	}
	userID := parseID(userIDStr)
	defeatedMonsters, err := controller.Interactor.Repo.FindDefeatedMonsters(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}

	targetEnemyMonsterID := parseID(targetEnemyMonsterIDStr)
	for _, enemy := range defeatedMonsters {
		if enemy.ID == targetEnemyMonsterID {

			daysUntilMonday := int(time.Monday - time.Now().Weekday())
			if daysUntilMonday > 0 {
				daysUntilMonday -= 7
			}
			monday := time.Now().AddDate(0, 0, daysUntilMonday)
			mondayMidnight := time.Date(monday.Year(), monday.Month(), monday.Day(), 0, 0, 0, 0, monday.Location())
			// 月曜日の午前0時を過ぎているか確認する、過ぎていた場合はcountを1にする。過ぎていない場合はincrementする
			count := enemy.Count + 1
			lastResetAts := enemy.LastResetAts
			if !time.Now().After(mondayMidnight) {
				count = 1
				lastResetAts = time.Now()
			}

			controller.Interactor.Repo.UpdateEnemyMonster(model.EnemyMonster{
				ID:           enemy.ID,
				MonsterID:    enemy.MonsterID,
				UserID:       userID,
				Name:         enemy.Name,
				Count:        count,
				DefeatedAt:   time.Now(), //最後のDefeatedAtが一週間以内だった確認する
				LastResetAts: lastResetAts,
			})
		}
	}
	return "ok"
}
