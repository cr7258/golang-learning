# 课后练习 4.1
## 安装 Ubuntu 虚拟机

### 下载 Ubuntu 镜像

镜像地址：http://mirrors.aliyun.com/ubuntu-releases/
![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220313215413.png)



### Vmware 新建虚拟机

安装下图所示一步一步操作即可。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220312114046.png)

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220312114003.png)

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220313205606.png)

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220313110631.png)c
![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220312140038.png)

### 操作系统设置

安装下图所示一步一步操作即可。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220312140225.png)

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220312140258.png)

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220312140319.png)

IP 地址先设置为 DHCP，自动获取，稍后通过 netplan 工具修改为静态 IP。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220312140404.png)

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220312140431.png)
设置镜像源地址，这个可以配置阿里云(http://mirrors.aliyun.com/ubuntu/)，提升依赖包下载速度。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220312140554.png)

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220312140640.png)

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220312140704.png)

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220312140734.png)

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220312140759.png)

设置主机名和用户。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220312140855.png)

启用 SSH 服务。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220312140924.png)

推荐的服务都不安装。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220313111219.png)

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220313114741.png)

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220313115405.png)

### 网络设置

Ubuntu 在升级到 17.10 版本之后，使用 netplan 来管理网络设置，Netplan 是 Canonical 推出的一款抽象网络配置生成器，可以通过一个或多个 yaml 格式的配置文件，生成所需的网络配置，并应用到系统中。编辑 /etc/netplan/00-installer-config.yaml 文件，为 Ubuntu 虚拟机设置静态 IP，根据实际情况进行设置。
```yaml
network:
  ethernets:
    ens33:  # 网卡名称
      addresses: [192.168.26.10/24]  # 静态 IP 地址
      routes:  # 路由, gateway4 参数的写法已经废弃
        - to: default
          via: 192.168.26.2
      nameservers:
        addresses: [114.114.114.114] # DNS 服务器
      dhcp4: no 
      optional: no
  version: 2
```

要设置的网络信息可以通过以下方式查看。
```bash
# 安装网络工具
root@k8s-node-1:~#sudo apt install net-tools

# 查看网卡名称, 192.168.26.128 是一开始 DHCP 动态分配的 IP 地址
root@k8s-node-1:~# ifconfig
ens33: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 192.168.26.128  netmask 255.255.255.0  broadcast 192.168.26.255
        inet6 fe80::20c:29ff:fee9:6bf4  prefixlen 64  scopeid 0x20<link>
        ether 00:0c:29:e9:6b:f4  txqueuelen 1000  (Ethernet)
        RX packets 6149  bytes 1083733 (1.0 MB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 5180  bytes 872586 (872.5 KB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

lo: flags=73<UP,LOOPBACK,RUNNING>  mtu 65536
        inet 127.0.0.1  netmask 255.0.0.0
        inet6 ::1  prefixlen 128  scopeid 0x10<host>
        loop  txqueuelen 1000  (Local Loopback)
        RX packets 1290  bytes 113305 (113.3 KB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 1290  bytes 113305 (113.3 KB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

# 查看路由, 默认网关是 192.168.26.2
root@k8s-node-1:~# route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         192.168.26.2    0.0.0.0         UG    100    0        0 ens33
192.168.26.0    0.0.0.0         255.255.255.0   U     0      0        0 ens33
192.168.26.2    0.0.0.0         255.255.255.255 UH    100    0        0 ens33
```

应用 netplan 设置。
```bash
sudo netplan apply
```

### 克隆虚拟机

使用 k8s-node-1 虚拟机克隆出另外两个虚拟机，CPU 至少要 2 核，否则 kubeadm 初始化集群时会报错。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220313142637.png)

克隆之前先要关闭虚拟机。


![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220313125428.png)

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220313125441.png)

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220313125456.png)

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220313125528.png)

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220313125545.png)



克隆完成以后可以看到总共有 3 台虚拟机，依次启动 3 台虚拟机。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220313125751.png)



登录另外两台虚拟机，修改主机名和网络配置。修改主机名命令如下，修改网络配置的方式前面已经介绍了。

```bash
sudo hostnamectl set-hostname <主机名>
```
全部修改完成以后用 SSH 软件登录 3台虚拟机。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220313130255.png)

## 环境准备
后面执行的所有命令都切换到 root 用户执行，除了 kubeadm 初始化集群的命令区分 master 和 worker 节点，其余命令在每个节点上都要执行。
```bash
sudo -i
# 输入 chengzw 用户的密码
```
### 关闭 Swap

```bash
swapoff -a && sed -ri 's/.*swap.*/#&/' /etc/fstab 
```

### 调整内核参数
```bash
cat <<EOF | sudo tee /etc/modules-load.d/containerd.conf
overlay
br_netfilter
EOF

modprobe overlay
modprobe br_netfilter

cat <<EOF | sudo tee /etc/sysctl.d/99-kubernetes-cri.conf
net.bridge.bridge-nf-call-iptables  = 1
net.ipv4.ip_forward                 = 1
net.bridge.bridge-nf-call-ip6tables = 1
EOF

sysctl --system
```

### 设置时区
```bash
cp /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime
```
## 安装 Containerd
```bash
apt-get install containerd
```

查看 containerd 进程状态。
```bash
root@k8s-node-1:~# systemctl status containerd.service 
● containerd.service - containerd container runtime
     Loaded: loaded (/lib/systemd/system/containerd.service; enabled; vendor preset: enabled)
     Active: active (running) since Sun 2022-03-13 05:15:28 UTC; 55s ago
       Docs: https://containerd.io
    Process: 1916 ExecStartPre=/sbin/modprobe overlay (code=exited, status=0/SUCCESS)
   Main PID: 1917 (containerd)
      Tasks: 7
     Memory: 15.6M
        CPU: 510ms
     CGroup: /system.slice/containerd.service
             └─1917 /usr/bin/containerd

Mar 13 05:15:28 k8s-node-1 containerd[1917]: time="2022-03-13T05:15:28.263452406Z" level=info msg=serving... address=/run/containerd/containerd.sock.ttrpc
Mar 13 05:15:28 k8s-node-1 containerd[1917]: time="2022-03-13T05:15:28.263621182Z" level=info msg=serving... address=/run/containerd/containerd.sock
Mar 13 05:15:28 k8s-node-1 systemd[1]: Started containerd container runtime.
Mar 13 05:15:28 k8s-node-1 containerd[1917]: time="2022-03-13T05:15:28.274255961Z" level=info msg="containerd successfully booted in 0.069267s"
Mar 13 05:15:28 k8s-node-1 containerd[1917]: time="2022-03-13T05:15:28.276889167Z" level=info msg="Start subscribing containerd event"
Mar 13 05:15:28 k8s-node-1 containerd[1917]: time="2022-03-13T05:15:28.277247382Z" level=info msg="Start recovering state"
Mar 13 05:15:28 k8s-node-1 containerd[1917]: time="2022-03-13T05:15:28.277431025Z" level=info msg="Start event monitor"
Mar 13 05:15:28 k8s-node-1 containerd[1917]: time="2022-03-13T05:15:28.277558591Z" level=info msg="Start snapshots syncer"
Mar 13 05:15:28 k8s-node-1 containerd[1917]: time="2022-03-13T05:15:28.277673612Z" level=info msg="Start cni network conf syncer"
Mar 13 05:15:28 k8s-node-1 containerd[1917]: time="2022-03-13T05:15:28.277785501Z" level=info msg="Start streaming server"
```

生成 containerd 默认配置。
```bash
mkdir /etc/containerd/
containerd config default > /etc/containerd/config.toml
```

编辑 containerd 配置文件 /etc/containerd/config.toml，设置 pause 容器使用的镜像仓库地址。
```bash
registry.aliyuncs.com/google_containers/pause:3.5
```

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220313170348.png)

保存退出，重启 containerd。
```bash
systemctl restart containerd
```
## 配置 Crictl
```bash
VERSION="v1.23.0"
wget https://github.com/kubernetes-sigs/cri-tools/releases/download/$VERSION/crictl-$VERSION-linux-amd64.tar.gz
tar zxvf crictl-$VERSION-linux-amd64.tar.gz -C /usr/local/bin
```

设置 crictl 连接 containerd。
```bash
cat > /etc/crictl.yaml << EOF
runtime-endpoint: unix:///var/run/containerd/containerd.sock
image-endpoint: unix:///var/run/containerd/containerd.sock
timeout: 2
debug: false
EOF
```
## 安装 Kubeadm, Kubectl, Kubelet

```bash
# 更新 apt 包索引，安装相关依赖
apt-get update
apt-get install -y apt-transport-https ca-certificates curl

# 下载 aliyun 公开签名秘钥
curl -s https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | sudo apt-key add -

# 添加 Kubernetes apt 仓库
tee /etc/apt/sources.list.d/kubernetes.list <<-'EOF'
deb https://mirrors.aliyun.com/kubernetes/apt kubernetes-xenial main
EOF

# 更新 apt 包索引
apt-get update
apt-get install -y kubelet kubeadm kubectl
apt-mark hold kubelet kubeadm kubectl
```


##  用 Kubeadm 安装 Kubernetes 集群。

在 master 节点上初始化 Kubernetes 集群。
```bash
kubeadm init \
 --image-repository registry.aliyuncs.com/google_containers \
 --kubernetes-version v1.23.4 \
 --pod-network-cidr=77.0.0.0/16 \
 --service-cidr=133.0.0.0/16 \
 --apiserver-advertise-address=192.168.26.10
```

提示以下内容表示 Kubernetes 集群初始化成功。
![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220313170724.png)

设置 kubectl 连接凭证。
```bash
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

worker 节点加入集群。
```bash

kubeadm join 192.168.26.10:6443 --token 623hmx.8l8suqe29hbipodj \
        --discovery-token-ca-cert-hash sha256:7681a324d0937d04634afa9a150383759e80014f5e5ad0e85c5e8ce492a502cd 
```

查看节点状态，此时节点还是 notReady 状态，因为还没安装网络插件。
```bash
$ kubectl get node -o wide
NAME         STATUS     ROLES                  AGE   VERSION   INTERNAL-IP     EXTERNAL-IP   OS-IMAGE             KERNEL-VERSION      CONTAINER-RUNTIME
k8s-node-1   NotReady   control-plane,master   17m   v1.23.4   192.168.26.10   <none>        Ubuntu 20.04.4 LTS   5.4.0-104-generic   containerd://1.5.5
k8s-node-2   NotReady   <none>                 10s   v1.23.4   192.168.26.11   <none>        Ubuntu 20.04.4 LTS   5.4.0-104-generic   containerd://1.5.5
k8s-node-3   NotReady   <none>                 17s   v1.23.4   192.168.26.12   <none>        Ubuntu 20.04.4 LTS   5.4.0-104-generic   containerd://1.5.5
```

## 安装 Calico 网络插件

calico 默认给 Pod 分配的 CIDR 是 192.168.0.0/16，在初始化集群的时候使用 `--pod-network-cidr` 参数自定义了 Pod 的 CIDR，因此需要修改 calico 默认的地址池配置。使用以下命令获取 calico 地址池设置的 yaml 文件。
```bash
 curl -o pool.yaml https://docs.projectcalico.org/manifests/custom-resources.yaml
```

编辑 pool.yaml 文件修改 CIDR。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220314223634.png)

修改完成后，部署 calico。

```bash
kubectl create -f https://docs.projectcalico.org/manifests/tigera-operator.yaml
kubectl create -f pool.yaml
```
确保所有容器已经正常运行。
```bash
$ kubectl get pod -n kube-system
root@k8s-node-1:~# kubectl  get pod -A
NAMESPACE         NAME                                      READY   STATUS            RESTARTS        AGE
calico-system     calico-kube-controllers-84d8677f4-7hvph   1/1     Running           0               101m
calico-system     calico-node-2sz45                         1/1     Running   0               101m
calico-system     calico-node-kdc5w                         1/1     Running           0               101m
calico-system     calico-node-xsqbb                         1/1     Running           0               101m
calico-system     calico-typha-6db8d7b4dc-bmvms             1/1     Running           0               101m
calico-system     calico-typha-6db8d7b4dc-nn6lz             1/1     Running           0               101m
kube-system       coredns-6d8c4cb4d-24c9r                   1/1     Running           0               11h
kube-system       coredns-6d8c4cb4d-mlcdf                   1/1     Running           0               11h
kube-system       etcd-k8s-node-1                           1/1     Running           2               11h
kube-system       kube-apiserver-k8s-node-1                 1/1     Running           2               11h
kube-system       kube-controller-manager-k8s-node-1        1/1     Running           12 (3h9m ago)   11h
kube-system       kube-proxy-5vz46                          1/1     Running           0               11h
kube-system       kube-proxy-8gmgs                          1/1     Running           0               11h
kube-system       kube-proxy-gbhjc                          1/1     Running           0               11h
kube-system       kube-scheduler-k8s-node-1                 1/1     Running           10 (3h9m ago)   11h
tigera-operator   tigera-operator-b876f5799-bvst7           1/1     Running           0               101m
```
查看节点状态。
```bash
$ kubectl get node -o wide
NAME         STATUS   ROLES                  AGE   VERSION
k8s-node-1   Ready    control-plane,master   11h   v1.23.4
k8s-node-2   Ready    <none>                 11h   v1.23.4
k8s-node-3   Ready    <none>                 11h   v1.23.4
```
## Kubectl 命令补全
```bash
echo 'source <(kubectl completion bash)' >> ~/.bashrc
```

# 课后练习 4.2
- 启动一个 Envoy Deployment。
- 要求 Envoy 的启动配置从外部的配置文件 Mount 进 Pod。
- 进入 Pod 查看 Envoy 进程和配置。
- 更改配置的监听端口并测试访问入口的变化。
- 通过非级联删除的方法逐个删除对象。

## 部署 Nginx 服务
```bash
kubectl apply -f nginx-deploy.yaml
```
## 根据 envoy.yaml 配置文件创建 Configmap
```bash
kubectl create configmap envoy-config --from-file=envoy.yaml
```

## 启动 Envoy Pod 挂载 Configmap
```bash
kubectl apply -f envoy-deploy.yaml
```

## 进入 Pod 查看 Envoy 进程和配置
```bash
# 进入 Pod
$ kubectl exec -it envoy-fb5d77cc9-276lt -- bash

# 查看进程
root@envoy-fb5d77cc9-276lt:/# ps -ef
UID         PID   PPID  C STIME TTY          TIME CMD
envoy         1      0  0 03:23 ?        00:00:00 envoy -c /etc/envoy/envoy.yaml
root         23      0  0 03:24 pts/0    00:00:00 bash
root         35     23  0 03:24 pts/0    00:00:00 ps -ef

# 查看配置
root@envoy-fb5d77cc9-276lt:/# cat /etc/envoy/envoy.yaml     
admin:
  address:
    socket_address: { address: 127.0.0.1, port_value: 9901 }

static_resources:
  listeners:
    - name: listener_0
      address:
        socket_address: { address: 0.0.0.0, port_value: 10000 }
      filter_chains:
        - filters:
            - name: envoy.filters.network.http_connection_manager
              typed_config:
                "@type": type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager
                stat_prefix: ingress_http
                codec_type: AUTO
                route_config:
                  name: local_route
                  virtual_hosts:
                    - name: local_service
                      domains: ["*"]
                      routes:
                        - match: { prefix: "/" }
                          route: { cluster: some_service }
                http_filters:
                  - name: envoy.filters.http.router
  clusters:
    - name: some_service
      connect_timeout: 0.25s
      type: LOGICAL_DNS
      lb_policy: ROUND_ROBIN
      load_assignment:
        cluster_name: some_service
        endpoints:
          - lb_endpoints:
              - endpoint:
                  address:
                    socket_address:
                      address: nginx
                      port_value: 80
```


## 更改配置的监听端口并测试访问入口的变化
当前 envoy 监听 10000 端口，将请求转发到后端的 nginx。
```bash
$ kubectl get pod -o wide  | grep envoy
envoy-fb5d77cc9-276lt         1/1     Running   0          3m13s   172.20.142.13   11.8.36.159   <none>           <none>

root@cluster01-1:/root #curl 172.20.142.13:10000
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
html { color-scheme: light dark; }
body { width: 35em; margin: 0 auto;
font-family: Tahoma, Verdana, Arial, sans-serif; }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>
```

修改 configmap 中 envoy 监听端口的配置，将监听端口改为 20000。
```bash
kubectl edit configmap envoy-config
```

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220312113043.png)

查看 envoy 中的配置可以看到监听端口已经更新。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220312113303.png)

但是 envoy 并不会对配置进行热加载，是否会热加载配置取决于应用程序，和 Kubernetes 无关。因此此时 envoy 监听的端口还是 10000。
```bash
$ curl 172.20.142.13:20000
curl: (7) Failed connect to 172.20.142.13:20000; Connection refused

$ curl 172.20.142.13:10000
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
html { color-scheme: light dark; }
body { width: 35em; margin: 0 auto;
font-family: Tahoma, Verdana, Arial, sans-serif; }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>
```

## 通过非级联删除的方法逐个删除对象
```bash
kubectl delete deployment envoy --cascade=orphan
kubectl delete replicaset envoy-xxxxx --cascade=orphan
kubectl delete pod envoy-xxxxx --cascade=orphan
```

# 参考资料
- [Netplan configuration examples](https://netplan.io/examples/)
- [Ubuntu 21.10 - Netplan: object has no attribute 'state'](https://askubuntu.com/questions/1369646/ubuntu-21-10-netplan-object-has-no-attribute-state)
- [在 Ubuntu 18.04 中使用 Netplan 配置网络](https://blog.xmuu.dev/2021/02/03/config-network-in-ubuntu-18-04-with-netplan/)
- [使用Kubeadm工具快速安装Kubernetes集群](https://www.yuque.com/xiamucc/io9h18/ia011v)
- [云原生2期 模块4](https://shimo.im/docs/F6ZzRT0qPZkcEa2n/read)
- [使用 kubeadm 创建集群](https://kubernetes.io/zh/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/#%E5%85%B3%E4%BA%8E-apiserver-advertise-address-%E5%92%8C-controlplaneendpoint-%E7%9A%84%E6%B3%A8%E6%84%8F%E4%BA%8B%E9%A1%B9?accessToken=eyJhbGciOiJIUzI1NiIsImtpZCI6ImRlZmF1bHQiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOiJhY2Nlc3NfcmVzb3VyY2UiLCJleHAiOjE2NDcxNDA4MjIsImciOiJGNlp6UlQwcVBaa2NFYTJuIiwiaWF0IjoxNjQ3MTQwNTIyLCJ1c2VySWQiOjY2ODcxNDMwfQ.z1rclpkkn1OQ8vdG6Ke495HgtFezXLd0iFrxrcTS-1c)
- [Container runtimes](https://kubernetes.io/docs/setup/production-environment/container-runtimes/#cri-o)
- [CRI Plugin Config Guide](https://github.com/containerd/cri/blob/master/docs/config.md)
- [kubeadm v1.14.2 init failed to pull pause image when using cri and private registry ( seems to pull 2 times ..)](https://github.com/kubernetes/kubernetes/issues/79422)