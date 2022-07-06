package db

import "context"

var _ userDatabaseImpl = &userDatabaseSpanner{}

type userDatabaseSpanner struct {
	*OTELTools
	next userDatabaseImpl
}

func (s *userDatabaseSpanner) GetUsers(ctx context.Context) (users []*User, err error) {
	ctx, span := s.Tracer.Start(s.Ctx, "userDatabase_GetUsers")
	s.Ctx = ctx
	defer span.End()
	return s.next.GetUsers(ctx)
}

func (s *userDatabaseSpanner) GetUserByID(ctx context.Context, id int64) (user *User, err error) {
	ctx, span := s.Tracer.Start(s.Ctx, "userDatabase_GetUserByID")
	s.Ctx = ctx
	defer span.End()
	return s.next.GetUserByID(ctx, id)
}

func (s *userDatabaseSpanner) CreateUser(ctx context.Context, req CreateUserRequest) (id int64, err error) {
	ctx, span := s.Tracer.Start(s.Ctx, "userDatabase_CreateUser")
	s.Ctx = ctx
	defer span.End()
	return s.next.CreateUser(ctx, req)
}

func (s *userDatabaseSpanner) DeleteUser(ctx context.Context, req DeleteUserRequest) (err error) {
	ctx, span := s.Tracer.Start(s.Ctx, "userDatabase_DeleteUser")
	s.Ctx = ctx
	defer span.End()
	return s.next.DeleteUser(ctx, req)
}

func (s *userDatabaseSpanner) UpdateUser(ctx context.Context, req UpdateUserRequest) (err error) {
	ctx, span := s.Tracer.Start(s.Ctx, "userDatabase_UpdateUser")
	s.Ctx = ctx
	defer span.End()
	return s.next.UpdateUser(ctx, req)
}
