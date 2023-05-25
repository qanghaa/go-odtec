package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go-odtec/domain/yasuo/entity"
	"go-odtec/utils/database"

	"github.com/jackc/pgtype"
	"go.uber.org/multierr"
)

type UserRepo struct{}

func (r *UserRepo) Create(ctx context.Context, db database.QueryExecer, user *entity.User) error {
	fields, _ := user.FieldMap()
	fieldNames := strings.Join(fields, ",")
	now := time.Now()
	err := multierr.Combine(
		user.CreatedAt.Set(now),
		user.UpdatedAt.Set(now),
	)
	if err != nil {
		return fmt.Errorf("create time now: %w", err)
	}

	placeHolders := database.GeneratePlaceholders(len(fields))
	stmt := "INSERT INTO " + user.TableName() + " (" + fieldNames + ") VALUES (" + placeHolders + ");"
	_, err = db.Exec(ctx, stmt, db)
	if err != nil {
		return fmt.Errorf("insert: %w", err)
	}

	return nil
}

func (r *UserRepo) Delete(ctx context.Context, db database.QueryExecer, userId pgtype.Text) error {
	user := &entity.User{}
	stmt := "DELETE FROM " + user.TableName() + " WHERE " + "user_id" + " = $1;"
	_, err := db.Exec(ctx, stmt, &userId)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	return nil
}

func (r *UserRepo) RetrieveOne(ctx context.Context, db database.QueryExecer, userID pgtype.Text) (*entity.User, error) {
	user := &entity.User{}
	fields, values := user.FieldMap()
	query := fmt.Sprintf("SELECT %s FROM %s WHERE user_id = $1", strings.Join(fields, ","), user.TableName())

	err := db.QueryRow(ctx, query, &userID).Scan(values...)
	if err != nil {
		return nil, fmt.Errorf("db.QueryRow: %w", err)
	}

	return user, nil
}
