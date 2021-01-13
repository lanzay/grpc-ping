package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func client(srv string) {

	log.Println("[I] Run as CLIENT connect to", srv)

	ping := &PingMsg{
		Id:      1,
		Tag:     "My ping",
		Payload: []byte(PayLoad),
		Type:    MsgType_Ping,
	}

	// grpc.WithTransportCredentials
	// grpc.WithPerRPCCredentials
	// grpc.WithCredentialsBundle
	// grpc.WithContextDialer(tcpConn)
	conn, err := grpc.Dial(srv, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := NewPingerClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for {
		pingData, _ := proto.Marshal(ping)
		t1 := time.Now()
		log.Println("[I] [CLIENT] ==>", ping.Id, ping.Type, ping.Tag, string(ping.Payload), string(pingData))
		if PrintDump {
			fmt.Println(hex.Dump(pingData))
		}

		pong, err := client.Ping(ctx, ping)
		if err != nil {
			panic(err)
		}
		pongData, _ := proto.Marshal(pong)
		log.Println("[I] [CLIENT] <==", pong.Id, pong.Type, pong.Tag, string(pong.Payload), string(pongData))
		if PrintDump {
			fmt.Println(hex.Dump(pongData))
		}
		log.Println("[I] [CLIENT] APDEX", time.Since(t1))
		ping.Id++
		<-time.NewTimer(1 * time.Second).C
	}
}
