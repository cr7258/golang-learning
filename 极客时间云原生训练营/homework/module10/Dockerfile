# golang 编译环境
FROM golang:1.17 AS build

# 工作目录
WORKDIR /usr/src/app
# 设置代理，提升依赖下载速度
ENV GOPROXY=https://goproxy.cn

#先复制 go.mod 和 go.sum ，然后 go mod download，
#是为了防止每次构建都会重新下载依赖包，利用 docker 构建缓存提升构建速度
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
COPY . .
RUN go build -o http-server http.go

FROM busybox:1.28
COPY --from=build /usr/src/app/http-server /bin/http-server
ENTRYPOINT ["/bin/http-server"]