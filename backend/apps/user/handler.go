package main

import (
	"context"
	v1 "github.com/XDWow/DouyinMall/backend/rpc_gen/kitex_gen/user/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *v1.RegisterReq) (resp *v1.RegisterResp, err error) {
	// TODO: Your code here...
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *v1.LoginReq) (resp *v1.LoginResp, err error) {
	// TODO: Your code here...
	return
}

// Logout implements the UserServiceImpl interface.
func (s *UserServiceImpl) Logout(ctx context.Context, req *emptypb.Empty) (resp *emptypb.Empty, err error) {
	// TODO: Your code here...
	return
}

// Update implements the UserServiceImpl interface.
func (s *UserServiceImpl) Update(ctx context.Context, req *v1.UpdateUserReq) (resp *emptypb.Empty, err error) {
	// TODO: Your code here...
	return
}

// Delete implements the UserServiceImpl interface.
func (s *UserServiceImpl) Delete(ctx context.Context, req *v1.DeleteUserReq) (resp *emptypb.Empty, err error) {
	// TODO: Your code here...
	return
}

// Query implements the UserServiceImpl interface.
func (s *UserServiceImpl) Query(ctx context.Context, req *v1.QueryUserReq) (resp *v1.QueryUserResp, err error) {
	// TODO: Your code here...
	return
}
