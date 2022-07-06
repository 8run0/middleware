package svc

import (
	"context"
	"fmt"

	"github.com/8run0/otel/backend/pkg/oteltools"
)

var (
	ErrCreateUser   = fmt.Errorf("failed to create user")
	ErrDeleteUser   = fmt.Errorf("failed to delete user")
	ErrUpdateUser   = fmt.Errorf("failed to update user")
	ErrNoUserFound  = fmt.Errorf("failed to get user")
	ErrNoUsersFound = fmt.Errorf("failed to get users")
)

type userServiceImpl interface {
	GetUsers(ctx context.Context) (users []*User, err error)
	GetUserByID(ctx context.Context, id int64) (user *User, err error)
	CreateUser(ctx context.Context, req CreateUserRequest) (id int64, err error)
	DeleteUser(ctx context.Context, req DeleteUserRequest) (err error)
	UpdateUser(ctx context.Context, req UpdateUserRequest) (err error)
}

type User struct {
	ID   int64
	Name string
}

type CreateUserRequest struct {
	Name     string
	Password string
}

type DeleteUserRequest struct {
	ID int64
}

type UpdateUserRequest struct {
	ID      int64
	Updates User
}

type UserService struct {
	userServiceImpl
	*oteltools.OTELTools
}

func NewUserService(tools *oteltools.TELTools) *UserService {
	userService := userServiceSpanner{
		OTELTools: tools,
		next:      &userService{},
	}
	return &UserService{
		userServiceImpl: &userService,
	}
}

var _ userServiceImpl = &userService{}

type userService struct {
}

func (*userService) GetUsers(ctx context.Context) (users []*User, err error) {
	//GetUsers business logic goes here
	return
}

func (*userService) GetUserByID(ctx context.Context, id int64) (user *User, err error) {
	//GetUserByID business logic goes here
	return
}

func (*userService) CreateUser(ctx context.Context, req CreateUserRequest) (id int64, err error) {
	//CreateUser business logic goes here
	return
}

func (*userService) DeleteUser(ctx context.Context, req DeleteUserRequest) (err error) {
	//DeleteUser business logic goes here
	return
}

func (*userService) UpdateUser(ctx context.Context, req UpdateUserRequest) (err error) {
	//UpdateUser business logic goes here
	return
}
