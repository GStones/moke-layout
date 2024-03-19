package dfx

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/tools"

	pb2 "github.com/gstones/moke-layout/api/gen/demo/api"
)

type DemoClientParams struct {
	fx.In

	DemoClient pb2.DemoServiceClient `name:"DemoClient"`
}

type DemoClientResult struct {
	fx.Out

	DemoClient pb2.DemoServiceClient `name:"DemoClient"`
}

func NewDemoClient(host string, logger *zap.Logger, sSetting sfx.SecuritySettingsParams) (pb2.DemoServiceClient, error) {
	if sSetting.TLSEnable {
		if conn, err := tools.DialWithSecurity(
			host,
			sSetting.ClientCert,
			sSetting.ClientKey,
			sSetting.ServerName,
			sSetting.ServerCaCert,
		); err != nil {
			return nil, err
		} else {
			return pb2.NewDemoServiceClient(conn), nil
		}
	} else {
		if conn, err := tools.DialInsecure(host); err != nil {
			return nil, err
		} else {
			return pb2.NewDemoServiceClient(conn), nil
		}
	}
}

// DemoClientModule  you can inject it to other module
var DemoClientModule = fx.Provide(
	func(
		setting SettingsParams,
		sSetting sfx.SecuritySettingsParams,
		logger *zap.Logger,
	) (out DemoClientResult, err error) {
		if cli, e := NewDemoClient(setting.DemoUrl, logger, sSetting); e != nil {
			err = e
		} else {
			out.DemoClient = cli
		}
		return
	},
)
