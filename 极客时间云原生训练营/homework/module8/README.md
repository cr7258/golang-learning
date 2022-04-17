## 1 第一部分

现在你对 Kubernetes 的控制面板的工作机制是否有了深入的了解呢？
是否对如何构建一个优雅的云上应用有了深刻的认识，那么接下来用最近学过的知识把你之前编写的 http 以优雅的方式部署起来吧，你可能需要审视之前代码是否能满足优雅上云的需求。
作业要求：编写 Kubernetes 部署脚本将 httpserver 部署到 Kubernetes 集群，以下是你可以思考的维度。

- 优雅启动
- 优雅终止
- 资源需求和 QoS 保证
- 探活
- 日常运维需求，日志等级
- 配置和代码分离

创建资源:

```bash
kubectl apply -f 8.1
```

查看创建的资源:

```bash
# Configmap
$ kubectl describe  cm game-demo
Name:         game-demo
Namespace:    default
Labels:       <none>
Annotations:  <none>

Data
====
game.properties:
----
enemy.types=aliens,monsters
player.maximum-lives=5

loglevel:
----
debug
player_initial_lives:
----
3
ui_properties_file_name:
----
user-interface.properties
user-interface.properties:
----
color.good=purple
color.bad=yellow
allow.textmode=true


BinaryData
====

Events:  <none>


# Pod
$ kubectl get pod -A -o yaml
httpserver-6889d79f96-867rw   1/1     Running       0          57s
httpserver-6889d79f96-9tlh2   1/1     Running       0          57s
httpserver-6889d79f96-rwfts   1/1     Running       0          57s
```

查看环境变量：

```bash
$ kubectl exec httpserver-6889d79f96-867rw  \
-- env | grep PLAYER_INITIAL_LIVES

PLAYER_INITIAL_LIVES=3
```

查看 downwardAPI 卷挂载的文件：

```bash
$ kubectl exec httpserver-6889d79f96-867rw  \
-- cat /etc/podinfo/labels

app="httpserver"
pod-template-hash="6889d79f96"
```

查看 ConfigMap 卷挂载的文件：

```bash
$ kubectl exec httpserver-6889d79f96-867rw  \
-- cat /etc/config-volume/game.properties

enemy.types=aliens,monsters
player.maximum-lives=5
```

查看 Init 容器和 PostStart 阶段写入的日志：

```

```

```bash
$ kubectl exec httpserver-6889d79f96-867rw  \
-- cat /tmp/test.log

init-httpserver-1 start!
init-httpserver-2 start!
Hello from the postStart handler
```

删除资源，查看 Pod 删除前 PreStop 写入的日志：

```bash
$ kubectl delete -f 8.1
$ kubectl exec httpserver-6889d79f96-867rw  \
-- tail /tmp/test.log

# 新增一行
Bye from the preStop handler
```


## 2 第二部分


除了将 httpServer 应用优雅的运行在 Kubernetes 之上，我们还应该考虑如何将服务发布给对内和对外的调用方。
来尝试用 Service, Ingress 将你的服务发布给集群外部的调用方吧。
在第一部分的基础上提供更加完备的部署 spec，包括（不限于）：

- Service
- Ingress
- 可以考虑的细节
- 如何确保整个应用的高可用。
- 如何通过证书保证 httpServer 的通讯安全。


### 2.1 创建 Service

```bash
kubectl apply -f 8.2/service.yaml
```

### 2.2 部署 Nginx Ingress Controller


官方文档：[kubernetes/ingress-nginx](https://github.com/kubernetes/ingress-nginx/tree/main/charts/ingress-nginx?accessToken=eyJhbGciOiJIUzI1NiIsImtpZCI6ImRlZmF1bHQiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOiJhY2Nlc3NfcmVzb3VyY2UiLCJleHAiOjE2NDk1NTYyMzcsImciOiJJMlRlRElJdWs4MDFOVm1zIiwiaWF0IjoxNjQ5NTU1OTM3LCJ1c2VySWQiOjY2ODcxNDMwfQ.GUxwEnq41l41BDYE0BC_SzwEMUaM7HmpHBRmSSkMQ0w)

```bash
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm install ingress-nginx ingress-nginx/ingress-nginx --create-namespace --namespace ingress
```

查看资源：

```bash
$ kubectl get pod -n ingress
NAME                                       READY   STATUS    RESTARTS   AGE
ingress-nginx-controller-dd7cd7565-g9gdk   1/1     Running   0          13h
```

helm 安装 nginx ingress controller 时会创建一个 IngressClass 对象。在 Ingress 中设置 `ingressClassName: nginx` 字段或者 `kubernetes.io/ingress.class: nginx` 注解可以指定使用 nginx ingress controller。

```yaml
$ kubectl get ingressclass nginx -o yaml
apiVersion: networking.k8s.io/v1
kind: IngressClass
metadata:
  annotations:
    meta.helm.sh/release-name: ingress-nginx
    meta.helm.sh/release-namespace: ingress
  creationTimestamp: "2022-04-16T16:02:46Z"
  generation: 1
  labels:
    app.kubernetes.io/component: controller
    app.kubernetes.io/instance: ingress-nginx
    app.kubernetes.io/managed-by: Helm
    app.kubernetes.io/name: ingress-nginx
    app.kubernetes.io/part-of: ingress-nginx
    app.kubernetes.io/version: 1.1.3
    helm.sh/chart: ingress-nginx-4.0.19
```

### 2.3 部署 Cert-Manage
```bash
helm repo add jetstack https://charts.jetstack.io
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.7.1/cert-manager.crds.yaml
helm install cert-manager jetstack/cert-manager --namespace cert-manager --create-namespace --version v1.7.1
```

为了让 ACME（Automated Certificate Management Environment，自动证书管理环境）的 CA 服务器给客户端颁发证书，客户端必须完成 challenges（挑战），以验证客户端拥有该域名。cert-manager 提供了两种挑战验证的方式：
- **HTTP01**：cert-manage 将会自动创建一个用于 ACME 验证的 ingress，在http://<YOUR_DOMAIN>/.well-known/acme-challenge/<TOKEN>（用提供的令牌替换 <TOKEN>）路径上放置指定文件。该文件包含令牌以及帐户密钥的指纹。ACME 服务器会访问这个 ingress 的路径，验证成功后，就会颁发证书，因此要求 **ingress 在公网能够访问到**。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220417133652.png)

- **DNS01**：提供在 DNS 服务商的 apiKey，cert-manage 会使用 apiKey 在 DNS 服务中创建 ACME 服务器指定的 TXT 记录，ACME 服务器访问域名的 TXT 记录，验证成功后，就会颁发证书，**ingress 不需要能被公网访问到**，只要域名对应的 IP（可以不是 ingress 的 IP） 能被 ACME 服务器访问到即可。
### 2.4 HTTP01 申请证书

查看 nginx ingress controller 对外暴露的公网 IP。
```bash
$ kubectl get svc -n ingress
NAME                                 TYPE           CLUSTER-IP       EXTERNAL-IP     PORT(S)                      AGE
ingress-nginx-controller             LoadBalancer   172.16.254.58    129.226.99.21   80:31937/TCP,443:30615/TCP   12h
ingress-nginx-controller-admission   ClusterIP      172.16.253.185   <none>          443/TCP                      12h
```

进入腾讯云 DNS 解析控制台，点击域名设置解析记录。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220417123156.png)


添加一条 A 记录，记录值是 nginx ingress controller 对外暴露的 LoadBalancer 服务的公网 IP。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220417123334.png)


验证解析。
```bash
$ nslookup ingress.se7enshare.cn

# 返回结果
nslookup ingress.se7enshare.cn
Server:		192.168.1.1
Address:	192.168.1.1#53

Non-authoritative answer:
Name:	ingress.se7enshare.cn
Address: 129.226.99.21
```

创建 Issuer。
```yaml
kind: Issuer
metadata:
  name: nginx-letsencrypt
spec:
  acme:
    email: chengzw258@163.com
    server: https://acme-v02.api.letsencrypt.org/directory
    privateKeySecretRef: # 私钥
      name: nginx-tls
    solvers:
    - http01:
        ingress:
          # cert-manager 会自动创建 Ingress 资源，并自动修改 Ingress 的资源 prod/web，以暴露校验所需的临时路径。
          # 指定自动创建的 Ingress 的 ingress class
          class: nginx
```

查看创建的 Issuer。

```bash
$ kubectl get issuer
NAME                READY   AGE
nginx-letsencrypt   True    94s
```

创建 HTTPS 加密的 Ingress。

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
      cert-manager.io/issuer: nginx-letsencrypt # 让 Cert-Manager 自动申请证书
  name: https-ingress
spec:
  ingressClassName: nginx
  rules:
  - host: ingress.se7enshare.cn  # 公网能解析的域名
    http:
      paths:
      - backend:
          service:
            name: httpserver  # 后端服务
            port:
              number: 8080
        path: /
        pathType: Prefix
  tls:
   - hosts:
      - ingress.se7enshare.cn
     secretName: nginx-cert  # cert-manager 会自动创建 secret，将申请的证书存放在里面
```

查看创建的 certificate，如果看到 READY 为 true，表示证书已经成功签发。

```bash
$ kubectl  get certificate
NAME         READY   SECRET       AGE
nginx-cert   True    nginx-cert   2m40s
```

查看证书申请日志。
```bash
$ kubectl describe certificate nginx-cert
Spec:
  Dns Names:
    ingress.se7enshare.cn
  Issuer Ref:
    Group:      cert-manager.io
    Kind:       Issuer
    Name:       nginx-letsencrypt
  Secret Name:  nginx-cert
  Usages:
    digital signature
    key encipherment
Status:
  Conditions:
    Last Transition Time:  2022-04-17T04:58:30Z
    Message:               Certificate is up to date and has not expired
    Observed Generation:   1
    Reason:                Ready
    Status:                True
    Type:                  Ready
  Not After:               2022-07-16T03:58:29Z
  Not Before:              2022-04-17T03:58:30Z
  Renewal Time:            2022-06-16T03:58:29Z
  Revision:                1
Events:
  Type    Reason     Age    From          Message
  ----    ------     ----   ----          -------
  Normal  Issuing    2m46s  cert-manager  Issuing certificate as Secret does not exist
  Normal  Generated  2m46s  cert-manager  Stored new private key in temporary Secret resource "nginx-cert-8zbsx"
  Normal  Requested  2m46s  cert-manager  Created new CertificateRequest resource "nginx-cert-6r9zb"
  Normal  Issuing    2m15s  cert-manager  The certificate has been successfully issued
```

证书和私钥将会保存在 nginx-cert 这个 Secert 中。

```bash
$ kubectl get secret nginx-cert -o yaml
```

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220417125034.png)

浏览器输入 https://ingress.se7enshare.cn/healthz 访问 httpserver 服务。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220417130232.png)



### 2.5 查看证书

根证书机构。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220417130320.png)

Let's Encrypt 是中级证书颁发机构。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220417130335.png)

签发给服务器的证书。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220417130410.png)



## 参考资料

- [Understanding Kubernetes Probes](https://blog.devgenius.io/understanding-kubernetes-probes-5daaff67599a)
- [kubernetes中的内存表示单位Mi和M的区别](https://www.jianshu.com/p/f798b02363e8)
- [云原生2期 模块8](https://shimo.im/docs/I2TeDIIuk801NVms/read)
- [使用metalLB来实现负载均衡Load Balancer](https://blog.cnscud.com/k8s/2021/09/17/k8s-metalb.html)
- [Securing NGINX-ingress](https://cert-manager.io/v0.13-docs/tutorials/acme/ingress/#step-3-assign-a-dns-name)
- [How to Set Up an Nginx Ingress with Cert-Manager on DigitalOcean Kubernetes](https://www.digitalocean.com/community/tutorials/how-to-set-up-an-nginx-ingress-with-cert-manager-on-digitalocean-kubernetes)
- [使用 cert-manager 签发免费证书](https://cloud.tencent.com/document/product/457/49368)
- [k8s 上利用 cert-manager 自动签发 TLS 证书](https://hadb.me/k8s-cert-manager-tls/)
- [ACME](https://cert-manager.io/docs/configuration/acme/)
- [Let's Encrypt 的运作方式](https://letsencrypt.org/zh-cn/how-it-works/)