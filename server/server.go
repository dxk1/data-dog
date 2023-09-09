/**
 * @Author: tbb
 * @Email: 645381379@qq.com
 * @Date: 2023/9/6 23:14
 */
package server

import (
	"context"
	"data-dog/conf"
	"data-dog/log"
	"data-dog/pb"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

// Server is the bootstrap module
type Server struct {
	gRPCServer *grpc.Server
	stop       chan struct{}
}

// Init init all modules of server
func Init() *Server {
	ser := &Server{
		stop: make(chan struct{}),
	}
	ser.gRPCServer = grpc.NewServer(
		grpc.MaxConcurrentStreams(10000),
		grpc.MaxRecvMsgSize(1024*1024*64),
	)
	return ser
}

func (s *Server) Run() error {
	pb.RegisterFuncMetricReportServer(s.gRPCServer, s)
	addr := fmt.Sprintf("%s:%d", "0.0.0.0", conf.GetConfig().ListenPort)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	err = s.gRPCServer.Serve(ln)
	if err != nil {
		log.Log.Error().Msgf("gRpc server Serve returns error: %+v", err)
	}
	return err
}

func (s *Server) InvokeMetricReport(ctx context.Context, request *pb.InvokeMetricRequest) (*pb.Reply, error) {
	logEvent := log.Log.Error()
	for k, v := range request.CustomFields {
		logEvent.Str(k, v)
	}
	//输出日志
	logEvent.Send()
	return &pb.Reply{
		Code: 200,
	}, nil
}
