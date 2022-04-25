## 第一部分
- 1.为 HTTPServer 添加 0-2 秒的随机延时；
- 2.为 HTTPServer 项目添加延时 Metric；
- 3.将 HTTPServer 部署至测试集群，并完成 Prometheus 配置；
- 4.从 Promethus 界面中查询延时指标数据；
- 5.（可选）创建一个 Grafana Dashboard 展现延时分配情况。


### 部署 HTTPServer
```bash
kubectl apply -f httpserver.yaml
```
查看 HTTPServer：
```bash
$ kubectl get pod
root@cluster02-1:/root #kubectl get pod -n cloudnative
NAME                          READY   STATUS    RESTARTS   AGE
httpserver-69755968b5-98mjv   1/1     Running   0          7s
httpserver-69755968b5-rx2px   1/1     Running   0          7s
httpserver-69755968b5-st8b9   1/1     Running   0          7s
```
### 部署 Prometheus
```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm -n prometheus-stack install kube-prometheus-stack prometheus-community/kube-prometheus-stack --create-namespace
```
查看 Prometheus：
```bash
$ kubectl get pod -n prometheus-stack
NAME                                                        READY   STATUS    RESTARTS   AGE
alertmanager-kube-prometheus-stack-alertmanager-0           2/2     Running   0          18m
kube-prometheus-stack-grafana-697c5ff6bf-6ln9w              3/3     Running   0          18m
kube-prometheus-stack-kube-state-metrics-6c5d86887c-krnk9   1/1     Running   0          18m
kube-prometheus-stack-operator-768f459cd8-g69rn             1/1     Running   0          18m
kube-prometheus-stack-prometheus-node-exporter-j8w25        1/1     Running   0          18m
kube-prometheus-stack-prometheus-node-exporter-jbmk4        1/1     Running   0          18m
kube-prometheus-stack-prometheus-node-exporter-xl5nh        1/1     Running   0          18m
prometheus-kube-prometheus-stack-prometheus-0               2/2     Running   0          18m
```
将 Prometheus 和 Grafana Service 通过 NodePort 方式暴露：
```bash
kubectl patch -n prometheus-stack svc kube-prometheus-stack-prometheus  -p '{"spec":{"type":"NodePort"}}'
kubectl patch -n prometheus-stack svc kube-prometheus-stack-grafana  -p '{"spec":{"type":"NodePort"}}'
```
查看服务：
```bash
$ kubectl  get svc -n prometheus-stack
NAME                                             TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                      AGE
alertmanager-operated                            ClusterIP   None            <none>        9093/TCP,9094/TCP,9094/UDP   28m
kube-prometheus-stack-alertmanager               ClusterIP   133.0.189.56    <none>        9093/TCP                     28m
kube-prometheus-stack-grafana                    NodePort    133.0.96.25     <none>        80:30338/TCP                 28m
kube-prometheus-stack-kube-state-metrics         ClusterIP   133.0.195.91    <none>        8080/TCP                     28m
kube-prometheus-stack-operator                   ClusterIP   133.0.183.50    <none>        443/TCP                      28m
kube-prometheus-stack-prometheus                 NodePort    133.0.180.147   <none>        9090:30009/TCP               28m
kube-prometheus-stack-prometheus-node-exporter   ClusterIP   133.0.3.117     <none>        9100/TCP                     28m
prometheus-operated                              ClusterIP   None            <none>        9090/TCP                     28m
```

## Promethues 创建发现规则。
```bash
kubectl create secret generic additional-configs --from-file=prometheus-additional.yaml -n prometheus-stack
```

**注入 Prometheus**：然后我们需要在声明 prometheus 的资源对象文件中通过 additionalScrapeConfigs 属性添加上这个额外的配置：
```yaml
apiVersion: v1
items:
- apiVersion: monitoring.coreos.com/v1
  kind: Prometheus
  metadata:
    annotations:
      meta.helm.sh/release-name: kube-prometheus-stack
      meta.helm.sh/release-namespace: prometheus-stack
    labels:
      app: kube-prometheus-stack-prometheus
      app.kubernetes.io/instance: kube-prometheus-stack
      app.kubernetes.io/managed-by: Helm
      app.kubernetes.io/part-of: kube-prometheus-stack
      app.kubernetes.io/version: 34.10.0
      chart: kube-prometheus-stack-34.10.0
      heritage: Helm
      release: kube-prometheus-stack
    name: kube-prometheus-stack-prometheus
    namespace: prometheus-stack
  spec:
    additionalScrapeConfigs:  # 自动发现的配置文件
      key: prometheus-additional.yaml
      name: additional-configs
    alerting:
      alertmanagers:
      - apiVersion: v2
        name: kube-prometheus-stack-alertmanager
        namespace: prometheus-stack
        pathPrefix: /
        port: http-web
    enableAdminAPI: false
    externalUrl: http://kube-prometheus-stack-prometheus.prometheus-stack:9090
    image: quay.io/prometheus/prometheus:v2.34.0
    listenLocal: false
    logFormat: logfmt
    logLevel: info
    paused: false
    podMonitorNamespaceSelector: {}
    podMonitorSelector:
      matchLabels:
        release: kube-prometheus-stack
    portName: http-web
    probeNamespaceSelector: {}
    probeSelector:
      matchLabels:
        release: kube-prometheus-stack
    replicas: 1
    retention: 10d
    routePrefix: /
    ruleNamespaceSelector: {}
    ruleSelector:
      matchLabels:
        release: kube-prometheus-stack
    securityContext:
      fsGroup: 2000
      runAsGroup: 2000
      runAsNonRoot: true
      runAsUser: 1000
    serviceAccountName: kube-prometheus-stack-prometheus
    serviceMonitorNamespaceSelector: {}
    serviceMonitorSelector:
      matchLabels:
        release: kube-prometheus-stack
    shards: 1
    version: v2.34.0
kind: List
metadata:
  resourceVersion: ""
  selfLink: ""
```

在 Prometheus 的 Target 应当能看到 HTTPServer。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220425124845.png)

在 Prometheus 上查询指标。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220425125223.png)
