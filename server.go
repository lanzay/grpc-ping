package main

import (
	"context"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type grpcServer struct {
	UnimplementedPingerServer
}

func (s *grpcServer) Ping(ctx context.Context, ping *PingMsg) (*PongMsg, error) {

	pingData, _ := proto.Marshal(ping)
	log.Println("[I] [SERVER] <==", ping.Id, ping.Type, ping.Tag, string(ping.Payload), string(pingData))

	pong := &PongMsg{
		Id:      ping.Id,
		Tag:     ping.Tag,
		Payload: ping.Payload,
		Type:    MsgType_Pong,
	}

	pongData, _ := proto.Marshal(pong)
	log.Println("[I] [SERVER] ==>", pong.Id, pong.Type, pong.Tag, string(pong.Payload), string(pongData))
	return pong, nil
}

func server(addr string) {

	_, port, _ := net.SplitHostPort(addr)
	port = ":" + port
	log.Println("[I] Run as SERVER listen on", port)

	srv, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	grpcSrv := grpc.NewServer()
	RegisterPingerServer(grpcSrv, &grpcServer{})
	if err := grpcSrv.Serve(srv); err != nil {
		panic(err)
	}
}
