package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// 增删改查
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

	// 添加 key
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = cli.Put(ctx, "key1", "value1")
	fmt.Printf("put key to etcd\n")
	cancel()
	if err != nil {
		fmt.Printf("put key to etcd failed, err:%v\n", err)
		return
	}

	// 获取 key
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "key1")
	cancel()
	if err != nil {
		fmt.Printf("get key from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("key:%s, value:%s\n", ev.Key, ev.Value)
	}

	// 删除 key
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Delete(ctx, "key1")
	fmt.Printf("delete key from etcd\n")
	cancel()
	if err != nil {
		fmt.Printf("delete from etcd failed, err:%v\n", err)
		return
	}
}
