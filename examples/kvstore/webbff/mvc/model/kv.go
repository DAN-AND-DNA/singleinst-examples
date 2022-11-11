package model

import (
	"context"
	"log"
	pb "singleinst/examples/kvstore/pb/kvservice"
	"singleinst/modules/mvc"
	"singleinst/modules/mvc/base"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	kvSingleInst *KV = nil
	kvOnce       sync.Once
)

type KV struct {
	base.BaseModel

	kvServiceConn   *grpc.ClientConn
	kvServiceClient pb.KVServiceClient
}

func (kv *KV) OnRun() {
	kv.BaseModel.OnRun()

	var err error
	kv.kvServiceConn, err = grpc.Dial("127.0.0.1:3731", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("[webbff] connect to UserService failed: %s\n", err.Error())
	}

	kv.kvServiceClient = pb.NewKVServiceClient(kv.kvServiceConn)
}

func (kv *KV) OnStop() {
	defer kv.BaseModel.OnStop()
	if kv.kvServiceConn != nil {
		kv.kvServiceConn.Close()
	}
}

func (kv *KV) OnClean() {
	defer kv.BaseModel.OnClean()
	kv.kvServiceConn = nil
	kv.kvServiceClient = nil
}

func (kv *KV) Get(ctx context.Context, key string) (*pb.KeyValue, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	resp, err := kv.kvServiceClient.Get(ctx, &pb.GetReq{Key: key})
	if err != nil {
		return nil, err
	}

	return resp.GetValue(), nil
}

func (kv *KV) Set(ctx context.Context, key, value string) (*pb.KeyValue, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	resp, err := kv.kvServiceClient.Set(ctx, &pb.SetReq{
		NewValue: &pb.KeyValue{
			Key:   key,
			Value: value,
		},
	})
	if err != nil {
		return nil, err
	}

	return resp.GetOldValue(), nil
}

func GetKV() *KV {
	if kvSingleInst == nil {
		kvOnce.Do(func() {
			kvSingleInst = new(KV)
		})
	}

	return kvSingleInst
}

func init() {
	mvc.RegisterModel("kv", GetKV())
}
