package grpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"walletsvc/api/protobuf/transaction/v1"
	pb "walletsvc/api/protobuf/transaction/v1"
	"walletsvc/internal/service"
)

type TransactionHandler struct {
	pb.UnimplementedTransactionsServer
	transactionService service.TransactionService
}

func (h *TransactionHandler) GetAllTransaction(ctx context.Context, req *transaction.GetAllTransactionReq) (*transaction.GetAllTransactionsRes, error) {
	transactions, exc := h.transactionService.GetAllTransactions(ctx, service.GetAllTransactionsReq{UserID: req.UserId})
	if exc != nil {
		return nil, status.Error(codes.Code(exc.GetGrpcCode()), fmt.Sprint(exc.Message))
	}

	var res []*pb.Transaction
	for _, t := range transactions {
		res = append(res, &pb.Transaction{
			Id:          t.ID,
			WalletId:    t.WalletID,
			UserId:      t.UserID,
			ReferenceId: t.ReferenceID,
			Type:        t.Type,
			Amount:      t.Amount,
			Status:      t.Status,
			Description: t.Description,
			CreatedAt:   timestamppb.New(*t.CreatedAt),
			UpdatedAt:   timestamppb.New(*t.UpdatedAt),
		})
	}

	return &transaction.GetAllTransactionsRes{
		Transactions: res,
	}, nil
}

func (h *TransactionHandler) GetDetailTransaction(ctx context.Context, req *transaction.GetDetailTransactionReq) (*transaction.GetDetailTransactionRes, error) {
	trans, err := h.transactionService.GetTransaction(ctx, service.GetTransactionReq{UserID: req.UserId, TransactionID: req.Id})
	if err != nil {
		return nil, status.Error(codes.Code(err.GetGrpcCode()), fmt.Sprint(err.Message))
	}

	res := &transaction.Transaction{
		Id:          trans.ID,
		WalletId:    trans.WalletID,
		UserId:      trans.UserID,
		ReferenceId: trans.ReferenceID,
		Type:        trans.Type,
		Amount:      trans.Amount,
		Status:      trans.Status,
		Description: trans.Description,
		CreatedAt:   timestamppb.New(*trans.CreatedAt),
		UpdatedAt:   timestamppb.New(*trans.UpdatedAt),
	}

	return &transaction.GetDetailTransactionRes{Transaction: res}, nil
}

func (h *TransactionHandler) CreateTransaction(ctx context.Context, req *transaction.CreateTransactionReq) (*transaction.MutationRes, error) {
	// TODO implement me
	panic("implement me")
}
