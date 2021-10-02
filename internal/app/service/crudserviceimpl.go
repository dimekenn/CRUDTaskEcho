package service

import (
	"context"
	"crudTaskEcho/internal/app/model"
	"crudTaskEcho/internal/app/repository"
)

type CRUDServiceImpl struct {
	repo repository.CRUDRepository
}

func NewCRUDService(repo repository.CRUDRepository) CRUDService {
	return &CRUDServiceImpl{repo: repo}
}

func (C CRUDServiceImpl) CreateUser(ctx context.Context, req *model.User) (*model.CreateUserRes, error){
	id, err := C.repo.CreateUser(ctx, req)
	if err != nil{
		return nil, err
	}
	res := &model.CreateUserRes{Id: id}

	return res, nil
}

func (C CRUDServiceImpl) UpdateUser(ctx context.Context, req *model.UpdateUserReq) error{
	return C.repo.UpdateUser(ctx, req)
}

func (C CRUDServiceImpl) GetUserById(ctx context.Context, id int) (*model.User, error) {
	return C.repo.GetUserById(ctx, id)
}

