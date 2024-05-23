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

func (d *Demo) Hi(uid, message string) error {
	// nosqlDb create
	if data, err := d.nosqlDb.LoadOrCreateDemo(uid); err != nil {
		return err
	} else {
		if err := data.Update(func() bool {
			data.SetMessage(message)
			return true
		}); err != nil {
			return err
		}
	}
	if err := db_sql.FirstOrCreate(d.gormDb, uid, message); err != nil {
		return err
	}
	d.redis.Set(context.Background(), "demo", message, time.Minute)
	//nats mq publish
	natsTopic := common.NatsHeader.CreateTopic("demo")
	if err := d.mq.Publish(
		natsTopic,
		miface.WithBytes([]byte(fmt.Sprintf("nats mq: %s", message))),
	); err != nil {
		return err
	}
	// local(channel) mq publish
	localTopic := common.LocalHeader.CreateTopic("demo")
	if err := d.mq.Publish(
		localTopic,
		miface.WithBytes([]byte(fmt.Sprintf("local mq: %s", message))),
	); err != nil {
		return err
	}

	return nil
}

func (d *Demo) Watch(ctx context.Context, topic string, callback func(message string) error) error {
	//nats mq subscribe
	natsTopic := common.NatsHeader.CreateTopic(topic)
	if _, err := d.mq.Subscribe(
		ctx,
		natsTopic,
		func(msg miface.Message, err error) common.ConsumptionCode {
			if err := callback(string(msg.Data())); err != nil {
				return common.ConsumeNackPersistentFailure
			}
			return common.ConsumeAck
		}); err != nil {
		return err
	}

	//local(channel) mq subscribe
	localTopic := common.LocalHeader.CreateTopic(topic)
	if _, err := d.mq.Subscribe(
		ctx,
		localTopic,
		func(msg miface.Message, err error) common.ConsumptionCode {
			if err := callback(string(msg.Data())); err != nil {
				return common.ConsumeNackPersistentFailure
			}
			return common.ConsumeAck
		}); err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}
