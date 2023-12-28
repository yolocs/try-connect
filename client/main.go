package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"connectrpc.com/connect"
	talkerv1 "github.com/yolocs/try-connect/gen/talker/v1"
	"github.com/yolocs/try-connect/gen/talker/v1/talkerv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

var (
	address  = flag.String("address", "http://localhost:8080", "address of the server")
	target   = flag.String("target", "Jane", "target of the hello message")
	token    = flag.String("token", "", "ID token")
	protocol = flag.String("protocol", "connect", "Allowed values are 'connect', 'grpc', 'h2'")
)

func main() {
	flag.Parse()

	var err error
	ctx := context.Background()

	switch *protocol {
	case "connect":
		err = viaConnect(ctx)
	case "grpc":
		err = viaGRPC(ctx)
	case "h2":
		err = viaHTTP(ctx)
	default:
		log.Fatalln("Protocol must be one of 'connect', 'grpc', 'h2'")
	}

	if err != nil {
		log.Fatal(err)
	}
}

func viaHTTP(ctx context.Context) error {
	addr := strings.TrimPrefix(*address, "https://")
	addr = addr + ":443"

	conf := &tls.Config{
		NextProtos: []string{"h2"},
		MinVersion: tls.VersionTLS12,
	}
	tcpConn, err := tls.Dial("tcp", addr, conf)
	if err != nil {
		return fmt.Errorf("tls.Dial: %w", err)
	}
	defer tcpConn.Close()

	cc, err := new(http2.Transport).NewClientConn(tcpConn)
	if err != nil {
		return fmt.Errorf("NewClientConn: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		*address+"/talker.v1.TalkerService/Hello",
		bytes.NewBufferString(`{"target":"`+*target+`"}`),
	)
	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	if *token != "" {
		req.Header.Set("Authorization", "Bearer "+*token)
	}
	res, err := cc.RoundTrip(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %d", res.StatusCode)
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	log.Println(res.Header)
	log.Println(string(b))
	return nil
}

func viaGRPC(ctx context.Context) error {
	oauthToken := &oauth2.Token{AccessToken: *token}
	rpcCreds := oauth.TokenSource{
		TokenSource: oauth2.StaticTokenSource(oauthToken),
	}

	pool, err := x509.SystemCertPool()
	if err != nil {
		return fmt.Errorf("failed to load system cert pool: %w", err)
	}
	creds := credentials.NewClientTLSFromCert(pool, "")

	// The address is input in a way that causes an error within grpc.
	// example input: https://my-example-server.a.run.app
	// expected url format: my-example-server.a.run.app:443
	addr := strings.TrimPrefix(*address, "https://")
	addr = addr + ":443"
	conn, err := grpc.DialContext(ctx, addr, grpc.WithPerRPCCredentials(rpcCreds), grpc.WithTransportCredentials(creds))
	if err != nil {
		return fmt.Errorf("failed to dial: %w", err)
	}
	defer conn.Close()

	client := talkerv1.NewTalkerServiceClient(conn)
	res, err := client.Hello(ctx, &talkerv1.HelloRequest{Target: *target})
	if err != nil {
		return fmt.Errorf("failed to say hello: %w", err)
	}

	log.Println("Can't print header in gRPC mode")
	log.Println(res.Message)
	return nil
}

func viaConnect(ctx context.Context) error {
	client := talkerv1connect.NewTalkerServiceClient(
		http.DefaultClient,
		*address,
	)

	req := connect.NewRequest(&talkerv1.HelloRequest{Target: *target})
	// This is how I authenticate. I don't need to worry about the SSL cert.
	if *token != "" {
		req.Header().Set("Authorization", "Bearer "+*token)
	}

	res, err := client.Hello(
		context.Background(),
		req,
	)
	if err != nil {
		return fmt.Errorf("connect request: %w", err)
	}
	log.Println(req.Header())
	log.Println(res.Msg.Message)
	return nil
}
