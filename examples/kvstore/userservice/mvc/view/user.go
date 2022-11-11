// 处理用户信息相关的操作
package view

import (
	pb "singleinst/examples/kvstore/pb/userservice"
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

// 提示错误
func (view *User) Error(err error, code codes.Code) error {
	if _, ok := status.FromError(err); ok {
		return err
	}
	return status.New(code, err.Error()).Err()
}

// 返回协议结果
func (view *User) LoginOK(authInfo *pb.AuthInfo, userInfo *pb.UserInfo) *pb.LoginResp {
	pbResp := &pb.LoginResp{}

	if authInfo == nil || userInfo == nil {
		return pbResp
	}

	pbResp.Token = authInfo.Token
	pbResp.UserInfo = userInfo
	pbResp.UserInfo.Password = ""

	return pbResp
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
