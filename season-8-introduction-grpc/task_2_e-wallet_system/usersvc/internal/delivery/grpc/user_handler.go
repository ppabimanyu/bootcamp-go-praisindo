package grpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "usersvc/api/protobuf/users/v1"
	"usersvc/internal/service"
)

type UserHandler struct {
	pb.UnimplementedUsersServer
	userService service.UserService
}

func NewUserHandler(
	userService service.UserService,
) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetAllUser(ctx context.Context, empty *emptypb.Empty) (*pb.GetAllUsersRes, error) {
	users, exc := h.userService.GetAllUsers(ctx)
	if exc != nil {
		return nil, status.Error(codes.Code(exc.GetGrpcCode()), fmt.Sprint(exc.Message))
	}

	var res []*pb.User
	for _, u := range users {
		res = append(res, &pb.User{
			Id:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			CreatedAt: timestamppb.New(*u.CreatedAt),
			UpdatedAt: timestamppb.New(*u.UpdatedAt),
		})
	}

	return &pb.GetAllUsersRes{Users: res}, nil
}

func (h *UserHandler) GetDetailUser(ctx context.Context, req *pb.GetDetailUserReq) (*pb.GetDetailUserRes, error) {
	user, exc := h.userService.GetDetailUser(ctx, service.GetDetailUserReq{UserID: req.Id})
	if exc != nil {
		return nil, status.Error(codes.Code(exc.GetGrpcCode()), fmt.Sprint(exc.Message))
	}

	return &pb.GetDetailUserRes{User: &pb.User{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: timestamppb.New(*user.CreatedAt),
		UpdatedAt: timestamppb.New(*user.UpdatedAt),
	}}, nil
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.MutationRes, error) {
	if exc := h.userService.CreateUser(ctx, service.CreateUserReq{
		Name:  req.Name,
		Email: req.Email,
	}); exc != nil {
		return nil, status.Error(codes.Code(exc.GetGrpcCode()), fmt.Sprint(exc.Message))
	}
	return &pb.MutationRes{Message: "success"}, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.MutationRes, error) {
	if exc := h.userService.UpdateUser(ctx, service.UpdateUserReq{
		UserID: req.Id,
		Name:   req.Name,
		Email:  req.Email,
	}); exc != nil {
		return nil, status.Error(codes.Code(exc.GetGrpcCode()), fmt.Sprint(exc.Message))
	}
	return &pb.MutationRes{Message: "success"}, nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (*pb.MutationRes, error) {
	if exc := h.userService.DeleteUser(ctx, service.DeleteUserReq{
		UserID: req.Id,
	}); exc != nil {
		return nil, status.Error(codes.Code(exc.GetGrpcCode()), fmt.Sprint(exc.Message))
	}
	return &pb.MutationRes{Message: "success"}, nil
}
