package ctrl

import (
	"context"
	"errors"
	gingrpc "github.com/dan-and-dna/gin-grpc"
	pb "singleinst/examples/kvstore/pb/userservice"
	"singleinst/examples/kvstore/userservice/mvc/model"
	"singleinst/examples/kvstore/userservice/mvc/view"
	"singleinst/modules/mvc"
	"singleinst/modules/mvc/base"
	"sync"

	"google.golang.org/grpc/codes"
)

var (
	authSingleAuth *Auth = nil
	authOnce       sync.Once
)

type Auth struct {
	base.BaseCtrl
}

func (auth *Auth) OnRun() {
	auth.BaseCtrl.OnRun()

	// 处理用户登录
	auth.HandleProto("userservice", "UserService", "IsAuthorized", &pb.UserService_ServiceDesc, gingrpc.Handler{Proto: &pb.IsAuthorizedReq{}, HandleProto: auth.onIsAuthorized})
}

func (auth *Auth) OnStop() {
	defer auth.BaseCtrl.OnStop()
	// 取消处理用户登录
	auth.StopHandleProto("userservice", "UserService", "IsAuthorized")
	auth.reset()
}

func (auth *Auth) reset() {
}

func (auth *Auth) onIsAuthorized(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto, ok := req.(*pb.IsAuthorizedReq)
	if !ok {
		return nil, view.GetUser().Error(errors.New("协议错误"), codes.Internal)
	}

	token := reqProto.Token
	_, ok, err := model.GetAuth().GetAuthInfo(token)
	if err != nil {
		return nil, view.GetUser().Error(err, codes.InvalidArgument)
	}

	return view.GetAuth().AuthResult(ok), nil
}

func GetAuth() *Auth {
	if authSingleAuth == nil {
		authOnce.Do(func() {
			authSingleAuth = new(Auth)
		})
	}
	return authSingleAuth
}

func init() {
	mvc.RegisterCtrl("auth", GetAuth())
}
