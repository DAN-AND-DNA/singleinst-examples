// 处理授权相关数据模型
package model

import (
	"context"
	pb "singleinst/examples/kvstore/pb/userservice"
	"singleinst/modules/mvc"
	"singleinst/modules/mvc/base"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	authSingleInst *Auth
	authOnce       sync.Once
)

type Auth struct {
	base.BaseModel
	authInfoCache map[string]*pb.AuthInfo
	sync.RWMutex
}

func (model *Auth) Initialize() error {
	return nil
}

// 根据token获得授权信息
func (model *Auth) GetAuthInfo(token string) (*pb.AuthInfo, bool, error) {
	if info, ok := model.getAuthInfoFromLocal(token); ok {
		return info, true, nil
	}

	if info, ok, err := model.getAuthInfoFromRemote(token); ok {
		return info, true, nil
	} else if err != nil {
		return nil, false, err
	}

	return nil, false, nil
}

// 从本地拿
func (model *Auth) getAuthInfoFromLocal(token string) (*pb.AuthInfo, bool) {
	model.RLock()
	defer model.RUnlock()
	if info, ok := model.authInfoCache[token]; ok {
		return info, true
	}

	return nil, false
}

// TODO 从远程缓存拿
func (model *Auth) getAuthInfoFromRemote(token string) (*pb.AuthInfo, bool, error) {
	return nil, false, nil
}

// 保存授权信息
func (model *Auth) SetAuthInfo(ctx context.Context, authInfo *pb.AuthInfo) error {
	if err := model.setAuthInfoToRemote(ctx, authInfo); err != nil {
		return status.New(codes.Internal, err.Error()).Err()
	}

	model.setAuthInfoToLocal(authInfo)
	return nil
}

// TODO 存到远程缓存
func (model *Auth) setAuthInfoToRemote(ctx context.Context, authInfo *pb.AuthInfo) error {
	return nil
}

// 存到本地
func (model *Auth) setAuthInfoToLocal(authInfo *pb.AuthInfo) {
	if authInfo == nil {
		return
	}
	if model.authInfoCache == nil {
		model.authInfoCache = make(map[string]*pb.AuthInfo)
	}
	model.authInfoCache[authInfo.Token] = authInfo
}

// 创建授权信息
func (model *Auth) CreateAuthInfo(ctx context.Context, userInfo *pb.UserInfo) (*pb.AuthInfo, error) {
	nowUnix := uint64(time.Now().Unix())

	newAuthInfo := &pb.AuthInfo{
		Token:       "login_abcdef1234567",
		CreatedUnix: nowUnix,
		ExpiredUnix: nowUnix,
		UserUid:     userInfo.Uid,
		Username:    userInfo.Username,
	}

	err := model.SetAuthInfo(ctx, newAuthInfo)
	if err != nil {
		return nil, err
	}

	return newAuthInfo, nil
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
	mvc.RegisterModel("auth", GetAuth())
}
