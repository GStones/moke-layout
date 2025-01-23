package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/gstones/moke-kit/mq/common"
	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-layout/internal/services/demo/db_nosql"
	"github.com/gstones/moke-layout/internal/services/demo/db_sql"
)

type Demo struct {
	logger  *zap.Logger
	nosqlDb db_nosql.Database
	mq      miface.MessageQueue
	gormDb  *gorm.DB
	redis   *redis.Client
}

func NewDemo(
	logger *zap.Logger,
	database db_nosql.Database,
	mq miface.MessageQueue,
	gormDb *gorm.DB,
	redis *redis.Client,
) *Demo {
	return &Demo{
		logger:  logger,
		nosqlDb: database,
		mq:      mq,
		gormDb:  gormDb,
		redis:   redis,
	}
}

func (d *Demo) Hi(ctx context.Context, uid, topic, message string) error {
	// Handle nosqlDb operation
	data, err := d.nosqlDb.LoadOrCreateDemo(ctx, uid)
	if err != nil {
		return fmt.Errorf("failed to load/create demo: %w", err)
	}

	if err := data.Update(func() bool {
		data.SetMessage(message)
		return true
	}); err != nil {
		return fmt.Errorf("failed to update demo: %w", err)
	}

	// Handle SQL operation
	if err := db_sql.FirstOrCreate(d.gormDb, uid, message); err != nil {
		return fmt.Errorf("failed to create SQL record: %w", err)
	}

	// Handle Redis operation
	if err := d.redis.Set(ctx, topic, message, time.Minute).Err(); err != nil {
		return fmt.Errorf("failed to set redis key: %w", err)
	}

	// Publish to message queues
	for _, header := range []common.Header{common.NatsHeader, common.LocalHeader} {
		mqTopic := header.CreateTopic(topic)
		prefix := "nats"
		if header == common.LocalHeader {
			prefix = "local"
		}
		if err := d.mq.Publish(
			mqTopic,
			miface.WithBytes([]byte(fmt.Sprintf("%s mq: %s", prefix, message))),
		); err != nil {
			return fmt.Errorf("failed to publish to %s queue: %w", prefix, err)
		}
	}

	return nil
}

func (d *Demo) Watch(ctx context.Context, topic string, callback func(message string) error) error {
	// Subscribe to both queues
	for _, header := range []common.Header{common.NatsHeader, common.LocalHeader} {
		mqTopic := header.CreateTopic(topic)
		if _, err := d.mq.Subscribe(
			ctx,
			mqTopic,
			func(msg miface.Message, err error) common.ConsumptionCode {
				if err := callback(string(msg.Data())); err != nil {
					return common.ConsumeNackPersistentFailure
				}
				return common.ConsumeAck
			}); err != nil {
			return fmt.Errorf("failed to subscribe to topic %s: %w", mqTopic, err)
		}
	}

	<-ctx.Done()
	return nil
}
