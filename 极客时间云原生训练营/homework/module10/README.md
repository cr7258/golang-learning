## 第一部分
1.为 HTTPServer 添加 0-2 秒的随机延时；
2.为 HTTPServer 项目添加延时 Metric；
3.将 HTTPServer 部署至测试集群，并完成 Prometheus 配置；
4.从 Promethus 界面中查询延时指标数据；
5。（可选）创建一个 Grafana Dashboard 展现延时分配情况。

### 部署 Prometheus
```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm -n prometheus-stack install kube-prometheus-stack prometheus-community/kube-prometheus-stack
```