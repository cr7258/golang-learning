## 1. 第一部分
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

## 2. 第二部分
除了将 httpServer 应用优雅的运行在 Kubernetes 之上，我们还应该考虑如何将服务发布给对内和对外的调用方。
来尝试用 Service, Ingress 将你的服务发布给集群外部的调用方吧。
在第一部分的基础上提供更加完备的部署 spec，包括（不限于）：
- Service
- Ingress
- 可以考虑的细节
- 如何确保整个应用的高可用。
- 如何通过证书保证 httpServer 的通讯安全。

### 部署 Metallb 作为 LoadBalancer Provider

确认 strictARP 模式：如果你的网络是运行在 IPVS 模式下（默认是 iptables）, 那么需要设置 strictARP 模式，修改其中的 strictARP 为 true:
```bash
kubectl edit configmap -n kube-system kube-proxy
```

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220410162424.png)

修改完以后重建 kube-proxy 容器：
```bash
kubectl rollout restart -n kube-system daemonset kube-proxy
```
安装 Metallb：
```bash
kubectl apply -f 8.2/namespace.yaml
kubectl apply -f 8.2/metallb.yaml
```
查看创建的资源：
```bash
$ kubectl get pod -n metallb-system
NAME                         READY   STATUS    RESTARTS   AGE
controller-57fd9c5bb-kzprr   1/1     Running   0          2m49s
speaker-g5blx                1/1     Running   0          2m49s
speaker-qck42                1/1     Running   0          2m49s
speaker-sm4cg                1/1     Running   0          2m49s
```

配置地址池：
```bash
kubectl apply -f 8.2/pool.yaml
```


## 安装 Nginx Ingress Controller
官方文档：[kubernetes/ingress-nginx](https://github.com/kubernetes/ingress-nginx/tree/main/charts/ingress-nginx?accessToken=eyJhbGciOiJIUzI1NiIsImtpZCI6ImRlZmF1bHQiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOiJhY2Nlc3NfcmVzb3VyY2UiLCJleHAiOjE2NDk1NTYyMzcsImciOiJJMlRlRElJdWs4MDFOVm1zIiwiaWF0IjoxNjQ5NTU1OTM3LCJ1c2VySWQiOjY2ODcxNDMwfQ.GUxwEnq41l41BDYE0BC_SzwEMUaM7HmpHBRmSSkMQ0w)

```bash
helm install my-ingress -n ingress --create-namespace 8.2/ingress-nginx
```

查看资源：
```bash
$ kubectl get pod -n ingress
NAME                                                   READY   STATUS    RESTARTS   AGE
my-ingress-ingress-nginx-controller-7575d9d586-z577w   1/1     Running   0          69s
```

nginx ingress controller 的服务默认是 LoadBalancer 类型的，可以看到 Metallb 已经自动为 nginx ingress controller 分配了一个 IP 地址。
```bash
$ kubectl get svc -n ingress
NAME                                            TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)                      AGE
my-ingress-ingress-nginx-controller             LoadBalancer   133.0.77.13    11.8.38.247   80:30917/TCP,443:32686/TCP   13m
```
## 安装 Cert Manager
```bash
helm repo add jetstack https://charts.jetstack.io
helm repo update

kubectl apply -f 8.2/cert-manager.crds.yaml

helm install \
cert-manager 8.2/cert-manager \
--namespace cert-manager \
--create-namespace \
--version v1.7.1
```

查看资源：
```bash
$ kubectl get pod -n cert-manager
NAME                                       READY   STATUS    RESTARTS   AGE
cert-manager-76578c9687-x6bd7              1/1     Running   0          108s
cert-manager-cainjector-5c55bb7cb4-dpntw   1/1     Running   0          108s
cert-manager-webhook-556f979d7f-jbshk      1/1     Running   0          108s
```


## 创建 Service
```bash
kubectl apply -f 8.2/service.yaml
``` 

查看资源：
```bash
$ kubectl get svc
NAME         TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
httpserver   ClusterIP   133.0.31.143   <none>        8080/TCP   14s
```

## 创建 Ingress 
helm 安装 nginx ingress controller 时会创建一个 IngressClass 对象。在 Ingress 中设置 `ingressClassName: nginx` 字段或者 `kubernetes.io/ingress.class: nginx` 注解可以指定使用 nginx ingress controller。
```yaml
$ kubectl get ingressclass nginx -o yaml
apiVersion: networking.k8s.io/v1
kind: IngressClass
metadata:
  annotations:
    meta.helm.sh/release-name: my-ingress
    meta.helm.sh/release-namespace: ingress
  creationTimestamp: "2022-04-10T09:23:38Z"
  generation: 1
  labels:
    app.kubernetes.io/component: controller
    app.kubernetes.io/instance: my-ingress
    app.kubernetes.io/managed-by: Helm
    app.kubernetes.io/name: ingress-nginx
    app.kubernetes.io/part-of: ingress-nginx
    app.kubernetes.io/version: 1.1.3
    helm.sh/chart: ingress-nginx-4.0.19
  name: nginx
  resourceVersion: "2542082"
  uid: fd701cc7-78a7-41ee-8e8d-101ce6ef3bd5
spec:
  controller: k8s.io/ingress-nginx
```
### HTTP 访问
创建 Ingress：
```yaml
kubectl apply -f 8.2/http-ingress.yaml
```

查看资源：
```bash
$ kubectl describe ingress http-ingress
Name:             http-ingress
Namespace:        default
Address:          11.8.38.247
Default backend:  default-http-backend:80 (<error: endpoints "default-http-backend" not found>)
Rules:
  Host             Path  Backends
  ----             ----  --------
  www.example.com  
                   /   httpserver:8080 (77.0.108.22:8080,77.0.47.95:8080,77.0.47.96:8080)
Annotations:       <none>
Events:
  Type    Reason  Age                 From                      Message
  ----    ------  ----                ----                      -------
  Normal  Sync    68s (x3 over 7m8s)  nginx-ingress-controller  Scheduled for sync
```

客户端访问，注意不需要带 30917 端口，直接访问 nginx ingress controller 的 80 端口。
```bash
curl -H "host:www.example.com" http://11.8.38.247/healthz
# 返回结果
working
```

### HTTPS 访问
先删除 http ingress：
```bash
kubectl delete -f 8.2/http-ingress.yaml
```
签发证书 CA 的配置：
```bash
kubectl apply -f 8.2/issuer.yaml
```
查看资源：
```bash
$ kubectl get issuers -n ingress
NAME                READY   AGE
nginx-letsencrypt   True    5s
```

创建 Ingress：
```bash
kubectl apply -f 8.2/https-ingress.yaml
```


客户端访问：
```bash
curl -H "host:www.example.com" -k https://11.8.38.247/healthz
# 返回结果
working
```
## 参考资料
- [Understanding Kubernetes Probes](https://blog.devgenius.io/understanding-kubernetes-probes-5daaff67599a)
- [kubernetes中的内存表示单位Mi和M的区别](https://www.jianshu.com/p/f798b02363e8)
- [云原生2期 模块8](https://shimo.im/docs/I2TeDIIuk801NVms/read)
- [使用metalLB来实现负载均衡Load Balancer](https://blog.cnscud.com/k8s/2021/09/17/k8s-metalb.html)
- [Securing NGINX-ingress](https://cert-manager.io/v0.13-docs/tutorials/acme/ingress/#step-3-assign-a-dns-name)
- [How to Set Up an Nginx Ingress with Cert-Manager on DigitalOcean Kubernetes](https://www.digitalocean.com/community/tutorials/how-to-set-up-an-nginx-ingress-with-cert-manager-on-digitalocean-kubernetes)