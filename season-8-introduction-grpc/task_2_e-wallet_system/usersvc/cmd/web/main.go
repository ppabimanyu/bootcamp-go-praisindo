package main

import (
	"context"
	"net"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	handler "usersvc/internal/delivery/grpc"
	"usersvc/migration"

	pb "usersvc/api/protobuf/users/v1"
	"usersvc/internal/repository"
	"usersvc/internal/service"
	"usersvc/pkg/database"
	"usersvc/pkg/logger"
	"usersvc/pkg/validator"
)

func main() {
	var (
		validate = validator.NewValidator()
		db       = database.NewDatabase(&database.GormConfig{
			DbHost:   "156.67.218.177",
			DbUser:   "root",
			DbPass:   "234524",
			DbName:   "intro-grpc",
			DbPort:   "3306",
			DbDriver: "mysql",
			Debug:    true,
		})
		log = logger.NewSlog(&logger.SlogConfig{
			LogPath: "./logs",
			Debug:   true,
		})
	)

	userRepo := repository.NewUserRepositoryImpl(log)

	userService := service.NewUserServiceImpl(validate, db.GetDB(), userRepo)

	userHandler := handler.NewUserHandler(userService)

	grpcServer := grpc.NewServer()
	pb.RegisterUsersServer(grpcServer, userHandler)

	migration.AutoMigration(db)

	s, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
	go func() {
		if err := grpcServer.Serve(s); err != nil {
			panic(err)
		}
	}()
	time.Sleep(1 * time.Second)

	conn, err := grpc.NewClient(
		"0.0.0.0:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Error("failed to dial: %v", err)
		os.Exit(1)
	}

	gmMux := runtime.NewServeMux()
	if err := pb.RegisterUsersHandler(context.Background(), gmMux, conn); err != nil {
		log.Error("failed to register gateway: %v", err)
		os.Exit(1)
	}

	gwServer := gin.Default()
	gwServer.Group("v1/*{grpc_gateway}").Any("", gin.WrapH(gmMux))
	log.Info("Starting server at port 8081")
	_ = gwServer.Run(":8081")
}
