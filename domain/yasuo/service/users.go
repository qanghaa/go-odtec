package service

import (
	"context"

	"go-odtec/domain/yasuo/entity"
	"go-odtec/utils/database"

	"github.com/jackc/pgtype"
)

type UserRepository interface {
	Create(ctx context.Context, db database.QueryExecer, user *entity.User) error
	Delete(ctx context.Context, db database.QueryExecer, userId pgtype.Text) error
	RetrieveOne(ctx context.Context, db database.QueryExecer, userID pgtype.Text) (*entity.User, error)
}
type userService struct {
	repo UserRepository
	db   database.QueryExecer
}

func NewUserService(repo UserRepository, db database.QueryExecer) *userService {
	return &userService{repo: repo, db: db}
}

func (s *userService) CreateNewUser(ctx context.Context, u *entity.User) error {
	return s.repo.Create(ctx, s.db, u)
}

func (s *userService) DeleteUser(ctx context.Context, userId pgtype.Text) error {
	return s.repo.Delete(ctx, s.db, userId)
}

func (s *userService) GetUser(ctx context.Context, userId pgtype.Text) (*entity.User, error) {
	return s.repo.RetrieveOne(ctx, s.db, userId)
}
