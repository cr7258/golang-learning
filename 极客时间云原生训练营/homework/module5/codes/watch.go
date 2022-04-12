package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// 监听
func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"11.8.36.25:20000"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Printf("connect to etcd failed, err: %v\n", err)
		return
	}

	fmt.Print("connect to etcd success\n")
	defer cli.Close()

	// 监听
	rch := cli.Watch(context.Background(), "key1")
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}

/**
# 在 etcd 操作
etcdctl put key1 value1
etcdctl put key1 value2
etcdctl del key1

# 在控制台可以看到以下输出
Type: PUT Key:key1 Value:value1
Type: PUT Key:key1 Value:value2
Type: DELETE Key:key1 Value:
**/
