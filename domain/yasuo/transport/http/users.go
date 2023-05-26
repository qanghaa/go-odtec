package http

import (
	"context"

	"go-odtec/domain/yasuo/entity"
	"go-odtec/utils/database"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgtype"
)

type UserService interface {
	CreateNewUser(ctx context.Context, u *entity.User) error
	DeleteUser(ctx context.Context, userId pgtype.Text) error
	GetUser(ctx context.Context, userId pgtype.Text) (*entity.User, error)
}

type UserHttp struct {
	srv UserService
	db  database.QueryExecer
}

func NewUserHttp(srv UserService, db database.QueryExecer) *UserHttp {
	return &UserHttp{srv: srv, db: db}
}

func (h *UserHttp) CreateNewUserAccount(ctx *gin.Context) {

}

func (h *UserHttp) GetUserProfile(ctx *gin.Context) {

}

func (h *UserHttp) DeleteUserAccount(ctx *gin.Context) {

}
