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
	"time"

	"google.golang.org/grpc/codes"
)

var (
	kvSingleInst *KV = nil
	kvOnce       sync.Once
)

type KV struct {
	base.BaseCtrl
}

func (kv *KV) Initialize() error {
	kv.BaseCtrl.Initialize()

	return nil
}

func (kv *KV) OnRun() {
	kv.BaseCtrl.OnRun()

	// 设置key value
	kv.HandleProto("webbff", "WebBFF", "Set", &pb.WebBFF_ServiceDesc, gingrpc.Handler{Proto: &pb.SetReq{}, HandleProto: kv.onSet})
	// 获取key value
	kv.HandleProto("webbff", "WebBFF", "Get", &pb.WebBFF_ServiceDesc, gingrpc.Handler{Proto: &pb.GetReq{}, HandleProto: kv.onGet})
}

func (kv *KV) OnStop() {
	defer kv.BaseCtrl.OnStop()

	kv.StopHandleProto("webbff", "WebBFF", "Set")
	kv.StopHandleProto("webbff", "WebBFF", "Get")

	kv.reset()
}

// 设置key value
func (kv *KV) onSet(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto, ok := req.(*pb.SetReq)
	if !ok {
		return nil, view.GetUser().Error(errors.New("协议错误"), codes.Internal)
	}

	key := reqProto.GetNewValue().GetKey()
	value := reqProto.GetNewValue().GetValue()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// 检查授权
	if ok, err := model.GetAuth().IsAuthorized(ctx); err != nil {
		return nil, view.GetUser().Error(err, codes.Internal)
	} else if !ok {
		return nil, view.GetUser().Error(errors.New("未授权"), codes.PermissionDenied)
	}

	// 执行key value 操作
	oldValue, err := model.GetKV().Set(ctx, key, value)
	if err != nil {
		return nil, view.GetUser().Error(err, codes.Internal)
	}

	return &pb.SetResp{OldValue: oldValue}, nil
}

// 设置key value
func (kv *KV) onGet(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto, ok := req.(*pb.GetReq)
	if !ok {
		return nil, view.GetUser().Error(errors.New("协议错误"), codes.Internal)
	}

	key := reqProto.GetKey()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// 检查授权
	if ok, err := model.GetAuth().IsAuthorized(ctx); err != nil {
		return nil, view.GetUser().Error(err, codes.Internal)
	} else if !ok {
		return nil, view.GetUser().Error(errors.New("未授权"), codes.PermissionDenied)
	}

	// 执行key value 操作
	value, err := model.GetKV().Get(ctx, key)
	if err != nil {
		return nil, view.GetUser().Error(err, codes.Internal)
	}

	return &pb.GetResp{Value: value}, nil
}

func (kv *KV) reset() {
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
