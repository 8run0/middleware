package db

import "context"

type userDatabaseImpl interface {
	GetUsers(ctx context.Context) (users []*User, err error)
	GetUserByID(ctx context.Context, id int64) (user *User, err error)
	CreateUser(ctx context.Context, req CreateUserRequest) (id int64, err error)
	DeleteUser(ctx context.Context, req DeleteUserRequest) (err error)
	UpdateUser(ctx context.Context, req UpdateUserRequest) (err error)
}
