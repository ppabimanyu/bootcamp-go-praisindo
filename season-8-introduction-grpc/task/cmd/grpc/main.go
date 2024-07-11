package main

import (
	"net"

	"google.golang.org/grpc"

	handler "intro-grpc-task/internal/delivery/grpc"
	"intro-grpc-task/migration"

	pb "intro-grpc-task/api/protobuf/users/v1"
	"intro-grpc-task/internal/repository"
	"intro-grpc-task/internal/service"
	"intro-grpc-task/pkg/database"
	"intro-grpc-task/pkg/logger"
	"intro-grpc-task/pkg/validator"
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

	s, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(s); err != nil {
		panic(err)
	}
}
