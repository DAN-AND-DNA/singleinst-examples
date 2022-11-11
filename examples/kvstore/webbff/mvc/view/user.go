package view

import (
	pb_userservice "singleinst/examples/kvstore/pb/userservice"
	pb "singleinst/examples/kvstore/pb/webbff"
	"singleinst/modules/mvc"
	"singleinst/modules/mvc/base"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	userSingleInst *User = nil
	userOnce       sync.Once
)

type User struct {
	base.BaseView
}

func (user *User) OnRun() {
}

func (user *User) OnStop() {
}

func (user *User) OnClean() {
}

// 提示服务器错误
func (user *User) Error(err error, code codes.Code) error {
	if err == nil {
		return status.Error(codes.OK, "")
	}

	if s, ok := status.FromError(err); ok {
		if s.Code() != codes.InvalidArgument {
			// 内部错误
			return status.Error(codes.Internal, "服务器内部错误")
		}

		// 普通业务错误
		return err
	}
	return status.Error(code, err.Error())
}

func (user *User) LoginOk(token string, userInfo *pb_userservice.UserInfo) *pb.LoginResp {
	resp := &pb.LoginResp{
		Token:        token,
		BaseUserinfo: userInfo,
	}

	resp.BaseUserinfo.Password = ""

	return resp
}

func GetUser() *User {
	if userSingleInst == nil {
		userOnce.Do(func() {
			userSingleInst = new(User)
		})
	}

	return userSingleInst
}

func init() {
	mvc.RegisterView("user", GetUser())
}
