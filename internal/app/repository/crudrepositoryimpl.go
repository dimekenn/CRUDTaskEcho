package repository

import (
	"context"
	"crudTaskEcho/internal/app/model"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

//CRUDRepositoryImpl implements the CRUDRepository interface
type CRUDRepositoryImpl struct {
	db *pgxpool.Pool
}

//NewCRUDRepository creates  new instance of CRUDRepository API using sql.db
func NewCRUDRepository(db *pgxpool.Pool) CRUDRepository {
	return &CRUDRepositoryImpl{db: db}
}

//CreateUser function creates new user in database with fields username, lastname, age, email
func (c CRUDRepositoryImpl) CreateUser(ctx context.Context, req *model.User) (int, error) {
	var id int

	err := c.db.QueryRow(
		ctx,
		"insert into users (firstname, lastname, email, age) values ($1, $2, $3, $4) returning id",
		req.FirstName, req.LastName, req.Email, req.Age,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("user executing error: %v", err.Error())
	}

	return id, nil
}

//UpdateUserByUUID takes all fields of user entity and updates all columns in users table by user id
func (c CRUDRepositoryImpl) UpdateUser(ctx context.Context, req *model.UpdateUserReq) error {
	_, err := c.db.Query(
		ctx,
		"update users set firstname = $1, lastname = $2, email = $3, age = $4 where id = $5",
		req.FirstName, req.LastName, req.Email, req.Age, req.Id,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("user is not found: %v", err.Error())
		}
		return fmt.Errorf("update user failed: %v", err.Error())
	}
	//err = tx.Commit(ctx)
	//if err != nil{
	//	fmt.Println("commit err ",err.Error())
	//}
	return nil
}

//GetUserByUUID takes id of user and returns user entity
func (c CRUDRepositoryImpl) GetUserById(ctx context.Context, id int) (*model.User, error) {
	res := &model.User{}
	err := c.db.QueryRow(
		ctx,
		"select * from users where id = $1",
		id,
	).Scan(&res.Id, &res.FirstName, &res.LastName, &res.Email, &res.Age, &res.CreatedDate)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user is not found: %v", err.Error())
		}
		return nil, fmt.Errorf("get user error: %v", err.Error())
	}

	return res, nil
}
