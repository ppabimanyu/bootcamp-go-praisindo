package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	handler "task_1_crud_user/internal/delivery/grpc"
	"task_1_crud_user/migration"

	pb "task_1_crud_user/api/protobuf/users/v1"
	"task_1_crud_user/internal/repository"
	"task_1_crud_user/internal/service"
	"task_1_crud_user/pkg/database"
	"task_1_crud_user/pkg/logger"
	"task_1_crud_user/pkg/validator"
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
	)

	logger.SetupLogger(&logger.SlogConfig{
		LogPath: "./logs",
		Debug:   true,
	})

	userRepo := repository.NewUserRepositoryImpl()

	userService := service.NewUserServiceImpl(validate, db.GetDB(), userRepo)

	userHandler := handler.NewUserHandler(userService)

	grpcServer := grpc.NewServer()
	pb.RegisterUsersServer(grpcServer, userHandler)

	migration.AutoMigration(db)

	s, err := net.Listen("tcp", "0.0.0.0:50051")
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
		"0.0.0.0:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	gmMux := runtime.NewServeMux()
	if err := pb.RegisterUsersHandler(context.Background(), gmMux, conn); err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}

	gwServer := gin.Default()
	gwServer.Group("v1/*{grpc_gateway}").Any("", gin.WrapH(gmMux))
	log.Println("Starting server at port 8080")
	_ = gwServer.Run(":8080")
}
