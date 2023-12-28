package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/connect"
	talkerv1 "github.com/yolocs/try-connect/gen/talker/v1"
	"github.com/yolocs/try-connect/gen/talker/v1/talkerv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	t := &TalkerServer{}
	mux := http.NewServeMux()
	path, handler := talkerv1connect.NewTalkerServiceHandler(t)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

type TalkerServer struct{}

func (s *TalkerServer) Hello(ctx context.Context, req *connect.Request[talkerv1.HelloRequest]) (*connect.Response[talkerv1.HelloResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&talkerv1.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.Msg.Target),
	})
	res.Header().Set("Talker-Version", "v1")
	return res, nil
}
