package db

import "context"

type UserDatabase struct {
	userDatabaseImpl
	*OTELTools
}

type userDatabaseImpl interface {
	GetUsers(ctx context.Context) (users []*User, err error)
	GetUserByID(ctx context.Context, id int64) (user *User, err error)
	CreateUser(ctx context.Context, req CreateUserRequest) (id int64, err error)
	DeleteUser(ctx context.Context, req DeleteUserRequest) (err error)
	UpdateUser(ctx context.Context, req UpdateUserRequest) (err error)
}

func NewUserDatabase(tools *OTELTools) *UserDatabase {
	userDatabase := userDatabaseSpanner{
		OTELTools: tools,
		next:      &userDatabase{},
	}
	return &UserDatabase{
		userDatabaseImpl: &userDatabase,
	}
}

var _ userDatabaseImpl = &userDatabase{}

type userDatabase struct {
}

func (*userDatabase) GetUsers(ctx context.Context) (users []*User, err error) {
	//GetUsers business logic goes here
	return
}

func (*userDatabase) GetUserByID(ctx context.Context, id int64) (user *User, err error) {
	//GetUserByID business logic goes here
	return
}

func (*userDatabase) CreateUser(ctx context.Context, req CreateUserRequest) (id int64, err error) {
	//CreateUser business logic goes here
	return
}

func (*userDatabase) DeleteUser(ctx context.Context, req DeleteUserRequest) (err error) {
	//DeleteUser business logic goes here
	return
}

func (*userDatabase) UpdateUser(ctx context.Context, req UpdateUserRequest) (err error) {
	//UpdateUser business logic goes here
	return
}
