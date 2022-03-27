## 课后练习 5.1
## 本地搭建 Etcd 集群
### 填写集群信息
etcd 官方提供了图形化生成配置文件的工具 [Install and deploy etcd](http://play.etcd.io/install)。


在页面上填写相关信息，在本地创建一个 3 节点的 etcd 集群，使用 HTTPS 加密，注意 etcd 节点的数据目录、证书路径。
![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220327131249.png)


### 创建证书

安装页面提示的命令进行操作。
![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220327131905.png)

#### 安装 cfssl 证书管理工具
```bash
rm -f /tmp/cfssl* && rm -rf /tmp/certs && mkdir -p /tmp/certs

curl -L https://pkg.cfssl.org/R1.2/cfssl_linux-amd64 -o /tmp/cfssl
chmod +x /tmp/cfssl
sudo mv /tmp/cfssl /usr/local/bin/cfssl

curl -L https://pkg.cfssl.org/R1.2/cfssljson_linux-amd64 -o /tmp/cfssljson
chmod +x /tmp/cfssljson
sudo mv /tmp/cfssljson /usr/local/bin/cfssljson

/usr/local/bin/cfssl version
/usr/local/bin/cfssljson -h

mkdir -p /tmp/certs
```

#### 生成 CA 证书
```bash
mkdir -p /tmp/certs

cat > /tmp/certs/etcd-root-ca-csr.json <<EOF
{
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "names": [
    {
      "O": "etcd",
      "OU": "etcd Security",
      "L": "San Francisco",
      "ST": "California",
      "C": "USA"
    }
  ],
  "CN": "etcd-root-ca"
}
EOF
cfssl gencert --initca=true /tmp/certs/etcd-root-ca-csr.json | cfssljson --bare /tmp/certs/etcd-root-ca

# verify
openssl x509 -in /tmp/certs/etcd-root-ca.pem -text -noout


# cert-generation configuration
cat > /tmp/certs/etcd-gencert.json <<EOF
{
  "signing": {
    "default": {
        "usages": [
          "signing",
          "key encipherment",
          "server auth",
          "client auth"
        ],
        "expiry": "87600h"
    }
  }
}
EOF
```
执行完成以后会生成以下文件。
```bash
# CSR configuration
/tmp/certs/etcd-root-ca-csr.json

# CSR
/tmp/certs/etcd-root-ca.csr

# self-signed root CA public key
/tmp/certs/etcd-root-ca.pem

# self-signed root CA private key
/tmp/certs/etcd-root-ca-key.pem

# cert-generation configuration for other TLS assets
/tmp/certs/etcd-gencert.json
```
#### 生成服务器证书
节点 1：
```bash
mkdir -p /tmp/certs

cat > /tmp/certs/s1-ca-csr.json <<EOF
{
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "names": [
    {
      "O": "etcd",
      "OU": "etcd Security",
      "L": "San Francisco",
      "ST": "California",
      "C": "USA"
    }
  ],
  "CN": "s1",
  "hosts": [
    "127.0.0.1",
    "localhost"
  ]
}
EOF
cfssl gencert \
  --ca /tmp/certs/etcd-root-ca.pem \
  --ca-key /tmp/certs/etcd-root-ca-key.pem \
  --config /tmp/certs/etcd-gencert.json \
  /tmp/certs/s1-ca-csr.json | cfssljson --bare /tmp/certs/s1

# verify
openssl x509 -in /tmp/certs/s1.pem -text -noout
```

节点 2：
```bash
mkdir -p /tmp/certs

cat > /tmp/certs/s2-ca-csr.json <<EOF
{
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "names": [
    {
      "O": "etcd",
      "OU": "etcd Security",
      "L": "San Francisco",
      "ST": "California",
      "C": "USA"
    }
  ],
  "CN": "s2",
  "hosts": [
    "127.0.0.1",
    "localhost"
  ]
}
EOF
cfssl gencert \
  --ca /tmp/certs/etcd-root-ca.pem \
  --ca-key /tmp/certs/etcd-root-ca-key.pem \
  --config /tmp/certs/etcd-gencert.json \
  /tmp/certs/s2-ca-csr.json | cfssljson --bare /tmp/certs/s2

# verify
openssl x509 -in /tmp/certs/s2.pem -text -noout
```

节点 3：
```bash
mkdir -p /tmp/certs

cat > /tmp/certs/s3-ca-csr.json <<EOF
{
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "names": [
    {
      "O": "etcd",
      "OU": "etcd Security",
      "L": "San Francisco",
      "ST": "California",
      "C": "USA"
    }
  ],
  "CN": "s3",
  "hosts": [
    "127.0.0.1",
    "localhost"
  ]
}
EOF
cfssl gencert \
  --ca /tmp/certs/etcd-root-ca.pem \
  --ca-key /tmp/certs/etcd-root-ca-key.pem \
  --config /tmp/certs/etcd-gencert.json \
  /tmp/certs/s3-ca-csr.json | cfssljson --bare /tmp/certs/s3

# verify
openssl x509 -in /tmp/certs/s3.pem -text -noout
```

生成的服务器证书文件如下所示：
```bash
/data/etcd/certs/s1-ca-csr.json
/data/etcd/certs/s1.csr
/data/etcd/certs/s1-key.pem
/data/etcd/certs/s1.pem
/data/etcd/certs/s2-ca-csr.json
/data/etcd/certs/s2.csr
/data/etcd/certs/s2-key.pem
/data/etcd/certs/s2.pem
/data/etcd/certs/s3-ca-csr.json
/data/etcd/certs/s3.csr
/data/etcd/certs/s3-key.pem
/data/etcd/certs/s3.pem
```

#### 安装 Etcd
根据网页提示的内容安装 etcd，选择对应的操作系统，我使用的是 Linux。
```bash
ETCD_VER=v3.3.8

# choose either URL
GOOGLE_URL=https://storage.googleapis.com/etcd
GITHUB_URL=https://github.com/coreos/etcd/releases/download
DOWNLOAD_URL=${GOOGLE_URL}

rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
rm -rf /tmp/test-etcd && mkdir -p /tmp/test-etcd

curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
tar xzvf /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz -C /tmp/test-etcd --strip-components=1

# sudo cp /tmp/test-etcd/etcd* [YOUR_EXEC_DIR]
# sudo mkdir -p /usr/local/bin/ && sudo cp /tmp/test-etcd/etcd* /usr/local/bin/

/tmp/test-etcd/etcd --version
ETCDCTL_API=3 /tmp/test-etcd/etcdctl version
```

#### 初始化 Etcd 集群
运行以下命令创建 etcd 集群，请确保端口和目录没有冲突。

节点 1：
```
# after transferring certs to remote machines
mkdir -p /data/etcd/certs
cp /tmp/certs/* /data/etcd/certs


# make sure etcd process has write access to this directory
# remove this directory if the cluster is new; keep if restarting etcd
# rm -rf /data/etcd/s1


/tmp/test-etcd/etcd --name s1 \
  --data-dir /data/etcd/s1 \
  --listen-client-urls https://localhost:2379 \
  --advertise-client-urls https://localhost:2379 \
  --listen-peer-urls https://localhost:2380 \
  --initial-advertise-peer-urls https://localhost:2380 \
  --initial-cluster s1=https://localhost:2380,s2=https://localhost:22380,s3=https://localhost:32380 \
  --initial-cluster-token tkn \
  --initial-cluster-state new \
  --client-cert-auth \
  --trusted-ca-file /data/etcd/certs/etcd-root-ca.pem \
  --cert-file /data/etcd/certs/s1.pem \
  --key-file /data/etcd/certs/s1-key.pem \
  --peer-client-cert-auth \
  --peer-trusted-ca-file /data/etcd/certs/etcd-root-ca.pem \
  --peer-cert-file /data/etcd/certs/s1.pem \
  --peer-key-file /data/etcd/certs/s1-key.pem
```

节点 2：
```bash
# after transferring certs to remote machines
mkdir -p /data/etcd/certs
cp /tmp/certs/* /data/etcd/certs


# make sure etcd process has write access to this directory
# remove this directory if the cluster is new; keep if restarting etcd
# rm -rf /data/etcd/s2


/tmp/test-etcd/etcd --name s2 \
  --data-dir /data/etcd/s2 \
  --listen-client-urls https://localhost:22379 \
  --advertise-client-urls https://localhost:22379 \
  --listen-peer-urls https://localhost:22380 \
  --initial-advertise-peer-urls https://localhost:22380 \
  --initial-cluster s1=https://localhost:2380,s2=https://localhost:22380,s3=https://localhost:32380 \
  --initial-cluster-token tkn \
  --initial-cluster-state new \
  --client-cert-auth \
  --trusted-ca-file /data/etcd/certs/etcd-root-ca.pem \
  --cert-file /data/etcd/certs/s2.pem \
  --key-file /data/etcd/certs/s2-key.pem \
  --peer-client-cert-auth \
  --peer-trusted-ca-file /data/etcd/certs/etcd-root-ca.pem \
  --peer-cert-file /data/etcd/certs/s2.pem \
  --peer-key-file /data/etcd/certs/s2-key.pem
```

节点 3：
```bash
# after transferring certs to remote machines
mkdir -p /data/etcd/certs
cp /tmp/certs/* /data/etcd/certs


# make sure etcd process has write access to this directory
# remove this directory if the cluster is new; keep if restarting etcd
# rm -rf /data/etcd/s3


/tmp/test-etcd/etcd --name s3 \
  --data-dir /data/etcd/s3 \
  --listen-client-urls https://localhost:32379 \
  --advertise-client-urls https://localhost:32379 \
  --listen-peer-urls https://localhost:32380 \
  --initial-advertise-peer-urls https://localhost:32380 \
  --initial-cluster s1=https://localhost:2380,s2=https://localhost:22380,s3=https://localhost:32380 \
  --initial-cluster-token tkn \
  --initial-cluster-state new \
  --client-cert-auth \
  --trusted-ca-file /data/etcd/certs/etcd-root-ca.pem \
  --cert-file /data/etcd/certs/s3.pem \
  --key-file /data/etcd/certs/s3-key.pem \
  --peer-client-cert-auth \
  --peer-trusted-ca-file /data/etcd/certs/etcd-root-ca.pem \
  --peer-cert-file /data/etcd/certs/s3.pem \
  --peer-key-file /data/etcd/certs/s3-key.pem
```

查看集群状态：
```bash
ETCDCTL_API=3 /tmp/test-etcd/etcdctl \
  --endpoints localhost:2379,localhost:22379,localhost:32379 \
  --cacert /data/etcd/certs/etcd-root-ca.pem \
  --cert /data/etcd/certs/s1.pem \
  --key /data/etcd/certs/s1-key.pem \
  endpoint health

# 返回结果
localhost:32379 is healthy: successfully committed proposal: took = 2.529ms
localhost:2379 is healthy: successfully committed proposal: took = 2.118402ms
localhost:22379 is healthy: successfully committed proposal: took = 2.165366ms
```

#### 基本操作
**由于我们的 etcd 集群使用的是自签名证书，因此在客户端在发送请求时需要指定 ca 证书，服务证书和私钥。**
查看节点信息：
```bash
# 方式一：--cacert 参数指定 CA 证书
ETCDCTL_API=3 /tmp/test-etcd/etcdctl  --cacert=/data/etcd/certs/etcd-root-ca.pem \
--cert=/data/etcd/certs/s1.pem \
--key=/data/etcd/certs/s1-key.pem \
--endpoints=https://localhost:2379 \
member list

# 方式二：--insecure-skip-tls-verify 参数不校验 CA
ETCDCTL_API=3 /tmp/test-etcd/etcdctl --cert=/data/etcd/certs/s1.pem \
--key=/data/etcd/certs/s1-key.pem \
--insecure-skip-tls-verify \
member list

# 返回结果
1a7febac07525fe7: name=s1 peerURLs=https://localhost:2380 clientURLs=https://localhost:2379 isLeader=true
8824476cc9eecda8: name=s3 peerURLs=https://localhost:32380 clientURLs=https://localhost:32379 isLeader=false
a39b736457d9b634: name=s2 peerURLs=https://localhost:22380 clientURLs=https://localhost:22379 isLeader=false
```



写一条数据：
```bash
ETCDCTL_API=3 /tmp/test-etcd/etcdctl  --cacert=/data/etcd/certs/etcd-root-ca.pem \
--cert=/data/etcd/certs/s1.pem \
--key=/data/etcd/certs/s1-key.pem \
--endpoints=https://localhost:2379 \
put k1 v1
```
查看数据细节：
```bash
ETCDCTL_API=3 /tmp/test-etcd/etcdctl  --cacert=/data/etcd/certs/etcd-root-ca.pem \
--cert=/data/etcd/certs/s1.pem \
--key=/data/etcd/certs/s1-key.pem \
--endpoints=https://localhost:2379 \
get k1 --write-out=fields

# 返回结果
"ClusterID" : 8792901916212013094
"MemberID" : 1909503891118120935
"Revision" : 2
"RaftTerm" : 13
"Key" : "k1"
"CreateRevision" : 2
"ModRevision" : 2
"Version" : 1
"Value" : "v1"
"Lease" : 0
"More" : false
"Count" : 1
```

列出所有 k 开头的键
```bash
ETCDCTL_API=3 /tmp/test-etcd/etcdctl  --cacert=/data/etcd/certs/etcd-root-ca.pem \
--cert=/data/etcd/certs/s1.pem \
--key=/data/etcd/certs/s1-key.pem \
get k --prefix  --keys-only
```
删除数据：
```bash
ETCDCTL_API=3 /tmp/test-etcd/etcdctl  --cacert=/data/etcd/certs/etcd-root-ca.pem \
--cert=/data/etcd/certs/s1.pem \
--key=/data/etcd/certs/s1-key.pem \
--endpoints=https://localhost:2379 \
del k1
```
监控数据：
```bash
ETCDCTL_API=3 /tmp/test-etcd/etcdctl  --cacert=/data/etcd/certs/etcd-root-ca.pem \
--cert=/data/etcd/certs/s1.pem \
--key=/data/etcd/certs/s1-key.pem \
--endpoints=https://localhost:2379 \
watch k1 --rev=2

# 返回结果
# 可以看到 k1 的操作记录，--rev 指定的是 Revision， Revision 在 etcd 集群中全局唯一递增
PUT
k1
v1
DELETE
k1
```

回收数据，为了防止空间不够用，必须定期释放一些用户已经声明删除的数据，这个动作就叫做 compact。compact 参数后面需要指定一个版本号。这个版本号就是写事务递增的那个版本号，compact 3，就是说把版本 3 以前的**标记删除**了的数据释放掉。

```bash
ETCDCTL_API=3 /tmp/test-etcd/etcdctl  --cacert=/data/etcd/certs/etcd-root-ca.pem \
--cert=/data/etcd/certs/s1.pem \
--key=/data/etcd/certs/s1-key.pem \
--endpoints=https://localhost:2379 \
compact 3

# 这次再 watch 查看 k1 就看不到之前的数据了
ETCDCTL_API=3 /tmp/test-etcd/etcdctl  --cacert=/data/etcd/certs/etcd-root-ca.pem \
--cert=/data/etcd/certs/s1.pem \
--key=/data/etcd/certs/s1-key.pem \
--endpoints=https://localhost:2379 \
watch k1 --rev=2

# 返回结果
watch was canceled (etcdserver: mvcc: required revision has been compacted)
Error: watch is canceled by the server
```

## 课后练习 5.2
在 Kubernetes 集群中创建一个高可用的 etcd 集群

### Helm 部署 Etcd 集群
添加 bitnami 官方 helm 仓库。
```bash
helm repo add bitnami https://charts.bitnami.com/bitnami
```
参数设置参见 [Etcd packaged by Bitnami](https://github.com/bitnami/charts/tree/master/bitnami/etcd/)。

创建 Etcd 集群，设置副本数为 3，设置 root 用户名密码为 czw123456。一旦 Etcd 部署完成，就无法通过 helm upgrade/install 命令重新修改密码，需要删除 PVC，再重新部署。
```bash
helm install -n etcd my-etcd bitnami/etcd \
  --set replicaCount=3 \
  --set auth.rbac.rootPassword=czw123456 \
  --create-namespace

# 返回结果
NAME: my-etcd
LAST DEPLOYED: Sun Mar 27 11:11:26 2022
NAMESPACE: etcd
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
CHART NAME: etcd
CHART VERSION: 6.13.7
APP VERSION: 3.5.2

** Please be patient while the chart is being deployed **

etcd can be accessed via port 2379 on the following DNS name from within your cluster:

    my-etcd.etcd.svc.cluster.local

To create a pod that you can use as a etcd client run the following command:

    kubectl run my-etcd-client --restart='Never' --image docker.io/bitnami/etcd:3.5.2-debian-10-r40 --env ROOT_PASSWORD=$(kubectl get secret --namespace etcd my-etcd -o jsonpath="{.data.etcd-root-password}" | base64 --decode) --env ETCDCTL_ENDPOINTS="my-etcd.etcd.svc.cluster.local:2379" --namespace etcd --command -- sleep infinity

Then, you can set/get a key using the commands below:

    kubectl exec --namespace etcd -it my-etcd-client -- bash
    etcdctl --user root:$ROOT_PASSWORD put /message Hello
    etcdctl --user root:$ROOT_PASSWORD get /message

To connect to your etcd server from outside the cluster execute the following commands:

    kubectl port-forward --namespace etcd svc/my-etcd 2379:2379 &
    echo "etcd URL: http://127.0.0.1:2379"

 * As rbac is enabled you should add the flag `--user root:$ETCD_ROOT_PASSWORD` to the etcdctl commands. Use the command below to export the password:

    export ETCD_ROOT_PASSWORD=$(kubectl get secret --namespace etcd my-etcd -o jsonpath="{.data.etcd-root-password}" | base64 --decode)
```

查看 etcd pod 和 service，etcd 的 pod 是通过 statefulSet 管理的，service 是 headless service。
```
$ kubectl get pod -n etcd
NAME        READY   STATUS    RESTARTS   AGE
my-etcd-0   1/1     Running   0          2m15s
my-etcd-1   1/1     Running   0          2m15s
my-etcd-2   1/1     Running   0          2m15s

$ kubectl get svc -n etcd
NAME               TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)             AGE
my-etcd            ClusterIP   10.68.122.157   <none>        2379/TCP,2380/TCP   9m57s
my-etcd-headless   ClusterIP   None            <none>        2379/TCP,2380/TCP   9m57s
```

### 查看集群状态
进入 etcd 容器。
```bash
kubectl exec -it -n etcd my-etcd-0 bash
```

健康检查:
```bash
$ etcdctl endpoint health --user root --password czw123456
127.0.0.1:2379 is healthy: successfully committed proposal: took = 22.564289ms
```

查看 member：
```bash
$ etcdctl member list
1b3a421a65cc82be, started, my-etcd-1, http://my-etcd-1.my-etcd-headless.etcd.svc.cluster.local:2380, http://my-etcd-1.my-etcd-headless.etcd.svc.cluster.local:2379,http://my-etcd.etcd.svc.cluster.local:2379, false
366d3c0e3458345e, started, my-etcd-0, http://my-etcd-0.my-etcd-headless.etcd.svc.cluster.local:2380, http://my-etcd-0.my-etcd-headless.etcd.svc.cluster.local:2379,http://my-etcd.etcd.svc.cluster.local:2379, false
ac7506143bd30c02, started, my-etcd-2, http://my-etcd-2.my-etcd-headless.etcd.svc.cluster.local:2380, http://my-etcd-2.my-etcd-headless.etcd.svc.cluster.local:2379,http://my-etcd.etcd.svc.cluster.local:2379, false
```
查看集群节点状态：
```bash
# etcd 节点通过 <pod-name>.<headless-service-name> 来访问
$ etcdctl --endpoints \ 
my-etcd-0.my-etcd-headless:2379,my-etcd-1.my-etcd-headless:2379,my-etcd-2.my-etcd-headless:2379 \
endpoint status --write-out=table
+---------------------------------+------------------+---------+---------+-----------+------------+-----------+------------+--------------------+--------+
|            ENDPOINT             |        ID        | VERSION | DB SIZE | IS LEADER | IS LEARNER | RAFT TERM | RAFT INDEX | RAFT APPLIED INDEX | ERRORS |
+---------------------------------+------------------+---------+---------+-----------+------------+-----------+------------+--------------------+--------+
| my-etcd-0.my-etcd-headless:2379 | 366d3c0e3458345e |   3.5.2 |   20 kB |     false |      false |         3 |        597 |                597 |        |
| my-etcd-1.my-etcd-headless:2379 | 1b3a421a65cc82be |   3.5.2 |   20 kB |     false |      false |         3 |        597 |                597 |        |
| my-etcd-2.my-etcd-headless:2379 | ac7506143bd30c02 |   3.5.2 |   20 kB |      true |      false |         3 |        597 |                597 |        |
+---------------------------------+------------------+---------+---------+-----------+------------+-----------+------------+--------------------+--------+
```
## 参考资料
- [Etcd packaged by Bitnami](https://bitnami.com/stack/etcd/helm)
- [云原生2期 模块五](https://shimo.im/docs/roiHSLzshJMj5EoO/read)