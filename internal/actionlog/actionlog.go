package actionlog

import (
	"context"
	"encoding/json"
	"log"
	"my_project/internal/database"
	"my_project/internal/redis"
	"sync"
	"time"
)

type ActionLog struct {
	UserID   int
	Username string
	Action   string
	TargetID int
	Time     time.Time
}


var (
	once    sync.Once
)

func Log(entry ActionLog) {
	data, _ := json.Marshal(entry)
	redis.GetRedisClient().RPush(context.Background(), "action_logs", data)
}

func StartWorker() {
	once.Do(func() {
		db := database.New()
		go func() {
			for {
				res, err := redis.GetRedisClient().BLPop(context.Background(), 0, "action_logs").Result()
				if err != nil || len(res) < 2 {
					continue
				}
				var logEntry ActionLog
				if err := json.Unmarshal([]byte(res[1]), &logEntry); err != nil {
					continue
				}
				_, err = db.GetDB().Exec(
					`INSERT INTO user_action_logs (user_id,username, action, target_id, created_at) VALUES (?, ?, ?, ?, NOW())`,
					logEntry.UserID,
					logEntry.Username,
					logEntry.Action,
					logEntry.TargetID,
				)
				if err != nil {
					log.Println("Error logging to database:", err)
				}
			}
		}()
	})
}
