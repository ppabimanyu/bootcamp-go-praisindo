package main

import (
	"net"

	"google.golang.org/grpc"

	pbtransaction "walletsvc/api/protobuf/transaction/v1"
	pbwallet "walletsvc/api/protobuf/wallet/v1"
	handler "walletsvc/internal/delivery/grpc"
	"walletsvc/migration"

	"walletsvc/internal/repository"
	"walletsvc/internal/service"
	"walletsvc/pkg/database"
	"walletsvc/pkg/logger"
	"walletsvc/pkg/validator"
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
			DbPrefix: "",
			DbDriver: "mysql",
			Debug:    true,
			Logger:   log,
		})
	)

	var (
		walletRepo      = repository.NewWalletRepositoryImpl(log)
		transactionRepo = repository.NewTransactionRepositoryImpl(log)
	)

	var (
		walletService      = service.NewWalletServiceImpl(validate, db.GetDB(), walletRepo)
		transactionService = service.NewTransactionServiceImpl(validate, db.GetDB(), walletRepo, transactionRepo)
	)

	var (
		walletHandler      = handler.NewWalletHandler(walletService)
		transactionHandler = handler.NewTransactionHandler(transactionService)
	)

	grpcServer := grpc.NewServer()
	pbwallet.RegisterWalletsServer(grpcServer, walletHandler)
	pbtransaction.RegisterTransactionsServer(grpcServer, transactionHandler)

	migration.AutoMigration(db)

	s, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		panic(err)
	}
	if err := grpcServer.Serve(s); err != nil {
		panic(err)
	}
}
