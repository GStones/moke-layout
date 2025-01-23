package demo

import (
	"context"
	"fmt"

	"github.com/gstones/zinx/ziface"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"

	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/nosql/diface"
	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/siface"

	pb "github.com/gstones/moke-layout/api/gen/demo/api"
	"github.com/gstones/moke-layout/internal/services/demo/db_nosql"
	"github.com/gstones/moke-layout/internal/services/demo/domain"
	"github.com/gstones/moke-layout/pkg/dfx"
)

type Service struct {
	logger      *zap.Logger
	demoHandler *domain.Demo
}

// ---------------- grpc ----------------

func (s *Service) Watch(request *pb.WatchRequest, server pb.DemoService_WatchServer) error {
	topic := request.GetTopic()
	s.logger.Info("Watch request received", zap.String("topic", topic))

	return s.demoHandler.Watch(server.Context(), topic, func(message string) error {
		return server.Send(&pb.WatchResponse{Message: message})
	})
}

func (s *Service) Hi(ctx context.Context, request *pb.HiRequest) (*pb.HiResponse, error) {
	s.logger.Info("Hi request received",
		zap.String("uid", request.GetUid()),
		zap.String("topic", request.GetTopic()),
		zap.String("message", request.GetMessage()))

	if err := s.demoHandler.Hi(ctx, request.GetUid(), request.GetTopic(), request.GetMessage()); err != nil {
		s.logger.Error("Hi handler failed", zap.Error(err))
		return nil, err
	}

	return &pb.HiResponse{Message: "response: " + request.GetMessage()}, nil
}

func (s *Service) RegisterWithGrpcServer(server siface.IGrpcServer) error {
	pb.RegisterDemoServiceServer(server.GrpcServer(), s)
	return nil
}

// ---------------- gateway ----------------

func (s *Service) RegisterWithGatewayServer(server siface.IGatewayServer) error {
	return pb.RegisterDemoServiceHandlerFromEndpoint(
		context.Background(),
		server.GatewayRuntimeMux(),
		server.Endpoint(),
		server.GatewayOption(),
	)
}

//---------------- zinx ----------------

func (s *Service) PreHandle(_ ziface.IRequest) {

}

func (s *Service) Handle(request ziface.IRequest) {
	switch request.GetMsgID() {
	case 1:
		s.handleHiRequest(request)
	case 2:
		s.handleWatchRequest(request)
	}
}

func (s *Service) handleHiRequest(request ziface.IRequest) {
	req := &pb.HiRequest{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		s.logger.Error("Failed to unmarshal Hi request", zap.Error(err))
		return
	}

	if err := s.demoHandler.Hi(request.GetConnection().Context(),
		req.GetUid(), req.GetTopic(), req.GetMessage()); err != nil {
		s.logger.Error("Hi handler failed", zap.Error(err))
	}
}

func (s *Service) handleWatchRequest(request ziface.IRequest) {
	req := &pb.WatchRequest{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		s.logger.Error("Failed to unmarshal Watch request", zap.Error(err))
		return
	}

	err := s.demoHandler.Watch(request.GetConnection().Context(), req.GetTopic(),
		func(message string) error {
			resp := &pb.WatchResponse{Message: message}
			data, err := proto.Marshal(resp)
			if err != nil {
				return fmt.Errorf("failed to marshal response: %w", err)
			}

			if err := request.GetConnection().SendMsg(2, data); err != nil {
				return fmt.Errorf("failed to send message: %w", err)
			}
			return nil
		})

	if err != nil {
		s.logger.Error("Watch handler failed", zap.Error(err))
	}
}

func (s *Service) PostHandle(_ ziface.IRequest) {

}

func (s *Service) RegisterWithServer(server siface.IZinxServer) {
	server.ZinxServer().AddRouter(1, s)
	server.ZinxServer().AddRouter(2, s)
}

func NewService(
	logger *zap.Logger,
	coll diface.ICollection,
	mq miface.MessageQueue,
	gdb *gorm.DB,
	redis *redis.Client,
) (result *Service, err error) {
	handler := domain.NewDemo(
		logger,
		db_nosql.OpenDatabase(logger, coll),
		mq,
		gdb,
		redis,
	)

	result = &Service{
		logger:      logger,
		demoHandler: handler,
	}
	return
}

var GrpcService = fx.Provide(
	func(
		l *zap.Logger,
		dProvider ofx.DocumentStoreParams,
		setting dfx.SettingsParams,
		mqParams mfx.MessageQueueParams,
		gParams ofx.GormParams,
		redisClient ofx.RedisParams,
	) (out sfx.GrpcServiceResult, err error) {
		if coll, err := dProvider.DriverProvider.OpenDbDriver(setting.DbName); err != nil {
			return out, err
		} else if s, err := NewService(
			l,
			coll,
			mqParams.MessageQueue,
			gParams.GormDB,
			redisClient.Redis,
		); err != nil {
			return out, err
		} else {
			out.GrpcService = s
		}
		return
	},
)

var HttpService = fx.Provide(
	func(
		l *zap.Logger,
		dProvider ofx.DocumentStoreParams,
		setting dfx.SettingsParams,
		mqParams mfx.MessageQueueParams,
		gParams ofx.GormParams,
		redisClient ofx.RedisParams,
	) (out sfx.GatewayServiceResult, err error) {
		if coll, err := dProvider.DriverProvider.OpenDbDriver(setting.DbName); err != nil {
			return out, err
		} else if s, err := NewService(
			l,
			coll,
			mqParams.MessageQueue,
			gParams.GormDB,
			redisClient.Redis,
		); err != nil {
			return out, err
		} else {
			out.GatewayService = s
		}
		return
	},
)

var TcpService = fx.Provide(
	func(
		l *zap.Logger,
		dProvider ofx.DocumentStoreParams,
		setting dfx.SettingsParams,
		mqParams mfx.MessageQueueParams,
		gParams ofx.GormParams,
		redisClient ofx.RedisParams,
	) (out sfx.ZinxServiceResult, err error) {
		if coll, err := dProvider.DriverProvider.OpenDbDriver(setting.DbName); err != nil {
			return out, err
		} else if s, err := NewService(
			l,
			coll,
			mqParams.MessageQueue,
			gParams.GormDB,
			redisClient.Redis,
		); err != nil {
			return out, err
		} else {
			out.ZinxService = s
		}
		return
	},
)
