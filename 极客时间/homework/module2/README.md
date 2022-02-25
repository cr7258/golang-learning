编写一个 HTTP 服务器：
1.接收客户端 request，并将 request 中带的 header 写入 response header。
2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header。
3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出。
4.当访问 /healthz 时，应返回 200。
5.针对 http 服务的 pprof。



启动之前先设置 VERSION 环境变量。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220225224048.png)

启动 HTTP 服务。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220225224127.png)

客户端发起请求设置 HTTP header。

```bash
$ curl 192.168.1.115 -H "name:chengzw" -H "Accept:Application/json" -I
# 响应结果
HTTP/1.1 200 OK
Accept: Application/json
Name: chengzw
User-Agent: curl/7.71.1
Version: 1
Date: Fri, 25 Feb 2022 14:37:10 GMT

```

HTTP 服务端输出客户端 IP 和 HTTP 返回码。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220225224229.png)

客户端请求 /healthz 接口。

```bash
$ curl 192.168.1.115/healthz
# 响应结果
ok
```

浏览器输入 http://192.168.1.115/debug/pprof 访问 pprod 界面。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220225224501.png)

点击 profile 可以生成并下载 profile 文件，在 speedscope UI(https://www.speedscope.app/) 网站上可以进行直观的分析。

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220225224736.png)

![](https://chengzw258.oss-cn-beijing.aliyuncs.com/Article/20220225224746.png)