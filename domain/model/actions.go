package model

import (
	"encoding/json"
	"fmt"
)

type ActionType string

const (
	Attack        ActionType = "ATTACK"
	LevelUp       ActionType = "LEVEL_UP"
	LogIN         ActionType = "LOG_IN"
	MissionStatus ActionType = "MISSION_STATUS"
)

type ActionRequest struct {
	UserID            string     `json:"userId"`
	ActionType        ActionType `json:"actionType"`
	OpponentMonsterID string     `json:"opponentMonsterId,omitempty"`
	MyMonsterID       string     `json:"myMonsterId,omitempty"`
	Amount            int        `json:"amount,omitempty"`
	CreatedAt         int64      `json:"createdAt"`
}

func NewActionRequest() *ActionRequest {
	return &ActionRequest{}
}

func (request *ActionRequest) Unmarshall(actionJson string) {
	err := json.Unmarshal([]byte(actionJson), request)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
