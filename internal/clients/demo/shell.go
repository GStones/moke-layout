package demo

import (
	"net"

	"github.com/abiosoft/ishell"

	"github.com/gstones/moke-kit/logging/slogger"
	"github.com/gstones/moke-kit/server/tools"
)

func RunGrpc(url string) {
	sh := ishell.New()
	slogger.Info(sh, "interactive demo connect to "+url)

	if conn, err := tools.DialInsecure(url); err != nil {
		slogger.Die(sh, err)
	} else {
		demoGrpc := NewDemoGrpcCli(conn)
		sh.AddCmd(demoGrpc.GetCmd())

		sh.Interrupt(func(c *ishell.Context, count int, input string) {
			if count >= 2 {
				c.Stop()
			}
			if count == 1 {
				conn.Close()
				slogger.Done(c, "interrupted, press again to exit")
			}
		})
	}
	sh.Run()
}

func RunTcp(url string) {
	sh := ishell.New()
	slogger.Info(sh, "interactive demo tcp connect to "+url)
	if conn, err := net.Dial("tcp", url); err != nil {
		slogger.Die(sh, err)
	} else {
		demoTcp := NewTcpCli(conn)
		sh.AddCmd(demoTcp.GetCmd())

		sh.Interrupt(func(c *ishell.Context, count int, input string) {
			if count >= 2 {
				c.Stop()
			}
			if count == 1 {
				conn.Close()
				slogger.Done(c, "interrupted, press again to exit")
			}
		})
	}
	sh.Run()
}
