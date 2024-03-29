// Code generated by goctl. DO NOT EDIT!
// Source: systemuserget.proto

//go:generate mockgen -destination ./systemusergeter_mock.go -package systemusergeter -source $GOFILE

package systemusergeter

import (
	"context"

	"backend/service/backend/cmd/rpc/systemuserget/systemuserget"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	Response = systemuserget.Response
	Request  = systemuserget.Request

	Systemusergeter interface {
		GetSystemuser(ctx context.Context, in *Request) (*Response, error)
	}

	defaultSystemusergeter struct {
		cli zrpc.Client
	}
)

func NewSystemusergeter(cli zrpc.Client) Systemusergeter {
	return &defaultSystemusergeter{
		cli: cli,
	}
}

func (m *defaultSystemusergeter) GetSystemuser(ctx context.Context, in *Request) (*Response, error) {
	client := systemuserget.NewSystemusergeterClient(m.cli.Conn())
	return client.GetSystemuser(ctx, in)
}
