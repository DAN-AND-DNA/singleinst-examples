package view

import (
	"singleinst/modules/mvc"
	"singleinst/modules/mvc/base"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	kvSingleInst *KV = nil
	kvOnce       sync.Once
)

type KV struct {
	base.BaseView
}

func (kv *KV) OnRun() {

}

func (kv *KV) OnStop() {

}

// 提示错误
func (kv *KV) Error(err error, code codes.Code) error {
	if _, ok := status.FromError(err); ok {
		return err
	}
	return status.New(code, err.Error()).Err()
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
	mvc.RegisterView("kv", GetKV())
}
