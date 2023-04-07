package usecase

import (
	"gaudy_code/domain/model"
	"gaudy_code/domain/repository"
)

type MonsterInteractor struct {
	Repo repository.MonsterRepository
}

func (interactor *MonsterInteractor) GetAllMonsters(userID int64) (myMonsters []model.MyMonster, err error) {
	monsters, err := interactor.Repo.FindMyMonsters(userID)
	if err != nil {
		return myMonsters, err
	}
	return monsters, nil
}

func (interactor *MonsterInteractor) GetFilteredMyMonsterIDsByLevel(userID int64, level int) (monsterIDs map[int64]bool, err error) {
	monsters, err := interactor.Repo.FindMyMonsters(userID)
	if err != nil {
		return monsterIDs, err
	}
	monsterIDs = make(map[int64]bool, len(monsters))
	for _, monster := range monsters {
		if monster.Level >= level {
			monsterIDs[monster.ID] = true
		}
	}
	return monsterIDs, nil
}

func (interactor *MonsterInteractor) GetCountAllDefeatedMonstersMap(userID int64) (countMap map[int64]int, err error) {
	monsters, err := interactor.Repo.FindDefeatedMonsters(userID)
	if err != nil {
		return countMap, err
	}
	countMap = make(map[int64]int, len(monsters))
	for _, monster := range monsters {
		countMap[monster.ID]++
	}
	return
}

// 特定のモンスターをDBから選択する場合、ここを変更する
func (interactor *MonsterInteractor) GetTargetMonsterIDs() (monsterIDs map[int64]bool, err error) {
	// 今回は "特定のモンスター" の IDを-1として設定する
	return map[int64]bool{-1: true}, nil
}

func (interactor *MonsterInteractor) GetCountDefeatedTargetMonstersMap(userID int64) (countMap map[int64]int, err error) {
	monsters, err := interactor.Repo.FindDefeatedMonsters(userID)
	if err != nil {
		return countMap, err
	}
	targetMonstersMap, err := interactor.GetTargetMonsterIDs()
	if err != nil {
		return countMap, err
	}
	countMap = make(map[int64]int, len(targetMonstersMap))
	for _, monster := range monsters {
		if targetMonstersMap[monster.ID] {
			countMap[monster.ID]++
		}
	}
	return
}

func (interactor *MonsterInteractor) UpdateMyMonsterLevel(userID int64, monsterID int64, level int) (err error) {
	monsters, err := interactor.Repo.FindMyMonsters(userID)
	if err != nil {
		return err
	}
	for _, monster := range monsters {
		if monster.ID == monsterID {
			if err := interactor.Repo.UpdateMyMonster(model.MyMonster{
				ID:        monster.ID,
				MonsterID: monster.MonsterID,
				UserID:    userID,
				Name:      monster.Name,
				Level:     level,
			}); err != nil {
				return err
			}
			return
		}
	}
	return
}
