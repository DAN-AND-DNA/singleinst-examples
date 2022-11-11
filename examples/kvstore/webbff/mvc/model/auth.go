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
	"google.golang.org/grpc/metadata"
)

var (
	authSingleInst *Auth = nil
	authOnce       sync.Once
)

type Auth struct {
	base.BaseModel

	userServiceConn   *grpc.ClientConn
	userServiceClient pb.UserServiceClient
}

func (auth *Auth) OnRun() {
	auth.BaseModel.OnRun()

	var err error
	auth.userServiceConn, err = grpc.Dial("127.0.0.1:3730", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("[webbff] connect to UserService failed: %s\n", err.Error())
	}

	auth.userServiceClient = pb.NewUserServiceClient(auth.userServiceConn)
}

func (auth *Auth) OnStop() {
	defer auth.BaseModel.OnStop()
	if auth.userServiceConn != nil {
		auth.userServiceConn.Close()
	}
}

func (auth *Auth) OnClean() {
	defer auth.BaseModel.OnClean()
	auth.userServiceClient = nil
	auth.userServiceConn = nil
}

func (auth *Auth) IsAuthorized(ctx context.Context) (bool, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false, nil
	}

	tokens := md.Get("token")
	if len(tokens) == 0 {
		return false, nil
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	resp, err := auth.userServiceClient.IsAuthorized(ctx, &pb.IsAuthorizedReq{
		Token: tokens[0],
	})

	if err != nil {
		return false, err
	}

	return resp.IsAuthorized, nil
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
