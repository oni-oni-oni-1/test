package controller

import (
	"encoding/json"
	"fmt"
	"gaudy_code/config"
	"gaudy_code/domain/repository"
	"gaudy_code/usecase"
	"net/http"
	"strconv"
	"time"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(repo repository.UserRepository) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			Repo: repo,
		},
	}
}

func (controller *UserController) GetUser(c Context, useIDStr string) string {
	userID, err := strconv.ParseInt(useIDStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
	user, err := controller.Interactor.GetUser(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
	return string(jsonBytes)
}

func (controller *UserController) UpdateLastLoginAndCheckRewards(c Context, useIDStr string) string {
	userID, err := strconv.ParseInt(useIDStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
	user, err := controller.Interactor.GetUser(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
	fmt.Println("user.LastLoginDate.Hour()", user.LastLoginDate.Hour())
	// reset後、初のログイン
	if user.LastLoginDate.Hour() < config.DAILY_RESET_TIME {
		// TODO rewardを申請する
		user.IsLoginRewaredNeed = true
	}
	user.LastLoginDate = time.Now()

	if err := controller.Interactor.Repo.UpdateUser(*user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
	return "ok"
}
