## 课后练习 3.1
- Memory 子系统练习

编译 malloc 程序，编译完成后在该目录下会生成名为 malloc 的二进制文件。
```bash
go build -o malloc
```

在不限制内存的情况下，启动 malloc 程序。
```bash
./malloc

# 输出如下，在不断申请内存
Allocating 100Mb memory, raw memory is 104960000
Allocating 200Mb memory, raw memory is 209920000
Allocating 300Mb memory, raw memory is 314880000
```

查看内存使用情况。
```bash
top
```

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220306204950.png)

通过 cgroup 限制 memory。
首先在 cgroup memory 子系统目录中创建目录结构。
```bash
cd /sys/fs/cgroup/memory
mkdir memorydemo
cd memorydemo
```
然后把进程添加到 cgroup 进程配置组。
```bash
ps -ef | grep malloc | grep -v grep | awk '{print $2}' > cgroup.procs
```
设置 memory.limit_in_bytes，限制内存使用量为 100Mb。 
```bash
echo 104960000 > memory.limit_in_bytes
```
等待进程被 OOM Kill，如果不起作用可能是使用了 swap 的原因，解决方法是将 swappiness 设置为 0。
```bash
echo 0 > memory.swappiness
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

## 3.3 chroot
chroot 命令用来在指定的根目录下运行指令。chroot，即 change root directory （更改 root 目录）。在 linux 系统中，系统默认的目录结构都是以 `/`，即是以根 (root) 开始的。而在使用 chroot 之后，系统的目录结构将以指定的位置作为 `/` 位置。

```bash
# 建立 chroot 所需的目录
export path=/tmp/mypath
mkdir -p $path
mkdir -p $path/{bin,lib/x86_64-linux-gnu,lib64,etc,var}

# 拷贝所需命令都目录里
cp -vf /bin/{bash,ls} $path/bin

# 拷贝依赖库
list=`ldd /bin/ls | egrep -o '/lib.*\.[0-9]'`
for i in $list; do sudo cp -vf $i $path/$i; done
list=`ldd /bin/bash | egrep -o '/lib.*\.[0-9]'`
for i in $list; do sudo cp -vf $i $path/$i; done

# chroot 进入环境，只能执行 ls 和 bash 命令
chroot $path /bin/bash
```

通过解压容器镜像获取运行时文件系统包（filesystem bundle）。
```bash
docker pull busybox
docker image save busybox -o busybox.tar
mkdir -p /tmp/chengzw/busybox
tar -xf busybox.tar -C /tmp/chengzw/busybox/
cd /tmp/chengzw/busybox/2461e071255f4a5cc48606b1b674e9113f7595485867bcc837e24d9ee13ef06b
# layer 目录存放的是运行时文件系统包，解压 layer 文件
tar -xf layer.tar
```

查看解压完的目录。
```
$ls 
bin  dev  etc  home  json  layer.tar  root  tmp  usr  var  VERSION
```
chroot 切换根目录。
```
chroot /tmp/chengzw/busybox/2461e071255f4a5cc48606b1b674e9113f7595485867bcc837e24d9ee13ef06b /bin/sh
```

## 参考资料
- [构建 Golang 应用最小 Docker 镜像](https://juejin.cn/post/6844904174396637197)
- [go语言编译真正的静态可执行文件](https://rocket049.cn/static-go.md)
- [基于 Alpine 的 Docker 镜像编译的程序无法在云函数环境运行](https://cloud.tencent.com/developer/article/1536308)
- [如何通过 Cgroups 机制实现资源限制](https://my.oschina.net/u/4923278/blog/4980857)
- [Linux CGroup 基础](https://wudaijun.com/2018/10/linux-cgroup/)
- [Go 编译 binutils 库问题](https://www.cnblogs.com/xuelisheng/p/10452111.html)
- [cgroup内存限制不起作用的原因](https://segmentfault.com/a/1190000037504275)
- [云原生2期 模块3](https://shimo.im/docs/yBL3u6MFMiw6OTLT/read)
- [chroot的用法](https://cloud.tencent.com/developer/article/1603505)
