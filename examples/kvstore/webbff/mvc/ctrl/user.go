package ctrl

import (
	"context"
	"errors"
	"github.com/dan-and-dna/gin-grpc"
	pb "singleinst/examples/kvstore/pb/webbff"
	"singleinst/examples/kvstore/webbff/mvc/model"
	"singleinst/examples/kvstore/webbff/mvc/view"

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

	// 用户登录
	user.HandleProto("webbff", "WebBFF", "Login", &pb.WebBFF_ServiceDesc, gingrpc.Handler{Proto: &pb.LoginReq{}, HandleProto: user.onLogin})
}

func (userCtrl *UserCtrl) OnStop() {
	defer userCtrl.BaseCtrl.OnStop()

	userCtrl.StopHandleProto("webbff", "WebBFF", "Login")

	userCtrl.reset()
}

// user login
func (userCtrl *UserCtrl) onLogin(ctx context.Context, req interface{}) (interface{}, error) {
	pbReq, ok := req.(*pb.LoginReq)
	if !ok {
		return nil, view.GetUser().Error(errors.New("protocol error"), codes.Internal)
	}

	username := pbReq.GetName()
	password := pbReq.GetPassword()

	// zap log
	l := ctxzap.Extract(ctx)
	l.Info("user try to login", zap.String("name", username), zap.String("password", password))

	// login
	userInfo, token, err := model.GetUser().Login(ctx, username, password)
	if err != nil {
		return nil, view.GetUser().Error(err, codes.InvalidArgument)
	}

	// done
	return view.GetUser().LoginOk(token, userInfo), nil
}

func (userCtrl *UserCtrl) reset() {
}

func GetUserSingleInst() *UserCtrl {
	if userSingleInst == nil {
		userOnce.Do(func() {
			userSingleInst = new(UserCtrl)
		})
	}

	return userSingleInst
}

func init() {
	// 注册到mvc
	mvc.RegisterCtrl("user", GetUserSingleInst())
}
