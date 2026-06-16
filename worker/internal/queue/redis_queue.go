package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type VideoProcessingTask struct {
	VideoID  uint  `json:"video_id"`
	QueuedAt int64 `json:"queued_at"`
}

type RedisQueue struct {
	client      *redis.Client
	streamName  string
	consumerGroup string
	consumerName  string
}

func NewRedisQueue(client *redis.Client, streamName, consumerGroup, consumerName string) *RedisQueue {
	return &RedisQueue{
		client:       client,
		streamName:   streamName,
		consumerGroup: consumerGroup,
		consumerName:  consumerName,
	}
}

func (q *RedisQueue) Initialize(ctx context.Context) error {
	// Create consumer group if it doesn't exist
	err := q.client.XGroupCreateMkStream(ctx, q.streamName, q.consumerGroup, "0").Err()
	if err != nil && err.Error() != "BUSYGROUP Consumer Group name already exists" {
		return fmt.Errorf("failed to create consumer group: %w", err)
	}

	log.Printf("Consumer group '%s' initialized for stream '%s'", q.consumerGroup, q.streamName)
	return nil
}

func (q *RedisQueue) Consume(ctx context.Context, handler func(task VideoProcessingTask) error) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// Read from stream
			streams, err := q.client.XReadGroup(ctx, &redis.XReadGroupArgs{
				Group:    q.consumerGroup,
				Consumer: q.consumerName,
				Streams:  []string{q.streamName, ">"},
				Count:    1,
				Block:    5 * time.Second,
			}).Result()

			if err != nil {
				if err == redis.Nil {
					// No new messages, continue
					continue
				}
				log.Printf("Error reading from stream: %v", err)
				time.Sleep(1 * time.Second)
				continue
			}

			// Process messages
			for _, stream := range streams {
				for _, message := range stream.Messages {
					task, err := q.parseTask(message.Values["data"].(string))
					if err != nil {
						log.Printf("Failed to parse task: %v", err)
						q.ackMessage(ctx, message.ID)
						continue
					}

					log.Printf("Processing task for video ID: %d", task.VideoID)

					// Execute handler
					if err := handler(task); err != nil {
						log.Printf("Failed to process video %d: %v", task.VideoID, err)
						// In production, implement retry logic or dead letter queue
					}

					// Acknowledge message
					q.ackMessage(ctx, message.ID)
				}
			}
		}
	}
}

func (q *RedisQueue) parseTask(data string) (VideoProcessingTask, error) {
	var task VideoProcessingTask
	if err := json.Unmarshal([]byte(data), &task); err != nil {
		return task, err
	}
	return task, nil
}

func (q *RedisQueue) ackMessage(ctx context.Context, messageID string) {
	if err := q.client.XAck(ctx, q.streamName, q.consumerGroup, messageID).Err(); err != nil {
		log.Printf("Failed to acknowledge message %s: %v", messageID, err)
	}
}
