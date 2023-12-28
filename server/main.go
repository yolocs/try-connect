package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"connectrpc.com/connect"
	talkerv1 "github.com/yolocs/try-connect/gen/talker/v1"
	"github.com/yolocs/try-connect/gen/talker/v1/talkerv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/abcxyz/pkg/serving"
)

func main() {
	ctx, done := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM)
	defer done()

	if err := realMain(ctx); err != nil {
		done()
		log.Fatal(err)
	}
	log.Println(ctx, "successful shutdown")
}

func realMain(ctx context.Context) error {
	t := &TalkerServer{}
	mux := http.NewServeMux()
	path, handler := talkerv1connect.NewTalkerServiceHandler(t)
	mux.Handle(path, handler)
	h := h2c.NewHandler(mux, &http2.Server{})

	hs := &http.Server{
		Addr:              ":8080",
		Handler:           h,
		ReadHeaderTimeout: 5 * time.Second,
	}

	srv, err := serving.New("8080")
	if err != nil {
		return fmt.Errorf("failed to new server: %w", err)
	}

	return srv.StartHTTP(ctx, hs)
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
