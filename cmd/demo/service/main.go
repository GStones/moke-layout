package main

import (
	"github.com/gstones/moke-kit/fxmain"
	"github.com/gstones/moke-kit/mq/pkg/mfx"

	"github.com/gstones/moke-layout/internal/demo"
	"github.com/gstones/moke-layout/pkg/dfx"
)

func main() {
	fxmain.Main(
		dfx.SettingsModule,
		// config sqlite db module
		//(you can customize it as gorm support db adapter)
		dfx.SqliteDriverModule,
		// nats message queue
		mfx.NatsModule,
		// grpc server
		demo.GrpcModule,
		// http server
		demo.GatewayModule,
		// tcp/websocket server
		demo.ZinxModule,
	)
}
