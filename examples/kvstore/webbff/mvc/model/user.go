package model

import (
	"context"
	"log"
	pb "singleinst/examples/kvstore/pb/userservice"
	"singleinst/modules/mvc"
	"singleinst/modules/mvc/base"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	singleInst *User = nil
	once       sync.Once
)

type User struct {
	base.BaseModel

	userServiceConn   *grpc.ClientConn
	userServiceClient pb.UserServiceClient
}

func (user *User) Initialize() error {
	user.BaseModel.Initialize()
	return nil
}

func (user *User) OnRun() {
	user.BaseModel.OnRun()
	var err error
	user.userServiceConn, err = grpc.Dial("127.0.0.1:3730", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("[webbff] connect to UserService failed: %s\n", err.Error())
	}

	user.userServiceClient = pb.NewUserServiceClient(user.userServiceConn)
}

func (user *User) OnStop() {
	defer user.BaseModel.OnStop()

	if user.userServiceConn != nil {
		user.userServiceConn.Close()
	}
}

func (user *User) OnClean() {
	defer user.BaseModel.OnClean()
	user.reset()
}

func (user *User) reset() {
	user.userServiceClient = nil
	user.userServiceConn = nil
}

// 登录
func (user *User) Login(ctx context.Context, username, password string) (*pb.UserInfo, string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resp, err := user.userServiceClient.Login(ctx, &pb.LoginReq{
		Name:     username,
		Password: password,
	})

	if err != nil {
		return nil, "", err
	}

	return resp.UserInfo, resp.Token, nil
}

func GetUser() *User {
	if singleInst == nil {
		once.Do(func() {
			singleInst = new(User)
		})
	}

	return singleInst
}

func init() {
	// 注册
	mvc.RegisterModel("user", GetUser())
}
