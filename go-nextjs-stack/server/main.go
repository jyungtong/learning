package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	greetv1 "github.com/jyungtong/learning/go-nextjs-stack/gen/greet/v1"
	"github.com/jyungtong/learning/go-nextjs-stack/gen/greet/v1/greetv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const defaultAddress = "localhost:8080"

func main() {
	address := serverAddress()
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(&greetServer{})
	mux.Handle(path, handler)

	log.Printf("Listening on %s", address)
	log.Fatal(http.ListenAndServe(address, h2c.NewHandler(mux, &http2.Server{})))
}

func serverAddress() string {
	if address := os.Getenv("ADDRESS"); address != "" {
		return address
	}

	return defaultAddress
}

type greetServer struct {
	greetv1connect.UnimplementedGreetServiceHandler
}

func (s *greetServer) SayHello(
	_ context.Context,
	req *greetv1.SayHelloRequest,
) (*greetv1.SayHelloResponse, error) {
	name := req.GetName()
	if name == "" {
		name = "world"
	}

	return &greetv1.SayHelloResponse{
		Message: fmt.Sprintf("Hello, %s!", name),
	}, nil
}
