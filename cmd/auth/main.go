package main

import (
	"fmt"
	"log"
	"net"

	"github.com/CatGitBon/auth-service/internal/handler"
	"github.com/CatGitBon/auth-service/pkg/auth"
	"google.golang.org/grpc"
)

func main() {

	err := run()

	fmt.Println(err)
}

func run() error {

	// configPath := flag.String("config", "./config", "path to the config file")

	// flag.Parse()

	// cfg, err := config.LoadConfig(*configPath)
	// if err != nil {
	// 	return fmt.Errorf("error loading config: %v", err)
	// }

	// fmt.Println(cfg)

	// Создание gRPC сервера
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// Регистрацию вашего AuthHandler
	authServer := grpc.NewServer()
	authHandler := handler.NewAuthHandler()
	auth.RegisterAuthServiceServer(authServer, authHandler)

	go func() {
		// Запуск сервера на порту
		fmt.Printf("Auth service gRPC server listening at %v\n", lis.Addr())
		if err := authServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	select {}
}
