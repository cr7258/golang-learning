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

# go build 是默认启用 CGO 的，
# 查看命令：go env CGO_ENABLED，
# 在 CGO 开启情况下，无论代码有没有用 CGO，都会有库依赖文件。
# 大部分的操作系统都有这些依赖，但是下面启动程序时使用的 scratch 镜像没有（为了追求小的体积），通过以下命令关闭 CGO
# 如果是使用例如 centos 这种完整操作系统的镜像，可以忽略这个步骤

# 构建镜像，简单方式，镜像大小：7.03MB
# RUN go CGO_ENABLED=0 build -o http-server

# 优化镜像大小，去除构建包的调试信息，减小 go 构建后程序体积，镜像大小：4.96MB
# -ldflags "-s -w"
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o http-server

FROM scratch
# 设置镜像标签，配合 label filter 可以查看过滤结果
LABEL company="geek" project="cloudnative"
COPY --from=build /usr/src/app/http-server /bin/http-server
ENTRYPOINT ["/bin/http-server"]