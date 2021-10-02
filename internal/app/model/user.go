package model

import "time"


type UpdateUserReq struct {
	Id          int32     `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Age         int32     `json:"age"`
}

//user structure
type User struct {
	Id          int32     `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Age         int32     `json:"age"`
	CreatedDate time.Time `json:"created_date"`
}

type CreateUserRes struct {
	Id int `json:"id"`
}
