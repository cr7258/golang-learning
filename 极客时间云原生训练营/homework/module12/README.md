## 第一部分
把我们的 httpserver 服务以 Istio Ingress Gateway 的形式发布出来。以下是你需要考虑的几点：
- 1.如何实现安全保证；
- 2.七层路由规则；
- 3.考虑 open tracing 的接入。

### 下载 Istio
```bash
curl -L https://istio.io/downloadIstio | sh -
```
添加环境变量：
```bash
echo "export PATH="$PATH:/software/istio-1.13.3/bin"" >> /etc/profile
source /etc/profile
```
查看版本：
```bash
$ istioctl version
no running Istio pods in "istio-system"
1.13.3
```
### 安装 Istio
```bash
istioctl install --set profile=demo -y
```

给命名空间添加标签，指示 Istio 在部署应用的时候，自动注入 Envoy 边车代理：
```bash
kubectl create ns istio-demo
kubectl label namespace istio-demo istio-injection=enabled
```

### 部署服务
```bash
kubectl apply -f httpserver/service0-v1/specs/deployment.yaml -n istio-demo
kubectl apply -f httpserver/service1/specs/deployment.yaml -n istio-demo
kubectl apply -f httpserver/service2/specs/deployment.yaml -n istio-demo
kubectl apply -f httpserver/nginx/deployment.yaml -n istio-demo
```
查看服务
```bash
$ kubectl get pod -n istio-demo
NAME                           READY   STATUS    RESTARTS   AGE
nginx-6574bb6c7f-skxbc         2/2     Running   0          2m59s
service0-v1-87798b85b-sfp74    2/2     Running   0          3m29s
service1-d88db9c98-gckkn       2/2     Running   0          2m49s
service2-6cbc7f745-g6zrc       2/2     Running   0          32s

$ kubectl get svc -n istio-demo
NAME       TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
nginx      ClusterIP   133.0.250.120   <none>        80/TCP    3m10s
service0   ClusterIP   133.0.131.207   <none>        80/TCP    3m40s
service1   ClusterIP   133.0.169.24    <none>        80/TCP    3m
service2   ClusterIP   133.0.219.11    <none>        80/TCP    43s
```
### 实现七层路由规则
#### 简易版本
```bash
kubectl apply -f istio/1.ingressgateway-simple.yaml -n istio-demo
```

查看 istio-ingressgateway 的地址：
```bash
$ kubectl get svc -n istio-system
NAME                   TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)                                                                      AGE
istio-egressgateway    ClusterIP      133.0.233.229   <none>        80/TCP,443/TCP                                                               3h45m
istio-ingressgateway   LoadBalancer   133.0.56.148    11.8.38.246   15021:31222/TCP,80:30708/TCP,443:30809/TCP,31400:31356/TCP,15443:31948/TCP   3h45m
istiod                 ClusterIP      133.0.216.152   <none>        15010/TCP,15012/TCP,443/TCP,15014/TCP                                        3h45m
```
客户端访问：
```bash
# 访问 service0 的 v1 版本
kubectl exec -it centos -- curl http://133.0.56.148 -H "Host: httpserver.io"
```
#### 复杂版本
```bash
kubectl apply -f istio/2.ingressgateway.yaml -n istio-demo
```

客户端访问：
```bash
# 访问 service0 的 v1 版本
kubectl exec -it centos -- curl http://133.0.56.148/hello -H "Host: httpserver.io"
# 访问 nginx 服务
kubectl exec -it centos -- curl http://133.0.56.148/nginx -H "Host: httpserver.io"
```

## 金丝雀发布
创建 service0 服务的 v2 版本。
```bash
kubectl apply -f httpserver/service0-v2/specs/deployment.yaml -n istio-demo
```
金丝雀发布。
```bash
kubectl apply -f istio/3.canay.yaml -n istio-demo
```
客户端访问。
```bash
# 访问 service0 的 v1 版本
kubectl exec -it centos -- curl http://133.0.56.148/hello -H "Host: httpserver.io"
# 访问 service0 的 v2 版本
kubectl exec -it centos -- curl http://133.0.56.148/hello -H "Host: httpserver.io" -H "user: chengzw"
```

## HTTPS 加密
创建证书。
```bash
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=httpserver.io Inc./CN=httpserver.io' -keyout istio/4.httpserver.io.key -out istio/4.httpserver.io.crt
kubectl create -n istio-system secret tls httpserver-credential --key=istio/4.httpserver.io.key --cert=istio/4.httpserver.io.crt --dry-run -o yaml > istio/4.secret.yaml
kubectl apply -f istio/4.secret.yaml
kubectl apply -f istio/4.ingressgateway-https.yaml -n istio-demo
```
部署 HTTPS 加密的 Gateway 和对应的 VirtualService。
```bash
kubectl apply -f istio/4.ingressgateway-https.yaml -n istio-demo
```
客户端访问。
```bash
# 访问 service0 的 v1 版本
kubectl exec -it centos -- curl https://133.0.56.148 -k -H "Host: httpsserver.io"
```

## 链路追踪
部署 Jaeger。
```bash
kubectl apply -f istio/5.jaeger.yaml
```
启动 Jaeger Dashboard。
```bash
istioctl dashboard jaeger
```

客户端访问。
```bash
# 访问 service0 的 v2 版本
kubectl exec -it centos -- curl http://133.0.56.148/hello -H "Host: httpserver.io" -H "user: chengzw"
```
查看调用链。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220514213034.png)