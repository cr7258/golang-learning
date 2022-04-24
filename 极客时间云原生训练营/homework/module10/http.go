package main

import (
	"fmt"
	"github.com/Am2901/httpserver/src/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
	"time"
)

/**
1.为 httpserver 添加 0-2 秒的随机延时
2.为 httpserver 添加延时的 metric
*/
func main() {
	//flag.Set("v", "2") // glog 读取 v 来决定日志级别
	//glog.V(2).Info("Starting http server...")
	fmt.Println("Starting http server...")
	// 注册 handle 处理函数
	//http.HandleFunc("/", rootHandler)
	//http.HandleFunc("/healthz", healthz)
	// 监听本地 80 端口
	//err := http.ListenAndServe(":80", nil)
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthz)
	// 5.针对 http 服务的 pprof
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	// 增加延时
	mux.HandleFunc("/delay", delay)
	// metric
	mux.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal(err)
	}
}

// 定义 handle 处理函数
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// 1.将 request 中带的 header 写入 response header
	fmt.Println("======== Get request header ========")
	for k, v := range r.Header {
		fmt.Printf("%s=%s\n", k, v)
		w.Header().Set(k, v[0])
	}

	// 2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	version := os.Getenv("VERSION")
	w.Header().Set("Version", version)
	fmt.Println("====================================")

	// 3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	remote := strings.Split(r.RemoteAddr, ":")
	ip := remote[0]
	port := remote[1]
	// 设置响应码
	w.WriteHeader(200)
	fmt.Println("source_address: "+ip+" source_port: "+port, " status_code: 200")
}

// 4.当访问 localhost/healthz 时，应返回 200
func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok\n")
}

// 增加延时
func delay(w http.ResponseWriter, r *http.Request) {
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	randInt := rand.Intn(2000)
	time.Sleep(time.Millisecond * time.Duration(randInt))
	w.Write([]byte(fmt.Sprintf("duration: %d", randInt)))
}
