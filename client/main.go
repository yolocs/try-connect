package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"connectrpc.com/connect"
	talkerv1 "github.com/yolocs/try-connect/gen/talker/v1"
	"github.com/yolocs/try-connect/gen/talker/v1/talkerv1connect"
)

var (
	address = flag.String("address", "http://localhost:8080", "address of the server")
	target  = flag.String("target", "Jane", "target of the hello message")
)

func main() {
	flag.Parse()

	client := talkerv1connect.NewTalkerServiceClient(
		http.DefaultClient,
		*address,
	)
	res, err := client.Hello(
		context.Background(),
		connect.NewRequest(&talkerv1.HelloRequest{Target: *target}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Msg.Message)
}
