// 用户数据相关数据模型
package model

import (
	"context"
	"errors"
	pb "singleinst/examples/kvstore/pb/userservice"
	"singleinst/modules/mvc"
	"singleinst/modules/mvc/base"
	"sync"
)

var (
	userSingleInst *User = nil
	userOnce       sync.Once
)

type User struct {
	base.BaseModel

	// 在线用户
	//onlineUsers map[uint32]*pb.OnlineUserInfo
}

func (userModel *User) Initialize() error {
	userModel.BaseModel.Initialize()
	return nil
}

func (userModel *User) OnRun() {
	userModel.BaseModel.OnRun()
}

func (userModel *User) OnStop() {
	defer userModel.BaseModel.OnStop()
}

func (userModel *User) OnClean() {
	defer userModel.BaseModel.OnClean()
	userModel.reset()
}

func (userModel *User) reset() {
}

func (userModel *User) GetUserInfo(ctx context.Context, name string) (*pb.UserInfo, error) {
	// TODO 从cache拿

	// TODO 没有则从db拿

	// TODO 然后写到cache中

	// mock 数据
	var userinfo *pb.UserInfo

	if name == "Dan" {
		userinfo = &pb.UserInfo{
			Username: "Dan",
			Password: "12345678",
			Age:      30,
			Uid:      "u10001",
		}
		return userinfo, nil
	}

	return nil, errors.New("no such user")
}

func (userModel *User) SetUserInfo(ctx context.Context, name string) error {
	// TODO 清理cache

	// TODO 写到db

	// TODO 写到cache

	return nil
}

// 检查用户的登录信息
func (model *User) CheckUser(username, password string, userInfo *pb.UserInfo) error {
	if userInfo == nil || userInfo.Username != username || userInfo.Password != password {
		return errors.New("账户或密码错误")
	}

	return nil
}

func (model *User) CheckLoginInfo(ctx context.Context, username, password string) (*pb.UserInfo, error) {
	// 拿数据
	userInfo, err := model.GetUserInfo(ctx, username)
	if err != nil {
		// 提示内部错误
		return nil, err
	}

	// 检查用户信息是否正确
	err = model.CheckUser(username, password, userInfo)
	if err != nil {
		// 提示业务错误
		return nil, err
	}

	return userInfo, nil
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
	// 注册
	mvc.RegisterModel("user", GetUser())
}
