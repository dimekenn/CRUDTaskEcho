package repository

import (
	"context"
	"crudTaskEcho/internal/app/model"
)

type CRUDRepository interface {
	CreateUser(ctx context.Context, req *model.User) (int, error)
	UpdateUser(ctx context.Context, req *model.UpdateUserReq) error
	GetUserById(ctx context.Context, id int) (*model.User, error)
}
