package usecase

import (
	"gaudy_code/domain/model"
	"gaudy_code/domain/repository"
)

type UserInteractor struct {
	Repo repository.UserRepository
}

func (interactor *UserInteractor) GetUser(userID int64) (user *model.User, err error) {
	user, err = interactor.Repo.FindUser(userID)
	if err != nil {
		return
	}
	return
}

func (interactor *UserInteractor) UpdateUserData(userID int64, user model.User) error {
	// 何らかの更新処理 または 別の場所で更新したデータを入れるだけにするか
	if err := interactor.Repo.UpdateUser(user); err != nil {
		return err
	}
	return nil
}
