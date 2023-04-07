package usecase

import (
	"gaudy_code/config"
	"gaudy_code/domain/model"
	"gaudy_code/domain/repository"
)

type MissionInteractor struct {
	Repo              repository.MissionRepository
	userInteractor    UserInteractor
	monsterInteractor MonsterInteractor
}

func (interactor *MissionInteractor) GetAllMissions() (allMissions []model.Mission, err error) {
	missions, err := interactor.Repo.FindAllMissions()
	if err != nil {
		return
	}
	return missions, nil
}

func (interactor *MissionInteractor) FindCompletedMissions(userID int64) (completedMissions []model.Mission, err error) {
	allMissions, err := interactor.Repo.FindAllMissions()
	if err != nil {
		return
	}
	user, err := interactor.userInteractor.GetUser(userID)
	if err != nil {
		return
	}
	monsterID, err := interactor.monsterInteractor.GetFilteredMyMonsterIDsByLevel(userID, config.MONSTER_A_LEVEL)
	if err != nil {
		return completedMissions, err
	}
	missions, err := interactor.Repo.FindAllMyMissions(userID)
	if err != nil {
		return completedMissions, err
	}
	for _, mission := range missions {
		if mission.IsRewarded {
			user.RewardedMissionIDs[mission.ID] = true
		}
	}

	for _, mission := range allMissions {
		// 既に付与されているミッションはリセットされるものを除いてスキップする
		if user.RewardedMissionIDs[mission.ID] {
			continue
		}
		// ミッション解放条件を満たしていない場合はスキップする
		if !isDependencyOK(completedMissions, mission.DependencyIDs) {
			continue
		}
		switch mission.ID {
		case 1:
			targetMonstersMap, err := interactor.monsterInteractor.GetTargetMonsterIDs()
			if err != nil {
				return completedMissions, err
			}
			if targetMonstersMap[config.TARGET_MONSTER_ID] {
				completedMissions = append(completedMissions, mission)
				user.RewardedMissionIDs[mission.ID] = true
			}
		case 2:
			if user.Coin >= 2000 {
				completedMissions = append(completedMissions, mission)
				user.RewardedMissionIDs[mission.ID] = true
			}
		case 3:
			if monsterID[config.MONSTER_A_ID] {
				completedMissions = append(completedMissions, mission)
				user.RewardedMissionIDs[mission.ID] = true
			}
		case 4:
			if len(monsterID) >= 2 {
				completedMissions = append(completedMissions, mission)
				user.RewardedMissionIDs[mission.ID] = true
			}
		case 5:
			defeatedMonstersMap, err := interactor.monsterInteractor.GetCountAllDefeatedMonstersMap(userID)
			if err != nil {
				return completedMissions, err
			}
			totalDefeatedCount := 0
			for _, cnt := range defeatedMonstersMap {
				totalDefeatedCount += cnt
			}
			if totalDefeatedCount >= config.ANY_MONSTER_DEFEATED_COUNT {
				completedMissions = append(completedMissions, mission)
				user.RewardedMissionIDs[mission.ID] = false // 一定の周期でリセットされるものはrewardedに記録しない
			}
		case 6:
			if user.IsLoginRewaredNeed {
				user.IsLoginRewaredNeed = false
				completedMissions = append(completedMissions, mission)
				user.RewardedMissionIDs[mission.ID] = false // 一定の周期でリセットされるものはrewardedに記録しない
			}
		case 7:
			// for _, item := range user.Items {
			// 	if item == config.ITEM_A_ID {
			// 		completedMissions = append(completedMissions, mission)
			// 		user.RewardedMissionIDs[mission.ID] = true
			// 		break
			// 	}
			// }
		}
	}
	if err := interactor.userInteractor.UpdateUserData(userID, *user); err != nil {
		return completedMissions, err
	}
	return
}

func isDependencyOK(missions []model.Mission, dependencyIDs []int64) bool {
	if len(dependencyIDs) < 1 {
		return true
	}
	dependencyMap := make(map[int64]bool)
	for _, id := range dependencyIDs {
		dependencyMap[id] = true
	}
	for _, mission := range missions {
		if dependencyMap[mission.ID] {
			delete(dependencyMap, mission.ID)
		}
	}
	// 全てのdependencyが解決されればミッションが解放される
	return len(dependencyMap) == 0
}
