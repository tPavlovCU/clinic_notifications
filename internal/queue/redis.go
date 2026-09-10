package queue

import (
	"clinic-notifications/internal/domain"
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisQueue struct {
	redis *redis.Client
}

func InitRedis(addr string) (*RedisQueue, error) {
	options := redis.Options{
		Addr:     addr,
		PoolSize: 10,
	}

	client := redis.NewClient(&options)
	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}
	result := RedisQueue{client}
	return &result, nil
}

func (r *RedisQueue) PushTask(task domain.NotificationTask) error {
	jsonTask, err := json.Marshal(task)
	if err != nil {
		return err
	}

	err = r.redis.LPush(ctx, "notification_tasks", string(jsonTask)).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisQueue) PopTask() (*domain.NotificationTask, error) {
	result, err := r.redis.BRPop(ctx, 0, "notification_tasks").Result()
	if err != nil {
		return nil, err
	}
	jsonTask := result[1]
	var task domain.NotificationTask
	err = json.Unmarshal([]byte(jsonTask), &task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *RedisQueue) Close() error {
	err := r.redis.Close()
	if err != nil {
		return err
	}
	return nil
}
