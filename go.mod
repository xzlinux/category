module github.com/xzlinux/category

go 1.15

require (
	github.com/golang/protobuf v1.4.3
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/config/source/consul/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/prometheus/common v0.6.0
	github.com/xzlinux/common v0.0.0-20210216125852-8d4edfe894ee
	google.golang.org/protobuf v1.25.0

)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
