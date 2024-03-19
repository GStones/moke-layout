package modules

import (
	"go.uber.org/fx"

	"github.com/gstones/moke-layout/internal/services/demo"
	"github.com/gstones/moke-layout/pkg/dfx"
)

// GrpcModule  grpc service module
// can inject it to any fxmain.Main(), if you want to start a demo grpc service
var GrpcModule = fx.Module("grpcService",
	dfx.SqliteDriverModule,
	dfx.SettingsModule,
	dfx.AuthModule,
	demo.GrpcModule,
)

// HttpModule  http service module
// can inject it to any fxmain.Main(), if you want to start a demo http service
var HttpModule = fx.Module("httpService",
	dfx.SettingsModule,
	dfx.SqliteDriverModule,
	dfx.AuthModule,
	demo.GrpcModule,
	demo.GatewayModule,
)

// TcpModule  tcp service module
// can inject it to any fxmain.Main(), if you want to start a demo tcp service
var TcpModule = fx.Module("tcpService",
	dfx.SettingsModule,
	dfx.SqliteDriverModule,
	demo.TcpModule,
)

// AllModule  all service module
// can inject it to any fxmain.Main(), if you want to start demo all type services
var AllModule = fx.Module("allService",
	dfx.AuthModule,
	dfx.SettingsModule,
	dfx.SqliteDriverModule,
	demo.GrpcModule,
	demo.GatewayModule,
	demo.TcpModule,
)

// GrpcClientModule  grpc client module
// can inject it to any fxmain.Main(), if you want a demo grpc client to rpc demo service
var GrpcClientModule = fx.Module("grpcClient",
	dfx.SettingsModule,
	dfx.DemoClientModule,
)
