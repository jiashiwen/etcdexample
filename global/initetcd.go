package global

import (
	"github.com/coreos/etcd/clientv3"
)

var (
	etcdClient *clientv3.Client
)

// 初始化etcdClient
func InitEtcd() {
	//初始化etcdclient
	var etcdCfg clientv3.Config
	err := RSPViper.UnmarshalKey("etcd", &etcdCfg)

	if err != nil {
		panic(err)
	}

	etcdClient, err = clientv3.New(etcdCfg)
	if err != nil {
		panic(err)
	}
}
