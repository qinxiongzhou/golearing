package setup

import (
	"fmt"
	conf "ch45-seckill/pkg/config"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

//初始化Etcd
func InitZk() {
	var hosts = []string{"39.98.179.73:2181"}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		fmt.Println(err)
		return
	}
	conf.Zk.ZkConn = conn
	conf.Zk.SecProductKey = "/product"
}
