package main

import (
	"net"

	"google.golang.org/grpc"

	handler "walletsvc/internal/delivery/grpc"
	"walletsvc/migration"

	pb "walletsvc/api/protobuf/wallet/v1"
	"walletsvc/internal/repository"
	"walletsvc/internal/service"
	"walletsvc/pkg/database"
	"walletsvc/pkg/logger"
	"walletsvc/pkg/validator"
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

	walletRepo := repository.NewWalletRepositoryImpl(log)

	walletService := service.NewUserServiceImpl(validate, db.GetDB(), walletRepo)

	walletHandler := handler.NewWalletHandler(walletService)

	grpcServer := grpc.NewServer()
	pb.RegisterWalletsServer(grpcServer, walletHandler)

	migration.AutoMigration(db)

	s, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(s); err != nil {
		panic(err)
	}
}
