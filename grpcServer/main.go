package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/ciiiii/echo-server/grpcServer/grpcServer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
)

func main() {
	addr := "0.0.0.0:8000"
	err := NewServer(addr).Listen()
	if err != nil {
		fmt.Printf("Failed to launch grpcServer\n%v\n", err)
	}
}

// Server ...
type Server struct {
	addr string
}

// NewServer ...
func NewServer(addr string) *Server {
	return &Server{addr: addr}
}

// Listen ...
func (s *Server) Listen() error {
	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	g := grpc.NewServer()
	grpcServer.RegisterEchoServer(g, s)
	reflection.Register(g)
	fmt.Printf("grpc run %s\n", s.addr)
	return g.Serve(l)
}

// Ping ...
func (s *Server) Ping(ctx context.Context, _ *grpcServer.Empty) (*grpcServer.Content, error) {
	p, _ := peer.FromContext(ctx)
	fmt.Printf("ping from %s\n", p.Addr.String())
	host, _ := os.Hostname()
	return &grpcServer.Content{Text: fmt.Sprintf("pong from %s", host)}, nil
}
