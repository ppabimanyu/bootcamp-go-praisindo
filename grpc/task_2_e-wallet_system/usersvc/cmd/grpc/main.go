package main

import (
	"net"

	"google.golang.org/grpc"

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
		log      = logger.NewSlog(&logger.SlogConfig{
			LogPath: "./logs",
			Debug:   true,
		})
		db = database.NewDatabase(&database.GormConfig{
			DbHost:   "156.67.218.177",
			DbUser:   "root",
			DbPass:   "234524",
			DbName:   "intro-grpc",
			DbPort:   "3306",
			DbDriver: "mysql",
			Debug:    true,
			Logger:   log,
		})
	)

	userRepo := repository.NewUserRepositoryImpl(log)

	userService := service.NewUserServiceImpl(validate, db.GetDB(), userRepo)

	userHandler := handler.NewUserHandler(userService)

	grpcServer := grpc.NewServer()
	pb.RegisterUsersServer(grpcServer, userHandler)

	migration.AutoMigration(db)

	s, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(s); err != nil {
		panic(err)
	}
}
