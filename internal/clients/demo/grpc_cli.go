package demo

import (
	"context"
	"fmt"

	"github.com/abiosoft/ishell"
	mm "github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"github.com/gstones/moke-kit/logging/slogger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	pb "github.com/gstones/moke-layout/api/gen/demo/api"
)

type DemoGrpc struct {
	client pb.DemoServiceClient
	cmd    *ishell.Cmd
}

func NewDemoGrpcCli(conn *grpc.ClientConn) *DemoGrpc {
	cmd := &ishell.Cmd{
		Name:    "demo",
		Help:    "demo interactive",
		Aliases: []string{"D"},
	}
	p := &DemoGrpc{
		client: pb.NewDemoServiceClient(conn),
		cmd:    cmd,
	}
	p.initSubShells()
	return p
}

func (p *DemoGrpc) GetCmd() *ishell.Cmd {
	return p.cmd
}

func (p *DemoGrpc) initSubShells() {
	p.cmd.AddCmd(&ishell.Cmd{
		Name:    "hi",
		Help:    "say hi",
		Aliases: []string{"hi"},
		Func:    p.sayHi,
	})
	p.cmd.AddCmd(&ishell.Cmd{
		Name:    "watch",
		Help:    "watch topic",
		Aliases: []string{"w"},
		Func:    p.watch,
	})

}

func (p *DemoGrpc) sayHi(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	msg := "hello"
	in := slogger.ReadLine(c, "message(default:hello): ")
	if in != "" {
		msg = in
	}
	topic := "demo"
	t := slogger.ReadLine(c, "topic(default:demo): ")
	if t != "" {
		topic = t
	}
	md := metadata.Pairs("authorization", fmt.Sprintf("%s %v", "bearer", "test"))
	ctx := mm.MD(md).ToOutgoing(context.Background())
	if response, err := p.client.Hi(ctx, &pb.HiRequest{
		Uid:     "10000",
		Message: msg,
		Topic:   topic,
	}); err != nil {
		slogger.Warn(c, err)
	} else {
		slogger.Infof(c, "Response: %s", response.Message)
	}
}

func (p *DemoGrpc) watch(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	topic := "demo"
	t := slogger.ReadLine(c, "topic(default:demo): ")
	if t != "" {
		topic = t
	}

	md := metadata.Pairs("authorization", fmt.Sprintf("%s %v", "bearer", "test"))
	ctx := mm.MD(md).ToOutgoing(context.Background())
	if stream, err := p.client.Watch(ctx, &pb.WatchRequest{
		Topic: topic,
	}); err != nil {
		slogger.Warn(c, err)
	} else {
		for {
			if response, err := stream.Recv(); err != nil {
				slogger.Warn(c, err)
				break
			} else {
				slogger.Infof(c, "Response: %s \r\n", response.Message)
			}
		}
	}
}
