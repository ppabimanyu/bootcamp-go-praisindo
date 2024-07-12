package grpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "walletsvc/api/protobuf/wallet/v1"
	"walletsvc/internal/service"
)

type WalletHandler struct {
	pb.UnimplementedWalletsServer
	walletService service.WalletService
}

func NewWalletHandler(
	walletService service.WalletService,
) *WalletHandler {
	return &WalletHandler{
		walletService: walletService,
	}
}

func (h *WalletHandler) GetAllWallet(ctx context.Context, empty *emptypb.Empty) (*pb.GetAllWalletsRes, error) {
	wallets, exc := h.walletService.GetAllWallets(ctx)
	if exc != nil {
		return nil, status.Error(codes.Code(exc.GetGrpcCode()), fmt.Sprint(exc.Message))
	}

	var res []*pb.Wallet
	for _, u := range wallets {
		res = append(res, &pb.Wallet{
			Id:        u.ID,
			UserId:    u.UserID,
			Balance:   u.Balance,
			CreatedAt: timestamppb.New(*u.CreatedAt),
			UpdatedAt: timestamppb.New(*u.UpdatedAt),
		})
	}

	return &pb.GetAllWalletsRes{Wallets: res}, nil
}

func (h *WalletHandler) GetDetailWallet(ctx context.Context, req *pb.GetDetailWalletReq) (*pb.GetDetailWalletRes, error) {
	wallet, exc := h.walletService.GetDetailWallet(ctx, service.GetDetailWalletReq{WalletID: req.Id})
	if exc != nil {
		return nil, status.Error(codes.Code(exc.GetGrpcCode()), fmt.Sprint(exc.Message))
	}

	return &pb.GetDetailWalletRes{Wallet: &pb.Wallet{
		Id:        wallet.ID,
		UserId:    wallet.UserID,
		Balance:   wallet.Balance,
		CreatedAt: timestamppb.New(*wallet.CreatedAt),
		UpdatedAt: timestamppb.New(*wallet.UpdatedAt),
	}}, nil
}

func (h *WalletHandler) CreateWallet(ctx context.Context, req *pb.CreateWalletReq) (*pb.MutationRes, error) {
	if exc := h.walletService.CreateWallet(ctx, service.CreateWalletReq{
		UserID: req.UserId,
	}); exc != nil {
		return nil, status.Error(codes.Code(exc.GetGrpcCode()), fmt.Sprint(exc.Message))
	}
	return &pb.MutationRes{Message: "success"}, nil
}

func (h *WalletHandler) UpdateWallet(ctx context.Context, req *pb.UpdateWalletReq) (*pb.MutationRes, error) {
	if exc := h.walletService.UpdateWallet(ctx, service.UpdateWalletReq{
		WalletID: req.Id,
		UserID:   req.UserId,
		Balance:  req.Balance,
	}); exc != nil {
		return nil, status.Error(codes.Code(exc.GetGrpcCode()), fmt.Sprint(exc.Message))
	}
	return &pb.MutationRes{Message: "success"}, nil
}

func (h *WalletHandler) DeleteWallet(ctx context.Context, req *pb.DeleteWalletReq) (*pb.MutationRes, error) {
	if exc := h.walletService.DeleteWallet(ctx, service.DeleteWalletReq{
		WalletID: req.Id,
	}); exc != nil {
		return nil, status.Error(codes.Code(exc.GetGrpcCode()), fmt.Sprint(exc.Message))
	}
	return &pb.MutationRes{Message: "success"}, nil
}
