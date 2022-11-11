// 处理授权相关操作
package view

import (
	pb "singleinst/examples/kvstore/pb/userservice"
	"singleinst/modules/mvc"
	"singleinst/modules/mvc/base"
	"sync"
)

var (
	authSingleInst *Auth
	authOnce       sync.Once
)

type Auth struct {
	base.BaseView
}

// 返回协议结果
func (view *Auth) AuthResult(ok bool) *pb.IsAuthorizedResp {
	pbResp := &pb.IsAuthorizedResp{
		IsAuthorized: ok,
	}

	return pbResp
}

func GetAuth() *Auth {
	if authSingleInst == nil {
		authOnce.Do(func() {
			authSingleInst = new(Auth)
		})
	}

	return authSingleInst
}

func init() {
	mvc.RegisterView("auth", GetAuth())
}
