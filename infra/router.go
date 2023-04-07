package infra

import (
	"gaudy_code/config"
	"gaudy_code/controller"
	"gaudy_code/domain/model"
	"gaudy_code/infra/gateway"
	"net/http"

	"github.com/gin-gonic/gin"
)

func bindJSONMiddleware(req *model.ActionRequest) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.Next()
	}
}

func Router() {
	r := gin.Default()
	// controller := controller.NewController(&articleRepo, &articleClient)
	var userRepo gateway.UserRepository
	userController := controller.NewUserController(&userRepo)
	var monsterRepo gateway.MonsterRepository
	monsterController := controller.NewMonsterController(&monsterRepo)

	r.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	r.GET("/test", func(c *gin.Context) {
		c.String(200, userController.GetUser(c, c.Query("id")))
	})

	var actionReq model.ActionRequest
	eventGroup := r.Group("/event")
	eventGroup.Use(bindJSONMiddleware(&actionReq))
	{
		eventGroup.POST("/attack", func(c *gin.Context) {
			userController.UpdateLastLoginAndCheckRewards(c, actionReq.UserID)
			c.String(http.StatusOK, monsterController.AttackToEnemy(c, actionReq.UserID, actionReq.OpponentMonsterID))
		})

		eventGroup.POST("/level_up", func(c *gin.Context) {
			userController.UpdateLastLoginAndCheckRewards(c, actionReq.UserID)
			c.String(http.StatusOK,
				monsterController.UpdateMonsterLevel(c, actionReq.UserID, actionReq.MyMonsterID, config.LEVEL_UP_INCREMENT))
		})
	}

	r.POST("/receive_event", func(c *gin.Context) {
		// TODO 非同期
	})
	r.GET("/mission/status", func(c *gin.Context) {
		// TODO 同期
	})
	r.POST("/mission/reward", func(c *gin.Context) {
		// TODO 同期
	})
	r.Run(":8080")
}
