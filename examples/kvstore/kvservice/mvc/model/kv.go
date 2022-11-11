package model

import (
	"context"
	pb "singleinst/examples/kvstore/pb/kvservice"
	"singleinst/modules/mvc"
	"singleinst/modules/mvc/base"
	"sync"
)

var (
	kvSingleInst *KV = nil
	kvOnce       sync.Once
)

type KV struct {
	base.BaseModel

	localStore map[string]string
}

func (kv *KV) Get(ctx context.Context, key string) (*pb.KeyValue, error) {
	ret := &pb.KeyValue{}
	// 查本地
	if value, ok := kv.localStore[key]; ok {
		ret.Key = key
		ret.Value = value
		return ret, nil
	}

	// 查远程cache
	value, ok, err := kv.getFromCache(key)
	if err != nil {
		return nil, err
	}

	if ok {
		return value, nil
	}

	// 查db
	value, ok, err = kv.getFromDB(key)
	if err != nil {
		return nil, err
	}

	if ok {
		return value, nil
	}

	// 没找到
	return ret, nil
}

// TODO 查远程cache 写到本地
func (kv *KV) getFromCache(key string) (*pb.KeyValue, bool, error) {

	return nil, false, nil
}

// TODO 查db 写到缓存 写到本地
func (kv *KV) getFromDB(key string) (*pb.KeyValue, bool, error) {
	return nil, false, nil
}

func (kv *KV) Set(ctx context.Context, newValue *pb.KeyValue) (*pb.KeyValue, error) {
	key := newValue.GetKey()

	oldValue, err := kv.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	// 写db
	err = kv.setFromDB(newValue)
	if err != nil {
		return nil, err
	}

	// 写远程cache
	err = kv.setFromCache(newValue)
	if err != nil {
		return nil, err
	}

	// 写本地
	if kv.localStore == nil {
		kv.localStore = make(map[string]string)
	}
	kv.localStore[key] = newValue.GetValue()

	return oldValue, nil
}

// TODO 查远程cache
func (kv *KV) setFromCache(newValue *pb.KeyValue) error {
	return nil
}

// TODO 查db
func (kv *KV) setFromDB(newValue *pb.KeyValue) error {
	return nil
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
