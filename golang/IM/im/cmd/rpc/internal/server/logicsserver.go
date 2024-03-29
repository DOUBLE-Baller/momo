// Code generated by goctl. DO NOT EDIT!
// Source: logic.proto

package server

import (
	"context"

	"backend/service/im/cmd/rpc/internal/logic"
	"backend/service/im/cmd/rpc/internal/svc"
	"backend/service/im/cmd/rpc/pb"
)

type LogicsServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedLogicsServer
}

func NewLogicsServer(svcCtx *svc.ServiceContext) *LogicsServer {
	return &LogicsServer{
		svcCtx: svcCtx,
	}
}

// 返送单条消息
func (s *LogicsServer) Push(ctx context.Context, in *pb.Send) (*pb.LGSuccessReply, error) {
	l := logic.NewPushLogic(ctx, s.svcCtx)
	return l.Push(in)
}
