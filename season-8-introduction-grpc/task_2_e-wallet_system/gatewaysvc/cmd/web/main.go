package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pbtransaction "usersvc/api/protobuf/transaction/v1"
	pbuser "usersvc/api/protobuf/users/v1"
	pbwallet "usersvc/api/protobuf/wallet/v1"
)

func main() {
	connUser, err := grpc.NewClient(
		"0.0.0.0:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	connWallet, err := grpc.NewClient(
		"0.0.0.0:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	gmMux := runtime.NewServeMux()
	if err := pbuser.RegisterUsersHandler(context.Background(), gmMux, connUser); err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}
	if err := pbwallet.RegisterWalletsHandler(context.Background(), gmMux, connWallet); err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}
	if err := pbtransaction.RegisterTransactionsHandler(context.Background(), gmMux, connWallet); err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}

	gwServer := gin.Default()
	gwServer.Group("v1/*{grpc_gateway}").Any("", gin.WrapH(gmMux))
	log.Println("Starting server at port 8081")
	_ = gwServer.Run(":8081")
}
