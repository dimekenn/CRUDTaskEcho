package model

import "time"

//user structure
type CreateUserReq struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Age         int32     `json:"age"`
}

type UpdateUserReq struct {
	Id          int32     `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Age         int32     `json:"age"`
}

type GetUserRes struct {
	Id          int32     `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Age         int32     `json:"age"`
	CreatedDate time.Time `json:"created_date"`
}
