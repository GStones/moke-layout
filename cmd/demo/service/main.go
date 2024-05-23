package main

import (
	"github.com/gstones/moke-kit/fxmain"
	"github.com/gstones/moke-kit/mq/pkg/mfx"

	"github.com/gstones/moke-layout/pkg/modules"
)

func main() {
	fxmain.Main(
		// infrastructures
		// you can find them from moke-kit https://github.com/GStones/moke-kit
		// nats message queue
		mfx.NatsModule,
		// local(channel) message queue
		mfx.LocalModule,

		// service
		modules.AllModule,
	)
}
