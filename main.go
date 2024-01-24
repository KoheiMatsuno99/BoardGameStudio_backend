package main

import (
	"fmt"
	server "geister/server"
	"log"
	"net"
	"os"
	"os/signal"

	geisterpb "github.com/KoheiMatsuno99/BoardGameStudio_gRPC/pkg/geister/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := 50051
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	// gRPCサーバーを作成
	s := grpc.NewServer()
	// gRPCサーバーにGeisterServerを登録
	geisterpb.RegisterGeisterServiceServer(s, server.NewGeisterServiceServer())
	// サーバーリフレクションの設定
	reflection.Register(s)
	// gRPCサーバーを指定のポートで起動
	go func() {
		log.Printf("start gRPC server port: %v", port)
		err := s.Serve(listener)
		if err != nil {
			panic(err)
		}
	}()

	// Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
