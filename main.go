package main

import (
	"flag"
	"strings"
)

var (
	MODE        = "c"
	SERVER_ADDR = "127.0.0.1:3300"
	PayLoad     = "<script>alert(112233)</script>11111111111111111111111111111111111"
	PrintDump   = false
)

//go:generate protoc --go_out=. --go-grpc_out=. *.proto
// https://tools.ietf.org/html/rfc7540

func main() {
	flag.StringVar(&MODE, "mode", MODE, "server/client")
	flag.StringVar(&SERVER_ADDR, "server", SERVER_ADDR, "mode: server => listen, mode: client => connect to")
	flag.StringVar(&PayLoad, "pl", PayLoad, "PayLoad")
	flag.BoolVar(&PrintDump, "d", PrintDump, "Print dump")
	flag.PrintDefaults()
	flag.Parse()

	switch strings.ToUpper(MODE) {
	case "S":
		server(SERVER_ADDR)
	default:
		client(SERVER_ADDR)
	}
}
