## 课后练习 3.1
- Memory 子系统练习。
- 在 cgroup memory 子系统目录中创建目录结构。
```bash
cd /sys/fs/cgroup/memory
mkdir memorydemo
cd memorydemo
```
- 运行 malloc（在 linux 机器 make build）。
- 查看内存使用情况。
```bash
watch 'ps -aux | grep malloc | grep -v grep'
```
- 通过 cgroup 限制 memory
把进程添加到 cgroup 进程配置组。
```bash
echo ps -ef | grep malloc | grep -v grep | awk '{print $2}' > cgroup.procs
```
设置 memory.limit_in_bytes
```bash
echo 104960000 > memory.limit_in_bytes
```

## 课后练习 3.2
- 构建本地镜像。
- 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化。
- 将镜像推送至 docker 官方镜像仓库。
- 通过 docker 命令本地启动 httpserver。
- 通过 nsenter 进入容器查看 IP 配置。
- 作业需编写并提交 Dockerfile 及源代码。

构建命令。
```bash
docker build -t http-server:v1 .
```
启动命令。
```bash
docker run -d  -p 8080:8080 --name http-server http-server:v1
```
客户端访问 http server。
```bash
curl localhost:8080
```
查看容器日志。
```bash
$ docker logs http-server
os version: v0.0.1 
Header key: User-Agent, Header value: [curl/7.71.1] 
Header key: Accept, Header value: [*/*] 
2022/03/06 06:15:55 Success! Response code: 200
2022/03/06 06:15:55 Success! clientip: 172.17.0.1
```
将镜像推送至 docker 官方镜像仓库。
```bash
docker tag http-server:v1 cr7258/http-server:v1
docker push cr7258/http-server:v1
```

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220306142909.png)

查看容器 IP。
```bash
# 获取容器进程 id
$lsns -t net | grep http-server
4026533216 net       1 34844 root /bin/http-server
# 使用 nsenter 命令进入容器命名空间查看 IP
# -t 指定进程号，-n 表示网络命名空间
$nsenter -t 34844 -n ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN qlen 1
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
176: eth0@if177: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP 
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```
## 参考资料
- [构建 Golang 应用最小 Docker 镜像](https://juejin.cn/post/6844904174396637197)
- [go语言编译真正的静态可执行文件](https://rocket049.cn/static-go.md)
- [基于 Alpine 的 Docker 镜像编译的程序无法在云函数环境运行](https://cloud.tencent.com/developer/article/1536308)