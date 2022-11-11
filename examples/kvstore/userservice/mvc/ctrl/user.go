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

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
)

var (
	userSingleInst *UserCtrl = nil
	userOnce       sync.Once
)

type UserCtrl struct {
	base.BaseCtrl
}

func (userCtrl *UserCtrl) Initialize() error {
	userCtrl.BaseCtrl.Initialize()
	userCtrl.reset()

	return nil
}

func (user *UserCtrl) OnRun() {
	user.BaseCtrl.OnRun()

	// 处理用户登录
	user.HandleProto("userservice", "UserService", "Login", &pb.UserService_ServiceDesc, gingrpc.Handler{Proto: &pb.LoginReq{}, HandleProto: user.onLogin})
}

func (userCtrl *UserCtrl) OnStop() {
	defer userCtrl.BaseCtrl.OnStop()
	// 取消处理用户登录
	userCtrl.StopHandleProto("userservice", "UserService", "Login")
	userCtrl.reset()
}

// 用户登录
func (userCtrl *UserCtrl) onLogin(ctx context.Context, req interface{}) (interface{}, error) {
	pbReq, ok := req.(*pb.LoginReq)
	if !ok {
		return nil, view.GetUser().Error(errors.New("协议错误"), codes.InvalidArgument)
	}

	username := pbReq.Name
	password := pbReq.Password

	// 日志
	l := ctxzap.Extract(ctx)
	l.Info("user try to login", zap.String("name", username), zap.String("password", password))

	// 检查登录信息
	userInfo, err := model.GetUser().CheckLoginInfo(ctx, username, password)
	if err != nil {
		return nil, view.GetUser().Error(err, codes.InvalidArgument)
	}

	// 创建授权信息
	authInfo, err := model.GetAuth().CreateAuthInfo(ctx, userInfo)
	if err != nil {
		// 提示内部错误
		return nil, view.GetUser().Error(err, codes.InvalidArgument)
	}

	// 登录成功
	return view.GetUser().LoginOK(authInfo, userInfo), nil
}

func (userCtrl *UserCtrl) reset() {
}

func GetUser() *UserCtrl {
	if userSingleInst == nil {
		userOnce.Do(func() {
			userSingleInst = new(UserCtrl)
		})
	}

	return userSingleInst
}

func init() {
	// 注册到mvc
	mvc.RegisterCtrl("user", GetUser())
}
