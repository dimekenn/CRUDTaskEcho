package repository

import (
	"context"
	"crudTaskEcho/internal/app/model"
)

type CRUDRepository interface {
	CreateUser(ctx context.Context, req *model.GetUserRes) error
	UpdateUser(ctx context.Context, req *model.UpdateUserReq) error
	GetUserById(ctx context.Context, id int) (*model.GetUserRes, error)
}
