package grpc

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go-odtec/domain/yasuo/entity"
	ypb "go-odtec/pkg/protobuf/yasuo"
	"go-odtec/utils/database"
	"go-odtec/utils/idutil"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx"
	"go.uber.org/multierr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService interface {
	CreateNewUser(ctx context.Context, u *entity.User) error
	DeleteUser(ctx context.Context, userId pgtype.Text) error
	GetUser(ctx context.Context, userId pgtype.Text) (*entity.User, error)
}

type userGRPC struct {
	srv UserService
	ypb.UnimplementedUserServiceServer
}

var _ ypb.UserServiceServer = &userGRPC{}

func (s *userGRPC) Register(server *grpc.Server) {
	ypb.RegisterUserServiceServer(server, s)
}

func NewUserGRPC(srv UserService) *userGRPC {
	return &userGRPC{srv: srv}
}

func (u *userGRPC) CreateNewUserAccount(ctx context.Context, req *ypb.CreateNewUserAccountRequest) (*ypb.CreateNewUserAccountResponse, error) {
	user := entity.User{}
	userID := idutil.ULIDNow()
	now := time.Now()
	if err := multierr.Combine(
		user.UserID.Set(userID),
		user.Email.Set(req.GetEmail()),
		user.Name.Set(req.GetName()),
		user.Avatar.Set(req.GetAvatar()),
		user.Birthday.Set(req.GetBirthday()),
		user.Gender.Set(req.Gender.String()),
		user.Street.Set(req.GetAddress().GetStreet()),
		user.City.Set(req.GetAddress().GetCity()),
		user.State.Set(req.GetAddress().GetState()),
		user.Country.Set(req.GetAddress().GetCountry()),
		user.CreatedAt.Set(now),
		user.UpdatedAt.Set(now),
		user.DeletedAt.Set(nil),
	); err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Errorf("multierr.Combine: %w", err).Error())
	}
	err := u.srv.CreateNewUser(ctx, &user)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("srv.CreateNewUser: %w", err).Error())
	}
	return &ypb.CreateNewUserAccountResponse{UserId: user.UserID.String}, nil
}

func (u *userGRPC) DeleteUserAccount(ctx context.Context, req *ypb.DeleteUserAccountRequest) (*ypb.DeleteUserAccountResponse, error) {
	err := u.srv.DeleteUser(ctx, database.Text(req.UserId))
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("srv.DeleteUser: %w", err).Error())
	}
	return &ypb.DeleteUserAccountResponse{}, nil
}

func (u *userGRPC) GetUserProfile(ctx context.Context, req *ypb.GetUserProfileRequest) (*ypb.GetUserProfileResponse, error) {
	user, err := u.srv.GetUser(ctx, database.Text(req.UserId))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Error(codes.NotFound, fmt.Errorf("srv.GetUser: %w", err).Error())
		}
		return nil, status.Error(codes.Internal, fmt.Errorf("srv.GetUser: %w", err).Error())
	}
	return &ypb.GetUserProfileResponse{
		UserId: user.UserID.String,
		Email:  user.Email.String,
		Name:   user.Name.String,
		Avatar: user.Avatar.String,
		Gender: ypb.Gender(ypb.Gender_value[user.Gender.String]),
		Address: &ypb.Address{
			Street:  user.Street.String,
			City:    user.City.String,
			State:   user.State.String,
			Country: user.Country.String,
		},
	}, nil
}
