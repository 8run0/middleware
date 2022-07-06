package svc

import (
	"context"

	"github.com/8run0/otel/backend/pkg/oteltools"
)

var _ userServiceImpl = &userServiceSpanner{}

type userServiceSpanner struct {
	*oteltools.OTELTools
	next userServiceImpl
}

func (s *userServiceSpanner) GetUsers(ctx context.Context) (users []*User, err error) {
	ctx, span := s.OTELTools.Tracer.Start(ctx, "userService_GetUsers")
	ctx = ctx
	defer span.End()
	return s.next.GetUsers(ctx)
}

func (s *userServiceSpanner) GetUserByID(ctx context.Context, id int64) (user *User, err error) {
	ctx, span := s.OTELTools.Tracer.Start(ctx, "userService_GetUserByID")
	ctx = ctx
	defer span.End()
	return s.next.GetUserByID(ctx, id)
}

func (s *userServiceSpanner) CreateUser(ctx context.Context, req CreateUserRequest) (id int64, err error) {
	ctx, span := s.OTELTools.Tracer.Start(ctx, "userService_CreateUser")
	ctx = ctx
	defer span.End()
	return s.next.CreateUser(ctx, req)
}

func (s *userServiceSpanner) DeleteUser(ctx context.Context, req DeleteUserRequest) (err error) {
	ctx, span := s.OTELTools.Tracer.Start(ctx, "userService_DeleteUser")
	ctx = ctx
	defer span.End()
	return s.next.DeleteUser(ctx, req)
}

func (s *userServiceSpanner) UpdateUser(ctx context.Context, req UpdateUserRequest) (err error) {
	ctx, span := s.OTELTools.Tracer.Start(ctx, "userService_UpdateUser")
	ctx = ctx
	defer span.End()
	return s.next.UpdateUser(ctx, req)
}
