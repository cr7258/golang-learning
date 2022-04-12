module github.com/cncamp/golang

go 1.16

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.6
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
require (
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529
	github.com/stretchr/testify v1.7.0
)
