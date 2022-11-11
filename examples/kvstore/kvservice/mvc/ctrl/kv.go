package ctrl

import (
	"context"
	"errors"
	gingrpc "github.com/dan-and-dna/gin-grpc"
	"singleinst/examples/kvstore/kvservice/mvc/model"
	"singleinst/examples/kvstore/kvservice/mvc/view"
	pb "singleinst/examples/kvstore/pb/kvservice"
	"singleinst/modules/mvc"
	"singleinst/modules/mvc/base"
	"sync"

	"google.golang.org/grpc/codes"
)

var (
	kvSingleInst *KV = nil
	kvOnce       sync.Once
)

type KV struct {
	base.BaseCtrl
}

func (kv *KV) OnRun() {
	kv.HandleProto("kvservice", "KVService", "Set", &pb.KVService_ServiceDesc, gingrpc.Handler{Proto: &pb.SetReq{}, HandleProto: kv.onSet})
	kv.HandleProto("kvservice", "KVService", "Get", &pb.KVService_ServiceDesc, gingrpc.Handler{Proto: &pb.GetReq{}, HandleProto: kv.onGet})
}

func (kv *KV) OnStop() {
	kv.StopHandleProto("kvservice", "KVService", "Set")
	kv.StopHandleProto("kvservice", "KVService", "Get")
}

// 存key value
func (kv *KV) onSet(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto, ok := req.(*pb.SetReq)
	if !ok {
		return nil, view.GetKV().Error(errors.New("协议错误"), codes.Internal)
	}

	oldValue, err := model.GetKV().Set(ctx, reqProto.GetNewValue())
	if err != nil {
		return nil, view.GetKV().Error(err, codes.Internal)
	}

	respProto := &pb.SetResp{
		OldValue: oldValue,
	}

	return respProto, nil
}

// 拿key value
func (kv *KV) onGet(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto, ok := req.(*pb.GetReq)
	if !ok {
		return nil, view.GetKV().Error(errors.New("协议错误"), codes.Internal)
	}

	value, err := model.GetKV().Get(ctx, reqProto.GetKey())
	if err != nil {
		return nil, view.GetKV().Error(err, codes.Internal)
	}

	respProto := &pb.GetResp{
		Value: value,
	}

	return respProto, nil
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
	mvc.RegisterCtrl("kv", GetKV())
}
